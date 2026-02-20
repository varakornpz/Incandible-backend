package providers

import (
	"os"
	"reflect"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)


type Config struct {
	JWTSecret		string	`validate:"required"`
	DBDsn			string	`validate:"required"`
	GGClientSecret	string	`validate:"required"`
	GGClientID		string	`validate:"required"`
	GGRedirectUrl	string	`validate:"required"`
	FrontendErrPage		string	`validate:"required"`
}

var AppConf	*Config

func InitAppConf(){
	dotEnvErr := godotenv.Load()
    if dotEnvErr != nil {
		log.Info().Msg("No .env file found, using system environment variables")
    }

	AppConf = &Config{
		JWTSecret: os.Getenv("JWT_SECRET"),
		DBDsn: os.Getenv("DB_DSN"),
		GGClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		GGClientID: os.Getenv("GOOGLE_CLIENT_ID"),
		GGRedirectUrl: os.Getenv("GOOGLE_REDIRECT_URL"),
		FrontendErrPage: os.Getenv("FE_ERROR_PAGE"),
	}

	v := reflect.ValueOf(*AppConf)
	t := v.Type()
	for i := 0 ; i < v.NumField() ; i++ {
		field := t.Field(i)
		val	:= v.Field(i).Interface()

		validateTag	:= field.Tag.Get("validate")
		
		if validateTag == "required" {
			valStringed	, ok := val.(string)
			if !ok {
				log.Fatal().Msg("Cant convert intterface to string called by providers/config.go")
			}
			if valStringed == ""{
				log.Fatal().Msgf("Missing required ENV: %s called by providers/config.go", field.Name)
			}
		}
	}

	log.Info().Msg("✅ Config loaded.")
}