package main

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// PropertyFacingDirection
type PropertyFacingDirection string

const PFDNorth PropertyFacingDirection = "north"
const PFDNorthEast PropertyFacingDirection = "north-east"
const PFDEast PropertyFacingDirection = "east"
const PFDSouthEast PropertyFacingDirection = "south-east"
const PFDSouth PropertyFacingDirection = "south"
const PFDSouthWest PropertyFacingDirection = "south-west"
const PFDWest PropertyFacingDirection = "west"
const PFDNorthWest PropertyFacingDirection = "north-west"

// RentalPeriodUnit
type RentalPeriodUnit string

const RPUMonths RentalPeriodUnit = "months"
const RPUYears RentalPeriodUnit = "years"
const RPUDays RentalPeriodUnit = "days"

// PricingUnit
type PricingUnit string

const PUVietnamDong PricingUnit = "VND"
const PUUSD PricingUnit = "USD"
const PUEuro PricingUnit = "EURO"

type TranslatableText struct {
	Language string `bson:"language" json:"language"`
	Text     string `bson:"text" json:"text"`
}

type PropertyCategory struct {
	ID   string             `bson:"_id" json:"id"`
	Name []TranslatableText `bson:"name" json:"name"`

	CAt time.Time `bson:"c_at" json:"c_at"`
	UAt time.Time `bson:"u_at" json:"u_at"`
}

type Property struct {
	ID     bson.ObjectId      `bson:"_id" json:"id"`
	Name   []TranslatableText `bson:"name" json:"name"`
	Images struct {
		ThumbnailURL string `bson:"thumbnail_url" json:"thumbnail_url"`
		MainURL      string `bson:"main_url" json:"main_url"`
	} `bson:"images" json:"images"`
	Desc []TranslatableText `bson:"desc" json:"desc"`

	Type           string    `bson:"type" json:"type"` // PropertyCategory.ID
	AvailableUntil time.Time `bson:"available_until" json:"available_until"`
	Size           struct {
		Width float32 `bson:"width" json:"width"`
		Depth float32 `bson:"depth" json:"depth"`
	} `bson:"size" json:"size"`
	Address struct {
		Name        []TranslatableText `bson:"name" json:"name"`
		Coordinates struct {
			lat float64 `bson:"lat" json:"lat"`
			lon float64 `bson:"lon" json:"lon"`
		} `bson:"coordinates" json:"coordinates"`
	} `bson:"address" json:"address"`
	BedCount        int                     `bson:"bed_count" json:"bed_count"`
	FacingDirection PropertyFacingDirection `bson:"facing_direction" json:"facing_direction"`

	RentalPeriod struct {
		Digits float32          `bson:"digits" json:"digits"`
		Unit   RentalPeriodUnit `bson:"unit" json:"unit"`
	} `bson:"rental_period" json:"rental_period"`
	Price struct {
		Digits float32     `bson:"digits" json:"digits"`
		Unit   PricingUnit `bson:"unit" json:"unit"`
	} `bson:"price" json:"price"`

	CAt time.Time `bson:"c_at" json:"c_at"`
	UAt time.Time `bson:"u_at" json:"u_at"`
}
