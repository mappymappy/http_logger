package logger

import (
	"fmt"
	"net/http"
)

type ResponseWriter struct {
	writer http.ResponseWriter
	status int
}

func WrapResponseWriter(w http.ResponseWriter) *ResponseWriter {
	return &ResponseWriter{w, 0}
}

func (w *ResponseWriter) Header() http.Header {
	return w.writer.Header()
}

func (w *ResponseWriter) Write(bs []byte) (int, error) {
	return w.writer.Write(bs)
}

func (w *ResponseWriter) WriteHeader(status int) {
	w.status = status
	w.writer.WriteHeader(status)
}

func (w *ResponseWriter) Status() string {
	return fmt.Sprintf("%d", w.status)
}
