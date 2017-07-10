package logger

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/mappymappy/http_logger/formatter"
)

var sortOrders = []string{
	"request_uri",
	"method",
	"remote",
	"user_agent",
	"pid",
	"latency",
	"body",
	"status",
	"tag",
}

type LoggerMiddleware struct {
	formatter FormatterInterface
	logger    LoggerInterface
}

func Custom(f FormatterInterface, l LoggerInterface) *LoggerMiddleware {
	return &LoggerMiddleware{
		formatter: f,
		logger:    l,
	}

}

func Default() *LoggerMiddleware {
	logger := log.New(os.Stdout, "", log.LstdFlags)
	return &LoggerMiddleware{
		formatter: &formatter.Ltsv{},
		logger:    logger,
	}
}

func (l *LoggerMiddleware) createDataByRequest(r *http.Request) map[string]string {
	data := map[string]string{}
	data["request_uri"] = r.RequestURI
	data["method"] = r.Method
	data["remote"] = r.RemoteAddr
	data["user_agent"] = r.UserAgent()
	data["pid"] = strconv.Itoa(os.Getpid())
	body := new(bytes.Buffer)
	body.ReadFrom(r.Body)
	data["body"] = body.String()

	return data
}

func (l *LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	data := l.createDataByRequest(r)
	data["tag"] = "start_serve_request"
	l.outputData(data)
	rw := WrapResponseWriter(w)

	startTime := time.Now()
	next(rw, r)
	endTime := time.Now()

	data["tag"] = "finish_serve_request"
	latency := endTime.Sub(startTime)
	data["latency"] = latency.String()
	data["status"] = rw.Status()
	l.outputData(data)
}

func (l *LoggerMiddleware) outputData(data map[string]string) {
	targets := make([]formatter.FormatTargetInterface, len(sortOrders))
	i := 0
	for _, name := range sortOrders {
		val, ok := data[name]
		if !ok {
			continue
		}
		targets[i] = l.formatter.ConvertToFormatTarget(name, val)
		i++
	}
	byteOutput := l.formatter.Format(targets[:i])
	l.logger.Print(string(byteOutput))
}
