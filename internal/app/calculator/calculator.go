package calculator

import (
	"embed"
	"flag"
	"fmt"

	"github.com/maldan/go-restserver"
)

var DataDir string

func Start(frontFs embed.FS) {
	// General flags
	var host = flag.String("host", "127.0.0.1", "Server Hostname")
	var port = flag.Int("port", 16000, "Server Port")
	var dataDir = flag.String("dataDir", "db", "Data Directory")

	// Additional flags
	_ = flag.Int("clientPort", 8000, "Client Port")
	_ = flag.String("appId", "id", "App id")
	var cmd = flag.String("cmd", "", "Command")

	// Parse flags
	flag.Parse()
	DataDir = *dataDir

	// Run command
	if *cmd != "" {
		fmt.Println(*cmd)
		return
	}

	restserver.Start(fmt.Sprintf("%s:%d", *host, *port), map[string]interface{}{
		"/": restserver.VirtualFs{Root: "frontend/build/", Fs: frontFs},
	})
}
