package config

type Arguments struct {
	Output    string `env:"OUT" envDefault:"out"`
	Config    string `env:"CONFIG" envDefault:"vanitygen.yaml"`
	Formatter string `env:"FORMAT" envDefault:"plain"`
}
