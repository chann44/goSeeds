package main

import (
	"log"
)

func main() {
	store , err := NewPostgresStore()

	if(err != nil) {
		log.Fatal(err)
	}
	
	if err := store.init(); err != nil {
		log.Fatal(err)
	}


	server := ApiServer{":3000", store}
	server.Run()
}