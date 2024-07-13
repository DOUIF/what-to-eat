package models

type Center struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

type Circle struct {
	Center Center  `json:"center"`
	Radius float64 `json:"radius"`
}

type LocationRestriction struct {
	Circle Circle `json:"circle"`
}

type PlaceRequest struct {
	IncludedTypes       []string            `json:"includedTypes"`
	LanguageCode        string              `json:"languageCode"`
	RegionCode          string              `json:"regionCode"`
	MaxResultCount      int                 `json:"maxResultCount"`
	RankPreference      string              `json:"rankPreference"`
	LocationRestriction LocationRestriction `json:"locationRestriction"`
}
