package main

import "log"

func main() {
	appConfig, err := NewConfig()
	if err != nil {
		log.Printf("Create config error: %s", err)
		return
	}
	taskStorage := NewTaskContainer()
	workQueue := make(chan Task, appConfig.MaxQueueSize)
	api := NewAPI(taskStorage, workQueue)

	for i := 0; i < appConfig.WorkPoolSize; i++ {
		fetcher := NewFetcher(taskStorage, workQueue)
		go fetcher.Start()
	}

	StartServer(appConfig, api)
}
// Lesha was here