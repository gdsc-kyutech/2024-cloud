package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
)

var client *firestore.Client

func main() {
	var err error

	// initialize Firestore client
	ctx := context.Background()
	client, err = firestore.NewClient(ctx, os.Getenv("GOOGLE_CLOUD_PROJECT"))
	if err != nil {
		log.Fatalf("Error initializing Cloud Firestore client: %v", err)
	}

	// Get all customers and their treatments
	customers, err := getAllCustomers(ctx)
	if err != nil {
		log.Fatalf("Error getting customers: %v", err)
	}

	// Create JSON file
	err = createJSONFile("customers.json", customers)
	if err != nil {
		log.Fatalf("Error creating JSON file: %v", err)
	}

	fmt.Println("JSON file created successfully")
}

type Customer struct {
	Email      string                   `firestore:"email"`
	ID         string                   `firestore:"id"`
	Name       string                   `firestore:"name"`
	Phone      string                   `firestore:"phone"`
	Treatments []map[string]interface{} `json:"treatments"`
}

func getAllCustomers(ctx context.Context) ([]Customer, error) {
	var customers []Customer
	iter := client.Collection("customers").Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		var customer Customer
		err = doc.DataTo(&customer)
		if err != nil {
			return nil, err
		}

		// Fetch treatments for each customer
		treatments, err := getTreatments(ctx, customer.Email)
		if err != nil {
			return nil, err
		}
		customer.Treatments = treatments

		customers = append(customers, customer)
	}
	return customers, nil
}

func getTreatments(ctx context.Context, customerEmail string) ([]map[string]interface{}, error) {
	var treatments []map[string]interface{}
	iter := client.Collection(fmt.Sprintf("customers/%s/treatments", customerEmail)).Documents(ctx)
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return nil, err
		}
		treatments = append(treatments, doc.Data())
	}
	return treatments, nil
}

func createJSONFile(filename string, data interface{}) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(data)
	if err != nil {
		return err
	}
	return nil
}
