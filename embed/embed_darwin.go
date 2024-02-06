package embed

import (
	_ "embed"
)

//go:embed yt-dlp_darwin
var YtDlpBinary []byte

//go:embed ffmpeg_darwin
var FfmpegBinary []byte