package logger

type Logging interface {
	Debug(op string, req string, reqID string, sID string, msg string)
	Info(op string, req string, reqID string, sID string, msg string)
	Error(op string, req string, reqID string, sID string, msg string, err error)
	Fatal(op string, req string, reqID string, sID string, msg string, err error)
}
