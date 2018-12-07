package main

import (
	"OldStatisticms/api"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main(){
	fmt.Println("This is a Server")
	// create a listener on TCP port 7777
	listener, err := net.Listen("tcp", "127.0.0.1:8085")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server is listening...")
	// create a server instance
	s := api.StructgRPC{}
		// create a gRPC server object
	grpcServer := grpc.NewServer()
	//регистрируем сервис Statistics на сервере
	api.RegisterStatisticsServer(grpcServer, &s)
	//// start the server
	//if err := grpcServer.Serve(listener); err != nil {
	//	log.Fatalf("failed to serve: %s", err)
	//}
	var counter int
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("Accept(), error:'%v'\n", err)
			conn.Close()
			continue
		}
		go handleConnection(conn, counter)
	}
}

func handleConnection(conn net.Conn, counter int){
	defer conn.Close()
	for{
		input := make([]byte, (1024*4))
		data, err := conn.Read(input)
		if err != nil{
			fmt.Println("Read error:", err)
			break
		}
		fmt.Printf("data = '%v'\n", data)
		fmt.Println("We get a connection")
		counter++
		fmt.Printf("Counter ov Evets: '%v'\n", counter)
	}
}
