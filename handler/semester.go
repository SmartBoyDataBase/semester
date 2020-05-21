package handler

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"sbdb-semester/model"
	"strconv"
)

func getHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	semesterId := r.URL.Query().Get("id")
	userId, _ := strconv.ParseUint(semesterId, 10, 64)
	semester, err := model.Get(userId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	resp, _ := json.Marshal(semester)
	_, _ = w.Write(resp)
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	body, _ := ioutil.ReadAll(r.Body)
	var toCreate model.Semester
	_ = json.Unmarshal(body, &toCreate)
	result, err := model.Create(toCreate)
	if err != nil {
		log.Println("Create semester failed")
		_, _ = w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	} else {
		log.Println("Semester ", result.Name, "created")
	}
	response, err := json.Marshal(result)
	_, _ = w.Write(response)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		getHandler(w, r)
	case "POST":
		postHandler(w, r)
	}
}

func AllHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	all, err := model.All()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	var body []byte
	if len(all) != 0 {
		body, _ = json.Marshal(all)
	} else {
		body = []byte("[]")
	}
	_, _ = w.Write(body)
}
