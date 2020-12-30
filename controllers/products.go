package controllers
import (
	"log"
	"net/http"
	//"time"
	"github.com/gin-gonic/gin"
	orm "github.com/go-pg/pg/v9/orm"
	"github.com/go-pg/pg/v9"
	//guuid "github.com/google/uuid"
)

type Products struct {
	Id               string    `json:"id"`
	Product_Name     string    `json:"product_name"`
	Price            string    `json:"price"`
}
// Create User Table
func CreateProductTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Products{}, opts)
	if createError != nil {
		log.Printf("Error while creating products table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Products table created")
	return nil
}

var dbConnect *pg.DB
func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func InsertProducts(c *gin.Context) {
	var products Products
	c.BindJSON(&products)
	product_name := products.Product_Name
	price := products.Price
	id := products.Id
	insertError := dbConnect.Insert(&Products{
		Id:                 id,
		Product_Name:       product_name,
		Price:              price,
	})
	if insertError != nil {
		log.Printf("Error while inserting new product into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Product Added Successfully",
	})
	return
}

func GetAllProducts(c *gin.Context) {
	var products []Products
	err := dbConnect.Model(&products).Select()
	if err != nil {
		log.Printf("Error while getting all Products, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Products",
		"data": products,
	})
	return
}

func GetSingleProduct(c *gin.Context) {
	productId := c.Param("productId")
	products := &Products{Id: productId}
	err := dbConnect.Select(products)
	if err != nil {
		log.Printf("Error while getting a single product, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Product not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Product",
		"data": products,
	})
	return
}

func UpdateProduct(c *gin.Context) {
	productId := c.Param("productId")
	var products Products
	c.BindJSON(&products)
	price := products.Price
	_, err := dbConnect.Model(&Products{}).Set("price = ?", price).Where("id = ?", productId).Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message":  "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Product Updated Successfully",
	})
	return
}

func DeleteProduct(c *gin.Context) {
	productId := c.Param("productId")
	products := &Products{Id: productId}
	err := dbConnect.Delete(products)
	if err != nil {
		log.Printf("Error while deleting a single product, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Product deleted successfully",
	})
	return
}