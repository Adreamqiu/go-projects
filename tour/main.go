package main

import (
	"log"
	"github.com/go-projects/tour/cmd"
)

func main () {
	err := cmd.Execute()
	if err != nil {
		log.Fatalf("cmd.Execute err:%v", err)
	}
}
