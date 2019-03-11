package main

import (
	"fmt"
	"github.com/nats-io/go-nats"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"os"
	"time"
)

type NatsyConfig struct {
	Url     string
	Subject string
	Message string
	Request bool
	Timeout time.Duration
}

func main() {

	// Config
	var cfg NatsyConfig
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("No config file found")
	}

	pflag.String("url", "", "nats url")
	pflag.String("subject", "", "nats subject")
	pflag.String("message", "", "nats message")
	pflag.Bool("request", false, "nats request")
	pflag.Duration("timeout", time.Second, "nats timeout (request only)")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)
	err = viper.Unmarshal(&cfg)
	exitIfErrWithConfig(err, cfg)
	fmt.Printf("%s - %s > %s\n", cfg.Url, cfg.Subject, cfg.Message)


	// Connect Options.
	opts := []nats.Option{nats.Name("natsy client")}

	// Connect to NATS
	nc, err := nats.Connect(cfg.Url, opts...)
	exitIfErrWithConfig(err, cfg)
	defer nc.Close()


	if cfg.Request {
		fmt.Printf("%s - %s < ", cfg.Url, cfg.Subject)
		reply, err := nc.Request(cfg.Subject, []byte(cfg.Message), cfg.Timeout)
		exitIfErrWithConfig(err, cfg)
		fmt.Printf("%s\n", string(reply.Data))
		return
	}

	err = nc.Publish(cfg.Subject, []byte(cfg.Message))
	exitIfErrWithConfig(err, cfg)
	fmt.Printf("%s - %s < published \n", cfg.Url, cfg.Subject)
}

func exitIfErrWithConfig(err error, cfg NatsyConfig) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s - %s < err: %v\n", cfg.Url, cfg.Subject, err)
		os.Exit(1)
	}
}