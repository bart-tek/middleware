package captorint

type Captor interface {
	generateIDCapteur() int
	generateIDAeroport() string
	generateValeur() float32
}
