package main

import "authApp/server"

func main() {
	server := server.NewServer()
	server.GrpcListen()
}
