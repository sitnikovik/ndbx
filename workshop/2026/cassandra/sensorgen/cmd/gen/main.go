package main

import (
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/sitnikovik/ndbx/sensorgen/internal/config"
	"github.com/sitnikovik/ndbx/sensorgen/internal/infra/client/cassandra"
	"github.com/sitnikovik/ndbx/sensorgen/service/sensor/generator"
)

func main() {
	cfg := config.Load()
	cli, err := cassandra.QuorumSession(
		cfg.Keyspace,
		cfg.Port,
		cfg.Consistency,
		cfg.Hosts...,
	)
	if err != nil {
		log.Fatalf("failed to connect Cassandra: %v", err)
		return
	}
	defer cli.Close()
	n := cfg.Amount
	startedAt := time.Now()
	jobs := make(chan struct{}, 100)
	var wg sync.WaitGroup
	for wid := 1; wid <= 5; wid++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			generator.Generate(
				id,
				cfg.Keyspace,
				cli,
				jobs,
			)
		}(wid)
	}
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Println("Shutdown signal received. Waiting for workers to finish...")
		close(jobs)
	}()
	if n == -1 {
		log.Println("Running in infinite loop. Press Ctrl+C to stop.")
		for {
			jobs <- struct{}{}
		}
	} else {
		log.Printf("Inserting %d records...", n)
		for range n {
			jobs <- struct{}{}
		}
		close(jobs)
	}
	wg.Wait()
	log.Printf(
		"Done in %s\n",
		time.Since(startedAt),
	)
}
