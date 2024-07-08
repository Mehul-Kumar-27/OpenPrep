package database

import (
	"time"

	"github.com/Mehul-Kumar-27/OpenBackend/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var PostgresConnHandler *DBHandler

type DatabaseConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	Retry    int           // Number of retries for connection
	Interval time.Duration // Interval between retries
}

type DBManager interface {
	Connect() (*gorm.DB, error)
	Close()
}

type DBHandler struct {
	dbManager DBManager
	config    DatabaseConfig
	DB        *gorm.DB
}

func InitializeDBHandler() {

	dbConfig := &DatabaseConfig{
		Host:     config.AppConfig.Postgreshost,
		Port:     config.AppConfig.PostgresPort,
		User:     config.AppConfig.PostgresUser,
		Password: config.AppConfig.PostgresPassword,
		Name:     config.AppConfig.PostgresDB,
		Retry:    5,
		Interval: 5 * time.Second,
	}

	PostgresConnHandler = &DBHandler{
		config: *dbConfig,
	}

	PostgresConnHandler.DB, _ = PostgresConnHandler.Connect()

	if PostgresConnHandler.DB == nil {
		config.Log.Fatal("Failed to connect to database")
	}

	config.Log.Info("Connected to database")
}

func (db *DBHandler) getDSN() string {
	return "host=" + db.config.Host +
		" port=" + db.config.Port +
		" user=" + db.config.User +
		" password=" + db.config.Password +
		" dbname=" + db.config.Name +
		" sslmode=disable"

}

func (db *DBHandler) Connect() (*gorm.DB, error) {
	var err error
	for i := 0; i < db.config.Retry; i++ {
		db.DB, err = gorm.Open(postgres.Open(db.getDSN()), &gorm.Config{})
		if err == nil {
			return db.DB, nil
		}
		config.Log.Info("Failed to connect to database. Retrying...")
		time.Sleep(db.config.Interval)
	}
	return nil, err
}


func (db *DBHandler) InitalizeDatabase(){
	
}
