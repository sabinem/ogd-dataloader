package main

import (
    "os"
    "net/http"
    "io"
    "log"
    "log/slog" 
    "gopkg.in/yaml.v2"
//    "errors"
)

const CONFIG_PATH = "config.yml"

type Config struct {
    FilePath string `yaml:"file_path"`
    Url string `yaml:"download_url"`
}

func ReadConfig() (Config, error) {
    var config Config
    // Open YAML file
    file, err := os.Open(CONFIG_PATH)
    if err != nil {
        log.Println(err.Error())
    }
    defer file.Close()

    // Decode YAML file to struct
    if file != nil {
        decoder := yaml.NewDecoder(file)
        if err := decoder.Decode(&config); err != nil {
            return config, err
        }
    }
    return config, nil
}

func dataloader(filepath string, url string) (err error) {
  // Create the file
  out, err := os.Create(filepath)
  if err != nil  {
    return err
  }
  defer out.Close()
  // Get the data
  resp, err := http.Get(url)
  if err != nil {
    return err
  }
  defer resp.Body.Close()
  // Writer the body to file
  _, err = io.Copy(out, resp.Body)
  if err != nil  {
    return err
  }

  return nil
}

func main() {
  config, err := ReadConfig()
  if (err != nil) {
    slog.Error("An error while reading the configuration", "msg", err)
    return
  }
  err =  dataloader(config.FilePath, config.Url)
  if (err != nil) {
    slog.Error("An error occured while downloading", "msg", err)
    return
  }
  slog.Info("Download successful from", "config", config)
}