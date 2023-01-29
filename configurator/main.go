package configurator

import "github.com/spf13/viper"

func InitConfiguration(config interface{}, version string) error {
	err := initConfig()
	if err != nil {
		return err
	}
	path, filename, ext := splitPath(*PathToConfig)
	viper.AddConfigPath(path)
	viper.SetConfigName(filename)
	viper.SetConfigType(ext[1:])
	err = viper.ReadInConfig()
	if err != nil {
		return err
	}
	err = viper.Unmarshal(config)
	if err != nil {
		return err
	}
	setVersion(version)
	initInfo()
	printInfo()
	return nil
}
