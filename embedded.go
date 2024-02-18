package embedded_openapi

import (
	"embed"
	"io/fs"
	"log"
)

//go:embed swagger-ui
var swaggerUIdir embed.FS

func Fs() fs.FS {
	swaggerAssets, fsErr := fs.Sub(swaggerUIdir, "swagger-ui")
	if fsErr != nil {
		log.Fatalf("Failed to load embedded Swagger UI assets: %v", fsErr)
	}
	return swaggerAssets
}
