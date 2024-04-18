package main

import (
	"log"
)

func main() {

	store, err := NewPostgresStore()

	if err != nil {
		log.Fatal(err)
	}

	if err = store.Init(); err != nil {
		log.Fatal(err)
	}

	// cronObj := NewCronObject(store)

	// go cronObj.FetchYoutubeData()

	server := NewApiServer(":3000", store)
	server.Run()
}
