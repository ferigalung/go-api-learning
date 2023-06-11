package main

import "github.com/gofiber/fiber/v2"

type User struct {
	Name string `json:"name"`
}

func main() {
	app := fiber.New()
	// user := User{
	// 	Name: "feri",
	// }
	var arr = []int{
		1, 2, 3,
	}

	app.Get("/", func(c *fiber.Ctx) error {
		// return c.SendString(c.BaseURL())
		// return c.SendFile("file-does-not-exist")
		return c.JSON(arr)
	})

	app.Listen(":3000")
}
