package sender

import (
	"iot/lib/formatter"
	"time"
	"iot/lib/utils"
	"iot/conf"
	"strings"
	"github.com/jrallison/go-workers"
)

type Response struct {
	Success bool ` json:"success" `
	Message string ` json:"message" `
	Data   interface{} ` json:"data" `
}

type Packet_3000 struct {
	SguId		uint64	` json:"sguid" `
	ScuId 		uint64	` json:"scuid" `
	GetSet		int	` json:"get_set" `
	Pwm 		int 	` json:"pwm" `
	Op1 		int 	` json:"op1" `
	Op2 		int 	` json:"op2" `
	Op3 		int 	` json:"op3" `
	Op4 		int 	` json:"op4" `
	Retry		int 	` json:"retry" `
	RetryDelay	int 	` json:"retry_delay" `

}

type Packet_8000 struct {
	SguId		uint64	` json:"sguid" `
	ScuId 		uint64	` json:"scuid" `
	GetSet		int	` json:"get_set" `
	SchedulingId 	int 	` json:"scheduling_id" `
	Pwm 		int 	` json:"pwm" `
	Expression 	string 	` json:"expression" `
}

func HandlePacket(packet_type int, params interface{}) Response{
	response := Response{}
	switch packet_type {
	case 0x3000:{
		packet := Packet_3000{}
		err := formatter.GetStructFromInterface(params, &packet)
		if err!=nil{
			response.Success = false
			response.Message = "Invalid Packet Structure"
		}
		conf.Retry_3000.Set(Get300Hash(packet), true)
		go send_with_retry_3000(packet)
		response.Success = true
		response.Message = "Packet Enqueued Successfully"
	}
	case 0x8000:{
		packet := Packet_8000{}
		err := formatter.GetStructFromInterface(params, &packet)
		if err!=nil{
			response.Success = false
			response.Message = "Invalid Packet Structure"
		}
		go send8000(packet)
		response.Success = true
		response.Message = "Packet Enqueued Successfully"
	}

	}
	return response
}

func send_with_retry_3000(params Packet_3000){
	job_params := make(map[string]interface{})

	sleep_du,_ := time.ParseDuration(utils.ToStr(params.RetryDelay)+"s")


	for i:=0; i< params.Retry+1 ;i++  {

		continue_retrial,_ := conf.Retry_3000.Get(Get300Hash(params))

		if !continue_retrial.(bool) {
			break
		}

		job_params["action"] = "send_3000"
		job_params["params"] = params

		workers.Enqueue("sender_queue", "send_3000", job_params)

		time.Sleep(sleep_du)
	}
}

func Get300Hash(params Packet_3000) string{
	var result []string

	result = append(result,utils.ToStr(params.SguId))
	result = append(result,utils.ToStr(params.ScuId))
	result = append(result,utils.ToStr(params.GetSet))

	return strings.Join(result,"#")
}

func send8000(params Packet_8000){
	job_params := make(map[string]interface{}, 2)
	job_params["action"] = "send_8000"
	job_params["params"] = params

	workers.Enqueue("sender_queue", "send_8000", job_params)


}