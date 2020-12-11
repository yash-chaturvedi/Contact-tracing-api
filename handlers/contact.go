package handlers

import (
	"contact/database"
	"contact/models"
	"context"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

// GetContact returns contact
func GetContact(c *fiber.Ctx) error {

	// Initialize Contact model slice
	var contacts []models.Contact

	// Parse Query Params
	userid := c.Query("user")
	timestamp, err := strconv.ParseInt(c.Query("infection_timestamp"), 10, 64)
	if err != nil {
		timestamp = time.Now().AddDate(0, 0, -14).UnixNano()
	}

	// Get the collection and handle error if any
	collection, err := database.GetCollection("contact", "contact")
	if err != nil {
		return err
	}

	// Filter Query
	filter := bson.M{
		"useridone": userid,
		"timestamp": bson.M{"$gt": timestamp},
	}

	// Find data in the collection and handle error if any
	cursor, err := collection.Find(context.Background(), filter)
	if err != nil {
		return err
	}

	// Append data in Contact model slice
	for cursor.Next(context.Background()) {
		var contact models.Contact
		cursor.Decode(&contact)
		contacts = append(contacts, contact)
	}

	// Finally return contact model slice
	return c.JSON(fiber.Map{
		"success":  true,
		"contacts": contacts,
	})
}

// AddContact adds contacts
func AddContact(c *fiber.Ctx) error {
	// Initialize Contact model
	contact := new(models.Contact)

	// Parse body and handle error if any
	if err := c.BodyParser(contact); err != nil {
		return err
	}

	// Generate current timestamp in epoch(mills)
	contact.Timestamp = time.Now().UnixNano()

	// Get Collection and handle error if any
	collection, err := database.GetCollection("contact", "contact")
	if err != nil {
		return err
	}

	// lock mutex to be thread safe
	database.Mu.Lock()

	// Insert to the collection and handle error if any
	res, err := collection.InsertOne(context.Background(), contact)
	if err != nil {
		return err
	}

	// Unlock mutex when done
	database.Mu.Unlock()

	// Finally return contact id
	return c.JSON(fiber.Map{
		"success":   true,
		"contactid": res.InsertedID,
	})

}
