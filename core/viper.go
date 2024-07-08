package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"mxaof_blog/dao/global"
)

func ViperInit() error {
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	err := v.ReadInConfig()
	if err != nil {
		log.Fatal(err)
		return err
	}
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config file changed:", in.Name)
		if err = v.Unmarshal(&global.GlobalConfig); err != nil {
			log.Fatal(err)
		}
	})
	if err = v.Unmarshal(&global.GlobalConfig); err != nil {
		return err
	}
	return nil
}
