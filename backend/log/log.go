package log

import (
	"flower-backend/config"
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type logEncoder struct {
	zapcore.Encoder
	errFile     *os.File
	file        *os.File
	currentDate string
}

const (
	colorRed    = "\033[31m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorReset  = "\033[0m"
)

func myEncodeLevel(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	switch level {
	case zapcore.InfoLevel:
		enc.AppendString(colorBlue + "INFO" + colorReset)
	case zapcore.WarnLevel:
		enc.AppendString(colorYellow + "WARN" + colorReset)
	case zapcore.ErrorLevel, zapcore.DPanicLevel, zapcore.PanicLevel, zapcore.FatalLevel:
		enc.AppendString(colorRed + "ERROR" + colorReset)
	default:
		enc.AppendString(level.CapitalString())
	}
}

func (e *logEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	buff, err := e.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}
	data := buff.String()
	buff.Reset()
	buff.AppendString("[flower-backend]" + data)
	data = buff.String()
	// time splice
	now := time.Now().Format("2006-01-02")
	if e.currentDate != now {
		os.MkdirAll(fmt.Sprintf("./logs/%s", now), 0755)
		name := fmt.Sprintf("./logs/%s/output.log", now)
		file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			return nil, err
		}
		e.file = file
		e.currentDate = now
	}

	switch entry.Level {
	case zapcore.ErrorLevel:
		if e.errFile == nil {
			name := fmt.Sprintf("./logs/%s/error.log", now)
			file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
			if err != nil {
				return nil, err
			}
			e.errFile = file
		}
		e.errFile.WriteString(buff.String())
		e.errFile.WriteString("\n")
	}

	if e.currentDate != now {
		e.file.WriteString(data)
	}
	return buff, nil
}

func InitLog() *zap.Logger {
	config := config.LoadConfig()

	if config.GO_ENV == "development" {
		cfg := zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
		cfg.EncoderConfig.EncodeLevel = myEncodeLevel

		encoder := &logEncoder{
			Encoder: zapcore.NewConsoleEncoder(cfg.EncoderConfig),
		}
		core := zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			zapcore.InfoLevel,
		)
		logger := zap.New(core, zap.AddCaller())

		zap.ReplaceGlobals(logger)
		return logger
	} else {
		cfg := zap.NewProductionConfig()
		cfg.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
		cfg.EncoderConfig.EncodeLevel = myEncodeLevel
		encoder := zapcore.NewJSONEncoder(cfg.EncoderConfig)
		core := zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			zapcore.InfoLevel,
		)
		logger := zap.New(core, zap.AddCaller())
		zap.ReplaceGlobals(logger)
		return logger
	}
}
