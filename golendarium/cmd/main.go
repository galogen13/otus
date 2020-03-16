package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"gopkg.in/yaml.v2"

	"go.uber.org/zap"
)

func main() {

	var config string
	flag.StringVar(&config, "config", "", "Config file address")
	flag.Parse()

	var c Config
	yamlFile, err := ioutil.ReadFile(config)
	if err != nil {
		log.Fatal(err)
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	//c := Config{"127.0.0.1", "9090", "./", "info"}

	logCfg := zap.NewDevelopmentConfig()
	logCfg.OutputPaths = []string{c.LogFile}
	logCfg.Encoding = "json"
	logger, err := logCfg.Build()
	if err != nil {
		log.Fatal(err)
	}
	defer logger.Sync()

	loggerFunc := func(msg string, fields ...zap.Field) {
		switch c.LogLevel {
		case "warn":
			logger.Warn(msg, fields...)
		case "error":
			logger.Error(msg, fields...)
		case "info":
			logger.Info(msg, fields...)
		case "debug":
			logger.Debug(msg, fields...)
		default:
			logger.Info(msg, fields...)
		}
		logger.Sync()
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello")
		loggerFunc("Hello comes")

	})
	http.ListenAndServe(c.HttpListenIp+":"+c.HttpListenPort, nil)
}

// Config .
type Config struct {
	HttpListenIp   string `yaml:"listen_ip"`
	HttpListenPort string `yaml:"listen_port"`
	LogFile        string `yaml:"log_file"`
	LogLevel       string `yaml:"log_level"`
}
