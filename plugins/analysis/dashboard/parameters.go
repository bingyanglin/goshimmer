package dashboard

import "github.com/iotaledger/hive.go/configuration"

type ParametersDefinition struct {
	// BindAddress defines the analysis dashboard binding address.
	BindAddress string `default:"0.0.0.0:8000" usage:"the bind address of the analysis dashboard"`
	// Dev defines the analysis dashboard dev mode.
	Dev bool `default:"false" usage:""whether the analysis dashboard runs in dev mode""`
	// BasicAuthEnabled defines  the analysis dashboard basic auth enabler.
	BasicAuthEnabled bool `default:"false" usage:"whether to enable HTTP basic auth"`
	// BasicAuthUsername defines the analysis dashboard basic auth username.
	BasicAuthUsername string `default:"goshimmer" usage:"HTTP basic auth username"`
	// BasicAuthPassword defines the analysis dashboard basic auth password.
	BasicAuthPassword string `default:"goshimmer" usage:"HTTP basic auth password"`
	// MongoDBEnabled defines the analysis dashboard to enable mongoDB.
	MongoDBEnabled bool `default:"false" usage:"whether to enable MongoDB"`
	// MongoDBUsername defines the analysis dashboard mongoDB username.
	MongoDBUsername string `default:"root" usage:"MongoDB username"`
	// MongoDBPassword defines the analysis dashboard mongoDB password.
	MongoDBPassword string `default:"password" usage:"MongoDB username"`
	// MongoDBHostAddress defines the analysis dashboard mongoDB binding address.
	MongoDBHostAddress string `default:"mongodb:27017" usage:"MongoDB host address"`
	// ManaDashboardAddress defines the mana dashboard address to stream mana info from.
	ManaDashboardAddress string `default:"http://127.0.0.1:8081" usage:"dashboard host address"`
}

// Parameters contains the configuration parameters of the logger plugin.
var Parameters = &ParametersDefinition{}

func init() {
	configuration.BindParameters(Parameters, "analysis.dashboard")
}
