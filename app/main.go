package main

import (
	"fmt"
	"github.com/stryn26/MiniProjectEvermosRakamin/internal/helper"
	"github.com/stryn26/MiniProjectEvermosRakamin/internal/infrastructure/repository/container"
	"github.com/stryn26/MiniProjectEvermosRakamin/internal/infrastructure/mysql"

	rest "github.com/stryn26/MiniProjectEvermosRakamin/internal/server/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})

	containerConf := container.InitContainer()
	defer mysql.CloseDatabaseConnection(containerConf.Mysqldb)

	app := fiber.New()
	app.Use(logger.New())

	rest.HTTPRouteInit(app, containerConf)

	port := fmt.Sprintf("%s:%d", containerConf.Apps.Host, containerConf.Apps.HttpPort)
	helper.Logger("main.go", helper.LoggerLevelFatal, app.Listen(port).Error())
}
