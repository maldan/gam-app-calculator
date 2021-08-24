package calculator

import (
	"embed"
	"flag"
	"fmt"

	"github.com/maldan/go-restserver"
)

var DataDir string

func Start(frontFs embed.FS) {
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16000, "Server Port")
	_ = flag.Int("clientPort", 8000, "Client Port")
	_ = flag.Bool("gui", false, "Use Gui")
	_ = flag.Bool("initDev", false, "Install dev")
	_ = flag.Int("width", 1100, "Window Width")
	_ = flag.Int("height", 900, "Window Height")
	var dataDir = flag.String("dataDir", "db", "Data Directory")
	_ = flag.String("folder", "", "Folder")
	_ = flag.String("appId", "id", "App id")
	flag.Parse()
	DataDir = *dataDir

	restserver.Start(fmt.Sprintf("%s:%d", *host, *port), map[string]interface{}{
		"/": restserver.VirtualFs{Root: "frontend/build/", Fs: frontFs},
	})
}
