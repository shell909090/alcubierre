package main

import (
	"encoding/json"
	"flag"
	stdlog "log"
	"os"

	logging "github.com/op/go-logging"
)

var logger = logging.MustGetLogger("")

var (
	ConfigFile string
)

type Config struct {
	Mode       string
	Logfile    string
	Loglevel   string
	Listen     string
	Upstream   string
	Forwarders []string
}

func init() {
	flag.StringVar(&ConfigFile, "config", "/etc/alcubierre/config.json", "config file")
	flag.Parse()
}

func LoadJson(configfile string, cfg interface{}) (err error) {
	file, err := os.Open(configfile)
	if err != nil {
		return
	}
	defer file.Close()

	dec := json.NewDecoder(file)
	err = dec.Decode(&cfg)
	return
}

func LoadConfig() (cfg *Config, err error) {
	cfg = &Config{}
	err = LoadJson(ConfigFile, cfg)
	if err != nil {
		return
	}
	return
}

func SetLogging(cfg *Config) (err error) {
	var file *os.File
	file = os.Stdout

	if cfg.Logfile != "" {
		file, err = os.OpenFile(cfg.Logfile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
		if err != nil {
			logger.Fatal(err)
		}
	}
	logBackend := logging.NewLogBackend(file, "",
		stdlog.LstdFlags|stdlog.Lmicroseconds|stdlog.Lshortfile)
	logging.SetBackend(logBackend)

	logging.SetFormatter(logging.MustStringFormatter("%{level}: %{message}"))

	lv, err := logging.LogLevel(cfg.Loglevel)
	if err != nil {
		panic(err.Error())
	}
	logging.SetLevel(lv, "")

	return
}

// func main() {
// 	cfg, err := LoadConfig()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	err = SetLogging(cfg)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}

// 	switch cfg.Mode {
// 	case "server":
// 		logger.Notice("server start.")

// 		err = RunServer(cfg)
// 	case "client":
// 		logger.Notice("client start.")
// 	default:
// 		logger.Info("unknown mode")
// 		return
// 	}
// 	if err != nil {
// 		logger.Error("%s", err)
// 	}
// 	logger.Info("server stopped")
// }
