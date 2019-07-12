package log

import(
    "fmt"
    "time"
    "strings"
    "os"
)

var LOGLEVEL = map[string]int{"debug" : 0,"info" : 1,"warn": 2,"error": 3,"fatal": 4}
var Level int

func Init(level string){
    Level = GetIntLogLevel(level)
}

func GetIntLogLevel (level string) int {
    return LOGLEVEL[strings.ToLower(level)]
}
func GetTime() string {
    return  time.Now().Format("2006-01-02 15:04:05")
}

func ShowMessage(level string, message interface{}){
    messageLevel :=  GetIntLogLevel(level)
    if Level <= messageLevel {
        fmt.Printf("%s [%s] %v\n",GetTime(), strings.ToUpper(level), message)
    }
    switch messageLevel  {
    case 4:
        os.Exit(1)
        break;
    }
}

func Debug(text ...interface{}){
    ShowMessage("DEBUG",text);
}

func Info(text ...interface{}){
    ShowMessage("INFO",text);
}

func Warn(text ...interface{}){
    ShowMessage("WARN",text);
}

func Error(text ...interface{}){
    ShowMessage("ERROR", text);
}

func Fatal(text ...interface{}){
    ShowMessage("FATAL",text);
}
