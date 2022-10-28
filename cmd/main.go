package cmd

import (
	"fmt"
	"os"

	"github.com/gapidobri/radigo/internal/config"
	"github.com/gapidobri/radigo/internal/radio"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "radigo",
	Short: "Headless 24/7 Radio",
	Run: func(cmd *cobra.Command, args []string) {
		radio.Start(
			config.C.OverlayPath,
			config.C.MusicFolderPath,
			config.C.RtmpUrl,
			config.C.VideoEncoder,
			config.C.VideoBitrate,
		)
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
