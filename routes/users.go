package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sachinsmc/hubuc-task/controllers"
	"github.com/sachinsmc/hubuc-task/model"
	"go.uber.org/zap"
)

func CreateUser(c *fiber.Ctx) error {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	u := new(model.User)
	if err := c.BodyParser(u); err != nil {
		logger.Info("failed to fetch URL")
		return err
	}

	return c.JSON(controllers.Add(u))
}
