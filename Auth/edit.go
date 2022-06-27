package Auth

import (
	"trashgo/Database"
	"trashgo/Models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

func Akun(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user Models.User

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "gagal terhubung",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	Database.DB.Model(&user).Where("Id_user = ?", claims.Issuer).Updates(Models.User{Title: data["title"], Foto: data["foto"]})

	return c.JSON(user)
}
func Edit(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user Models.User

	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "gagal terhubung",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	Database.DB.Model(&user).Where("Id_user = ?", claims.Issuer).Updates(Models.User{Username: data["username"], Title: data["title"]})

	return c.JSON(user)
}
