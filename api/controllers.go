package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type TodoController interface {
	List() http.Handler
	Get() http.Handler
	Insert() http.Handler
	Delete() http.Handler
}

func NewTodoController(ts TodoStore) TodoController {
	return &todoController{
		store: ts,
	}
}

type todoController struct {
	store TodoStore
}

func (tc *todoController) List() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if h := recover(); h != nil {
				log.Println(r)
				writeErrorMessage(http.StatusInternalServerError, fmt.Errorf("%v", h), w)
				return
			}
		}()

		res, err := tc.store.List()
		if err != nil {
			writeErrorMessage(http.StatusNotFound, err, w)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(res)
	})
}

func (tc *todoController) Get() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if h := recover(); h != nil {
				log.Println(r)
				writeErrorMessage(http.StatusInternalServerError, fmt.Errorf("%v", h), w)
				return
			}
		}()

		vls := mux.Vars(r)
		id := vls["id"]

		res, err := tc.store.Get(id)
		if err != nil {
			writeErrorMessage(http.StatusNotFound, err, w)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(res)
	})
}

func (tc *todoController) Insert() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if h := recover(); h != nil {
				log.Println(r)
				writeErrorMessage(http.StatusInternalServerError, fmt.Errorf("%v", h), w)
				return
			}
		}()

		var payload struct {
			Content string `json:"content"`
		}

		err := json.NewDecoder(r.Body).Decode(&payload)
		if err != nil {
			writeErrorMessage(http.StatusUnprocessableEntity, err, w)
			return
		}

		var data Todo

		data.Content = payload.Content
		data.CreateAt = time.Now().Unix()
		data.ID = uuid.NewString()

		err = tc.store.Insert(data)
		if err != nil {
			writeErrorMessage(http.StatusInternalServerError, err, w)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})
	})
}

func (tc *todoController) Delete() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if h := recover(); h != nil {
				log.Println(r)
				writeErrorMessage(http.StatusInternalServerError, fmt.Errorf("%v", h), w)
				return
			}
		}()

		vls := mux.Vars(r)
		id := vls["id"]

		err := tc.store.Delete(id)
		if err != nil {
			writeErrorMessage(http.StatusNotFound, err, w)
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		_ = json.NewEncoder(w).Encode(map[string]interface{}{"ok": true})
	})
}
