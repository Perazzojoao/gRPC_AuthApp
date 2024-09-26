package main

import "mail-service/server"

func main() {
	app := server.NewServer()
	app.GrpcListen()
}
