package csv_sub

import (
	"log"
	"net/url"
	client "github.com"
)

func main() {
	uri, err := url.Parse("mqtt://pvpuovcq:h56KR9mXu9Xu@farmer.cloudmqtt.com:15652/")
	if err != nil {
		log.Fatal(err)
	}

	topic := uri.Path[1:len(uri.Path)]
	if topic == "" {
		topic = "test"
	}

	go Listen
}
