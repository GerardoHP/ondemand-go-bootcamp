package config

// Represents the configuration needed for the application to start
type config struct {
	Port            string
	StorageFileName string
}

var singleInstance *config

// Gets a singleton of the configuration for the service
func New() *config {
	if singleInstance == nil {
		singleInstance = &config{
			Port:            "8080",
			StorageFileName: "pokemons.csv"}
	}

	return singleInstance
}
