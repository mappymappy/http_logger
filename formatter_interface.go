package logger

import "github.com/mappymappy/http_logger/formatter"

type FormatterInterface interface {
	Format([]formatter.FormatTargetInterface) []byte
	ConvertToFormatTarget(key, val string) formatter.FormatTargetInterface
}
