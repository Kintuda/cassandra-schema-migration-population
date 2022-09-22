package main

import (
	"log"

	"github.com/Kintuda/cassandra-schema-migration/cmd"
)

func main() {
	root := cmd.NewRootCmd()

	if err := root.Execute(); err != nil {
		log.Fatal("command resulted in error")
	}
}
