package donneestruct

import (
	"time"
)

// DonneesCapteur represents data sent by captors to MQTT
type DonneesCapteur struct {
	CapteurID  int       `json:"capteur_id"`
	AeroportID string    `json:"aeroport_id"`
	Nature     string    `json:"nature"`
	Valeur     float32   `json:"valeur"`
	Date       time.Time `json:"date"`
}

//Mesure represents a measure made by a captor
type Measure struct {
	CaptorID int       `json:"captor_id"`
	Value    float32   `json:"value"`
	Date     time.Time `json:"date"`
}

// Measures contains a list of measures to be converted in json
type Measures struct {
	Measures []Measure `json:"measures"`
}

//Moyenne represents the average of a ind of measure
type Average struct {
	AverageTemp float32 `json:"averageTemp"`
	AveragePres float32 `json:"averagePres"`
	AverageWind float32 `json:"averageWind"`
}

// Enum for nature of captor
const (
	TEMP = "temperature"
	PRES = "atmospheric_pressure"
	WIND = "wind_speed"
)
