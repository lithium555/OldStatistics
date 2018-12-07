package api

import (
	"fmt"
	"golang.org/x/net/context"
	"log"
	"math/rand"
	"time"
)

// StructgRPC represents the gRPC server
type StructgRPC struct {
	counter int
}

func (s *StructgRPC) GetStatistics(ctx context.Context, data *TaskMessage)(*TaskMessage, error){
	log.Println("function GetStatistics received message")
	currentTime := time.Now()
	s.counter++
	fmt.Printf("s.counter:= '%v'\n",s.counter)


	return &TaskMessage{
		Date: time.Now().Format("2006-01-02"),
		Time: currentTime.Format(time.RFC3339),
		EventId: rand.Int63(),
		PartnerId: data.PartnerId,
		Revenue: int32(s.counter),
	}, nil
}