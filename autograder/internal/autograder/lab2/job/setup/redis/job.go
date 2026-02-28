package redis

import (
	"context"
)

const (
	// Name is the name of the Redis setup job.
	Name = "Redis Setup"
	// Description is a brief explanation of what the Redis setup job does.
	Description = "Sets up the Redis server for Lab 2 by flushing all existing data."
)

// redisClient defines the interface for interacting with Redis to perform the setup job.
type redisClient interface {
	// Ping checks the connection to the Redis server.
	Ping(ctx context.Context) error
	// FlushAll removes all keys from the Redis server.
	FlushAll(ctx context.Context) error
}

// Job represents the Redis setup job in the autograder process.
type Job struct {
	// cli is the Redis client used to interact with the Redis server.
	cli redisClient
}

// NewJob creates a new Job instance with the provided Redis client.
func NewJob(cli redisClient) *Job {
	return &Job{
		cli: cli,
	}
}

// Name returns the name of the job.
func (j *Job) Name() string {
	return Name
}

// Description returns a brief explanation of what the job does.
func (j *Job) Description() string {
	return Description
}
