package main

type errorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type config struct {
	Listen            int          `yaml:"listen"`
}
