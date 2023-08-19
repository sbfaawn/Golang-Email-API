package env

import (
	"fmt"

	"github.com/spf13/viper"
)

var propReader EnvPropReader

func NewEnvPropReader() *EnvPropReader {
	propReader = EnvPropReader{
		FileName:    "",
		FileType:    "",
		Location:    "",
		EnvVariable: map[string]any{},
	}

	return &propReader
}

type EnvPropReader struct {
	FileName    string
	FileType    string
	Location    string
	EnvVariable map[string]any
}

func (reader *EnvPropReader) ReadEnv() {
	defer fmt.Println("ReadEnv is Finished")

	viper.SetConfigName(reader.FileName)
	viper.SetConfigType(reader.FileType)
	viper.AddConfigPath(reader.Location)
	// viper.AddConfigPath("../Golang-Gin-Gonic/")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("Config File is Missing")
		} else {
			fmt.Println("Another Error")
			fmt.Println(err)
		}
		panic(err)
	}

	viper.Unmarshal(&reader.EnvVariable)
	fmt.Println("Env Variable is Read Successfully")
}
