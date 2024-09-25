package main

import (
	"encoding/json"	
	"net/http"

	"os"
	"context"
	"log"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type Airport struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	IATA    string `json:"iata"`
	ImageURL string `json:"image_url"`
}

type AirportV2 struct {
	Airport
	RunwayLength int `json:"runway_length"`
}

// Mock data for airports in Bangladesh
var airports = []Airport{
	{"Hazrat Shahjalal International Airport", "Dhaka", "DAC", "https://storage.googleapis.com/bd-airport-data/dac.jpg"},
	{"Shah Amanat International Airport", "Chittagong", "CGP", "https://storage.googleapis.com/bd-airport-data/cgp.jpg"},
	{"Osmani International Airport", "Sylhet", "ZYL", "https://storage.googleapis.com/bd-airport-data/zyl.jpg"},
}

// Mock data for airports in Bangladesh (with runway length for V2)
var airportsV2 = []AirportV2{
	{Airport{"Hazrat Shahjalal International Airport", "Dhaka", "DAC", "https://storage.googleapis.com/bd-airport-data/dac.jpg"}, 3200},
	{Airport{"Shah Amanat International Airport", "Chittagong", "CGP", "https://storage.googleapis.com/bd-airport-data/cgp.jpg"}, 2900},
	{Airport{"Osmani International Airport", "Sylhet", "ZYL", "https://storage.googleapis.com/bd-airport-data/zyl.jpg"}, 2500},
}

// HomePage handler
func HomePage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Status: OK"))
}

// Airports handler for the first endpoint
func Airports(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(airports)
}

// AirportsV2 handler for the second version endpoint
func AirportsV2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(airportsV2)
}

// ##############################
// ## TODO: Edit this function ##
// ##############################

// UpdateAirportImage handler for updating airport images
func UpdateAirportImage(w http.ResponseWriter, r *http.Request) {
	// Parse the request body to get the airport name and image data
	r.ParseMultipartForm(1 << 20)
	img, metadata, err := r.FromFile("img")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer img.close()
	img_key := metadata.Filename
	airport_name := r.FromValue("name")
	if airport_name == "" {
		http.Error(w, "No Airport Name provided", http.StatusBadRequest)
		return
	}

	// Find the airport by name
	var Updated_airport AirportV2
	airport_idx := 0
	for ; airport_idx < length(airportsV2) ; airport_idx++ {
		if airportsV2[airport_idx].Airport.Name == airport_name {
			Updated_airport = airportsV2[airport_idx]
			break
		}
	}
	if airport_idx >= length(airportsV2) {
		http.Error(w, "No Airport found with the name", http.StatusNotFound)
		return
	}
	var Updated_airport AirportV2

	// Initialize GCS client
	// loading the shared configs from ~/.aws/config
	bucket_name := os.Getenv("BUCKET_TO_UPLOAD")
	ctx := context.TODO()
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		log.Fatal(err)
	}
	s3_client := s3.NewFromConfig(cfg) 

	// Upload image to GCS and update the airport's image URL
	_, err := s3_client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket_name),
		Key: aws.String(img_key),
		Body: img,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	Updated_airport.Airport.ImageURL = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", bucket_name, img_key)
	airports_v2[airport_idx] = Updated_airport
	airports[airport_idx] = Updated_airport.Airport

	// Respond with success/failure
	json.NewEncoder(w).Encode(Updated_airport)
	return
}

func main() {
	// Setup routes
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/airports", Airports)
	http.HandleFunc("/airports_v2", AirportsV2)

	// TODO: complete the UpdateAirportImage handler function
	http.HandleFunc("/update_airport_image_v2", UpdateAirportImage)

	// Start the server
	http.ListenAndServe(":8080", nil)
}
