package main

import (
	"log"
	"trashgo/Auth"
	"trashgo/Database"
	"trashgo/Home"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {

	Database.Connect()
	app := fiber.New()
	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Post("/auth/register", Auth.Register)
	app.Post("/auth/login", Auth.Login)
	app.Get("/auth/user", Auth.Users)
	app.Get("/auth/logout", Auth.Logout)
	app.Post("/home/laporan", Home.Laporan)
	app.Post("/auth/akun", Auth.Akun)
	app.Post("/profil/akun/edit_profil", Auth.Edit)
	app.Post("/reset", Auth.Reset)
	app.Post("/reset/v1", Auth.Forgot)
	app.Post("/resert", Auth.Resert)
	app.Get("/google",Auth.OauthClient)
	log.Fatal(app.Listen(":1234"))
}
