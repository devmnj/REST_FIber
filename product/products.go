package product

import ( 
"github.com/gofiber/fiber/v2"
"example/REST_API/database"
// "gorm.io/driver/sqlite"
"gorm.io/gorm"
)

type Product struct{
   gorm.Model
   Name string `json:"name"`
   Category string `json:category`
   Price int `json:"price"`
}

func GetProducts(c *fiber.Ctx) error {
	db:=database.DBConn
	var products []Product
    db.Find(&products)
   return  c.JSON(products)  
}

func GetProduct(c *fiber.Ctx) error {
	db:=database.DBConn
	id:=c.Params("id")
	var product Product
    db.Find(&product,id)
    return  c.JSON(product)  
}

func DemoProduct(c *fiber.Ctx) error {
	db:= database.DBConn
	var product Product
	product.Name="Mac Mini"
	product.Category="Computer"
	product.Price=10000
    db.Create(&product)

   return  c.JSON(product)  
}

func DelProduct(c *fiber.Ctx) error {
	db:=database.DBConn
	id:=c.Params("id")
	var product Product
    db.Find(&product,id)
	if product.Name==""{
		c.Status(500).SendString("No Product located")
	}
	db.Delete(product)
    return  c.SendString("Product Deleted")  
}


func AddProduct(c *fiber.Ctx) error {
	db:= database.DBConn
	product:=new (Product)
	if err:=c.BodyParser(product); err!=nil{
		c.Status(503)
		 
	}

    db.Create(&product)

   return  c.JSON(product)  
}
