package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
	add zap.logger to the context
	get the zap.logger from context
	setAtomic Value should change should be implemneted
*/

type LoggerField struct {
	Operation string
	Req       string
	ReqID     string
	SessionID string
	Msg       string
	Err       error
}

type Logger struct {
	log *zap.Logger
}

func GetLogger(logger *zap.Logger) Logging {
	return &Logger{log: logger}
}

func (l *Logger) loging(lvl zapcore.Level, fields *LoggerField) {
	if !l.log.Core().Enabled(lvl) {
		return
	}
	zapFeild := []zap.Field{}
	if lvl == zap.InfoLevel || lvl == zap.DebugLevel {
		zapFeild = append(zapFeild,
			zap.String("OP", fields.Operation),
			zap.String("Request", fields.Req),
			zap.String("RequestID", fields.ReqID),
			zap.String("SID", fields.SessionID),
			zap.String("Message", fields.Msg),
		)
	} else {
		zapFeild = append(zapFeild,
			zap.String("OP", fields.Operation),
			zap.String("Request", fields.Req),
			zap.String("RequestID", fields.ReqID),
			zap.String("SID", fields.SessionID),
			zap.String("Message", fields.Msg),
			zap.Error(fields.Err),
		)
	}

	l.log.Check(lvl, fields.Msg).Write(zapFeild...)
}

func (l *Logger) Debug(op string, req string, reqID string, sID string, msg string) {
	feild := &LoggerField{
		Operation: op,
		Req:       req,
		ReqID:     req,
		SessionID: sID,
		Msg:       msg,
		Err:       nil,
	}
	l.loging(zap.DebugLevel, feild)
}

func (l *Logger) Info(op string, req string, reqID string, sID string, msg string) {
	feild := &LoggerField{
		Operation: op,
		Req:       req,
		ReqID:     req,
		SessionID: sID,
		Msg:       msg,
		Err:       nil,
	}
	l.loging(zap.InfoLevel, feild)
}

func (l *Logger) Error(op string, req string, reqID string, sID string, msg string, err error) {
	feild := &LoggerField{
		Operation: op,
		Req:       req,
		ReqID:     req,
		SessionID: sID,
		Msg:       msg,
		Err:       err,
	}
	l.loging(zap.ErrorLevel, feild)
}

func (l *Logger) Fatal(op string, req string, reqID string, sID string, msg string, err error) {
	feild := &LoggerField{
		Operation: op,
		Req:       req,
		ReqID:     req,
		SessionID: sID,
		Msg:       msg,
		Err:       err,
	}
	l.loging(zap.FatalLevel, feild)
}
