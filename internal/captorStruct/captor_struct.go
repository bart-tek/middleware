package captorStruct

import "time"

type DonneesCapteur struct {
	idCapteur  int
	idAeroport string
	nature     string
	valeur     float32
	date       time.Time
}
