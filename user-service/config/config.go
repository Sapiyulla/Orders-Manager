package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type GRPC struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type REST struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Database struct {
	Type    string `yaml:"type"`
	Admin   string `yaml:"admin"`
	Host    string `yaml:"host"`
	Port    string `yaml:"port"`
	DB_name string `yaml:"db_name"`
}

type Config struct {
	GRPC     *GRPC     `yaml:"grpc"`
	REST     *REST     `yaml:"rest"`
	Database *Database `yaml:"database"`
}

func Load() (*Config, error) {
	// 1. Попробуем найти файл в разных возможных местах
	possiblePaths := []string{
		"config.yaml",                      // текущая директория
		"config/config.yaml",               // системная конфигурация
		filepath.Join("..", "config.yaml"), // на уровень выше
	}

	var filePath string
	var fileContent []byte
	var err error

	for _, path := range possiblePaths {
		fileContent, err = os.ReadFile(path)
		if err == nil {
			filePath = path
			break
		}
	}

	if filePath == "" {
		return nil, fmt.Errorf("не удалось найти config.yaml в следующих местах: %v", possiblePaths)
	}

	// 2. Парсим конфиг
	var cfg Config
	if err := yaml.Unmarshal(fileContent, &cfg); err != nil {
		return nil, fmt.Errorf("ошибка парсинга YAML в файле %s: %v", filePath, err)
	}

	fmt.Printf("%+v \n", cfg.Database)
	return &cfg, nil
}
