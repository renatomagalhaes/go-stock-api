package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/renatomagalhaes/go-stock-api/database"
	"log"
	"math/rand"
	"time"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	defer db.Close()

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 100; i++ {
		sku := fmt.Sprintf("SKU%d", i)
		name := fmt.Sprintf("Product%d", i)
		types := []int{1, 2, 3}
		typeIndex := rand.Intn(len(types))
		typeValue := types[typeIndex]
		quantity := rand.Intn(100) + 1

		_, err = db.Exec(`
			INSERT INTO stocks (sku, name, type, quantity)
			VALUES (?, ?, ?, ?)
		`, sku, name, typeValue, quantity)
		if err != nil {
			log.Fatal(err)
		}
	}
}
