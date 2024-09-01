package repository

import (
	"errors"
	"log"
	"time"
)

type Sensors struct {
	Sensor_id         int       `json:"sensor_id"`
	Sensor_type       string    `json:"sensor_type"`
	Sys_creation_date time.Time `json:"sys_creation_date"`
	Sys_update_date   time.Time `json:"sys_update_date"`
	Sensor_name       string    `json:"sensor_name"`
	Description       string    `json:"description"`
}

func (s *Sensors) CreateNewSensor(sId int, sType, sName, sDescription string) error {
	return repo.CreateNewSensor(sId, sType, sName, sDescription)
}

func (thisRepo *psqlRepo) CreateNewSensor(sId int, sType, sName, sDescription string) error {

	query := `INSERT INTO home.sensors(sensor_id, sensor_type, sensor_name, description) VALUES ($1,$2,$3,$4);`

	res, err := thisRepo.DB.Exec(query, sId, sType, sName, sDescription)
	if err != nil {
		log.Println("CreateNewSensor() :  PSQL : ", err)
		return errors.New("psql error")
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Println("CreateNewSensor() : Failed to read result rows")
	}
	if rowsAffected == 0 {
		return errors.New("failed to create new sensor")
	}
	return nil
}

func (s *Sensors) SelectSingleSensor(sId int) (Sensors, error) {
	return repo.SelectSingleSensor(sId)
}

func (thisRepo *psqlRepo) SelectSingleSensor(sId int) (Sensors, error) {
	query := `SELECT sensor_id, sensor_type, sensor_name, COALESCE(description,'')
	FROM home.sensors WHERE sensor_id=$1;`

	rows, err := thisRepo.DB.Query(query, sId)
	if err != nil {
		log.Println("Error in: SelectAllSensors()", err)
	}
	// fmt.Println("rows: ", rows)
	defer rows.Close()

	rows.Next()
	var sensor Sensors
	err = rows.Scan(
		&sensor.Sensor_id,
		&sensor.Sensor_type,
		&sensor.Sensor_name,
		&sensor.Description,
	)
	if err != nil {
		return sensor, err
	}
	return sensor, nil
}

func (s *Sensors) SelectAllSensors() ([]*Sensors, error) {
	return repo.SelectAllSensors()
}

func (thisRepo *psqlRepo) SelectAllSensors() ([]*Sensors, error) {
	query := `SELECT sensor_id, sensor_type, sensor_name, COALESCE(description,'')
	FROM home.sensors;`

	rows, err := thisRepo.DB.Query(query)
	if err != nil {
		log.Println("Error in: SelectAllSensors()", err)
	}
	// fmt.Println("rows: ", rows)
	defer rows.Close()

	var sensors []*Sensors

	for rows.Next() {
		var sensor Sensors
		err := rows.Scan(
			&sensor.Sensor_id,
			&sensor.Sensor_type,
			&sensor.Sensor_name,
			&sensor.Description,
		)
		if err != nil {
			return nil, err
		}
		sensors = append(sensors, &sensor)
		// fmt.Println("sensor:")
		// fmt.Println(sensor, sensor.Sensor_id, sensor.Sensor_name, sensor.Sensor_type, sensor.Description)
	}

	// fmt.Println("sensors:")
	// fmt.Println(sensors)
	return sensors, nil
}

func (s *Sensors) UpdateSensor(sId int, sType, sName, sDescription string) (int64, error) {
	return repo.UpdateSensor(sId, sType, sName, sDescription)
}

func (thisRepo *psqlRepo) UpdateSensor(sId int, sType, sName, sDescription string) (int64, error) {

	query := `UPDATE home.sensors
	SET sensor_type=$2, sensor_name=$3, description=$4
		where sensor_id=$1;`

	res, err := thisRepo.DB.Exec(query, sId, sType, sName, sDescription)
	if err != nil {
		log.Panicln("UpdateSensor error: ", err)
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		log.Panicln("error in UpdateSensor:RowsAffected()")
	}
	if rowsAffected == 0 {
		log.Println("INFO: sensor not found, creating new one")
		err = thisRepo.CreateNewSensor(sId, sType, sName, "created by update request")
		if err != nil {
			log.Panicln("error creating new sensor definition")
		}
	}
	// fmt.Println("CreateNewSensor result: ", rowsAffected)
	return rowsAffected, nil

}
