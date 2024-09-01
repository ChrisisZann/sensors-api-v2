package main

import "net/http"

func (api *api) router() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", api.rootHandler)
	// CRUD
	// Create
	/* Create New Sensor */

	mux.HandleFunc("/sensor", api.sensor)

	// mux.HandleFunc("/create-sensor", api.createNewSensor)
	mux.HandleFunc("/sensor-data", api.sensorData) // should change to insert-sensor-value

	return mux

}
