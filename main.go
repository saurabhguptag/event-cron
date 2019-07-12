package main

import (
    "flag"
    "github.com/saurabhguptag/event-cron/log"
    "github.com/saurabhguptag/event-cron/config"
    "github.com/saurabhguptag/event-cron/cron"
    "github.com/saurabhguptag/event-cron/server"
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
