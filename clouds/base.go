package clouds


type Cloud struct {
	// Name of the cloud
	Name string
	// Client is the client for the cloud
	Client interface{}
}


// Cloud is the interface that all clouds must implement
type Client interface {
	init(appid string,appkey string) error
	// Get the name of the cloud
	close() error
}
