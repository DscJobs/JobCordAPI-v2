package state

import (
	"context"
	"os"

	"jobcord/config"

	"github.com/go-playground/validator"
	"github.com/infinitybotlist/genconfig"
	"github.com/jackc/pgx/v5/pgxpool"
	"gopkg.in/yaml.v3"
)

var (
	Pool      *pgxpool.Pool
	Context   = context.Background()
	Validator = validator.New()
	Config    *config.Config
)

func Setup() {
	genconfig.GenConfig(config.Config{})

	cfg, err := os.ReadFile("config.yaml")

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(cfg, &Config)

	if err != nil {
		panic(err)
	}

	err = Validator.Struct(Config)

	if err != nil {
		panic("configError: " + err.Error())
	}

	p, err := pgxpool.New(Context, Config.PostgresURL)

	if err != nil {
		panic(err)
	}

	Pool = p
}
