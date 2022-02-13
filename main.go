package main

import  
("github.com/gofiber/fiber/v2"
"example/REST_API/product"
"example/REST_API/database"
"fmt"
"log"
"gorm.io/gorm"
"gorm.io/driver/sqlite"
)

func index(c *fiber.Ctx) error {
   return  c.SendString("<h1>This is a Fiber-GO API</h1>")  
}

func setupRoutes(app *fiber.App){
    app.Get("/",index )
    app.Get("/api/v1/products",product.GetProducts )
    app.Get("/api/v1/products/:id",product.GetProduct )
    app.Post("/api/v1/demoproducts",product.DemoProduct )
    app.Post("/api/v1/products",product.AddProduct )
    app.Delete("/api/v1/products/:id",product.DelProduct )
}

func initDB(){
    var err error
   database.DBConn, err = gorm.Open(sqlite.Open("products.db"), &gorm.Config{})
    if err != nil {
      panic("failed to connect database")
    }
    fmt.Println("Database connection successfully Opened !")
    database.DBConn.AutoMigrate(product.Product{})
    fmt.Println("Migration completed successfully")
}

func main(){
	app:= fiber.New()
    initDB()
    setupRoutes(app)
    fmt.Println("Server starting in a minute")
    log.Fatal(app.Listen(":3000"))
}