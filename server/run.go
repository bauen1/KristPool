package main

import "handler"

func main() {
	go handler.HandleHTTP()
	go handler.HandleTCP()
	for {}
}