package config

import (
	"database/sql"
	"sync"

	"github.com/ChrisisZann/sensors-api-v2/repository"
)

type Application struct {
	Models *repository.Models
}

var instance *Application
var once sync.Once
var db *sql.DB

func New(pool *sql.DB) *Application {
	db = pool
	return GetInstance()
}

func GetInstance() *Application {
	once.Do(func() {
		instance = &Application{
			Models: repository.New(db),
		}
	})
	return instance
}
