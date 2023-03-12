package state

import (
	"context"
	"os"

	"jobcord/config"

	"github.com/go-playground/validator/v10"
	"github.com/infinitybotlist/genconfig"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

var (
	Pool      *pgxpool.Pool
	Context   = context.Background()
	Validator = validator.New()
	Config    *config.Config
	Logger    *zap.SugaredLogger
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

	w := zapcore.AddSync(os.Stdout)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.DebugLevel,
	)

	Logger = zap.New(core).Sugar()
}
