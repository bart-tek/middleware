package captorint

// Captor Interface listant les méthdodes que doivent implémenter les capteurs
type Captor interface {
	generateCapteurID() int
	generateAeroportID() string
	generateValeur() float32
}
