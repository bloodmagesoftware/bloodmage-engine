// Bloodmage Engine
// Copyright (C) 2023 Frank Mayer
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <https://github.com/bloodmagesoftware/bloodmage-engine/blob/main/LICENSE.md>.

package engine

import (
	"os"
	"path/filepath"

	"github.com/charmbracelet/log"
	"github.com/go-yaml/yaml"
)

type Options struct {
	Vsync      bool      `yaml:"vsync"`
	Fullscreen bool      `yaml:"fullscreen"`
	PixelScale int32     `yaml:"pixel_scale"`
	LogLevel   log.Level `yaml:"log_level"`
}

var (
	dataPath = "./data"
	options  = Options{
		Vsync:      true,
		Fullscreen: true,
		PixelScale: 4,
		LogLevel:   log.ErrorLevel,
	}
)

// Ensures that options file exists and is initialized.
// If it doesn't exist, create it.
// If it does exist, load it.
func InitOptions() {
	// check if dataDirectory exists
	_, err := os.Stat(dataPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info("Creating data directory", "path", dataPath)
			err := os.MkdirAll(dataPath, 0775)
			if err != nil {
				log.Fatal("Error creating data directory", "path", dataPath)
			}
		} else {
			log.Fatal("Error checking data directory", "path", dataPath, "error", err)
		}
	}

	engineOptionsPath := filepath.Join(dataPath, "engine.yaml")
	_, err = os.Stat(engineOptionsPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Info("Creating engine options file", "path", engineOptionsPath)
			err = options.Save(engineOptionsPath)
			if err != nil {
				log.Fatal("Error creating engine options file", "path", engineOptionsPath, "error", err)
			}
		} else {
			log.Fatal("Error checking engine options file", "path", engineOptionsPath, "error", err)
		}
	} else {
		log.Info("Loading engine options file", "path", engineOptionsPath)
		err = options.Load(engineOptionsPath)
		if err != nil {
			log.Fatal("Error loading engine options file", "path", engineOptionsPath, "error", err)
		}
	}

    log.SetLevel(options.LogLevel)

    log.Info("Engine options loaded", "options", options)
}

func (self *Options) Save(path string) error {
	dat, err := yaml.Marshal(self)
	if err != nil {
		return err
	}
	err = os.WriteFile(path, dat, 0664)
	if err != nil {
		return err
	}
	return nil
}

func (self *Options) Load(path string) error {
	dat, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(dat, self)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
