package donneestruct

import (
	"fmt"
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

func (d DonneesCapteur) String() string {
	return fmt.Sprintf("{CapteurID: %d"+"AeroportID: %s"+"\nNature: %s"+"\nValeur: %f"+"\nDate: %s\n}\n", d.CapteurID, d.AeroportID, d.Nature, d.Valeur, d.Date.Format("2006-01-02 15:04:05"))
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
