package config

import (
   "encoding/json"
   "os"
   "log"
   "../system"
)

// description
// Config 
type Config struct {
    Protocal           string `json:"protocal"` 
    Host               string `json:"host"`
    Port               uint16 `json:"port"`
    Service            string `json:"service"`
    Strategy           string `json:"strategy"`
    RequestQueueSize   int `json:"requestqueuesize"` 
    MaxProcessor       int `json:"maxprocessor"`
    Hearbeat           int    `json:"heartbeat"`
    Keepalive          int    `json:"keepalive"` 
    Backends           []system.Backend `json:"backends"`
}

// description
// load config file
func Load(filename string) (*Config, error) {
   var config Config
   file, err := os.Open(filename)
   if err != nil {
       log.Println("load config error:", err)
   } else {
       buff := make([]byte , 1024)
       end, _ := file.Read(buff)
       err = json.Unmarshal(buff[:end], &config)
       if err != nil {
           log.Println("decode json error:", err)
       }
   }
   log.Println("success load config file:" , filename)
   return &config, err
}
