package handlers

import (
	"contact/database"
	"contact/models"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GetUser returns a user
func GetUser(c *fiber.Ctx) error {
	// initialize user model
	user := new(models.User)

	// Parse request
	userid, _ := primitive.ObjectIDFromHex(c.Params("id"))

	// Get the Collection and handle error if any
	collection, err := database.GetCollection("contact", "users")
	if err != nil {
		return err
	}

	// Insert to the collection and handle error if any
	err = collection.FindOne(context.Background(), models.User{ID: userid}).Decode(&user)
	if err != nil {
		return err
	}

	// Finally return user
	t := time.Unix(0,user.CreatedAt)

	return c.JSON(fiber.Map{
		"userId":  user.ID,
		"name": user.Name,
		"birthDate": user.BirthDate,
		"phone": user.PhoneNumber,
		"email": user.Email,
		"createdAt": t.Format("02/01-2006 15:04:05"),
	})
}

// CreateUser creates a user
func CreateUser(c *fiber.Ctx) error {
	// Initialize user model
	user := new(models.User)

	// Parse body and handle error if any
	if err := c.BodyParser(user); err != nil {
		return err
	}

	// Generate current timestamp in epoch(mills)
	user.CreatedAt = time.Now().UnixNano()

	// Get the collection and handle error if any
	collection, err := database.GetCollection("contact", "users")
	if err != nil {
		return err
	}

	// Insert the data and handle error if any
	res, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	t := time.Unix(0,user.CreatedAt)

	// Finally return user id
	return c.JSON(fiber.Map{
		"userId":  res.InsertedID,
		"name": user.Name,
		"birthDate": user.BirthDate,
		"phone": user.PhoneNumber,
		"email": user.Email,
		"createdAt": t.Format("02/01-2006 15:04:05"),
	})

}
