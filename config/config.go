package config

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type (
	ConfigData struct {
		DataBase  map[string]DB        `yaml:"database"`
		GoogleDoc map[string]GoogleDoc `yaml:"googledoc"`
	}

	DB struct {
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Database string `yaml:"database"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Driver   string `yaml:"driver"`
	}

	GoogleDoc struct {
		SpreadsheetsID string `yaml:"spreadsheetsID"`
		SheetName      string `yaml:"sheetName"`
	}
)

var Config *ConfigData

func NewConfig(path string) (*ConfigData, error) {
	if Config != nil {
		return Config, nil
	}

	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(yamlFile, &Config)
	if err != nil {
		return nil, err
	}

	return Config, nil
}

func (db DB) GetConnectionDSN() string {
	var dsn string

	switch db.Driver {
	case "mysql":
		dsn = fmt.Sprintf(
			"%s:%s@tcp(%s:%v)/%s?charset=utf8&parseTime=True&loc=Local",
			db.User,
			db.Password,
			db.Host,
			db.Port,
			db.Database,
		)
	case "postgres":
		dsn = fmt.Sprintf(
			"host=%s port=%v user=%s dbname=%s password=%s",
			db.Host,
			db.Port,
			db.User,
			db.Database,
			db.Password,
		)
	}

	return dsn
}
