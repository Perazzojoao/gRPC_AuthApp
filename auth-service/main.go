package main

import "auth-service/api/server"

func main() {
	server := server.NewServer()
	server.GrpcListen()
}
