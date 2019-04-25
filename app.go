package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"reflect"
	"time"

	"github.com/gorilla/mux"
)

//Schema - API DataType (struct -> Config)
type Schema struct {
	Config Config `json:"Config"`
}

// Config - Name (String) and Data (struct -> Data)
type Config struct {
	Name string `json:"Name"`
	Data Data   `json:"Data"`
}

//Data - ("Key": "Value" pairs)
type Data struct {
	Key1 string `json:"Key1"`
	Key2 string `json:"Key2"`
	Key3 string `json:"Key3"`
}

var schema []Schema

//List Config
func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(schema)
}

//Create Config
func Create(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post Schema
	json.NewDecoder(r.Body).Decode(&post)
	schema = append(schema, post)
}

//Get Config Attr Name
func Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, elem := range schema {
		if elem.Config.Name == params["name"] {
			json.NewEncoder(w).Encode(elem)
			return
		}
	}
}

//Update Config w.r.t corresponding Name Attr
func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, elem := range schema {
		if elem.Config.Name == params["name"] {
			schema = append(schema[:i], schema[i+1:]...)
			var updateSchema Schema
			json.NewDecoder(r.Body).Decode(&updateSchema)
			updateSchema.Config.Name = params["name"]
			schema = append(schema, updateSchema)
			json.NewEncoder(w).Encode(updateSchema)
			return
		}
	}
}

//Delete Config w.r.t corresponding Name Attr
func Delete(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for i, elem := range schema {
		if elem.Config.Name == params["name"] {
			schema = append(schema[:i], schema[i+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(schema)
}

//Query Search/Find
func Query(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := r.URL.Query()
	var data string
	for k := range params {
		if k != "name" {
			data = params.Get(k)
		}
	}
	for _, elem := range schema {
		if elem.Config.Name == params.Get("name") {
			v := reflect.ValueOf(elem.Config.Data)
			num := v.NumField()
			for i := 0; i < num; i++ {
				value := v.Field(i)
				v := value.String()
				if v == data {
					json.NewEncoder(w).Encode(elem)
				}
			}
		}
	}
}

// Kubernetes Liveliness Probe Endpoint
func healthcheck(w http.ResponseWriter, r *http.Request) {
	duration := time.Since(time.Now())
	if duration.Seconds() > 10 {
		w.WriteHeader(500)
		w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
	} else {
		w.WriteHeader(200)
		w.Header().Set("HOW-ARE-YOU", "FINE_THANK_YOU")
	}
}

func main() {
	// simulate sample data
	var d1, d2 Data
	d1 = Data{Key1: "Value1", Key2: "Value2", Key3: "Value3"}
	d2 = Data{Key1: "Value1", Key2: "Value2", Key3: "Value3"}
	var c1, c2 Config
	c1 = Config{Name: "Name1", Data: d1}
	c2 = Config{Name: "Name2", Data: d2}
	schema = append(schema, Schema{c1})
	schema = append(schema, Schema{c2})

	router := mux.NewRouter() // intialize router
	// endpoints
	router.HandleFunc("/configs", List).Methods("GET")
	router.HandleFunc("/configs", Create).Methods("POST")
	router.HandleFunc("/configs/{name}", Get).Methods("GET")
	router.HandleFunc("/configs/{name}", Update).Methods("PUT")
	router.HandleFunc("/configs/{name}", Delete).Methods("DELETE")
	router.HandleFunc("/search", Query).Methods("GET")
	router.HandleFunc("/healthz", healthcheck).Methods("GET")
	val, ok := os.LookupEnv("SERVE_PORT")
	if !ok {
		panic("SERVE_PORT not set") // Exit if SERVE_PORT is not set
	} else {
		log.Fatal(http.ListenAndServe(":"+val, router))
	}
}
