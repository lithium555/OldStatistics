package main

import (
	"OldStatisticms/api"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"sync"
	"time"
)

func main(){
	conn, err := grpc.Dial("127.0.0.1:8085", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := api.NewStatisticsClient(conn)
	wg := &sync.WaitGroup{}
	for i := 1; i < 1000; i++ {
		wg.Add(1)
		go func(index int, wg *sync.WaitGroup) {
			ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
			defer cancel()
			select {
			case <-time.After(1 * time.Second):
				fmt.Println("overslept")
			case <-ctx.Done():
				fmt.Println(ctx.Err()) // prints "context deadline exceeded"
			}

			response, err := c.GetStatistics(ctx, &api.TaskMessage{Date: time.Now().Format("2006-01-02"), Revenue: int32(i)}, grpc.FailFast(false))
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
			cancel()
			defer wg.Done()
		}(i, wg)
	}
	wg.Wait()
}
