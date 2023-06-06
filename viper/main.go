package main

import (
	"fmt"
	"viper/util"

	"github.com/spf13/viper"
)

// we can write  and read from files using viper
// writing will be done at runtime

func main() {

	vp := viper.New()
	vp.SetConfigName("test")
	vp.SetConfigType("json")
	vp.AddConfigPath(".")

	err := vp.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(vp.GetString("name"))

	// In order to write to the file specified
	vp.Set("age", 23)
	vp.WriteConfig()

	config,err := util.LoadConfig() 
	if err != nil {
		panic(err)
	}
	fmt.Println(config)
}
