package main

import (
	"log"

	"fsedano.net/test/ceclient"
	"fsedano.net/test/mocks"
)

func main() {

	// Lets create a real cloud event client
	c := ceclient.New("http://broker.com")

	// This is a mock that will always fail
	m_always_fails := mocks.New(true)

	// This is a mock that will always succeed
	m_always_success := mocks.New(false)

	// And let's send a few cloud events. All of them comply to 'Ceclient' interface,
	// so all of them are valid Ceclients
	//
	// The behaviour of each of them will be different, though
	doSendCE(c)
	doSendCE(m_always_fails)
	doSendCE(m_always_success)

}

// Send a cloud event and print if an error happened
func doSendCE(c ceclient.Ceclient) error {
	err := c.SendCE()
	if err != nil {
		log.Printf("OOPS an error happened: %s", err)
	} else {
		log.Printf("* Success *")
	}
	return err
}
