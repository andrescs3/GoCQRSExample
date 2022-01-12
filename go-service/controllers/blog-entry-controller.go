package controllers

import (
	"context"
	"encoding/json"
	blogEntryCommand "go-service/commands"
	"go-service/entity"
	"go-service/kafka"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllPerson get all person data
func GetAllBlogEntries(w http.ResponseWriter, r *http.Request) {
	blogEntries, err := blogEntryCommand.GetAll()
	w.Header().Set("Content-Type", "application/json")
	var status int
	if err == nil {
		status = http.StatusOK
	} else {
		status = http.StatusInternalServerError
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(blogEntries)
}

func GetBlogEntryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	i, convErr := strconv.Atoi(key)
	if convErr == nil {
		blogEntry, err := blogEntryCommand.GetByID(i)
		w.Header().Set("Content-Type", "application/json")
		var status int
		if err == nil {
			status = http.StatusOK
		} else {
			status = http.StatusInternalServerError
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(blogEntry)
	}
}

func CreateBlogEntry(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var blogEntry entity.BlogEntry
	json.Unmarshal(requestBody, &blogEntry)
	i, err := blogEntryCommand.CreateBlogEntry(blogEntry)
	w.Header().Set("Content-Type", "application/json")
	var status int
	if err == nil {
		status = http.StatusOK
		ctx := context.Background()
		go kafka.Publish(ctx, string(requestBody), i)
	} else {
		status = http.StatusInternalServerError
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(i)

}

func UpdateBlogEntry(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var blogEntry entity.BlogEntry
	json.Unmarshal(requestBody, &blogEntry)
	i, err := blogEntryCommand.UpdateBlogEntry(blogEntry)
	w.Header().Set("Content-Type", "application/json")
	var status int
	if err == nil {
		status = http.StatusOK
	} else {
		status = http.StatusInternalServerError
	}
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(i)

}

func DeleteBlogEntryByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	i, convErr := strconv.Atoi(key)
	if convErr == nil {
		rows, err := blogEntryCommand.DeleteBlogEntry(i)
		w.Header().Set("Content-Type", "application/json")
		var status int
		if err == nil {
			status = http.StatusOK
		} else {
			status = http.StatusInternalServerError
		}
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(rows)
	}
}
