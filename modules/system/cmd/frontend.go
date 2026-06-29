package cmd

import (
	"io/fs"
	"mime"
	"net/http"
	pathpkg "path"
	"strings"

	backendassets "devinggo/admin-ui/apps/backend"

	"github.com/gogf/gf/v2/net/ghttp"
)

var frontendFS fs.FS

func init() {
	var err error
	frontendFS, err = fs.Sub(backendassets.Dist, "dist")
	if err != nil {
		panic(err)
	}
}

func serveFrontend(r *ghttp.Request) {
	if !isFrontendRequest(r) {
		r.Middleware.Next()
		return
	}

	filePath := frontendFilePath(r.URL.Path)
	if serveFrontendFile(r, filePath) {
		r.ExitAll()
		return
	}

	if pathpkg.Ext(filePath) != "" && !acceptsHTML(r) {
		r.Response.WriteStatus(http.StatusNotFound)
		r.ExitAll()
		return
	}

	if serveFrontendFile(r, "index.html") {
		r.ExitAll()
		return
	}

	r.Middleware.Next()
}

func isFrontendRequest(r *ghttp.Request) bool {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		return false
	}

	requestPath := r.URL.Path
	return !hasAnyPathPrefix(requestPath,
		"/api",
		"/app",
		"/system",
		"/uploads",
		"/ws",
		"/health",
		"/swagger",
		"/api.json",
		"/debug",
	)
}

func hasAnyPathPrefix(requestPath string, prefixes ...string) bool {
	for _, prefix := range prefixes {
		if requestPath == prefix || strings.HasPrefix(requestPath, prefix+"/") {
			return true
		}
	}
	return false
}

func frontendFilePath(requestPath string) string {
	filePath := strings.TrimPrefix(pathpkg.Clean("/"+requestPath), "/")
	if filePath == "." {
		return "index.html"
	}
	return filePath
}

func acceptsHTML(r *ghttp.Request) bool {
	return strings.Contains(r.Header.Get("Accept"), "text/html")
}

func serveFrontendFile(r *ghttp.Request, filePath string) bool {
	data, err := fs.ReadFile(frontendFS, filePath)
	if err != nil {
		return false
	}

	if contentType := mime.TypeByExtension(pathpkg.Ext(filePath)); contentType != "" {
		r.Response.Header().Set("Content-Type", contentType)
	}
	if filePath == "index.html" {
		r.Response.Header().Set("Content-Type", "text/html; charset=utf-8")
	}
	if r.Method == http.MethodHead {
		r.Response.WriteHeader(http.StatusOK)
		return true
	}
	r.Response.Write(data)
	return true
}
