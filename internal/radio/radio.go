package radio

import (
	"fmt"
	"io"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Start(overlayPath string, musicFolderPath string, rtmpUrl string, videoEncoder string, videoBitrate string) {

	dstCmd := ffmpeg.Filter(
		[]*ffmpeg.Stream{
			// Video overlay input
			ffmpeg.Input(
				overlayPath,
				ffmpeg.KwArgs{"stream_loop": -1},
			),
			// Audio stream input
			ffmpeg.Input("pipe:").Audio(),
		},
		"amerge",
		ffmpeg.Args{"inputs=2"},
	).Output(
		rtmpUrl,
		ffmpeg.KwArgs{
			"map":    "0:v",
			"c:v":    videoEncoder,
			"b:v":    videoBitrate,
			"format": "flv",
		},
	).Compile()

	dstStdin, _ := dstCmd.StdinPipe()
	dstCmd.Start()

	dirs, _ := os.ReadDir(musicFolderPath)
	for _, e := range dirs {
		if !e.Type().IsRegular() {
			continue
		}

		fmt.Println("Playing " + e.Name())

		srcCmd := ffmpeg.Input(
			musicFolderPath+"/"+e.Name(),
			ffmpeg.KwArgs{"re": ""},
		).Output(
			"pipe:",
			ffmpeg.KwArgs{"format": "ogg"},
		).Compile()

		srcStdout, _ := srcCmd.StdoutPipe()
		srcCmd.Start()

		io.Copy(dstStdin, srcStdout)
	}

}
