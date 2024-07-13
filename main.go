// Copyright 2016 LINE Corporation
//
// LINE Corporation licenses this file to you under the Apache License,
// version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at:
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v8/linebot/messaging_api"
	"github.com/line/line-bot-sdk-go/v8/linebot/webhook"

	"github.com/DOUIF/what-to-eat/cmd"
	"github.com/DOUIF/what-to-eat/models"
)

func main() {
	// load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	initialLineBotHandler()
	initialPingHandler()
	startServer()
}

func initialPingHandler() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("ping")
		fmt.Fprintf(w, "pong")
	})
}

func initialLineBotHandler() {
	// initial handler
	handler, err := webhook.NewWebhookHandler(
		os.Getenv("LINE_CHANNEL_SECRET"),
	)
	if err != nil {
		log.Fatal(err)
	}

	// initial line bot api
	bot, err := messaging_api.NewMessagingApiAPI(os.Getenv("LINE_CHANNEL_TOKEN"))
	if err != nil {
		log.Fatal(err)
	}

	// Setup HTTP Server for receiving requests from LINE platform
	handler.HandleEvents(func(req *webhook.CallbackRequest, r *http.Request) {
		if err != nil {
			log.Print(err)
			return
		}
		log.Println("Handling events...")
		for _, event := range req.Events {
			log.Printf("/callback called%+v...\n", event)
			switch e := event.(type) {
			case webhook.MessageEvent:
				switch message := e.Message.(type) {
				case webhook.TextMessageContent:
					_, err = bot.ReplyMessage(
						&messaging_api.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []messaging_api.MessageInterface{
								&messaging_api.TextMessage{
									Text: message.Text,
								},
							},
						},
					)
					if err != nil {
						log.Print(err)
					}
				case webhook.LocationMessageContent:
					reponse, err := cmd.GetRestaurantByLocation(message.Latitude, message.Longitude)
					if err != nil {
						log.Fatal(err)
						return
					}

					humanReadableText := formatResponse(reponse)
					_, err = bot.ReplyMessage(
						&messaging_api.ReplyMessageRequest{
							ReplyToken: e.ReplyToken,
							Messages: []messaging_api.MessageInterface{
								&messaging_api.TextMessage{
									Text: humanReadableText,
								},
							},
						},
					)
					if err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	http.Handle("/callback", handler)
}

// Convert the struct data into human-readable text
func formatResponse(response models.PlaceResponse) string {
	var sb strings.Builder

	for _, place := range response.Places {
		sb.WriteString(fmt.Sprintf("Name: %s\n", place.DisplayName.Text))
		sb.WriteString(fmt.Sprintf("Primary Type: %s\n", place.PrimaryTypeDisplayName.Text))
		sb.WriteString(fmt.Sprintf("Address: %s\n", place.FormattedAddress))
		sb.WriteString(fmt.Sprintf("Phone: %s\n", place.NationalPhoneNumber))
		sb.WriteString(fmt.Sprintf("Rating: %.1f (based on %d reviews)\n", place.Rating, place.UserRatingCount))
		sb.WriteString(fmt.Sprintf("Google Maps URL: %s\n", place.GoogleMapsUri))
		sb.WriteString(fmt.Sprintf("Location: (Lat: %.6f, Lon: %.6f)\n", place.Location.Latitude, place.Location.Longitude))
		sb.WriteString(fmt.Sprintf("Business Status: %s\n", place.BusinessStatus))
		sb.WriteString("Types: " + strings.Join(place.Types, ", ") + "\n")
		sb.WriteString(fmt.Sprintf("Open Now: %t\n", place.RegularOpeningHours.OpenNow))
		sb.WriteString("Opening Hours:\n")
		for _, day := range place.RegularOpeningHours.WeekdayDescriptions {
			sb.WriteString(fmt.Sprintf("  - %s\n", day))
		}
		sb.WriteString("Reviews:\n")
		for _, review := range place.Reviews {
			sb.WriteString(fmt.Sprintf("  - %s\n", review.Text.Text))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func startServer() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "5002"
	}
	host := os.Getenv("HOST")
	if host == "" {
		host = "0.0.0.0"
	}
	addr := fmt.Sprintf("%s:%s", host, port)
	fmt.Printf("listen on: %s\n", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}
