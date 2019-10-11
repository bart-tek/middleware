package main

import (
	"math/rand"
	"middleware/internal/captorstruct"
	"time"
)

func main() {

}

func getDonnees() captorstruct.DonneesCapteur {

	return captorstruct.DonneesCapteur{
		IDCapteur:  1,
		IDAeroport: "",
		Nature:     "Atmospheric pressure",
		Valeur:     2.,
		Date:       time.Now()}

}

func generatePressureValue() float32 {

	return rand.Float32()

}
