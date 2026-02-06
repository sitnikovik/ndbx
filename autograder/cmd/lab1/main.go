package main

import (
	lab0 "github.com/sitnikovik/ndbx/autograder/internal/autograder/lab1"
	"github.com/sitnikovik/ndbx/autograder/internal/console"
	"github.com/sitnikovik/ndbx/autograder/internal/env"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			console.Panic(r)
		}
	}()
	appPort := env.MustGet("APP_PORT")
	err := lab0.Check("http://localhost:" + appPort.String() + "/health")
	if err != nil {
		console.Fatal("Healthcheck failed: %v", err)
	}
}
