package job_worker

import (
	"github.com/jrallison/go-workers"
)

func Init()  {
	workers.Configure(map[string]string{
		// location of redis instance
		"server":  "localhost:6379",
		// instance of the database
		"database":  "1",
		// number of connections to keep open with redis
		"pool":    "5",
		// unique process id for this instance of workers (for proper recovery of inprogress jobs on crash)
		"process": "iot",
	})

	workers.Process("packets", ProcessPacket,2)
	workers.Process("sender_queue", ProcessPacket,2)
	go workers.StatsServer(8080)

	// Blocks until process is told to exit via unix signal
	workers.Run()
}
