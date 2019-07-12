package server

import (
    "net/http"
    "github.com/saurabhguptag/event-cron/config"
    "github.com/saurabhguptag/event-cron/log"
    "github.com/saurabhguptag/event-cron/cron"
)

func handler(w http.ResponseWriter, r *http.Request) {
    keys, ok := r.URL.Query()["event"]
    var key string
    responseText := "OK"
    if !ok || len(keys[0]) < 1 {
       responseText = "Event is missing in Request"
    }else{
        key = keys[0]
        if cron.EventJobs[key] == nil {
            responseText = "Event [" + key + "] not found"
        }else{
            go cron.RunEventJob(cron.EventJobs[key])
        }
    }
    if responseText != "OK" {
        http.Error(w, responseText, 404)
        log.Error(responseText)
    }else {
        log.Debug("Event [" + key + "] Executed")
        w.Write([]byte(responseText))
    }
}

func Start(){
    http.HandleFunc("/run-cron", handler)
    log.Info("Starting " + config.Config.Name + " on " + config.Config.Server.Host + " at port " + config.Config.Server.Port)
    log.Fatal(http.ListenAndServe(config.Config.Server.Host+":"+config.Config.Server.Port, nil))
}
