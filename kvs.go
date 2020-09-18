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
	fmt.Fprintf(w, string(b))

}

func getValueHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decode := json.NewDecoder(r.Body)
	req := &KeyValue{}
	w.Header().Set("Content-Type", "application/json")

	if err := decode.Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	value := kv_store[req.Key]

	if value == nil {
		response := &ActionResp{Ok: false, Message: "unable to find key"}
		b, _ := json.Marshal(response)
		fmt.Fprintf(w, string(b))
		return
	}

	response := value
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(b))


}

func getValuesHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	b, err := json.Marshal(kv_store)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(b))
}

func dropValueHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	decode := json.NewDecoder(r.Body)
	req := &KeyValue{}

	if err := decode.Decode(req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	delete(kv_store, req.Key)
	w.Header().Set("Content-Type", "application/json")
	response := &ActionResp{Ok: true, Message: "success"}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Fprintf(w, string(b))
}

func main() {

	http.HandleFunc("/set", setValueHandler)
	http.HandleFunc("/get", getValueHandler)
	http.HandleFunc("/all", getValuesHandler)
	http.HandleFunc("/drop", dropValueHandler)

	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal(err)
	}
}
