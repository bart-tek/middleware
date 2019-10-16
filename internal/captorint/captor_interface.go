package captorint

type Captor interface {
	generateCapteurID() int
	generateAeroportID() string
	generateValeur() float32
}
