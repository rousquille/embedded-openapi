package embedded_openapi

import (
	"embed"
	"io/fs"
	"log"
	"net/http"
)

//go:embed swagger-ui
var swaggerUIdir embed.FS

func Assets() http.FileSystem {
	swaggerAssets, fsErr := fs.Sub(swaggerUIdir, "swagger-ui")
	if fsErr != nil {
		log.Fatalf("Failed to load embedded OpenAPI UI assets: %v", fsErr)
	}
	return http.FS(swaggerAssets)
}
