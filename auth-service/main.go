package main

import "auth-service/server"

func main() {
	app := server.NewServer()
	app.GrpcListen()
}
