package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type config struct {
	OverlayPath     string `mapstructure:"overlay_path"`
	MusicFolderPath string `mapstructure:"music_folder_path"`
	RtmpUrl         string `mapstructure:"rtmp_url"`
	VideoEncoder    string `mapstructure:"video_encoder"`
	VideoBitrate    string `mapstructure:"video_bitrate"`
}

var C config

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	viper.SetDefault("video_bitrate", "6000k")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		panic(fmt.Errorf("unable to parse config, %v", err))
	}

}
