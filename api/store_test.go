package main

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/go-memdb"
)

func TestNewDB(t *testing.T) {
	_, err := memdb.NewMemDB(newSchema())
	if err != nil {
		t.Error("It was not possible to create a new in memory db\nerr: ", err)
	}
}

func TestPutTodo(t *testing.T) {
	db, err := memdb.NewMemDB(newSchema())
	if err != nil {
		t.Error("It was not possible to create a new in memory db\nerr: ", err)
	}

	store := NewTodoStore(db)

	id, _ := uuid.NewRandom()

	todo := Todo{
		ID:       id.String(),
		CreateAt: time.Now().Unix(),
		Content:  "test",
	}

	err = store.Insert(todo)
	if err != nil {
		t.Error("It was not possible to create a new todo\nerr: ", err)
	}
}

func TestGetTodo(t *testing.T) {
	db, err := memdb.NewMemDB(newSchema())
	if err != nil {
		t.Error("It was not possible to create a new in memory db\nerr: ", err)
	}

	store := NewTodoStore(db)

	id, _ := uuid.NewRandom()

	todo := Todo{
		ID:       id.String(),
		CreateAt: time.Now().Unix(),
		Content:  "test",
	}

	err = store.Insert(todo)
	if err != nil {
		t.Error("It was not possible to create a new todo\nerr: ", err)
		return
	}

	_, err = store.Get(uuid.Nil.String())
	if err == nil {
		t.Error("Having success where it should be failing")
		return
	}

	_, err = store.Get(id.String())
	if err != nil {
		t.Error("It was not possible to a find an existing todo\nerr: ", err)
		return
	}
}

func TestDeleteTodo(t *testing.T) {
	db, err := memdb.NewMemDB(newSchema())
	if err != nil {
		t.Error("It was not possible to create a new in memory db\nerr: ", err)
	}

	store := NewTodoStore(db)

	id, _ := uuid.NewRandom()

	todo := Todo{
		ID:       id.String(),
		CreateAt: time.Now().Unix(),
		Content:  "test",
	}

	err = store.Insert(todo)
	if err != nil {
		t.Error("It was not possible to create a new todo\nerr: ", err)
		return
	}

	err = store.Delete(uuid.Nil.String())
	if err == nil {
		t.Error("Having success where it should be failing")
		return
	}

	err = store.Delete(id.String())
	if err != nil {
		t.Error("It was not possible to delete an existing todo\nerr: ", err)
		return
	}

	_, err = store.Get(id.String())
	if err == nil {
		t.Error("Being able to retrieve a deleted todo")
		return
	}
}

func TestListTodo(t *testing.T) {
	db, err := memdb.NewMemDB(newSchema())
	if err != nil {
		t.Error("It was not possible to create a new in memory db\nerr: ", err)
	}

	store := NewTodoStore(db)

	todos := []Todo{
		{ID: uuid.NewString(), Content: "test 1", CreateAt: time.Now().Unix()},
		{ID: uuid.NewString(), Content: "test 2", CreateAt: time.Now().Unix()},
		{ID: uuid.NewString(), Content: "test 3", CreateAt: time.Now().Unix()},
		{ID: uuid.NewString(), Content: "test 4", CreateAt: time.Now().Unix()},
	}

	for _, v := range todos {
		err = store.Insert(v)
		if err != nil {
			t.Error("It was not possible to create a new todo\nerr: ", err)
			return
		}
	}

	sodot, err := store.List()
	if err != nil {
		t.Error("It was not possible to retrieve all the todos\nerr: ", err)
		return
	}

	if len(todos) != len(sodot) {
		t.Error("The number of todos retrieved does not match the number of todos inserted")
		return
	}
}
