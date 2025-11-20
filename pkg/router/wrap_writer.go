package router

import "net/http"

type WrapWriter struct {
	http.ResponseWriter
	wroteCode bool
	code      int
}

func NewWrapWriter(w http.ResponseWriter) *WrapWriter {
	return &WrapWriter{ResponseWriter: w}
}

func (w *WrapWriter) WriteHeader(code int) {
	if !w.wroteCode {
		w.setCode(code)
	}

	w.ResponseWriter.WriteHeader(code)
}

func (w *WrapWriter) Write(data []byte) (int, error) {
	if !w.wroteCode {
		w.setCode(http.StatusOK)
	}

	return w.ResponseWriter.Write(data)
}

func (w *WrapWriter) Code() int {
	return w.code
}

func (w *WrapWriter) setCode(code int) {
	w.wroteCode = true
	w.code = code
}
