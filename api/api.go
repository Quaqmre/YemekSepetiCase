package api

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/Quaqmre/yemeksepetiCase/store"
)

var db, interChan = store.NewStore()

// requestLogger is middleware to log request path and method for every request.
func requestLogger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path, r.Method)
		next(w, r)
	}
}

func create(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var err error
		defer func() {
			if err != nil {
				log.Println("create finished with error")

				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		}()
		obJson, err := ioutil.ReadAll(r.Body)

		db.Put(string(obJson))
		resp := make(map[string]int32, 1)
		resp["key"] = *db.Ops
		jsonResp, err := json.Marshal(resp)
		w.Header().Set("Content-Type", "application/json")
		log.Println("create finished successfully")

		w.Write(jsonResp)
		w.WriteHeader(http.StatusCreated)

	}
}

func get(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		var err error
		defer func() {
			if err != nil {
				log.Println("get finished with error")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		}()

		p := strings.Split(r.URL.Path, "/")
		if len(p) == 3 {
			i, _ := strconv.ParseInt(p[2], 10, 32)
			value := db.Get(string(i))

			if value == "" {
				w.WriteHeader(http.StatusNotFound)
			}

			jsonResp, _ := json.Marshal(value)
			log.Println("get finished successfully")

			w.Write(jsonResp)
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusNotFound)

		}

	}
}

func flush(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		var err error
		defer func() {
			if err != nil {
				log.Println("flush finished with error")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(err.Error()))
			}
		}()

		db.Flush()
		db.Marshall()
		log.Println("flush finished successfully")

		w.WriteHeader(http.StatusOK)

	}
}

// New Api return bunch of handler for store objects
func NewApi() {
	log.Println("Api will increase soon..")
	http.HandleFunc("/create", requestLogger(create))
	http.HandleFunc("/get/", requestLogger(get))
	http.HandleFunc("/flush", requestLogger(flush))

	log.Println("Listening port 8080")
	http.ListenAndServe(":8080", nil)
}
