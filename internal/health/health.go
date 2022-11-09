package health

import (
	"caltoph/internal/db"
	"strings"
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

func getDbHealth() bool {
	return dbHealth
}

func GetHealth() (bool, string) {
	var sb strings.Builder
	var error = false
	if !getDbHealth() {
		error = true
		sb.WriteString("CAN'T REACH DATABASE\n")
	}
	if error {
		return false, sb.String()
	}
	return true, "HEALTHY"
}
