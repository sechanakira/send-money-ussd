package main

import (
	"time"

	"gorm.io/gorm"
)

type UssdSession struct {
	gorm.Model
	Msisdn           string
	SessionId        string
	State            SessionStage
	SessionStartTime time.Time
	SelectedCountry  string
	AmountToSend     float32
}

type MenuText struct {
	gorm.Model
	Text       string
	Language   string
	MenuNumber int
}

type ForeignCurrencyRates struct {
	gorm.Model
	BaseCurrency string
	ZarToKwsRate float32
	ZarToMwkRate float32
}

type SessionStage string

const (
	Complete SessionStage = "COMPLETE"
	Menu1    SessionStage = "MENU_1"
	Menu2    SessionStage = "MENU_2"
	Menu3    SessionStage = "MENU_3"
	Menu4    SessionStage = "MENU_4"
)
