package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var customers = []Customer{
	{
		Email: "amani_ohara@hassie.info",
		ID:    "22530",
		Name:  "Amani O'Hara",
		Phone: "(989) 239-0083",
		Treatments: []Treatment{
			{Cost: 1602, Description: "Bladder Stone Removal", Status: "proposed"},
			{Cost: 33, Description: "Vaccination", Status: "approved"},
			{Cost: 489, Description: "Benign Skin Mass Removal", Status: "rejected"},
			{Cost: 552, Description: "Tooth Extraction", Status: "approved"},
		},
	},
	{
		Email: "brooke_torp@lola.net",
		ID:    "34216",
		Name:  "Brooke Torp",
		Phone: "676.661.1456",
		Treatments: []Treatment{
			{Cost: 28, Description: "Vaccination", Status: "approved"},
			{Cost: 36, Description: "Deworming medication", Status: "proposed"},
		},
	},
	{
		Email:      "elroy_dickens88@nestor.biz",
		ID:         "70156",
		Name:       "Elroy Dickens",
		Phone:      "481-816-2004",
		Treatments: nil,
	},
	{
		Email:      "hank_ullrich@hillard.info",
		ID:         "82964",
		Name:       "Hank Ullrich",
		Phone:      "711.757.4837 x847",
		Treatments: nil,
	},
	{
		Email:      "lukas.dickinson@morgan.info",
		ID:         "63503",
		Name:       "Lukas Dickinson",
		Phone:      "381.878.7573 x7866",
		Treatments: nil,
	},
	{
		Email:      "orlando80@derrick.org",
		ID:         "31652",
		Name:       "Orlando Lubowitz",
		Phone:      "746.302.9473",
		Treatments: nil,
	},
	{
		Email:      "shanelle.gleason@verona.name",
		ID:         "63570",
		Name:       "Shanelle Gleason",
		Phone:      "1-216-134-2722 x5827",
		Treatments: nil,
	},
	{
		Email:      "timmothy.lang@stevie.org",
		ID:         "96978",
		Name:       "Timmothy Lang",
		Phone:      "(329) 651-8694",
		Treatments: nil,
	},
	{
		Email:      "verdie94@joanny.org",
		ID:         "99627",
		Name:       "Verdie Dibbert",
		Phone:      "(748) 047-3963 x81784",
		Treatments: nil,
	},
	{
		Email:      "willis.quigley@alfonzo.net",
		ID:         "45510",
		Name:       "Willis Quigley",
		Phone:      "1-656-365-9602",
		Treatments: nil,
	},
}

func main() {
	// get the port to listen on
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// create a new router
	r := mux.NewRouter()
	r.HandleFunc("/v1/", rootHandler)
	r.HandleFunc("/v1/customer/{id}", customerHandler)

	// start the server
	log.Println("Pets REST API listening on port", port)
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Authorization", "Origin"}),
		handlers.AllowedOrigins([]string{"https://storage.googleapis.com"}),
		handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "OPTIONS", "PATCH", "CONNECT"}),
	)

	if err := http.ListenAndServe(":"+port, cors(r)); err != nil {
		log.Fatalf("Error launching Pets REST API server: %v", err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "{status: 'running'}")
}

func customerHandler(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	customer, err := getCustomer(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": "fail", "data": '%s'}`, err)
		return
	}
	if customer == nil {
		w.WriteHeader(http.StatusNotFound)
		msg := fmt.Sprintf("`Customer \"%s\" not found`", id)
		fmt.Fprintf(w, fmt.Sprintf(`{"status": "fail", "data": {"title": %s}}`, msg))
		return
	}
	amount, err := getAmounts(customer)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": "fail", "data": "Unable to fetch amounts: %s"}`, err)
		return
	}
	data, err := json.Marshal(amount)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, `{"status": "fail", "data": "Unable to fetch amounts: %s"}`, err)
		return
	}
	fmt.Fprintf(w, fmt.Sprintf(`{"status": "success", "data": %s}`, data))
}

type Customer struct {
	Email      string      `json:"Email"`
	ID         string      `json:"ID"`
	Name       string      `json:"Name"`
	Phone      string      `json:"Phone"`
	Treatments []Treatment `json:"treatments"`
}

type Treatment struct {
	Cost        int64  `json:"cost"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func getCustomer(id string) (*Customer, error) {
	for _, customer := range customers {
		if customer.ID == id {
			return &customer, nil
		}
	}
	return nil, nil
}

func getAmounts(c *Customer) (map[string]int64, error) {
	if c == nil {
		return map[string]int64{}, fmt.Errorf("Customer should be non-nil: %v", c)
	}
	result := map[string]int64{
		"proposed": 0,
		"approved": 0,
		"rejected": 0,
	}
	if c.Treatments == nil {
		return result, nil
	}
	for _, treatment := range c.Treatments {
		switch treatment.Status {
		case "proposed":
			result["proposed"] += treatment.Cost
		case "approved":
			result["approved"] += treatment.Cost
		case "rejected":
			result["rejected"] += treatment.Cost
		}
	}
	return result, nil
}
