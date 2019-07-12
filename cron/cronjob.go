package cron

import (
    "../log"
    "bufio"
    "regexp"
    "strings"
    "os"
    "os/exec"
    "github.com/robfig/cron/v3"
    "../config"
)

type Job struct {
    Interval    string
    Task        string
}

var EventJobs map[string][]Job
var ScheduledJobs []Job
var Cronjob *cron.Cron

func (job *Job) Run(){
    go func(){
        joblabel := job.Interval + " => " + job.Task
        log.Debug(joblabel)
        out, err := exec.Command("sh", "-c", job.Task).Output()
        if err != nil {
            log.Error(joblabel, err)
        }else{
            log.Debug("[OUTPUT]", string(out))
        }
    }()
}

func FindAllSubString(pattern string, line string) []string {
    re := regexp.MustCompile(pattern)
    parsed := re.FindAllStringSubmatch(line,-1)[0]
    return parsed
}

func Parse(){
    EventJobs = make(map[string][]Job)
    file, err :=  os.Open(config.Config.CronFilePath)
    log.Info("Reading Jobs from " + config.Config.CronFilePath)
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.Trim(scanner.Text(), " ")
        if len(line) > 2 && string(line[0:1]) != "#" {
            if string(line[0:2]) == "@@"{
                parsed := FindAllSubString(`@@(\S+)\s(.*)$`, line)
                event := parsed[1]
                task  := parsed[2]
                EventJobs[event] = append(EventJobs[event], Job{event, task})
            }else if string(line[0:1]) == "@"{
                parsed := FindAllSubString(`@(\S+)\s(.*)$`, line)
                event := parsed[1]
                task  := parsed[2]
                if event == "every" {
                    parsed = FindAllSubString(`@(\S+)\s(\S+)\s(.*)$`, line)
                    event = event + " "+ parsed[2]
                    task  = parsed[3]
                }
                ScheduledJobs = append(ScheduledJobs, Job{"@"+event,task})
            } else {
                parsed := FindAllSubString(`(([0-9\*,\-]+\s){5})(.*)$`, line)
                if len(parsed) > 0 {
                    ScheduledJobs = append(ScheduledJobs, Job{parsed[1], parsed[3]})
                }
            }
        }
    }
    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    ScheduleJob()
}

func ScheduleJob(){
    Cronjob = cron.New()
    for i := 0; i < len(ScheduledJobs); i++ {
        job := ScheduledJobs[i]
        Cronjob.AddFunc(job.Interval, func(){job.Run()})
    }
    Cronjob.Start()
}

func RunEventJob(jobs []Job) {
    for i := 0; i < len(jobs); i++ {
        jobs[i].Run()
    }
}
