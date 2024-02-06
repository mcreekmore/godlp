package embed

import (
	_ "embed"
)

//go:embed yt-dlp-linux
var YtDlpBinary []byte

//go:embed ffmpeg-linux
var FfmpegBinary []byte