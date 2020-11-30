package service

import (
	"Go-000/Week02/db"
)

func Service() (string, error) {
	result, err := db.DB()
	//return result, errors.WithMessage(err, "Service error")
	return result, err

}
