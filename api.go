package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type API struct {
	storage   *TaskContainer
	workQueue chan Task
}

func NewAPI(storage *TaskContainer, queue chan Task) *API {
	return &API{
		storage:   storage,
		workQueue: queue,
	}
}

func (api *API) AddTask(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var rawReq = TaskRequest{}
	err := decoder.Decode(&rawReq)
	defer r.Body.Close()
	if err != nil {
		fmtErr := fmt.Errorf("decode body request error %s", err)
		GenerateResponse(w, nil, false, fmtErr)
		log.Println(fmtErr)
		return
	}

	err = rawReq.Validation()
	if err != nil {
		fmtErr := fmt.Errorf("request error %s", err)
		GenerateResponse(w, nil, false, fmtErr)
		log.Println(fmtErr)
		return
	}

	newTask := Task{
		Id:       GenerateID(),
		Status:   StatusInQueue,
		WorkUrls: rawReq.Urls,
		Result:   make([]Product, 0),
	}

	api.storage.Set(newTask.Id, newTask)

	go func() {
		api.workQueue <- newTask
	}()

	GenerateResponse(w, AddNewTaskResponse{Id: newTask.Id}, true, nil)
	return

}

func (api *API) Task(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	taskID, ok := vars["id"]
	if !ok {
		fmtErr := fmt.Errorf("not found id")
		GenerateResponse(w, nil, false, fmtErr)
		log.Println(fmtErr)
	}
	task, ok := api.storage.Get(taskID)
	if !ok {
		fmtErr := fmt.Errorf("task with id: %s doesn`t exist", taskID)
		GenerateResponse(w, nil, false, fmtErr)
		log.Println(fmtErr)
	}

	GenerateResponse(w, task, true, nil)
}

func StartServer(config *Config, api *API) {
	router := mux.NewRouter()
	router.HandleFunc("/task/{id}/", api.Task).Methods("GET")
	router.HandleFunc("/task/", api.AddTask).Methods("POST", "PUT")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", config.Host, config.Port), router)
	log.Printf("FATAL: %s", err)
}
