package persistence

import (
	"log"
	"lottoengine/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type SqlLayer struct {
	Session *gorm.DB
}

func NewSqlLayer(dsn string) *SqlLayer {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to open database: %v", err)
	}

	return &SqlLayer{Session: db}
}

func (sql *SqlLayer) PersistGameLog(gamelog model.GamePlayLog) (string, error) {
	err := sql.Session.Create(&gamelog).Error
	if err != nil {
		return "", err
	}
	return gamelog.Id, err
}

func (sql *SqlLayer) GetGameLog(options model.GamePlayLog) ([]model.GamePlayLog, error) {
	var result []model.GamePlayLog
	err := sql.Session.Where(options).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, err
}
