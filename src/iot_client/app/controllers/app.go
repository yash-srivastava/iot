package controllers

import (
	"github.com/revel/revel"
	"iot_client/lib/kafka"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	//go dial_tcp.Connect()
	go kafka.NewConsumer()
	return c.Render()
}
