/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/booscaaa/hamburguer-go/cmd/cmd"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`config/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}
func main() {
	cmd.Execute()
}
