package main

import (
	"fmt"
	"math/rand"
	"modak-rated-limited-challenge/internal/controller"
	"modak-rated-limited-challenge/internal/repository"
	"modak-rated-limited-challenge/internal/service"
	"time"
)

func main() {
	ctrl := controller.NewController(service.NewService(repository.NewRepository()))
	_ = ctrl.SetRule("news", "1s", 2)
	_ = ctrl.SetRule("Status", "1", 1)
	_ = ctrl.SetRule("Marketing", "1h", 3)
	tests := 100
	result := make(chan error, tests)
	defer close(result)

	for i := 1; i < tests; i++ {

		n := rand.Intn(3) // n will be between 0 and 3
		time.Sleep(time.Duration(n) * time.Second)

		go func(i int) {
			err := ctrl.SendNotification("news", fmt.Sprintf("user%d", i), "news 1")
			if err != nil {
				fmt.Println(err)
			}
			result <- err
		}(i)

	}
	for o := 1; o < tests; o++ {
		<-result
	}

}

/*
News: not more than 1 per day for each recipient
Status: not more than 2 per minute for each recipient
Marketing: not more than 3 per hour for each recipient

   service.send("news", "user", "news 1");
   service.send("news", "user", "news 2");
   service.send("news", "user", "news 3");
   service.send("news", "another user", "news 1");
   service.send("update", "user", "update 1");
*/
