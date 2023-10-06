package main

import (
	"context"
)

var ctx = context.Background()

func main() {
	ConnectToRedis()
	SendGetStatusInterval()
	ShowUi()	
	StopInterval()
}
