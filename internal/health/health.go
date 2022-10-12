package health

import (
	"caltoph/internal/db"
	"time"
)

const healthCheckIntervalDb int = 5

var dbHealth bool = false

func Init() {
	go checkDbHealth()
}

func checkDbHealth() {
	for {
		dbHealth = db.PingDB()
		time.Sleep(time.Duration(healthCheckIntervalDb) * time.Second)
	}
}

func GetDbHealth() bool {
	return dbHealth
}
