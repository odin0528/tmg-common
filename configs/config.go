package configs

import (
	"log"
	"os"

	"gopkg.in/ini.v1"
)

func HasSection(section string) bool {
	if envConfig == nil {
		return false
	}
	return envConfig.HasSection(section)
}

func Get(section, key, defaultValue string) string {
	configMutex.RLock()
	defer configMutex.RUnlock()

	if envConfig == nil {
		return defaultValue
	}

	var value = defaultValue
	if v := envConfig.Section(section).Key(key).String(); v != "" {
		value = v
	}

	return value
}

func GetInt(section, key string, defaultValue int) int {
	configMutex.RLock()
	defer configMutex.RUnlock()

	if envConfig == nil {
		return defaultValue
	}

	var value = defaultValue
	if v, err := envConfig.Section(section).Key(key).Int(); nil == err {
		value = v
	}

	return value
}

func GetBool(section, key string, defaultValue bool) bool {
	configMutex.RLock()
	defer configMutex.RUnlock()

	if envConfig == nil {
		return defaultValue
	}

	var value = defaultValue
	if v, err := envConfig.Section(section).Key(key).Bool(); nil == err {
		value = v
	}

	return value
}

func GetFloat64(section, key string, defaultValue float64) float64 {
	configMutex.RLock()
	defer configMutex.RUnlock()

	if envConfig == nil {
		return defaultValue
	}

	var value = defaultValue
	if v, err := envConfig.Section(section).Key(key).Float64(); nil == err {
		value = v
	}

	return value
}

func Init(prefix string) {
	err := loadDefaultConfig(prefix)

	if err != nil {
		log.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
}

func loadDefaultConfig(path string) error {
	pathBytes := []byte(path)
	if string(pathBytes[len(pathBytes)-1]) != "/" {
		path += "/"
	}

	configList := []string{
		"web_api.ini",
		"websocket.ini",
		"system.ini",
		"log.ini",
	}

	for index, config := range configList {
		config = path + config
		configList[index] = config
	}

	return LoadConfig(configList)
}

// LoadConfigFromBytes loads configuration from an in-memory byte slice.
// The data should be in INI format. If envConfig is already initialized,
// the data is appended; otherwise a new config is created.
// Primarily used in tests to load a minimal config without files on disk.
func LoadConfigFromBytes(data []byte) error {
	configMutex.Lock()
	defer configMutex.Unlock()

	var loadOptions = ini.LoadOptions{
		SpaceBeforeInlineComment: true,
	}

	if envConfig != nil {
		return envConfig.Append(data)
	}

	iniFile, err := ini.LoadSources(loadOptions, data)
	if err != nil {
		return err
	}
	envConfig = iniFile
	return nil
}

// ResetConfig clears the loaded configuration. Intended for use in tests only.
func ResetConfig() {
	configMutex.Lock()
	defer configMutex.Unlock()
	envConfig = nil
}

// path should include file name. Ex: "config/env.ini"
func LoadConfig(filePaths []string) error {
	var err error

	paths := []interface{}{}
	for _, path := range filePaths {
		paths = append(paths, path)
	}

	configMutex.Lock()
	defer configMutex.Unlock()

	if envConfig != nil {
		if len(paths) > 1 {
			err = envConfig.Append(paths[0], paths[1:]...)
		} else {
			err = envConfig.Append(paths[0])
		}
	} else {
		var iniFile *ini.File
		var loadOptions = ini.LoadOptions{
			SpaceBeforeInlineComment: true,
		}

		if len(paths) > 1 {
			iniFile, err = ini.LoadSources(loadOptions, paths[0], paths[1:]...)
		} else {
			iniFile, err = ini.LoadSources(loadOptions, paths[0])
		}

		if err != nil {
			return err
		}

		envConfig = iniFile
	}

	return err
}
