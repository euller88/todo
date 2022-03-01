package main

import (
	"github.com/hashicorp/go-memdb"
)

// Todo is an struct that represents things that should be done
type Todo struct {
	ID       string `json:"id"`
	CreateAt int64  `json:"created_at"`
	Content  string `json:"content"`
}

func newSchema() *memdb.DBSchema {
	schema := &memdb.DBSchema{
		Tables: map[string]*memdb.TableSchema{
			"todo": {
				Name: "todo",
				Indexes: map[string]*memdb.IndexSchema{
					"id": {
						Name:    "id",
						Unique:  true,
						Indexer: &memdb.UUIDFieldIndex{Field: "ID"},
					},
					"created_at": {
						Name:    "created_at",
						Unique:  false,
						Indexer: &memdb.IntFieldIndex{Field: "CreateAt"},
					},
					"content": {
						Name:    "content",
						Unique:  false,
						Indexer: &memdb.StringFieldIndex{Field: "Content"},
					},
				},
			},
		},
	}

	return schema
}
