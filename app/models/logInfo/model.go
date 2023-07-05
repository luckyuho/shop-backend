package logInfo

import "time"

type logInfoRecorder struct {
	Header string    `json:"header"`
	Error  string    `json:"error"`
	Time   time.Time `json:"time"`
}

func NewLogInfo(
	header string,
	err error,
) *logInfoRecorder {
	return &logInfoRecorder{
		Header: header,
		Error:  err.Error(),
		Time:   time.Now(),
	}
}
