A small utility to monitor n8n workers

![CleanShot 2023-10-06 at 17 11 00](https://github.com/flipswitchingmonkey/n8nworkerview/assets/6930367/2fb2fd2f-84d0-4c03-8606-3469331690cf)

To configure, create a `.env` file next to the executable or use the environment to set these options:

```
REDIS_HOST=127.0.0.1
REDIS_PORT=6379
REDIS_PASS=
REDIS_USER=
PUBSUB_COMMAND_CHANNEL=n8n.commands
PUBSUB_WORKER_RESPONSE_CHANNEL=n8n.worker-response
```

At a regular interval (1s by default) the tool will send a message to the command channel asking for `GetStatus`. Workers should then respon on the worker response channel with a corresponding message, containing a payload of:

```
			payload: {
				workerId: string;
				runningJobs: string[];
				runningJobsSummary: WorkerJobStatusSummary[];
				freeMem: number;
				totalMem: number;
				uptime: number;
				loadAvg: number[];
				cpus: string;
				arch: string;
				platform: NodeJS.Platform;
				hostname: string;
				net: string[];
			};
```

### Keyboard shortcuts

`Esc | q | ctrl-c` - Exit application
`Tab` - switch between worker list and details
`Enter | Space` - toggle jobs list and worker details view
`+ | -` - increase or decrease polling rate
