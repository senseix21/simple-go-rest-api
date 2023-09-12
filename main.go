package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title     string
	Completed bool
}

func InitDB() *gorm.DB {
	dsn := "user=postgres password=M-eHra2Ue?LYii6 dbname=postgres host=db.gdwrapaueubkavnclrrz.supabase.co port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to the database")
	}

	db.AutoMigrate(&Todo{})

	return db
}

func setupRoutes(r *gin.Engine, db *gorm.DB) {
	// Create a new todo
	r.POST("/todos", func(c *gin.Context) {
		var todo Todo
		if err := c.BindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Create(&todo)
		c.JSON(http.StatusCreated, todo)
	})

	// Get all todos
	r.GET("/todos", func(c *gin.Context) {
		var todos []Todo
		db.Find(&todos)
		c.JSON(http.StatusOK, todos)
	})

	// Get a single todo by ID
	r.GET("/todos/:id", func(c *gin.Context) {
		var todo Todo
		id := c.Param("id")
		if db.First(&todo, id).Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		c.JSON(http.StatusOK, todo)
	})

	// Update a todo by ID
	r.PUT("/todos/:id", func(c *gin.Context) {
		var todo Todo
		id := c.Param("id")
		if db.First(&todo, id).Error != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
			return
		}
		if err := c.BindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		db.Save(&todo)
		c.JSON(http.StatusOK, todo)
	})

	// Delete a todo by ID
	r.DELETE("/todos/:id", func(c *gin.Context) {
		id := c.Param("id")
		db.Delete(&Todo{}, id)
		c.Status(http.StatusNoContent)
	})
}

func main() {
	r := gin.Default()
	db := InitDB()

	setupRoutes(r, db)

	// Start the server on port 8080
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Failed to start the server:", err)
	}
}
