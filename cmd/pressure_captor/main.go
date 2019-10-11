package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Evrard-Nil/middleware/internal/confcaptorstruct"
	"github.com/Evrard-Nil/middleware/internal/donneestruct"
	"github.com/Evrard-Nil/middleware/internal/enumnature"
)

func main() {
	fmt.Println(getDonnees())
	fmt.Println(getDonnees())
	fmt.Println(getDonnees())
	fmt.Println(getDonnees())

}

func getDonnees() donneestruct.DonneesCapteur {

	return donneestruct.DonneesCapteur{
		IDCapteur:  generateIDCapteur(),
		IDAeroport: generateIDAeroport(),
		Nature:     enumnature.PRES,
		Valeur:     generateValeur(),
		Date:       time.Now(),
	}

}

func generateValeur() float32 {

	var min float32 = 950
	var max float32 = 1050
	return min + rand.Float32()*(max-min)

}

func generateIDCapteur() int {
	min := 1
	max := 5
	return rand.Intn(max-min) + min
}

func generateIDAeroport() string {
	min := 0
	max := 14

	var c confcaptorstruct.ConfCaptorStruct
	c.GetConf()

	return c.ListeAeroport[rand.Intn(max-min)+min]
}
