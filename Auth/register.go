package Auth

import (
	"strconv"
	"time"
	"trashgo/Database"
	"trashgo/Models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

const SecretKey = "secret"

func Register(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)
	user := Models.User{
		Email:          data["email"],
		Username:       data["username"],
		Password:       password,
		Address:        data["address"],
		Place_of_birth: data["place_of_birth"],
		Gender:         data["gender"],
		Foto:           data["Foto"],
		Title:          data["title"],
	}

	Database.DB.Create(&user)

	return c.JSON(fiber.Map{
		"massage": "Terima kasih telah mendaftar di TRASHGO",
	})
}

func Login(c *fiber.Ctx) error {

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	var user Models.User

	Database.DB.Where("Email = ?", data["email"]).First(&user)

	if user.IDUser == 0 {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "User not found!",
		})
	}

	if err := bcrypt.CompareHashAndPassword(user.Password, []byte(data["password"])); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect password!",
		})

	}

	claims := Models.Aku{StandardClaims: jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(user.IDUser)),
		ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
	},
		Username:       user.Username,
		Gender:         user.Gender,
		Email:          user.Email,
		Address:        user.Address,
		IDUser:         user.IDUser,
		Place_of_birth: user.Place_of_birth,
		Foto:           user.Foto,
		Title:          user.Title,
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := jwtToken.SignedString([]byte(SecretKey))

	if err != nil {
		c.Status(fiber.StatusInternalServerError)
		return c.JSON(fiber.Map{
			"message": "Tidak Bisa Login!",
		})
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"Massage": "Selamat Datang",
	})

}

func Users(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")

	token, err := jwt.ParseWithClaims(cookie, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})
	if err != nil {
		c.Status(fiber.StatusUnauthorized)
		return c.JSON(fiber.Map{
			"message": "unauthenticated",
		})
	}
	claims := token.Claims.(*jwt.StandardClaims)

	var user Models.User

	Database.DB.Where("Id_user = ?", claims.Issuer).First(&user)

	return c.JSON(user)
}
func Logout(c *fiber.Ctx) error {
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)

	return c.JSON(fiber.Map{
		"message": "success",
	})

}
