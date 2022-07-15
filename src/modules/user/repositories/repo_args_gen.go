// Code generated by "repogen"; DO NOT EDIT.
package repositories

import (
	"database/sql"
)

type Filter interface {
	Query() string
	Values() []interface{}
}

type Pagination interface {
	GetPage() int
	GetSize() int
}

type Order interface {
	Value() string
	Direction() string
}

type PaginationData struct {
	Page int
	Size int
}

func (p PaginationData) GetPage() int {
	return p.Page
}

func (p PaginationData) GetSize() int {
	return p.Size
}

type InsertResult struct {
	sql.Result
}

func excludeFields(excludedFields, allFields []string) []string {
	var selectedFields []string
	for _, field := range allFields {
		var isExField bool
		for _, exField := range excludedFields {
			if field == exField {
				isExField = true
				break
			}
		}

		if isExField {
			continue
		}
		selectedFields = append(selectedFields, field)
	}
	return selectedFields
}