package main

import (
	"fmt"
	"modak-rated-limited-challenge/internal/controller"
	"modak-rated-limited-challenge/internal/repository"
	"modak-rated-limited-challenge/internal/service"
	"testing"
)

func TestRangeLimited(t *testing.T) {
	ctrl := controller.NewController(service.NewService(repository.NewRepository()))
	_ = ctrl.SetRule("news", "2s", 2)
	_ = ctrl.SetRule("Status", "1s", 1)
	_ = ctrl.SetRule("Marketing", "1h", 3)

	tests := 10
	result := make(chan error, tests*2)
	defer close(result)

	for i := 1; i < tests; i++ {

		/*	n := rand.Intn(3) // n will be between 0 and 3
			time.Sleep(time.Duration(n) * time.Second)
		*/
		go func(i int) {
			err := ctrl.SendNotification("news", fmt.Sprintf("user%d", i), "news 1")
			if err != nil {
				fmt.Println(err)
			}
			result <- err
		}(i)

		go func(i int) {
			err := ctrl.SendNotification("Status", fmt.Sprintf("user%d", i), "Status 1")
			if err != nil {
				fmt.Println(err)
			}
			result <- err
		}(i)

	}
	for o := 1; o < 19; o++ {
		<-result
	}
}
