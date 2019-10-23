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
type Mesure struct {
	CapteurID int       `json:"capteur_id"`
	Valeur    float32   `json:"valeur"`
	Date      time.Time `json:"date"`
}

//Moyenne represents the average of a ind of measure
type Moyenne struct {
	Nature string  `json:"nature"`
	Valeur float32 `json:"valeur"`
}

//MonTest oisdsdqsd
type MonTest struct {
	Nature   string `json:"nature"`
	Aeroport string `json:"aeroport"`
}

// Enum for nature of captor
const (
	TEMP = "temperature"
	PRES = "atmospheric_pressure"
	WIND = "wind_speed"
)
