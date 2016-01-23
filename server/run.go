package main

import "handler"

func main() {
	go handler.HandleHTTP()
	for {}
}