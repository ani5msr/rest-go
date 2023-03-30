package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

type CmdRequest struct {
	Command string `json:"command"`
}
type ErrBody struct {
	Error string `json:"error"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func cmdprocess(words string[]){
    if(words[0] == "SET"){
		key := words[1]
		value := words[2]
		if(words[3] == "EX"){
			expiry := words[4]
			if(words[4] == "NX" || words[4] == "XX"){
            
			}
		}

	}else if(words[0] == "GET"){
		key := words[1]
		value := words[2]
		
	}else if(words[0] == "QPUSH"){

	}else if(words[0] == "QPOP"){

	}else if(words[0] == "BQPOP"){

	}
}
func createNewPost(w http.ResponseWriter, r *http.Request) {
	// get the body of our POST request
	// return the string response containing the request body
	reqBody, _ := ioutil.ReadAll(r.Body)
	var req CmdRequest
	json.Unmarshal(reqBody, &req)
	words := strings.Fields(req.Command)
	if words[0] == "SET" || words[0] == "GET" {
		fmt.Fprintf(w, "%+v", string(reqBody))
		cmdprocess(words)
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
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
