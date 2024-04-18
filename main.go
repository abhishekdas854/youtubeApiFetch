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

	cronObj := NewCronObject(store)

	go cronObj.FetchYoutubeData()

	// c := cron.New()

	// _, err := c.AddFunc("*/10 * * * *", func() {
	// 	// This function will be called every 10 minutes
	// 	fmt.Println("Running cron job every 10 minutes")
	// })

	// fmt.Printf("%+v\n", store)
	server := NewApiServer(":3000", store)
	server.Run()
}
