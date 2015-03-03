package main

import (
	"honey"
	"log"
)

func main() {
	var (
		err    error
		conf   *honey.Config
		server *honey.Server
	)

	if conf, err = honey.NewConfig(); err != nil {
		log.Fatal(err)
	}

	server = &honey.Server{conf}
	server.IncreaseRlimit()
	
	if err = server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
