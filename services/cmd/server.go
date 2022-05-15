package cmd

import (
	"log"
	"lottoengine/internals/persistence"
	"lottoengine/model"
	"lottoengine/services"
	"os"
)

var port = os.Getenv("HTTPPORT")

// RunServer -
func RunServer() error {
	var gameservices = services.GameHttpService{}

	//The Db
	dsn := persistence.Config(persistence.InitConfig())
	sqlDb := persistence.NewSqlLayer(dsn)
	sqlDb.Session.AutoMigrate(model.GamePlayLog{})
	err := gameservices.InitializeDb(sqlDb)
	if err != nil {
		log.Printf("RunServer() - Failed to initialize db with error %v", err)
		return err
	}

	//The Router
	gameservices.SetRoutes()

	err = gameservices.StartHttp(port)
	return err
}
