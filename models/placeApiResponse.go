package models

import "time"

type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type RegularOpeningHours struct {
	OpenNow             bool     `json:"openNow"`
	WeekdayDescriptions []string `json:"weekdayDescriptions"`
}

type Text struct {
	Text string `json:"text"`
}

type Review struct {
	Text Text `json:"text"`
}

type Place struct {
	Types                  []string            `json:"types"`
	NationalPhoneNumber    string              `json:"nationalPhoneNumber"`
	FormattedAddress       string              `json:"formattedAddress"`
	Location               Location            `json:"location"`
	Rating                 float64             `json:"rating"`
	GoogleMapsUri          string              `json:"googleMapsUri"`
	RegularOpeningHours    RegularOpeningHours `json:"regularOpeningHours"`
	BusinessStatus         string              `json:"businessStatus"`
	UserRatingCount        int                 `json:"userRatingCount"`
	DisplayName            Text                `json:"displayName"`
	PrimaryTypeDisplayName Text                `json:"primaryTypeDisplayName"`
	PrimaryType            string              `json:"primaryType"`
	Reviews                []Review            `json:"reviews"`
}

type PlaceResponse struct {
	Places []Place `json:"places"`
}

type PlaceOld struct {
	Types               []string `json:"types"`
	NationalPhoneNumber string   `json:"nationalPhoneNumber,omitempty"`
	FormattedAddress    string   `json:"formattedAddress"`
	Location            struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"location"`
	Rating              float64 `json:"rating"`
	GoogleMapsURI       string  `json:"googleMapsUri"`
	RegularOpeningHours struct {
		OpenNow             bool     `json:"openNow"`
		WeekdayDescriptions []string `json:"weekdayDescriptions"`
	} `json:"regularOpeningHours"`
	BusinessStatus  string `json:"businessStatus"`
	UserRatingCount int    `json:"userRatingCount"`
	DisplayName     struct {
		Text string `json:"text"`
	} `json:"displayName"`
	PrimaryTypeDisplayName struct {
		Text string `json:"text"`
	} `json:"primaryTypeDisplayName"`
	PrimaryType string `json:"primaryType"`
	Reviews     []struct {
		Text struct {
			Text string `json:"text"`
		} `json:"text"`
	} `json:"reviews"`
}

type PlaceApiResponseAll struct {
	Places []struct {
		Name                     string   `json:"name"`
		ID                       string   `json:"id"`
		Types                    []string `json:"types"`
		NationalPhoneNumber      string   `json:"nationalPhoneNumber,omitempty"`
		InternationalPhoneNumber string   `json:"internationalPhoneNumber,omitempty"`
		FormattedAddress         string   `json:"formattedAddress"`
		AddressComponents        []struct {
			LongText     string   `json:"longText"`
			ShortText    string   `json:"shortText"`
			Types        []string `json:"types"`
			LanguageCode string   `json:"languageCode"`
		} `json:"addressComponents"`
		PlusCode struct {
			GlobalCode   string `json:"globalCode"`
			CompoundCode string `json:"compoundCode"`
		} `json:"plusCode"`
		Location struct {
			Latitude  float64 `json:"latitude"`
			Longitude float64 `json:"longitude"`
		} `json:"location"`
		Viewport struct {
			Low struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"low"`
			High struct {
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
			} `json:"high"`
		} `json:"viewport"`
		Rating              float64 `json:"rating"`
		GoogleMapsURI       string  `json:"googleMapsUri"`
		WebsiteURI          string  `json:"websiteUri,omitempty"`
		RegularOpeningHours struct {
			OpenNow bool `json:"openNow"`
			Periods []struct {
				Open struct {
					Day    int `json:"day"`
					Hour   int `json:"hour"`
					Minute int `json:"minute"`
				} `json:"open"`
				Close struct {
					Day    int `json:"day"`
					Hour   int `json:"hour"`
					Minute int `json:"minute"`
				} `json:"close"`
			} `json:"periods"`
			WeekdayDescriptions []string `json:"weekdayDescriptions"`
		} `json:"regularOpeningHours"`
		UtcOffsetMinutes    int    `json:"utcOffsetMinutes"`
		AdrFormatAddress    string `json:"adrFormatAddress"`
		BusinessStatus      string `json:"businessStatus"`
		PriceLevel          string `json:"priceLevel,omitempty"`
		UserRatingCount     int    `json:"userRatingCount"`
		IconMaskBaseURI     string `json:"iconMaskBaseUri"`
		IconBackgroundColor string `json:"iconBackgroundColor"`
		DisplayName         struct {
			Text         string `json:"text"`
			LanguageCode string `json:"languageCode"`
		} `json:"displayName"`
		PrimaryTypeDisplayName struct {
			Text         string `json:"text"`
			LanguageCode string `json:"languageCode"`
		} `json:"primaryTypeDisplayName"`
		Takeout             bool `json:"takeout"`
		Delivery            bool `json:"delivery"`
		DineIn              bool `json:"dineIn"`
		CurbsidePickup      bool `json:"curbsidePickup,omitempty"`
		Reservable          bool `json:"reservable,omitempty"`
		ServesBreakfast     bool `json:"servesBreakfast"`
		ServesLunch         bool `json:"servesLunch"`
		ServesDinner        bool `json:"servesDinner"`
		ServesBeer          bool `json:"servesBeer"`
		ServesWine          bool `json:"servesWine,omitempty"`
		ServesBrunch        bool `json:"servesBrunch,omitempty"`
		CurrentOpeningHours struct {
			OpenNow bool `json:"openNow"`
			Periods []struct {
				Open struct {
					Day    int `json:"day"`
					Hour   int `json:"hour"`
					Minute int `json:"minute"`
					Date   struct {
						Year  int `json:"year"`
						Month int `json:"month"`
						Day   int `json:"day"`
					} `json:"date"`
				} `json:"open,omitempty"`
				Close struct {
					Day    int `json:"day"`
					Hour   int `json:"hour"`
					Minute int `json:"minute"`
					Date   struct {
						Year  int `json:"year"`
						Month int `json:"month"`
						Day   int `json:"day"`
					} `json:"date"`
				} `json:"close,omitempty"`
				Close0 struct {
					Day       int  `json:"day"`
					Hour      int  `json:"hour"`
					Minute    int  `json:"minute"`
					Truncated bool `json:"truncated"`
					Date      struct {
						Year  int `json:"year"`
						Month int `json:"month"`
						Day   int `json:"day"`
					} `json:"date"`
				} `json:"close,omitempty"`
				Open0 struct {
					Day       int  `json:"day"`
					Hour      int  `json:"hour"`
					Minute    int  `json:"minute"`
					Truncated bool `json:"truncated"`
					Date      struct {
						Year  int `json:"year"`
						Month int `json:"month"`
						Day   int `json:"day"`
					} `json:"date"`
				} `json:"open,omitempty"`
			} `json:"periods"`
			WeekdayDescriptions []string `json:"weekdayDescriptions"`
		} `json:"currentOpeningHours"`
		CurrentSecondaryOpeningHours []struct {
			OpenNow bool `json:"openNow"`
			Periods []struct {
				Open struct {
					Day    int `json:"day"`
					Hour   int `json:"hour"`
					Minute int `json:"minute"`
					Date   struct {
						Year  int `json:"year"`
						Month int `json:"month"`
						Day   int `json:"day"`
					} `json:"date"`
				} `json:"open,omitempty"`
				Close struct {
					Day    int `json:"day"`
					Hour   int `json:"hour"`
					Minute int `json:"minute"`
					Date   struct {
						Year  int `json:"year"`
						Month int `json:"month"`
						Day   int `json:"day"`
					} `json:"date"`
				} `json:"close,omitempty"`
				Close0 struct {
					Day       int  `json:"day"`
					Hour      int  `json:"hour"`
					Minute    int  `json:"minute"`
					Truncated bool `json:"truncated"`
					Date      struct {
						Year  int `json:"year"`
						Month int `json:"month"`
						Day   int `json:"day"`
					} `json:"date"`
				} `json:"close,omitempty"`
				Open0 struct {
					Day       int  `json:"day"`
					Hour      int  `json:"hour"`
					Minute    int  `json:"minute"`
					Truncated bool `json:"truncated"`
					Date      struct {
						Year  int `json:"year"`
						Month int `json:"month"`
						Day   int `json:"day"`
					} `json:"date"`
				} `json:"open,omitempty"`
			} `json:"periods"`
			WeekdayDescriptions []string `json:"weekdayDescriptions"`
			SecondaryHoursType  string   `json:"secondaryHoursType"`
		} `json:"currentSecondaryOpeningHours,omitempty"`
		RegularSecondaryOpeningHours []struct {
			OpenNow bool `json:"openNow"`
			Periods []struct {
				Open struct {
					Day    int `json:"day"`
					Hour   int `json:"hour"`
					Minute int `json:"minute"`
				} `json:"open"`
				Close struct {
					Day    int `json:"day"`
					Hour   int `json:"hour"`
					Minute int `json:"minute"`
				} `json:"close"`
			} `json:"periods"`
			WeekdayDescriptions []string `json:"weekdayDescriptions"`
			SecondaryHoursType  string   `json:"secondaryHoursType"`
		} `json:"regularSecondaryOpeningHours,omitempty"`
		PrimaryType           string `json:"primaryType"`
		ShortFormattedAddress string `json:"shortFormattedAddress"`
		Reviews               []struct {
			Name                           string `json:"name"`
			RelativePublishTimeDescription string `json:"relativePublishTimeDescription"`
			Rating                         int    `json:"rating"`
			Text                           struct {
				Text         string `json:"text"`
				LanguageCode string `json:"languageCode"`
			} `json:"text"`
			OriginalText struct {
				Text         string `json:"text"`
				LanguageCode string `json:"languageCode"`
			} `json:"originalText"`
			AuthorAttribution struct {
				DisplayName string `json:"displayName"`
				URI         string `json:"uri"`
				PhotoURI    string `json:"photoUri"`
			} `json:"authorAttribution"`
			PublishTime time.Time `json:"publishTime"`
		} `json:"reviews"`
		Photos []struct {
			Name               string `json:"name"`
			WidthPx            int    `json:"widthPx"`
			HeightPx           int    `json:"heightPx"`
			AuthorAttributions []struct {
				DisplayName string `json:"displayName"`
				URI         string `json:"uri"`
				PhotoURI    string `json:"photoUri"`
			} `json:"authorAttributions"`
		} `json:"photos"`
		OutdoorSeating  bool `json:"outdoorSeating"`
		LiveMusic       bool `json:"liveMusic"`
		MenuForChildren bool `json:"menuForChildren"`
		ServesCocktails bool `json:"servesCocktails,omitempty"`
		ServesDessert   bool `json:"servesDessert"`
		ServesCoffee    bool `json:"servesCoffee"`
		GoodForChildren bool `json:"goodForChildren"`
		AllowsDogs      bool `json:"allowsDogs,omitempty"`
		Restroom        bool `json:"restroom"`
		GoodForGroups   bool `json:"goodForGroups,omitempty"`
		PaymentOptions  struct {
			AcceptsCashOnly bool `json:"acceptsCashOnly"`
		} `json:"paymentOptions,omitempty"`
		ParkingOptions struct {
			FreeParkingLot    bool `json:"freeParkingLot"`
			FreeStreetParking bool `json:"freeStreetParking"`
		} `json:"parkingOptions,omitempty"`
		AccessibilityOptions struct {
			WheelchairAccessibleParking bool `json:"wheelchairAccessibleParking"`
			WheelchairAccessibleSeating bool `json:"wheelchairAccessibleSeating"`
		} `json:"accessibilityOptions,omitempty"`
		ParkingOptions0 struct {
			FreeParkingLot bool `json:"freeParkingLot"`
		} `json:"parkingOptions,omitempty"`
		PaymentOptions0 struct {
			AcceptsDebitCards bool `json:"acceptsDebitCards"`
			AcceptsCashOnly   bool `json:"acceptsCashOnly"`
			AcceptsNfc        bool `json:"acceptsNfc"`
		} `json:"paymentOptions,omitempty"`
		ParkingOptions1 struct {
			FreeParkingLot bool `json:"freeParkingLot"`
		} `json:"parkingOptions,omitempty"`
		AccessibilityOptions0 struct {
			WheelchairAccessibleParking  bool `json:"wheelchairAccessibleParking"`
			WheelchairAccessibleEntrance bool `json:"wheelchairAccessibleEntrance"`
			WheelchairAccessibleSeating  bool `json:"wheelchairAccessibleSeating"`
		} `json:"accessibilityOptions,omitempty"`
		ParkingOptions2 struct {
			FreeParkingLot bool `json:"freeParkingLot"`
		} `json:"parkingOptions,omitempty"`
		AccessibilityOptions1 struct {
			WheelchairAccessibleParking  bool `json:"wheelchairAccessibleParking"`
			WheelchairAccessibleEntrance bool `json:"wheelchairAccessibleEntrance"`
			WheelchairAccessibleSeating  bool `json:"wheelchairAccessibleSeating"`
		} `json:"accessibilityOptions,omitempty"`
		ParkingOptions3 struct {
			FreeStreetParking bool `json:"freeStreetParking"`
		} `json:"parkingOptions,omitempty"`
	} `json:"places"`
}
