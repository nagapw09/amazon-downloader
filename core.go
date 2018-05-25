package main

import (
	"github.com/cornelk/hashmap"
)

var (
	StatusInQueue = "waiting"
	StatusInWork  = "working"
	StatusDone    = "done"
	StatusError   = "error"
)

type Product struct {
	Url    string `json:"url"`
	Price  string `json:"price"`
	Image  string `json:"image"`
	Title  string `json:"title"`
	IsSale bool   `json:"is_sale"`
}

type Task struct {
	Id       string    `json:"id"`
	Status   string    `json:"status"`
	WorkUrls []string  `json:"urls"`
	Result   []Product `json:"result"`
}

// in-memory task storage
type TaskContainer struct {
	storage *hashmap.HashMap
}

func NewTaskContainer() *TaskContainer {
	return &TaskContainer{
		storage: &hashmap.HashMap{},
	}
}

func (container *TaskContainer) Get(key string) (Task, bool) {
	value, ok := container.storage.Get(key)
	if !ok {
		return Task{}, ok
	}
	return (value).(Task), ok
}

func (container *TaskContainer) Set(key string, task Task) {
	container.storage.Set(key, task)
}

