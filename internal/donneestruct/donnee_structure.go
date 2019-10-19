package donneestruct

import (
	"time"
)

// DonneesCapteur represente les données des capteurs reçus depuis MQTT
//
type DonneesCapteur struct {
	CapteurID  int       `json:"capteur_id"`
	AeroportID string    `json:"aeroport_id"`
	Nature     string    `json:"nature"`
	Valeur     float32   `json:"valeur"`
	Date       time.Time `json:"date"`
}

type Mesure struct {
	CapteurID int       `json:"capteur_id"`
	Valeur    float32   `json:"valeur"`
	Date      time.Time `json:"date"`
}

type Moyenne struct {
	Nature string  `json:"nature"`
	Valeur float32 `json:"valeur"`
}

type MonTest struct {
	Nature   string `json:"nature"`
	Aeroport string `json:"aeroport"`
}
