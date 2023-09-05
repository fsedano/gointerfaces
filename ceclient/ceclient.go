package ceclient

import "log"

type ceclient struct {
	URL string
}

// Create a new Ceclient that will send request to the URL in the parameter
func New(s string) Ceclient {
	return &ceclient{
		URL: s,
	}
}

// Sends a real cloud event to a very real broker
func (c *ceclient) SendCE() error {
	log.Printf("Sending CE to URL:%s", c.URL)
	return nil
}
