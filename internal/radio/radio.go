package radio

import (
	"fmt"
	"io"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

func Start() {

	dstCmd := ffmpeg.Filter(
		[]*ffmpeg.Stream{
			// Video overlay input
			ffmpeg.Input(
				"testing/bradio.mov",
				ffmpeg.KwArgs{"stream_loop": -1},
			),
			// Audio stream input
			ffmpeg.Input("pipe:").Audio(),
		},
		"amerge",
		ffmpeg.Args{"inputs=2"},
	).Output(
		"rtmp://127.0.0.1/live/rJMKs3tQj",
		ffmpeg.KwArgs{
			"map": "0:v",
			"c:v": "h264_videotoolbox",
			// "c:a":    "aac",
			"b:v":    "6000k",
			"format": "flv",
		},
	).Compile()

	dstStdin, _ := dstCmd.StdinPipe()

	dstCmd.Start()

	dirs, _ := os.ReadDir("testing/music")

	for _, e := range dirs {
		if !e.Type().IsRegular() {
			continue
		}

		fmt.Println("Playing " + e.Name())

		srcCmd := ffmpeg.Input("testing/music/"+e.Name(), ffmpeg.KwArgs{"re": ""}).Output("pipe:", ffmpeg.KwArgs{"format": "ogg"}).Compile()
		srcStdout, _ := srcCmd.StdoutPipe()
		srcCmd.Start()

		io.Copy(dstStdin, srcStdout)
	}

}
