package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "gopkg.in/yaml.v2"
)

type AuthScheme struct {
    Name string `yaml:"name"`
    Type string `yaml:"type"`
    Config struct {
        User     string `yaml:"user"`
        Password string `yaml:"password"`
        URL      string `yaml:"url"`
        Refresh  bool   `yaml:"refresh"`
        TTL      int    `yaml:"ttl"`
    } `yaml:"config"`
}

type Route struct {
    Name         string `yaml:"name"`
    Method       string `yaml:"method"`
    Path         string `yaml:"path"`
    UpstreamURL  string `yaml:"upstream_url"`
    UpstreamAuth string `yaml:"upstream_auth"`
    LocalAuth    string `yaml:"local_auth"`
}

type Config struct {
    Routes      []Route      `yaml:"routes"`
    AuthSchemes []AuthScheme `yaml:"auth_schemes"`
}

func loadConfig(filename string) (Config, error) {
    data, err := ioutil.ReadFile(filename)
    if err != nil {
        return Config{}, err
    }

    var config Config
    if err := yaml.Unmarshal(data, &config); err != nil {
        return Config{}, err
    }

    return config, nil
}

func getAndPrintConfig() (Config) {
    config, err := loadConfig("config.yaml")
    if err != nil {
        log.Fatalf("Error reading or parsing config: %v", err)
    }

    fmt.Printf("Routes:\n")
    for _, route := range config.Routes {
        fmt.Printf("Name: %s, Method: %s, Path: %s, Upstream URL: %s, Upstream Auth: %s, Local Auth: %s\n", route.Name, route.Method, route.Path, route.UpstreamURL, route.UpstreamAuth, route.LocalAuth)
    }

    fmt.Printf("\nAuth Schemes:\n")
    for _, scheme := range config.AuthSchemes {
        fmt.Printf("Name: %s, Type: %s\n", scheme.Name, scheme.Type)
        if scheme.Type == "BearerTokenCached" {
            fmt.Printf("User: %s, Password: %s, URL: %s, Refresh: %v, TTL: %d\n", scheme.Config.User, scheme.Config.Password, scheme.Config.URL, scheme.Config.Refresh, scheme.Config.TTL)
        }
    }

	return config
}
