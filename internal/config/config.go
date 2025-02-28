package config

const (
	dir      = "../../static"
	host     = "localhost"
	port     = "5432"
	dbname   = "bookstore"
	user     = "postgres"
	password = "1111"
	sslmode  = "disable"
)

type Config struct {
	Dir string

	BD struct {
		Host     string
		Port     string
		Username string
		Password string
		BDName   string
		SSLMode  string
	}
}

func NewConfig() *Config {
	return &Config{
		Dir: dir,
		BD: struct {
			Host     string
			Port     string
			Username string
			Password string
			BDName   string
			SSLMode  string
		}{
			Host:     host,
			Port:     port,
			Username: user,
			Password: password,
			BDName:   dbname,
			SSLMode:  sslmode,
		},
	}
}
