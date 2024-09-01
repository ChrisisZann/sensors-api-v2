package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ChrisisZann/sensors-api-v2/config"
)

const port = ":8888"

type api struct {
	config *config.Application
}

func main() {

	db_conf := dbConfig{
		db_user:     "postgres",
		db_password: "4tE_pale",
		db_host:     "192.168.1.5",
		db_name:     "chrisis_home",
	}

	// config.New will create new repository by
	// chr_api.config.Models is a repository of models

	chr_api := api{
		config: config.New(connectToDB(db_conf)),
	}
	// Testing
	// chr_api.config.Models.Sensors.CreateNewSensor(775, "string", "repo test", "testing go repo")
	// chr_api.config.Models.Sensor.InsertSensorData(775, "Hello there again")
	// chr_api.config.Models.Sensors.SelectAllSensors()
	// chr_api.config.Models.Sensor.GetSensorLatestValue(446)

	srv := &http.Server{
		Addr:              port,
		Handler:           chr_api.router(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       30 * time.Second,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}

	log.Println("Starting web application on port", port)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}
