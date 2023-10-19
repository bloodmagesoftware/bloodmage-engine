package engine

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
)

type Options struct {
	Vsync      bool  `yaml:"vsync"`
	Fullscreen bool  `yaml:"fullscreen"`
	PixelScale int32 `yaml:"pixel_scale"`
}

var (
	data_path = "."
	options   = Options{
		Vsync:      true,
		Fullscreen: true,
		PixelScale: 4,
	}
)

// Ensure that options file exists and is initialized.
// If it doesn't exist, create it.
// If it does exist, load it.
func InitOptions(ident string) {
	data_path = getDataPath(ident)
	if os.IsNotExist(os.ErrNotExist) {
		log.Println("Creating data directory", data_path)
		err := os.MkdirAll(data_path, 0775)
		if err != nil {
			log.Println("Error creating data directory", data_path)
			os.Exit(1)
		}
	}

	engine_options_path := filepath.Join(data_path, "engine.yaml")
	if os.IsExist(os.ErrExist) {
		log.Println("Loading options from", engine_options_path)
		err := options.Load(engine_options_path)
		if err != nil {
			log.Println("Error loading options", engine_options_path)
			log.Println("Saving default options to", engine_options_path)
			err := options.Save(engine_options_path)
			if err != nil {
				log.Println("Error saving options", engine_options_path)
			}
			os.Exit(1)
		}
	} else {
		log.Println("Saving default options to", engine_options_path)
		err := options.Save(engine_options_path)
		if err != nil {
			log.Println("Error saving options", engine_options_path)
			os.Exit(1)
		}
	}
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
		return err
	}
	return nil
}

func getDataPath(n string) string {
	if data_path != "." {
		return data_path
	}

	if os.Getenv("XDG_DATA_HOME") != "" {
		data_path = filepath.Join(os.Getenv("XDG_DATA_HOME"), n)
	} else if os.Getenv("HOME") != "" {
		data_path = filepath.Join(os.Getenv("HOME"), ".local", "share", n)
	} else if os.Getenv("USERPROFILE") != "" {
		data_path = filepath.Join(os.Getenv("USERPROFILE"), "AppData", "Local", n)
	} else if os.Getenv("APPDATA") != "" {
		data_path = filepath.Join(os.Getenv("APPDATA"), n)
	} else {
		data_path = "."
	}
	return data_path
}
