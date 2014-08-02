package config

// Typically, you read from the environment variable.
// Can't do that on AppEngine :(
var ENV = "dev"

var BaseURL = ""

func init() {
	switch ENV {
	case "dev":
		BaseURL = "http://localhost:8080"
	case "qa":
	case "prod":
		// Insert your app's production url.
		BaseURL = "http://testgaeappjcsiii.appspot.com/"
	default:
		BaseURL = ""
	}
}
