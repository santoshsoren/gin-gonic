package controllers
import (
	"log"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
	orm "github.com/go-pg/pg/v9/orm"
	"github.com/go-pg/pg/v9"
	guuid "github.com/google/uuid"
)

type Todo struct {
	Id        string    `json:"id"`
	Name     string    `json:"name"`
	Price      string    `json:"price"`
}
// Create User Table
func CreateTodoTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Todo{}, opts)
	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Todo table created")
	return nil
}

var dbConnect *pg.DB
func InitiateDB(db *pg.DB) {
	dbConnect = db
}