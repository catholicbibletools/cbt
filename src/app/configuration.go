package main

import (
	"bytes"
	"encoding/json"
	"os"
)

const (
	DEFAULT_BIBLE           = "en.ccb"
	DEFAULT_LANGUAGE        = "en"
	DEFAULT_OUTPUT_FORMAT   = "txt"
	DEFAULT_OUTER_SEPARATOR = ";"
	DEFAULT_MAJOR_SEPARATOR = ","
	DEFAULT_MINOR_SEPARATOR = "-"
)

var (
	config_file string
	config_data configuration
)

type Configuration interface {
	SetCurrentBible(value string)
	GetCurrentBible() string
	SetCurrentLanguage(value string)
	GetCurrentLanguage() string
	SetOutputFormat(value string)
	GetOutputFormat() string
	SetOuterSeparator(value string)
	GetOuterSeparator() string
	SetMajorSeparator(value string)
	GetMajorSeparator() string
	SetMinorSeparator(value string)
	GetMinorSeparator() string
}

type configuration struct {
	CurrentBible    string `json:"current-bible"`
	CurrentLanguage string `json:"current-language"`
	OutputFormat    string `json:"output-format"`
	OuterSeparator  string `json:"outer-separator"`
	MajorSeparator  string `json:"major-separator"`
	MinorSeparator  string `json:"minor-separator"`
}

func (me configuration) SetCurrentBible(value string) {
	config_data.CurrentBible = value
	SaveConfigurationToFile()
}

func (me configuration) GetCurrentBible() string {
	return config_data.CurrentBible
}

func (me configuration) SetCurrentLanguage(value string) {
	config_data.CurrentLanguage = value
	SaveConfigurationToFile()
}

func (me configuration) GetCurrentLanguage() string {
	return config_data.CurrentLanguage
}

func (me configuration) SetOutputFormat(value string) {
	config_data.OutputFormat = value
	SaveConfigurationToFile()
}

func (me configuration) GetOutputFormat() string {
	return config_data.OutputFormat
}

func (me configuration) SetOuterSeparator(value string) {
	config_data.OuterSeparator = value
	SaveConfigurationToFile()
}

func (me configuration) GetOuterSeparator() string {
	return config_data.OuterSeparator
}

func (me configuration) SetMajorSeparator(value string) {
	config_data.MajorSeparator = value
	SaveConfigurationToFile()
}

func (me configuration) GetMajorSeparator() string {
	return config_data.MajorSeparator
}

func (me configuration) SetMinorSeparator(value string) {
	config_data.MinorSeparator = value
	SaveConfigurationToFile()
}

func (me configuration) GetMinorSeparator() string {
	return config_data.MinorSeparator
}

func SaveConfigurationToFile() {
	b := new(bytes.Buffer)
	err1 := json.NewEncoder(b).Encode(&config_data)
	if err1 != nil {
		panic(err1)
	}
	file, err2 := os.Create(config_file)
	if err2 != nil {
		panic(err2)
	}
	defer file.Close()
	file.Truncate(0)
	file.Write(b.Bytes())
}

func LoadConfigurationFromFile(path string) Configuration {
	config_file = path
	file, err1 := os.Open(path)
	if err1 != nil {
		panic(err1)
	}
	defer file.Close()
	err2 := json.NewDecoder(file).Decode(&config_data)
	if err2 != nil {
		panic(err2)
	}
	if config_data.CurrentBible == "" {
		config_data.CurrentBible = DEFAULT_BIBLE
	}
	if config_data.CurrentLanguage == "" {
		config_data.CurrentLanguage = DEFAULT_LANGUAGE
	}
	if config_data.OuterSeparator == "" {
		config_data.OuterSeparator = DEFAULT_OUTER_SEPARATOR
	}
	if config_data.MajorSeparator == "" {
		config_data.MajorSeparator = DEFAULT_MAJOR_SEPARATOR
	}
	if config_data.MinorSeparator == "" {
		config_data.MinorSeparator = DEFAULT_MINOR_SEPARATOR
	}
	return config_data
}
