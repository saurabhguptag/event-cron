package config

import (
    "io/ioutil"
    "encoding/json"
)

type Configuration struct {
    Name string     `json:"name"`
	Server struct {
        Host string `json:"host"`
        Port string `json:"port"`
    } `json:"server"`
    Loglevel string `json:"loglevel"`
    CronFilePath string `json:"cronFilePath"`
}

var Config Configuration
func LoadConfig(ConfigPath string){
    json_file, _ := ioutil.ReadFile(ConfigPath)
    json.Unmarshal(json_file, &Config)

}
