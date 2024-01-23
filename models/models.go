package models

import (
	"errors"
	"fmt"
	"strings"

	"github.com/upper/db/v4"
)

var (
	ErrNoMoreRows     = errors.New("no record found")
	ErrDuplicateEmail = errors.New("email already in database")
	ErrUserNotActive  = errors.New("user not active")
	ErrInvalidLogin   = errors.New("invalid login")
)

type Models struct {
	Users UsersModel
	Post  PostsModel
}

func New(db db.Session) Models {
	return Models{
		Users: UsersModel{db: db},
		Post:  PostsModel{db: db},
	}
}

func convertUpperIDtoInt(id db.ID) int {
	idType := fmt.Sprintf("%T", id)
	if idType == "int64" {
		return int(id.(int64))
	}

	return id.(int)
}

func errHasDuplicate(err error, key string) bool {
	str := fmt.Sprintf(`ERROR: duplicate key value violates unique constraint "%s"`, key)
	return strings.Contains(err.Error(), str)
}
