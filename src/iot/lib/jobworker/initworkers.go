package jobworker

import (

	"github.com/benmanns/goworker"
)

//Init Initialize Goworker
func Init() {
	settings := goworker.WorkerSettings{
		URI:            "redis://localhost:6379/",
		Connections:    10,
		Queues:         []string{"packet_queue", "sender_queue"},
		UseNumber:      true,
		ExitOnComplete: false,
		Concurrency:    2,
		Namespace:      "goiot:",
		Interval:       5.0,
	}
	goworker.SetSettings(settings)
	goworker.Register("packets", ProcessPacket)
	if err := goworker.Work(); err != nil {
		revel.Error.Println("Error:", err)
	}
}
