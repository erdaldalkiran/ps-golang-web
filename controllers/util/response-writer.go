package util

import (
	"compress/gzip"
	"net/http"
	"strings"
)

func GetResponseWriter(w http.ResponseWriter, r *http.Request) CloseableResponseWriter {
	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gw := gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gzip.NewWriter(w),
		}

		return gw
	}

	return closeableResponseWriter{
		ResponseWriter: w,
	}
}

type CloseableResponseWriter interface {
	http.ResponseWriter
	Close() error
}

type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

func (gw gzipResponseWriter) Write(data []byte) (int, error) {
	return gw.Writer.Write(data)
}

func (gw gzipResponseWriter) Close() error {
	return gw.Writer.Close()
}

func (gw gzipResponseWriter) Header() http.Header {
	return gw.ResponseWriter.Header()
}

type closeableResponseWriter struct {
	http.ResponseWriter
}

func (cw closeableResponseWriter) Close() error {
	return nil
}
