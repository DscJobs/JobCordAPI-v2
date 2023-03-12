package config

type Config struct {
	Token       string `yaml:"token" comment:"Discord token" validate:"required"`
	PostgresURL string `yaml:"postgres_url" default:"postgresql:///dscjobs" comment:"Postgres URL" validate:"required"`
	RedisURL    string `yaml:"redis_url" default:"redis://localhost:6379/2" comment:"Redis URL" validate:"required"`
	Port        string `yaml:"port" default:":5848" comment:"Port to run the server on" validate:"required"`
	APIUrl      string `yaml:"api_url" default:"https://api.jobcord.co" comment:"URL of the API" validate:"required"`
}
