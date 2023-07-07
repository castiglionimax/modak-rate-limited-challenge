package main

import (
	"fmt"
	"math/rand"
	"modak-rated-limited-challenge/internal/controller"
	"modak-rated-limited-challenge/internal/repository"
	"modak-rated-limited-challenge/internal/service"
	"testing"
	"time"
)

func TestRangeLimited(t *testing.T) {
	ctrl := controller.NewController(service.NewService(repository.NewRepository()))
	_ = ctrl.SetRule("news", "ms", 2)
	_ = ctrl.SetRule("Status", "1", 1)
	_ = ctrl.SetRule("Marketing", "1h", 3)

	tests := 100
	result := make(chan error, tests)
	defer close(result)

	for i := 1; i < tests; i++ {

		n := rand.Intn(3) // n will be between 0 and 3
		time.Sleep(time.Duration(n) * time.Millisecond)

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
