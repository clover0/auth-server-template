package core

import (
	"github.com/jmoiron/sqlx"
)

type (
	User struct {
		ID        uint32 `db:"id"`
		UpdatedAt string `db:"created_at"`
		CreatedAt string `db:"updated_at"`
		Email     string `db:"email"`
		Password  string `db:"password"`
	}

	UserStore interface {
		Count(column string, param interface{}) (uint, error)
		Find(uint32) (*User, error)
		FindByEmail(string) (*User, error)
		Create(user *User) (uint32, error)
	}

	// improve
	UserStoreFunc func(session *sqlx.Tx) UserStore

	UserService interface {
		CheckDuplicateEmail(user *User) (bool,error)
		Register(user *User) error
	}
)
