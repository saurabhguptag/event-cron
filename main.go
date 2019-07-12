package main

import (
    "flag"
    "./log"
    "./config"
    "./cron"
    "./server"
)

func main(){
    ConfigPath := flag.String("config", "config/config.json", "config path")
    log.Info("config:", *ConfigPath)
    flag.Parse()
    config.LoadConfig(*ConfigPath)
    log.Init(config.Config.Loglevel)
    cron.Parse()
    server.Start()
}
