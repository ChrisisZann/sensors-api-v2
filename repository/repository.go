package repository

import "database/sql"

// ================================================================================
// Repository interface
// --------------------------------------------------------------------------------
type Repository interface {
	CreateNewSensor(sId int, sType, sName, sDescription string) error
	InsertSensorData(sId int, sValue string) error
	SelectAllSensors() ([]*Sensors, error)
	GetSensorLatestValue(sId int) (Sensor, error)
	UpdateSensor(sId int, sType, sName, sDescription string) (int64, error)
	SelectSingleSensor(sId int) (Sensors, error)
}

// ================================================================================
// postgres repository
// --------------------------------------------------------------------------------
type psqlRepo struct {
	DB *sql.DB
}

// psqlRepo implements Repository
func newPsqlRepo(conn *sql.DB) Repository {
	return &psqlRepo{
		DB: conn,
	}
}

// ================================================================================
// test repository
// --------------------------------------------------------------------------------
// type testRepository struct {
// 	DB *sql.DB
// }

// func newTestRepository(conn *sql.DB) Repository {
// 	return &testRepository{
// 		DB: nil,
// 	}
// }

// ================================================================================
