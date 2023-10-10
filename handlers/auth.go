package handlers

import (
	"go-series-app/configs"
	"go-series-app/models"
	"go-series-app/services"
	"go-series-app/utils"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type LoginInput struct {
	Identity string `json:"identity"`
	Password string `json:"password"`
}
type UserData struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(userService services.UserService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		input := new(LoginInput)
		var ud UserData

		if err := c.BodyParser(input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "error", "message": "Error on login request", "errors": err.Error()})
		}

		identity := input.Identity
		pass := input.Password
		user, email, err := new(models.User), new(models.User), *new(error)

		user, err = userService.GetUserByUsername(identity)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Error on username", "errors": err.Error()})
		}
		ud = UserData{
			ID:       user.ID,
			Username: user.Username,
			Email:    user.Email,
			Password: user.Password,
		}

		if email == nil && user == nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "User not found", "errors": err.Error()})
		}

		if !utils.CheckPasswordHash(pass, ud.Password) {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid password", "data": nil})
		}

		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = ud.Username
		claims["user_id"] = ud.ID
		claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(configs.EnvConfigs.JwtSecret))
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.JSON(fiber.Map{"status": "success", "message": "Success login", "data": t})
	}
}
