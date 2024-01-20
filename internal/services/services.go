package services

import (
	"github.com/Rizabekus/effective-mobile-rest-api/internal/models"
	"github.com/Rizabekus/effective-mobile-rest-api/internal/storage"
)

type Services struct {
	PersonService models.PersonService
}

func ServiceInstance(storage *storage.PersonStorage) *Services {
	return &Services{
		PersonService: CreatePersonService(storage),
	}
}
