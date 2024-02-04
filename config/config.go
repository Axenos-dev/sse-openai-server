package config

import "github.com/kelseyhightower/envconfig"

var Config config

type config struct {
	OpenAI OpenAIConfig
	Port   string `envconfig:"PORT" required:"true"`
}

type OpenAIConfig struct {
	ApiKey    string `envconfig:"OPEN_AI_API_KEY" required:"true"`
	MaxTokens int    `envconfig:"MAX_TOKENS"`
}

func FillConfig() error {
	if err := envconfig.Process("", &Config); err != nil {
		return err
	}

	if Config.OpenAI.MaxTokens == 0 {
		Config.OpenAI.MaxTokens = 20
	}

	return nil
}
