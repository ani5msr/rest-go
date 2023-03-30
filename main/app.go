package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/gorilla/mux"
)

type CmdRequest struct {
	Command string `json:"command"`
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
func processcmd(words string){
	slicedwords := strings.Fields(words)
	if slicedwords[0] == "SET" {
		
	}
	else if slicedwords[0] == "GET"{

	}
}
func createNewPost(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // return the string response containing the request body    
    reqBody, _ := ioutil.ReadAll(r.Body)
	var req CmdRequest
	json.Unmarshal(reqBody, &req)
	value := reflect.ValueOf(req).Interface().(string)
	processcmd(value)
    fmt.Fprintf(w, "%+v", string(reqBody))
}
func handleRequests() {
    // creates a new instance of a mux router
    myRouter := mux.NewRouter().StrictSlash(true)
    // replace http.HandleFunc with myRouter.HandleFunc
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/get", )
    // finally, instead of passing in nil, we want
    // to pass in our newly created router as the second
    // argument
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	handleRequests()
}
