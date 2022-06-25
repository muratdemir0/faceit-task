package errors

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Handler(logger *zap.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {
		if err != nil {

			res := makeResponse(err)
			logger.Error("an error: "+err.Error(), zap.Int("status code", res.StatusCode()))

			if err = c.Status(res.StatusCode()).JSON(res); err != nil {
				logger.Error("failed writing error making response: "+err.Error(), zap.Error(err))
				return err
			}
		}
		return nil
	}
}
