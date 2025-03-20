package logger

import (
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.Logger
}

type CustomEncoder struct {
	zapcore.Encoder
}

func (e *CustomEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	// Добавляем цвета для уровней логирования
	var levelColor string
	switch entry.Level {
	case zapcore.DebugLevel:
		levelColor = "\033[36m" // Cyan
	case zapcore.InfoLevel:
		levelColor = "\033[32m" // Green
	case zapcore.WarnLevel:
		levelColor = "\033[33m" // Yellow
	case zapcore.ErrorLevel, zapcore.FatalLevel, zapcore.PanicLevel:
		levelColor = "\033[31m" // Red
	default:
		levelColor = "\033[0m" // Reset
	}

	formatted := fmt.Sprintf(
		"%s[%s]%s \033[1m%s\033[0m\033[0m: %s\n",
		levelColor,
		entry.Time.Format("2006-01-02 15:04:05"),
		levelColor,
		entry.Level.CapitalString(),
		entry.Message,
	)

	buf := buffer.NewPool().Get()
	buf.AppendString(formatted)
	return buf, nil
}

func Init() *Logger {
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	config.EncoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	encoder := &CustomEncoder{
		Encoder: zapcore.NewConsoleEncoder(config.EncoderConfig),
	}

	core := zapcore.NewCore(encoder, zapcore.AddSync(zapcore.Lock(os.Stdout)), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())

	return &Logger{logger}
}