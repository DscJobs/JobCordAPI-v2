package state

import (
	"context"
	"os"

	"jobcord/config"

	"github.com/bwmarrin/discordgo"
	"github.com/go-playground/validator/v10"
	"github.com/infinitybotlist/eureka/dovewing"
	"github.com/infinitybotlist/eureka/genconfig"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/yaml.v3"
)

var (
	Pool      *pgxpool.Pool
	Redis     *redis.Client
	Discord   *discordgo.Session
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

	Pool, err = pgxpool.New(Context, Config.PostgresURL)

	if err != nil {
		panic(err)
	}

	rOptions, err := redis.ParseURL(Config.RedisURL)

	if err != nil {
		panic(err)
	}

	Redis = redis.NewClient(rOptions)

	Discord, err = discordgo.New("Bot " + Config.Token)

	if err != nil {
		panic(err)
	}

	Discord.Identify.Intents = discordgo.IntentsGuilds | discordgo.IntentGuildPresences | discordgo.IntentsGuildMembers

	go func() {
		err = Discord.Open()
		if err != nil {
			panic(err)
		}

		err = Discord.UpdateWatchStatus(0, Config.MainServer)

		if err != nil {
			panic(err)
		}
	}()

	w := zapcore.AddSync(os.Stdout)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.DebugLevel,
	)

	Logger = zap.New(core).Sugar()

	// Load dovewing state
	dovewing.SetState(&dovewing.State{
		Discord:        Discord,
		Pool:           Pool,
		Logger:         Logger,
		PreferredGuild: Config.MainServer,
		Context:        Context,
		Redis:          Redis,
	})
}
