package config

type Config struct {
	Token        string                 `yaml:"token" comment:"Discord token" validate:"required"`
	PostgresURL  string                 `yaml:"postgres_url" default:"postgresql:///dscjobs" comment:"Postgres URL" validate:"required"`
	RedisURL     string                 `yaml:"redis_url" default:"redis://localhost:6379/2" comment:"Redis URL" validate:"required"`
	Port         string                 `yaml:"port" default:":5848" comment:"Port to run the server on" validate:"required"`
	APIUrl       string                 `yaml:"api_url" default:"https://api.jobcord.co" comment:"URL of the API" validate:"required"`
	LoginClients map[string]LoginClient `yaml:"login_clients" comment:"Login clients" validate:"required"`
	MainServer   string                 `yaml:"main_server" comment:"Main server ID" validate:"required"`
}

type LoginClient struct {
	ClientID     string `yaml:"client_id" comment:"Discord client ID" validate:"required"`
	ClientSecret string `yaml:"client_secret" comment:"Discord client secret" validate:"required"`
	RedirectURL  string `yaml:"redirect_url" comment:"Redirect URL" validate:"required"`
	Note         string `yaml:"note" comment:"Note to add to the client" validate:"required"`
}
