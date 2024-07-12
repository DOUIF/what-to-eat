# What to eat

## Overview

This system uses the Google Maps API to find great restaurants nearby and communicates with users through a Line bot. Users can send their location to the Line bot, which responds with a list of highly-rated restaurants nearby.

## Features

- **Location-Based Search**: Users send their current location to the bot, which then finds nearby restaurants using the Google Maps API.
- **High-Rated Restaurants**: Shows only the best-rated restaurants.
- **Interactive Communication**: Simple interaction through the Line bot.

## Prerequisites

- Google Maps API Key
- Line Developers Account
- Go Programming Language

## Setup

1. **Clone the Repository**
    ```sh
    git clone <repository_url>
    cd <repository_directory>
    ```

2. **Set Up Environment Variables**
    - Create a `.env` file and add your Google Maps API key and Line bot credentials.
    ```env
    GOOGLE_MAPS_API_KEY=your_google_maps_api_key
    LINE_CHANNEL_SECRET=your_line_channel_secret
    LINE_CHANNEL_TOKEN=your_line_channel_token
    ```

3. **Install Dependencies**
    ```sh
    go mod tidy
    ```

4. **Run the Server**
    ```sh
    go run main.go
    ```

## Usage

1. **Add the Line Bot**: Add the Line bot to your Line account.
2. **Send Location**: Send your current location to the bot.
3. **Receive Recommendations**: The bot will respond with a list of nearby highly-rated restaurants.

## Project Structure

