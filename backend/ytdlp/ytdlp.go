package ytdlp

type YTDLP struct {
	Path string
}

type Download struct {
	YTDLP           YTDLP
	URL             string
	DestinationPath string
	FileName        string
}
