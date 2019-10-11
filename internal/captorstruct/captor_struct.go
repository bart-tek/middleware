package captorstruct

import "time"

type DonneesCapteur struct {
	IDCapteur  int
	IDAeroport string
	Nature     string
	Valeur     float32
	Date       time.Time
}
