package rest_utils

type RestConfig struct {
	Address string `env:"ADDRESS,default=0.0.0.0"`
	Port    string `env:"PORT,default=8080"`
}
