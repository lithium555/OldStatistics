package main

import (
	"OldStatisticms/api"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

func main(){
	var conn *grpc.ClientConn

	conn, err := grpc.Dial(":7777", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewStatisticsClient(conn)

	wg := &sync.WaitGroup{}
	for i := 1; i < 1000; i++ {
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup) {
			response, err := c.GetStatistics(context.Background(), &api.TaskMessage{Date: time.Now().Format("2006-01-02")})
			if err != nil{
				log.Fatalf("Error, when we calling function GetStatistics: '%v'", err)
			}
			log.Println("---------------------------------------------------------------")
			log.Println("Response from the server: ")
			log.Printf("response.Revenue = '%v'\n", response.Revenue)
			log.Printf("response.PartnerId = '%v'\n", response.PartnerId)
			log.Printf("response.EventId = '%v'\n", response.EventId)
			log.Printf("response.Date = '%v'\n", response.Date)
			log.Printf("response.Time = '%v'\n", response.Time)
			log.Println("---------------------------------------------------------------")

			defer wg.Done()
		}(i, wg)
	}
	wg.Wait()
}