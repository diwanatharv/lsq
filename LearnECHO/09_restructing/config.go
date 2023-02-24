package tronics

// based on the environment variables
// first know what is  env?
//this will not work automotically u need to install some package

type Configdatabase struct {
	AppName  string `env:"APP_NAME" env-default:"TRONICS"`
	AppEnv   string `env:"APP_ENV" env-default:"DEV"`
	Port     string `env:"MY_APP_PORT" env-default:"7500"`
	Host     string `env:"HOST" env-default:"localhost"`
	Loglevel string `env:"LOG_LEVEL" env-default:"ERROR"`
}

var cfg Configdatabase
