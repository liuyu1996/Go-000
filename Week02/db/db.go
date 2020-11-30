package db

import "github.com/pkg/errors"

func DB() (result string, err error) {
	err = errors.New("ErrNoRows")
	return "test result", errors.Wrap(err, "")
}
