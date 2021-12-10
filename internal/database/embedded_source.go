package database

import (
	"embed"
	"github.com/golang-migrate/migrate/v4/source"
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"io/ioutil"
	"log"
	"net/http"
)

type EmbbedSource struct {
	httpfs.PartialDriver
	fs embed.FS
}

func NewEmbeddedFileSource(path string) (source.Driver, error) {
	embeddedSource := &EmbbedSource{fs: embeddedFiles}
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		log.Print(file.Name(), file.IsDir())
	}

	sourceDriver, err := embeddedSource.Open(path)
	if err != nil {
		return nil, err
	}

	return sourceDriver, nil
}

func init() {
	source.Register("embed", &EmbbedSource{})
}

func (e *EmbbedSource) Open(path string) (source.Driver, error) {
	es := &EmbbedSource{fs: e.fs}

	if err := es.Init(http.FS(es.fs), path); err != nil {
		return nil, err
	}

	return es, nil
}
