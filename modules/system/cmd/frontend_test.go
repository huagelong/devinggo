package cmd

import (
	"testing"
	"testing/fstest"
)

func TestFrontendEncodedFilePathPrefersBrotli(t *testing.T) {
	testFS := fstest.MapFS{
		"js/app.js":    &fstest.MapFile{Data: []byte("plain")},
		"js/app.js.br": &fstest.MapFile{Data: []byte("brotli")},
		"js/app.js.gz": &fstest.MapFile{Data: []byte("gzip")},
	}

	filePath, encoding := frontendEncodedFilePath(testFS, "gzip, br", "js/app.js")

	if filePath != "js/app.js.br" {
		t.Fatalf("filePath = %q, want %q", filePath, "js/app.js.br")
	}
	if encoding != "br" {
		t.Fatalf("encoding = %q, want %q", encoding, "br")
	}
}

func TestFrontendEncodedFilePathFallsBackToGzip(t *testing.T) {
	testFS := fstest.MapFS{
		"js/app.js":    &fstest.MapFile{Data: []byte("plain")},
		"js/app.js.gz": &fstest.MapFile{Data: []byte("gzip")},
	}

	filePath, encoding := frontendEncodedFilePath(testFS, "br, gzip", "js/app.js")

	if filePath != "js/app.js.gz" {
		t.Fatalf("filePath = %q, want %q", filePath, "js/app.js.gz")
	}
	if encoding != "gzip" {
		t.Fatalf("encoding = %q, want %q", encoding, "gzip")
	}
}

func TestFrontendEncodedFilePathFallsBackToOriginal(t *testing.T) {
	testFS := fstest.MapFS{
		"js/app.js": &fstest.MapFile{Data: []byte("plain")},
	}

	filePath, encoding := frontendEncodedFilePath(testFS, "br, gzip", "js/app.js")

	if filePath != "js/app.js" {
		t.Fatalf("filePath = %q, want %q", filePath, "js/app.js")
	}
	if encoding != "" {
		t.Fatalf("encoding = %q, want empty", encoding)
	}
}
