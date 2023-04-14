package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	"github.com/ani5msr/rest-go/redis"
	"github.com/gammazero/deque"
	"github.com/gorilla/mux"
)

type CmdRequest struct {
	Command string `json:"command"`
}
type ErrBody struct {
	Error string `json:"error"`
}
type RespBody struct {
	Value string `json:"value"`
}
type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
type QueueInstance struct {
	Q    deque.Deque[string]
	Name string
}

var queuemap = make(map[string]*QueueInstance)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func createNewPost(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var req CmdRequest
	json.Unmarshal(reqBody, &req)
	words := strings.Fields(req.Command)
	w.Header().Set("Content-Type", "application/json")
	if words[0] == "SET" {
		key := words[1]
		value := words[2]
		if words[3] == "EX" {
			expiry := words[4]
			exp, _ := strconv.ParseInt(expiry, 10, 64)
			requestinstance := redis.SetInterface{
				Key:    key,
				Value:  value,
				Expiry: exp,
				Lock:   sync.Mutex{},
			}
			stat := redis.SetRedis(requestinstance)
			res := Response{
				Status:  stat,
				Message: "User inserted to redis successFully!!!",
			}
			json.NewEncoder(w).Encode(res)
		}

	} else {
		fmt.Fprintf(w, "invalid command")
	}
}
func getPost(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var req CmdRequest
	json.Unmarshal(reqBody, &req)
	words := strings.Fields(req.Command)
	w.Header().Set("Content-Type", "application/json")
	if words[0] == "GET" {
		fmt.Fprintf(w, "%+v", string(reqBody))
		key := words[1]
		requestinstance := redis.GetInterface{
			Key:  key,
			Lock: sync.Mutex{},
		}
		res := redis.GetRedis(requestinstance)
		respval := RespBody{Value: res}
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(respval)
	} else {
		fmt.Fprintf(w, "invalid command")
	}
}

func queuepush(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var req CmdRequest
	json.Unmarshal(reqBody, &req)
	words := strings.Fields(req.Command)
	w.Header().Set("Content-Type", "application/json")
	if words[0] == "QPUSH" {
		fmt.Fprintf(w, "%+v", string(reqBody))
		w.WriteHeader(http.StatusCreated)
	} else {
		fmt.Fprintf(w, "invalid command")
	}
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/post", createNewPost)
	myRouter.HandleFunc("/get", getPost)
	myRouter.HandleFunc("/pushqueue", queuepush)
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
