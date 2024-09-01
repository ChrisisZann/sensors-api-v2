package repository

import (
	"log"
	"time"
)

type Sensor struct {
	Sensor_id         string    `json:""`
	Sensor_value      string    `json:"Sensor_value"`
	Sys_creation_date time.Time `json:"sys_creation_date"`
}

func (s *Sensor) InsertSensorData(sId int, sValue string) error {
	return repo.InsertSensorData(sId, sValue)
}

func (thisRepo *psqlRepo) InsertSensorData(sId int, sValue string) error {

	query := `INSERT INTO home.sensor(sensor_id, sensor_value) VALUES ($1,$2);`

	// Execute the SQL query
	_, err := thisRepo.DB.Exec(query, sId, sValue)
	if err != nil {
		return err
	}
	log.Println("insertSensorData success")
	return nil
}

func (s *Sensor) GetSensorLatestValue(sId int) (Sensor, error) {
	return repo.GetSensorLatestValue(sId)
}

func (thisRepo *psqlRepo) GetSensorLatestValue(sId int) (Sensor, error) {

	query := `SELECT sensor_id, sensor_value, sys_creation_date 
	FROM home.sensor 
	WHERE sensor_id=$1
		AND sys_creation_date=(select max(sys_creation_date)
								FROM home.sensor 
								WHERE sensor_id=$1
		)
	;`

	rows, err := thisRepo.DB.Query(query, sId)
	if err != nil {
		log.Println("Error in: SelectAllSensors()", err)
	}
	// fmt.Println("rows: ", rows)
	defer rows.Close()

	// var sensor Sensor

	var sensorData Sensor
	if rows.Next() {
		err = rows.Scan(
			&sensorData.Sensor_id,
			&sensorData.Sensor_value,
			&sensorData.Sys_creation_date,
		)
		if err != nil {
			return sensorData, err
		}
	}

	return sensorData, nil
}
