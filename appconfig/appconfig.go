package appconfig

//AppConfig is a struct for configs
type AppConfig struct {
	Freshconn      string
	Seq_Server     string
	Seq_Server_Key string
}

var _configs AppConfig

// SetConfig sets configurations to conf
func SetConfig(conf AppConfig) {
	_configs = conf
}

//GetConfig retuns the configs
func GetConfig() AppConfig {
	return _configs
}
