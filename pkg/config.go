package pkg

import (
	"fmt"
	"net/url"

	"github.com/knadh/koanf/parsers/toml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
)

type Config struct {
	LogEnv      string `koanf:"env"`
	RabbitMQURL string `koanf:"rabbitmq_url"`
	WoC         struct {
		WebhookURL string `koanf:"webhook_url"`
		QueueName  string `koanf:"queue_name"`
	} `koanf:"woc"`
	AIVerse struct {
		WebhookURL string `koanf:"webhook_url"`
		QueueName  string `koanf:"queue_name"`
	} `koanf:"aiverse"`
}

var AppConfig *Config

func LoadConfig() error {
	var koa = koanf.New(".")

	if err := koa.Load(file.Provider("env.toml"), toml.Parser()); err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	var config Config
	if err := koa.Unmarshal("", &config); err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}

	if err := validateURL(config.RabbitMQURL); err != nil {
		return fmt.Errorf("invalid RabbitMQ URL: %w", err)
	}
	if err := validateURL(config.WoC.WebhookURL); err != nil {
		return fmt.Errorf("invalid WoC webhook URL: %w", err)
	}
	if err := validateURL(config.AIVerse.WebhookURL); err != nil {
		return fmt.Errorf("invalid AIVerse webhook URL: %w", err)
	}

	AppConfig = &config
	return nil
}

func validateURL(rawURL string) error {
	_, err := url.ParseRequestURI(rawURL)
	if err != nil {
		return err
	}
	return nil
}
