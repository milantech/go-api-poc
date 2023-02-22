package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//ConnectDB connects go to mysql database
func ConnectDB() *gorm.DB {
	//	errorENV := godotenv.Load()
	// if errorENV != nil {
	// 	panic("Failed to load env file")
	// }

	// dbUser := os.Getenv("DB_USER")
	// dbPass := os.Getenv("DB_PASS")
	// dbHost := os.Getenv("DB_HOST")
	// dbName := os.Getenv("DB_NAME")
	dsn := "postgresql://root:root@localhost:26257/sampledbtwo?sslmode=disable"
	//dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, errorDB := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if errorDB != nil {
		panic("Failed to connect mysql database")
	}

	return db
}

//DisconnectDB is stopping your connection to mysql database
func DisconnectDB(db *gorm.DB) {
	dbSQL, err := db.DB()
	if err != nil {
		panic("Failed to kill connection from database")
	}
	dbSQL.Close()
}
