package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type Fetcher struct {
	workQueue chan Task
	storage   *TaskContainer
	client    *http.Client
}

func NewFetcher(storage *TaskContainer, queue chan Task) *Fetcher {
	return &Fetcher{
		workQueue: queue,
		storage:   storage,
		client:    &http.Client{},
	}
}

func (ft *Fetcher) Start() {
	for task := range ft.workQueue {
		ft.handle(&task)
	}
}

func (ft *Fetcher) handle(task *Task) {
	log.Printf("task id %s in work \n", task.Id)
	task.Status = StatusInWork
	ft.storage.Set(task.Id, *task)
	var isError = false

	for _, url := range task.WorkUrls {
		ok := ft.work(task, url)
		if !ok{
			break
			isError = true
		}
	}

	if !isError {
		task.Status = StatusDone
		ft.storage.Set(task.Id, *task)
		log.Printf("task id=%s is done [%v]", task.Id, task)
	} else {
		log.Printf("task id=%s done with status error", task.Id)
	}

}

func (ft *Fetcher) work(task *Task, url string) bool{
	// request
	htmlContent, err := ft.doRequest(url)
	if err != nil {
		ft.handleError(task, err)
		return false

	}

	// parsing
	parsedProduct, err := ft.doParse(htmlContent, url)
	if err != nil {
		ft.handleError(task, err)
		return false
	}

	ft.handleResult(task, parsedProduct)
	return true
}

func (ft *Fetcher) doParse(htmlContent string, url string) (*Product, error) {
	parser := AmazonParser{
		body: htmlContent,
		product: &Product{
			Url: url,
		},
	}
	err := parser.Parse()
	if err != nil {
		return nil, err
	}

	return parser.product, nil
}

func (ft *Fetcher) doRequest(url string) (string, error) {
	resp, err := ft.client.Get(url)
	if err != nil {
		return "", err
	}

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(contents), nil

}

func (ft *Fetcher) handleResult(task *Task, product *Product){
	log.Printf("save result for task id=%s product=%v\n", task.Id, *product)
	task.Result = append(task.Result, *product)
	ft.storage.Set(task.Id, *task)
}

func (ft *Fetcher) handleError(task *Task, err error){
	task.Status = StatusError
	ft.storage.Set(task.Id, *task)
	log.Printf("task id %s have error %s\n", task.Id, err)
}

