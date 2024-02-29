package dto

import "time"

type T struct {
	Id              string    `json:"id"`
	Href            string    `json:"href"`
	Name            string    `json:"name"`
	Description     string    `json:"description"`
	Version         string    `json:"version"`
	ValidFor        ValidFor  `json:"validFor"`
	LifecycleStatus string    `json:"lifecycleStatus"`
	LastUpdate      time.Time `json:"lastUpdate"`
	IsRoot          bool      `json:"isRoot"`
	SubCategory     []struct {
		Href         string `json:"href"`
		Id           string `json:"id"`
		Name         string `json:"name"`
		Version      string `json:"version"`
		ReferredType string `json:"@referredType"`
		Type         string `json:"@type"`
	} `json:"subCategory"`
	ProductOffering []struct {
		Id           string `json:"id"`
		Href         string `json:"href"`
		Name         string `json:"name"`
		ReferredType string `json:"@referredType"`
		Type         string `json:"@type"`
	} `json:"productOffering"`
	Type   string `json:"@type"`
	Parent struct {
		Href         string `json:"href"`
		Id           string `json:"id"`
		Name         string `json:"name"`
		Version      string `json:"version"`
		ReferredType string `json:"@referredType"`
		Type         string `json:"@type"`
	} `json:"parent,omitempty"`
}

type ValidFor struct {
	StartDateTime time.Time `json:"startDateTime"`
	EndDateTime   time.Time `json:"endDateTime,omitempty"`
}
