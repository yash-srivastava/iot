package controllers

import (
	"github.com/revel/revel"
	"iot_client/lib/dial_tcp"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	go dial_tcp.Connect()
	return c.Render()
}
