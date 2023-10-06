package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)


var rdb *redis.Client
var workerPubSub *redis.PubSub
var intervalTicker *time.Ticker
var intervalDoneChan chan bool
var intervalDuration time.Duration = 1 * time.Second

type RedisServiceCommand struct {
	SenderId string `json:"senderId"`
	Command string `json:"command"`
	Payload string `json:"payload"`
	Targets []string `json:"targets"`
};

type RedisWorkerResponse struct {
	WorkerId string `json:"workerId"`
	Command string `json:"command"`
	Payload RedisWorkerResponsePayload `json:"payload"`
};

type RedisWorkerResponsePayload struct {
	WorkerId string `json:"workerId"`
	RunningJobs []string `json:"runningJobs"`
	RunningJobsSummary []ExecutionsCurrentSummary `json:"runningJobsSummary"`
	FreeMem float32 `json:"freeMem"`
	TotalMem float32 `json:"totalMem"`
	Uptime float32 `json:"uptime"`
	LoadAvg []float32 `json:"loadAvg"`
	Cpus string `json:"cpus"`
	Arch string `json:"arch"`
	Platform string `json:"platform"`
	Hostname string `json:"hostname"`
	Net []string `json:"net"`
};

type ExecutionsCurrentSummary struct {
	JobId string `json:"jobId"`
	ExecutionId string `json:"executionId"`
	WorkflowName string `json:"workflowName"`
	WorkflowId string `json:"workflowId"`
	RetryOf string `json:"retryOf"`
	StartedAt string `json:"startedAt"`
	Status string `json:"status"`
}

func connect() {
	errDotEnv := godotenv.Load()
	if errDotEnv != nil {
		log.Fatal("Could not load .env file")
	}

	rdb = redis.NewClient(&redis.Options{
		Addr:         getEnv("REDIS_HOST", "127.0.0.1")+":"+getEnv("REDIS_PORT", ":6379"),
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
		Username: 		getEnv("REDIS_USER", ""),
		Password: 		getEnv("REDIS_PASS", ""),
	})

	workerPubSub = rdb.Subscribe(ctx, getEnv("PUBSUB_WORKER_RESPONSE_CHANNEL", "n8n.worker-response"))
}

func listenToMessages() {
	for {
		msg, err := workerPubSub.ReceiveMessage(ctx)
		if err != nil  {
			log.Println(err)
		}
		var m RedisWorkerResponse
		json.Unmarshal([]byte(msg.Payload), &m);
		switch m.Command {
		case "getStatus":
			Ui.Send(m)
		case "getId":
			Ui.Send(m)
		}
	}
}

func SendGetStatusInterval() {
	intervalTicker = time.NewTicker(intervalDuration)
	intervalDoneChan = make(chan bool)
	go func() {
		for {
			select {
			case <-intervalDoneChan:
				return
			case <-intervalTicker.C:
				SendGetStatus()
			}
		}
	}()
}

func StopInterval() {
	intervalTicker.Stop()
	intervalDoneChan <- true
}

func IncreaseInterval() {
	intervalDuration += 100 * time.Millisecond
	ResetInterval()
}

func DecreaseInterval() {
	if (intervalDuration > 200*time.Millisecond) {
		intervalDuration -= 100 * time.Millisecond
		ResetInterval()
	}
}

func ResetInterval() {
	intervalTicker.Reset(intervalDuration)
}

func SendCommandMessage(msg RedisServiceCommand) {
	bytes, marshalErr := json.Marshal(msg)
	if marshalErr != nil {
		log.Fatal(marshalErr)
	}
	err := rdb.Publish(ctx, getEnv("PUBSUB_COMMAND_CHANNEL", "n8n.commands"), string(bytes)).Err()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Could not send command message")
	}
}

func SendGetId() {
		msg := RedisServiceCommand{
		SenderId: "workergo",
		Command: "getId",
		Payload: "",
	}
	SendCommandMessage(msg);
}

func SendGetStatus() {
		msg := RedisServiceCommand{
		SenderId: "workergo",
		Command: "getStatus",
	}
	SendCommandMessage(msg);
}

func SendGetStatusByName(target string) {
		msg := RedisServiceCommand{
		SenderId: "workergo",
		Command: "getStatus",
		Targets: []string{target},
	}
	SendCommandMessage(msg);
}

func ConnectToRedis() {
	connect()
	pong, err := rdb.Ping(ctx).Result()
	if err != nil || pong != "PONG" {
		log.Fatal("Could not contact Redis db (no pong to my ping)")
		log.Fatal(err)
	}
	go listenToMessages();
}
