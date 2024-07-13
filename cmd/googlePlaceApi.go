package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/DOUIF/what-to-eat/models"
)

func GetRestaurantByLocation(latitude float64, longitude float64) (models.PlaceResponse, error) {
	var placesResponse models.PlaceResponse

	payload := &models.PlaceRequest{
		IncludedTypes:  []string{"restaurant"},
		LanguageCode:   "zh-TW",
		RegionCode:     "TW",
		MaxResultCount: 1,
		RankPreference: "POPULARITY",
		LocationRestriction: models.LocationRestriction{
			Circle: models.Circle{
				Center: models.Center{
					Latitude:  latitude,
					Longitude: longitude,
				},
				Radius: 1000.0,
			},
		},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return placesResponse, err
	}

	request, err := http.NewRequest("POST", os.Getenv("GOOGLE_PLACE_API"), bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return placesResponse, err
	}

	request.Header.Set("X-Goog-Api-Key", os.Getenv("GOOGLE_MAPS_API_KEY"))
	request.Header.Set("X-Goog-FieldMask", os.Getenv("GOOGLE_PLACE_MASK_FILEDS"))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return placesResponse, err
	}
	defer response.Body.Close()

	// Read and print the response
	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return placesResponse, err
	}

	fmt.Println("Response: ", string(body))

	err = json.Unmarshal([]byte(body), &placesResponse)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return placesResponse, err
	}

	// Print the response
	fmt.Printf("%+v\n", response)
	return placesResponse, nil
}
