package main 

import ( 
	_ "strings"
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// TitleText("how are you today!")
	// CapitalizedText("hello world text")
	fmt.Println("Hello Main function")

	app := fiber.New()
	// GET http://localhost:8080/hello%20world
	// app.Get("/", func(c *fiber.Ctx) error {
	// 	return c.SendString("Hello SendString")
	// })


	// app.Get("/:value", func(c *fiber.Ctx) error {
	// 	return c.SendString("value: " + c.Params("value"))
	// })

	//wildcard simulation
	// GET http://localhost:3000/api/user/john
	app.Get("/api/*", func(c *fiber.Ctx) error {
		return c.SendString("API Path: " + c.Params("*"))
	})

	// app.Get("/login", func(c *fiber.Ctx) error {
	// 	return c.Render("login")
	// })

	app.Static("/", "./public")



	// app.Get("/:value", func)

	app.Listen(":3000")
}



// func TitleText(textToBeCapitalized string) string {
// 	titleOftheText := strings.Title(textToBeCapitalized)
// 	fmt.Println("entitle the text: " + titleOftheText)
// 	return titleOftheText
// }

// func CapitalizedText(paramText string) string {
// 	textToBeUppered := strings.ToUpper(paramText)
// 	fmt.Println("textToBeUppered: " + paramText)
// 	return textToBeCapitalized
// }