package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/tsawler/toolbox"
)

func (a *api) rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello There, this is the root response from sensor-api-v2!")
}

func (a *api) sensor(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		var t toolbox.Tools
		sID := r.URL.Query().Get("sensor_id")
		if sID == "all" {
			log.Println("getting all sensors")
			sensors, err := a.config.Models.Sensors.SelectAllSensors()
			if err != nil {
				_ = t.ErrorJSON(w, err, http.StatusBadRequest)
			}
			if len(sensors) == 0 {
				fmt.Fprintln(w, http.StatusNoContent, "No Content")
				return
			}
			_ = t.WriteJSON(w, http.StatusOK, sensors)
			return
		}
		i_sID, err := strconv.Atoi(r.FormValue("sensor_id"))
		if err != nil {
			log.Panic("error converitng sensor id to integer", err)
			http.Error(w, "bad argument", http.StatusBadRequest)
			return
		}
		sensor, err := a.config.Models.Sensors.SelectSingleSensor(i_sID)
		if err != nil {
			_ = t.ErrorJSON(w, err, http.StatusBadRequest) // http.NotFound(w, r)
		}
		if sensor.Sensor_name == "" {
			fmt.Fprintln(w, http.StatusNoContent, "No Content")
			return
		}
		_ = t.WriteJSON(w, http.StatusOK, sensor)
	// case "HEAD":
	case "POST":
		//192.168.1.9:8888/create-sensor?sensor_id=555&sensor_type=string&sensor_name=test&description=Created%20test%20sensor%20from%20URI%20parameters
		err := r.ParseForm()
		if err != nil {
			log.Panic("failed to parse form data", err)
			http.Error(w, "failed to parse form data", http.StatusBadRequest)
			return
		}
		sType := r.FormValue("sensor_type")
		sName := r.FormValue("sensor_name")
		sDescription := r.FormValue("description")
		i_sID, err := strconv.Atoi(r.FormValue("sensor_id"))
		if err != nil {
			log.Println("ERROR:", err)
			fmt.Fprintln(w, http.StatusBadRequest, "bad argument")
			// http.Error(w, "bad argument", http.StatusBadRequest)
			return
		}
		err = a.config.Models.Sensors.CreateNewSensor(i_sID, sType, sName, sDescription)
		if err != nil {
			log.Println("ERROR:", err)
			http.Error(w, "Conflict", http.StatusConflict)
			return
		}
		fmt.Fprintln(w, http.StatusCreated, "Created")
	case "PUT":
		var t toolbox.Tools
		sID := r.URL.Query().Get("sensor_id")
		sType := r.URL.Query().Get("sensor_type")
		sName := r.URL.Query().Get("sensor_name")
		sDescription := r.URL.Query().Get("description")
		i_sID, err := strconv.Atoi(sID)
		if err != nil {
			log.Panicln("error converitng sensor id to integer", err)
			http.Error(w, "Bad Argument", http.StatusBadRequest)
		}
		rowsAffected, err := a.config.Models.Sensors.UpdateSensor(i_sID, sType, sName, sDescription)
		if err != nil {
			log.Panicln("error updating sensor info")
		}
		if rowsAffected == 0 {
			fmt.Fprintln(w, http.StatusCreated, "Created")
			return
		}
		if rowsAffected > 1 {
			_ = t.ErrorJSON(w, err, http.StatusBadRequest) // http.NotFound(w, r)
		}
		fmt.Fprintln(w, http.StatusPartialContent, "Resource Updated")
	// case "PATCH": // RFC 5789
	// case "DELETE":
	// case "CONNECT":
	// case "OPTIONS":
	// case "TRACE":
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
}

func (a *api) sensorData(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		var t toolbox.Tools
		i_sID, err := strconv.Atoi(r.FormValue("sensor_id"))
		if err != nil {
			log.Panic("error converitng sensor id to integer", err)
			http.Error(w, "bad argument", http.StatusBadRequest)
			return
		}
		sensor, err := a.config.Models.Sensor.GetSensorLatestValue(i_sID)
		if err != nil {
			_ = t.ErrorJSON(w, err, http.StatusBadRequest) // http.NotFound(w, r)
		}
		log.Println("sensor value = ", sensor.Sensor_value)
		if sensor.Sensor_value == "" {
			fmt.Fprintln(w, http.StatusNoContent, "No Content")
			return
		}
		_ = t.WriteJSON(w, http.StatusOK, sensor)
	// case "HEAD":
	case "POST":
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		sID := r.FormValue("sensor_id")
		sValue := r.FormValue("sensor_value")
		i_sID, err := strconv.Atoi(sID)
		if err != nil {
			log.Panic("error converitng sensor id to integer", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		err = a.config.Models.Sensor.InsertSensorData(i_sID, sValue)
		if err != nil {
			fmt.Println("ERROR: ", err)
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}
		fmt.Fprintln(w, http.StatusCreated, "Created")
	// case "PUT":
	// case "PATCH": // RFC 5789
	// case "DELETE":
	// case "CONNECT":
	// case "OPTIONS":
	// case "TRACE":
	default:
		http.Error(w, "Bad Request", http.StatusBadRequest)
	}
}
