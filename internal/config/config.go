package config

const (
	dir      = "../../static"
	host     = "db"
	port     = "5432"
	dbname   = "bookstore"
	user     = "postgres"
	password = "1111"
	sslmode  = "disable"
	key      = "key"
)

type Config struct {
	Dir string
	Key string
	BD  struct {
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
		Key: key,
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
