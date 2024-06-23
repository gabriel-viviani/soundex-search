package main

import (
	"fmt"

	"github.com/Sigma-Ratings/sigma-code-challenges/api/config"
	"github.com/Sigma-Ratings/sigma-code-challenges/api/pkg/db"
	"github.com/Sigma-Ratings/sigma-code-challenges/api/pkg/service"
	"github.com/Sigma-Ratings/sigma-code-challenges/api/routes"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

// @title Sigma Backend Code Challenge
// @version 0.2
// @description This is a boiler plate API for the Sigma Backend Code Challenge

// @contact.name Sigma Team - Or anyone that you are in contact with during the hiring process
// @contact.url https://sigmaratings.com
// @contact.email support@sigmaratings.com
// @BasePath /
func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logLvl, err := logrus.ParseLevel(config.App.LogLevel)
	if err != nil {
		logrus.Warn("could not parse log level, using debug default")
		logLvl = logrus.DebugLevel
	}
	logrus.SetLevel(logLvl)
	e := createApi()
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.App.Port)))
}

func createApi() *echo.Echo {
	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	database, err := db.GetConnection(config.Database())
	if err != nil {
		logrus.WithError(err).Fatal("could not connect to database") // Fatal will exit the application
	}

	routes.AttachRoutes(e, database)
	ss := service.SanctionsImport{DB: database}
	go ss.ImportSanctions()
	return e
}
