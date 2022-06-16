package main

import (
	"fmt"

	"github.com/hackathon-22-spring-16/backend/model"
	"github.com/hackathon-22-spring-16/backend/router"
)

func main() {
	_, err := model.InitDB()
	if err != nil {
		panic(fmt.Errorf("db error: %w", err))
	}

	router.SetRouting()
}
