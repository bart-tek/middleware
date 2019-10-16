package donneestruct

import (
	"fmt"
	"time"
)

// DonneesCapteur represente les données reçus depuis MQTT
//
type DonneesCapteur struct {
	CapteurID  int       `json:"capteur_id"`
	AeroportID string    `json:"aeroport_id"`
	Nature     string    `json:"nature"`
	Valeur     float32   `json:"valeur"`
	Date       time.Time `json:"date"`
}

func (d DonneesCapteur) String() string {
	return fmt.Sprintf("{\"capteur_id\": %d"+"\n\"aeroport_id\": \"%s\""+"\n\"nature\": \"%s\""+"\n\"valeur\": %f"+"\n\"date\": %s\n}\n", d.CapteurID, d.AeroportID, d.Nature, d.Valeur, d.Date)
}
