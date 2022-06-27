package Auth

import (
	"fmt"
	"math/rand"
	"os"
	"trashgo/Database"
	"trashgo/Models"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
	"golang.org/x/crypto/bcrypt"
)

var (
	accountSid string
	authToken  string
	fromPhone  string
	toPhone    string
	client     *twilio.RestClient
)

func SendMessage(msg string) {

	params := openapi.CreateMessageParams{}
	params.SetTo(toPhone)
	params.SetFrom(fromPhone)
	params.SetBody(msg)

	response, err := client.Api.CreateMessage(&params)
	if err != nil {
		fmt.Printf("error creating and sending message: %s\n", err.Error())
		return
	}
	fmt.Printf("Message SID: %s\n", *response.Sid)
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("error loading .env: %s\n", err.Error())
		os.Exit(1)
	}

	accountSid = os.Getenv("ACCOUNT_SID")
	authToken = os.Getenv("AUTH_TOKEN")
	fromPhone = os.Getenv("FROM_PHONE")
	toPhone = os.Getenv("TO_PHONE")

	client = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

}

func Forgot(c *fiber.Ctx) error {

	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	email := data["email"]
	token := Randstring(12)

	resert := Models.Resetpass{
		Email: email,
		Token: token,
	}
	Database.DB.Create(&resert)

	url := "http://localhost:1234/resert"
	msg := fmt.Sprintf(os.Getenv("MSG"), url)
	SendMessage(msg)
	return c.JSON(fiber.Map{
		"massage": "cek your sms",
	})

}

func Resert(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] != data["password_confim"] {
		return c.Status(400).JSON(fiber.Map{
			"message": "Salah",
		})
	}
	passwordReset := Models.User{}
	Database.DB.Where(bcrypt.CompareHashAndPassword(passwordReset.Password, []byte(data["password"]))).Last(&passwordReset)
	password, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), bcrypt.DefaultCost)

	Database.DB.Model(&passwordReset).Where("Id_user = ?", passwordReset.IDUser).Updates(Models.User{Password: password})

	return c.JSON(fiber.Map{
		"massage": "Password Telah di ganti",
	})
}

func Randstring(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstucvxyzABCDEFGHIJKLMNOPQRSTUVWXYZ123456789")

	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]

	}
	return string(b)
}
