package main

import (
	"encoding/json"
	"fmt"
	"github.com/rs/xid"
	"net/http"
)

func GenerateID() string {
	id := xid.New()
	return id.String()
}

type TaskRequest struct {
	Urls []string `json:"urls"`
}

func (tr *TaskRequest) Validation() error {
	if len(tr.Urls) <= 0 {
		return fmt.Errorf("url list is empty")
	}

	// TODO: Add urls validation

	return nil
}

type AddNewTaskResponse struct {
	Id string `json:"id"`
}

type Response struct {
	Value  interface{} `json:"value"`
	Error  string      `json:"error"`
	Status bool        `json:"status"`
}

func GenerateResponse(writer http.ResponseWriter, value interface{}, status bool, err error) {
	var errorStatus string
	if err == nil {
		errorStatus = "None"
	} else {
		writer.WriteHeader(http.StatusBadRequest)
		errorStatus = fmt.Sprintf("%s", err.Error())
	}
	r := Response{
		Value:  value,
		Status: status,
		Error:  errorStatus,
	}
	jsonStructure, err := json.Marshal(&r)

	writer.Write(jsonStructure)
	return
}
