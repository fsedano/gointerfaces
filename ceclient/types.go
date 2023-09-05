package ceclient

// Any type that has a method 'SendCE' that takes no parameters and returns an error will be a valid Ceclient
type Ceclient interface {
	// Sends a CE to some client
	SendCE() error
}
