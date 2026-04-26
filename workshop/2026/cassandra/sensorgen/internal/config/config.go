package config

import (
	"flag"
	"strings"
)

type Config struct {
	Hosts       []string
	Keyspace    string
	Consistency string
	Port        int
	Amount      int
}

func Load() Config {
	n := flag.Int(
		"n",
		500,
		"Amount of records to create",
	)
	p := flag.Int(
		"port",
		9042,
		"Cassandra cluser main port",
	)
	hosts := flag.String(
		"hosts",
		"",
		"Comma-separated list of hosts",
	)
	ks := flag.String(
		"keyspace",
		"",
		"Keyspace to connect",
	)
	c9y := flag.String(
		"consistency",
		"QUORUM",
		"Consistency level (ONE, QUOURUM, ALL)",
	)
	flag.Parse()
	return Config{
		Amount:      *n,
		Keyspace:    *ks,
		Consistency: *c9y,
		Port:        *p,
		Hosts:       strings.Split(*hosts, ","),
	}
}
