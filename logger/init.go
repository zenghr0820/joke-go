package logger

import (
	"github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"joke-go/config"
	"joke-go/utils"
	"os"
	"strings"
	"time"
)

var (
	zapLogger *zap.Logger
	zapSugar  *zap.SugaredLogger
	// mainLogger default logger, with unit filed valued "main"
	mainLogger *Logger
)

func InitLog() {
	// 设置一些基本日志格式 具体含义还比较好理解，直接看zap源码也不难懂
	encoderConfig := zapcore.EncoderConfig{
		MessageKey:  "msg",
		LevelKey:    "level",
		EncodeLevel: zapcore.CapitalLevelEncoder,
		TimeKey:     "ts",
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		CallerKey:    "file",
		EncodeCaller: zapcore.ShortCallerEncoder,
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
		EncodeName: zapcore.FullNameEncoder,
	}
	var encoder zapcore.Encoder

	if config.EnvMode == "dev" {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	}

	// 实现两个判断日志等级的interface
	infoLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.WarnLevel
	})

	warnLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	// 获取 info、warn日志文件的io.Writer 抽象 getWriter() 在下方实现
	loggerPath := config.GetConfig("logger.dir")
	appName := config.GetConfig("app.name")
	if loggerPath == "" || config.EnvMode == "dev" {
		address, _ := os.Getwd()
		loggerPath = address + "/logs"
	}

	// 检查文件夹是否存在
	isExist := utils.CheckFileAndCreate(loggerPath)

	if isExist != nil {
		panic("[文件夹不存在] " + loggerPath)
	}

	infoWriter := getWriter(loggerPath + "/" + appName + ".log")
	warnWriter := getWriter(loggerPath + "/common-error.log")

	writeConsole := zapcore.AddSync(os.Stdout)
	writeInfoFile := zapcore.AddSync(infoWriter)
	writeWarnFile := zapcore.AddSync(warnWriter)

	var wsInfo zapcore.WriteSyncer
	var wsWarn zapcore.WriteSyncer
	if config.EnvMode == "dev" {
		wsInfo = zapcore.NewMultiWriteSyncer(writeConsole, writeInfoFile)
		wsWarn = zapcore.NewMultiWriteSyncer(writeConsole, writeWarnFile)
	} else {
		wsInfo = zapcore.NewMultiWriteSyncer(writeInfoFile)
		wsWarn = zapcore.NewMultiWriteSyncer(writeWarnFile)
	}

	// 最后创建具体的Logger
	core := zapcore.NewTee(
		zapcore.NewCore(encoder, wsInfo, infoLevel),
		zapcore.NewCore(encoder, wsWarn, warnLevel),
	)

	// 开启开发模式，堆栈跟踪 需要传入 zap.AddCaller() 才会显示打日志点的文件名和行数, 有点小坑
	caller := zap.AddCaller()
	// 防止zap始终将包装器代码报告为调用者( 需要跳过一个级别，否则打印的文件名和行号 是封装的文件名)
	skip := zap.AddCallerSkip(1)

	zapLogger = zap.New(core, caller, skip)
	zapSugar = zapLogger.Sugar()
	mainLogger = NewLoggerOf()
}

func getWriter(filename string) io.Writer {
	// 生成rotatelogs的Logger 实际生成的文件名 demo.log.YYmmddHH
	// demo.log是指向最新日志的链接
	// 保存7天内的日志，每天分割一次日志
	hook, err := rotatelogs.New(
		filename+".%Y-%m-%d", // 没有使用go风格反人类的format格式
		rotatelogs.WithLinkName(filename),
		rotatelogs.WithMaxAge(time.Hour*24*7),
		rotatelogs.WithRotationTime(time.Hour*24),
	)

	if err != nil {
		panic(err)
	}
	return hook
}

// 把字符串转换为日志级别（数字）
func convertLogLevel(levelStr string) zapcore.Level {
	// 不区分大小写
	levelStr = strings.ToLower(levelStr)
	var level zapcore.Level
	switch levelStr {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}
	return level
}
