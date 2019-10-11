package donneestruct

import (
	"fmt"
	"time"
)

type DonneesCapteur struct {
	IDCapteur  int
	IDAeroport string
	Nature     string
	Valeur     float32
	Date       time.Time
}

func (d DonneesCapteur) String() string {
	return fmt.Sprintf("{\nIDCapteur: %d"+"\nIDAeroport: %s"+"\nNature: %s"+"\nValeur: %f"+"\nDate: %s\n}\n", d.IDCapteur, d.IDAeroport, d.Nature, d.Valeur, d.Date.Format("2006-01-02 15:04:05"))
}
