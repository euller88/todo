package main

import (
	"errors"

	"github.com/hashicorp/go-memdb"
)

// TodoStore is a type responsible for the operations over the Todo struct
type TodoStore interface {
	List() ([]Todo, error)
	Get(key string) (Todo, error)
	Insert(data Todo) error
	Delete(key string) error
}

// NewTodoStore returns a new Storage for the Todo struct
func NewTodoStore(db *memdb.MemDB) TodoStore {
	return &todoStore{
		db: db,
	}
}

type todoStore struct {
	db *memdb.MemDB
}

func (ts *todoStore) List() ([]Todo, error) {
	result := make([]Todo, 0)

	txn := ts.db.Txn(false)
	defer txn.Abort()

	it, err := txn.GetReverse("todo", "created_at")
	if err != nil {
		return result, err
	}

	for obj := it.Next(); obj != nil; obj = it.Next() {
		t, ok := obj.(Todo)
		if !ok {
			continue
		}
		result = append(result, t)
	}

	return result, nil
}

func (ts *todoStore) Get(key string) (Todo, error) {
	result := Todo{}

	txn := ts.db.Txn(false)
	defer txn.Abort()

	res, err := txn.First("todo", "id", key)
	if err != nil {
		return result, err
	}

	result, ok := res.(Todo)

	if !ok {
		return result, errors.New("empty result")
	}

	txn.Commit()

	return result, nil
}

func (ts *todoStore) Insert(data Todo) error {
	txn := ts.db.Txn(true)
	defer txn.Abort()

	err := txn.Insert("todo", data)
	if err != nil {
		return err
	}

	txn.Commit()

	return nil
}

func (ts *todoStore) Delete(key string) error {
	item, err := ts.Get(key)
	if err != nil {
		return err
	}

	txn := ts.db.Txn(true)
	defer txn.Abort()

	err = txn.Delete("todo", item)
	txn.Commit()

	return err
}
