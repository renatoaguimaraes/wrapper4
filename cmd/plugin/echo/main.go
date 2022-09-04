package main

import "log"

func GetPlugin() interface{} {
	return &echoPlugin{}
}

type echoPlugin struct{}

func (p echoPlugin) Run() {
	log.Println("Echo demo plugin")
}
