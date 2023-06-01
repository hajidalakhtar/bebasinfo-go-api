package utils

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var logger *zap.Logger

// InitLogger digunakan untuk menginisialisasi logger pada awal aplikasi
func InitLogger() {
	// Konfigurasi encoder untuk format log
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		TimeKey:        "time",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		CallerKey:      "caller",
		EncodeCaller:   zapcore.ShortCallerEncoder,
		StacktraceKey:  "stacktrace",
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}

	// Konfigurasi level log menjadi debug
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(zapcore.DebugLevel)

	// Konfigurasi core log
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.Lock(os.Stdout),
		atomicLevel,
	)

	// Buat logger baru
	logger = zap.New(core)
}

// GetLogger digunakan untuk mendapatkan instance logger
func GetLogger() *zap.Logger {
	return logger
}
