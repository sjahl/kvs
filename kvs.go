package main

import (
	"encoding/json"
	"net/http"
	"fmt"
	"log"
)

var kv_store = make(map[string]interface{})

type KeyValue struct {
	Key string `json:"key"`
	Value interface{} `json:"value"`
}

type ActionResp struct {
	Ok bool `json:"ok"`
	Message string `json:"message"`
}


func setValueHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decode := json.NewDecoder(r.Body)
	req := &KeyValue{}

	if err := decode.Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	kv_store[req.Key] = req.Value
	w.Header().Set("Content-Type", "application/json")
	response := &ActionResp{Ok: true, Message: "success"}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))

}

//func getValueHandler
func getValueHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decode := json.NewDecoder(r.Body)
	req := &KeyValue{}

	if err := decode.Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value := kv_store[req.Key]

	if value == nil {
		response := &ActionResp{Ok: false, Message: "unable to find key"}
		b, _ := json.Marshal(response)
		fmt.Println(string(b))
	}

	response := value
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))


}

//func getValuesHandler

//func dropValueHandler

func main() {

	http.HandleFunc("/set", setValueHandler)
	http.HandleFunc("/get", getValueHandler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
