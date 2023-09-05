package mocks

import (
	"errors"
	"log"

	"fsedano.net/test/ceclient"
)

type mock_ceclient struct {
	ShouldReturnError bool
}

// Create a new Ceclient that will fail or not depending on the parameter
func New(s bool) ceclient.Ceclient {
	return &mock_ceclient{
		ShouldReturnError: s,
	}
}

func (c *mock_ceclient) SendCE() error {
	log.Printf("This is the mock. Returning err=%t", c.ShouldReturnError)
	if c.ShouldReturnError {
		return errors.New("a fake error")
	} else {
		return nil
	}
}
