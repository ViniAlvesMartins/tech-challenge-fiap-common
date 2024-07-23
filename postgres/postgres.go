package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

func NewConnection(host, user, pwd, dbName, port, dbSchema string) (*gorm.DB, error) {
	var err error
	var conn *gorm.DB

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s search_path=%s",
		host, user, pwd, dbName, port, dbSchema)

	conn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: gormlogger.Discard,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   fmt.Sprintf("%s.", dbSchema),
			SingularTable: false,
		},
	})

	if err != nil {
		return nil, err
	}

	db, err := conn.DB()

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
