package main

import (
  "net/http"
  "fmt"
  "log"
  "encoding/json"
	"io/ioutil"
  "time"
)

type obj struct {                //obj struct
  Timestamp    time.Time `json:"timestamp"`
	Key          string `json:"key"`
	Value        string `json:"value"`
}

type allObjs []obj

var objs = allObjs{             //dummy database for storing obj
	{
    Timestamp:  time.Now(),
		Key:        "testing",
		Value:      "testing",
	},
}

func addOne(w http.ResponseWriter, r *http.Request) {  //adding one obj into the dummy database
	var newObj obj
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Pls enter the key and value.")
	}

	json.Unmarshal(reqBody, &newObj)
  newObj.Timestamp = time.Now()
	objs = append(objs, newObj)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newObj)
}

func listAll(w http.ResponseWriter, r *http.Request) {  //listing all of the objs
	json.NewEncoder(w).Encode(objs)
}

func homePage(w http.ResponseWriter, r *http.Request) { //homepage
	fmt.Fprintf(w, "Welcome to my homepage!")
}

func main() {       //main
  fmt.Println("httpwebserver.go is running...")
  http.HandleFunc("/", homePage)
  http.HandleFunc("/add", addOne)
  http.HandleFunc("/list", listAll)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
