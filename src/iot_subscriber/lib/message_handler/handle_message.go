package message_handler

import (
	"iot_subscriber/lib/formatter"
	"github.com/revel/revel"
	"iot/lib/utils"
	"iot_subscriber/lib/dbutils"
	"time"
)

func Handle(input interface{})  {
	data := formatter.GetMessageFromInterface(input)
	revel.INFO.Println(data)


	packet_type := utils.ToInt((data["packet_type"].Value))
	revel.INFO.Println("packet_type=>",packet_type)
	db := dbutils.DBCONN

	//Update Status of SGU

	sgu := dbutils.Sgu{}
	incoming_sgu := dbutils.Sgu{}
	incoming_sgu.Sgu_id = utils.ToUint64(data["incoming_sgu_id"].Value)

	db.Where(incoming_sgu).Assign(dbutils.Sgu{Is_connected: 1}).FirstOrCreate(&sgu)

	switch packet_type {

	case 0x003: {
		iterate := utils.ToInt(data["iterate"].Value)

		incoming_scu := dbutils.Scu{}

		for i:=0;i<iterate;i++{
			if i==0{
				scu := dbutils.Scu{}
				incoming_scu.Scu_id = utils.ToUint64(data["scuid"].Value)
				db.Where(incoming_scu).Assign(dbutils.Scu{Sgu_id: sgu.Sgu_id, Is_connected: 1}).FirstOrCreate(&scu)
			}else{
				scu := dbutils.Scu{}
				incoming_scu.Scu_id = utils.ToUint64(data["scuid_" + utils.ToStr(i)].Value)
				db.Where(incoming_scu).Assign(dbutils.Scu{Sgu_id: sgu.Sgu_id, Is_connected: 1}).FirstOrCreate(&scu)
			}
		}
	}

	case 0x3001:{
		if utils.ToInt(data["status"].Value) != 0{
			revel.INFO.Println("Packet :", packet_type, " Received with INVALID Status :", data["status"].Value)
			return
		}
		incoming_scu := dbutils.Scu{}

		incoming_scu.Scu_id = utils.ToUint64(data["scuid"].Value)

		pwm := utils.ToInt(data["pwm"].Value)

		scu := dbutils.Scu{}
		db.Where(incoming_scu).Assign(dbutils.Scu{Sgu_id: sgu.Sgu_id, Pwm: pwm, Is_connected: 1, Status_updated_at: time.Now()}).FirstOrCreate(&scu)


	}

	case 0xe000:{
		db.Where(incoming_sgu).Assign(dbutils.Sgu{Is_connected: 1}).FirstOrCreate(&sgu)

		iterate := utils.ToInt(data["iterate"].Value)

		incoming_scu := dbutils.Scu{}

		for i:=0;i<iterate;i++{
			if i==0{
				scu := dbutils.Scu{}
				incoming_scu.Scu_id = utils.ToUint64(data["scuid"].Value)
				db.Where(incoming_scu).Assign(dbutils.Scu{Sgu_id: sgu.Sgu_id, Is_connected: 1}).FirstOrCreate(&scu)
			}else{
				scu := dbutils.Scu{}
				incoming_scu.Scu_id = utils.ToUint64(data["scuid_" + utils.ToStr(i)].Value)
				db.Where(incoming_scu).Assign(dbutils.Scu{Sgu_id: sgu.Sgu_id, Is_connected: 1}).FirstOrCreate(&scu)
			}
		}
	}
	case 0x8001:{
		if utils.ToInt(data["status"].Value) != 0{
			revel.INFO.Println("Packet :", packet_type, " Received with INVALID Status :", data["status"].Value)
			return
		}
		incoming_scu := dbutils.Scu{}

		incoming_scu.Scu_id = utils.ToUint64(data["scuid"].Value)
		scu := dbutils.Scu{}
		db.Where(incoming_scu).Assign(dbutils.Scu{Sgu_id: sgu.Sgu_id, Is_connected: 1}).FirstOrCreate(&scu)

		att_schedule := dbutils.Attached_Schedules{}
		att_schedule.Scheduling_id = utils.ToInt(data["scheduling_id"].Value)
		att_schedule.Pwm = utils.ToInt(data["pwm"].Value)
		att_schedule.Scu_id = incoming_scu.Scu_id

		expr := utils.ToStr(data["expression"].Value)

		saved_schedule := dbutils.Attached_Schedules{}

		db.Where(att_schedule).Assign(dbutils.Attached_Schedules{Expression: expr, Status_updated_at: time.Now()}).FirstOrCreate(&saved_schedule)
	}

	}
}
