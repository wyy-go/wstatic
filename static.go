package wstatic

import (
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	Index     = "index.html"
	Separator = "/"
)

type localFileSystem struct {
	http.FileSystem
	options    Options
	fileServer http.Handler
	fileMap    map[string]bool
}

func localFile(opts ...Option) *localFileSystem {
	local := localFileSystem{
		fileMap: make(map[string]bool),
	}

	for _, opt := range opts {
		opt(&local.options)
	}

	if local.options.embed {
		local.FileSystem = http.FS(local.options.fs)
	} else {
		if len(local.options.root) < 1 {
			panic("root is empty")
		}

		local.FileSystem = gin.Dir(local.options.root, local.options.indexes)
	}

	return &local
}

func (l *localFileSystem) parseUrl(req *http.Request, index bool) string {
	filepath := strings.TrimPrefix(req.URL.Path, l.options.urlPrefix)
	if !strings.HasPrefix(filepath, Separator) {
		filepath = l.options.root + Separator + filepath
	} else {
		filepath = l.options.root + filepath
	}

	if filepath == l.options.root+Separator && index {
		filepath = filepath + Index
	}
	return filepath
}

func (l *localFileSystem) Exists(ctx *gin.Context) bool {
	var filepath = ctx.Request.URL.Path

	if l.options.embed {

		filepath = l.parseUrl(ctx.Request, true)

		if _, ok := l.fileMap[filepath]; !ok {
			if f, err := l.FileSystem.Open(filepath); err != nil {
				return false
			} else if _, err := f.Stat(); err != nil {
				return false
			}
			l.fileMap[filepath] = true
		}

		return true
	} else if p := strings.TrimPrefix(filepath, l.options.urlPrefix); len(p) < len(filepath) {
		name := path.Join(l.options.root, p)
		stats, err := os.Stat(name)
		if err != nil {
			return false
		}
		if stats.IsDir() {
			if !l.options.indexes {
				index := path.Join(name, Index)
				_, err := os.Stat(index)
				if err != nil {
					return false
				}
			}
		}
		return true
	}
	return false
}

func (l *localFileSystem) request(req *http.Request) *http.Request {

	if l.options.embed {

		filepath := l.parseUrl(req, false)

		req.URL.Path = filepath
		req.URL.RawPath = filepath
		return req
	}
	return req
}

func (l *localFileSystem) GetFileServer() http.Handler {
	l.fileServer = http.FileServer(l.FileSystem)
	if len(l.options.urlPrefix) > 0 {
		l.fileServer = http.StripPrefix(l.options.urlPrefix, l.fileServer)
	}
	return l.fileServer
}

// New returns a middleware handler that serves static files in the given directory.
func New(opts ...Option) gin.HandlerFunc {
	fs := localFile(opts...)
	fileServer := fs.GetFileServer()

	return func(c *gin.Context) {
		if fs.Exists(c) {
			fileServer.ServeHTTP(c.Writer, fs.request(c.Request))
			c.Abort()
		}
	}
}
