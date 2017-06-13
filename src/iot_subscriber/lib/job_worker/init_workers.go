package job_worker

import (
	"github.com/benmanns/goworker"
	"fmt"
)

func Init()  {
	settings := goworker.WorkerSettings{
		URI:            "redis://localhost:6379/",
		Connections:    10,
		Queues:         []string{"packet_subscriber_queue"},
		UseNumber:      true,
		ExitOnComplete: false,
		Concurrency:    2,
		Namespace:      "goiot_subscriber:",
		Interval:       5.0,
	}
	goworker.SetSettings(settings)
	goworker.Register("subscribers", ProcessData)
	if err := goworker.Work(); err != nil {
		fmt.Println("Error:", err)
	}
}
