package repository

import (
	"database/sql"
)

var repo Repository

type Models struct {
	Sensors Sensors
	Sensor  Sensor
}

func New(conn *sql.DB) *Models {
	if conn != nil {
		repo = newPsqlRepo(conn)
	}
	// else {
	// 	repo = newTestRepository(conn)
	// }

	return &Models{
		Sensors: Sensors{},
		Sensor:  Sensor{},
	}
}
