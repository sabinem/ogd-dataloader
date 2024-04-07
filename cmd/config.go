package main

import "os"
import "log"
import "fmt"
import "gopkg.in/yaml.v3"

const CONFIG_PATH = "config.yml"

type Config struct {
    FilePath string `yaml:"file_path"`
    Url string `yaml:"download_url"`
}

func ReadConfig() Config {
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
            log.Println(err.Error())
        }
    }

    return config
}

func main() {
	config := ReadConfig()
	fmt.Println(config.Url)
}	