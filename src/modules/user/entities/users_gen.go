// Code generated by "repogen"; DO NOT EDIT.
package entities

import (
	"github.com/guregu/null"
)

type Users struct {
	UserId    int32    `db:"user_id"`
	Photo     string   `db:"photo"`
	Username  string   `db:"username"`
	Email     string   `db:"email"`
	Password  string   `db:"password"`
	Name      string   `db:"name"`
	CreatedAt int64    `db:"created_at"`
	UpdatedAt null.Int `db:"updated_at"`
}

type UsersList []*Users
