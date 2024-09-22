package main

import (
	"fmt"
	"log/slog"
	"os"
)

func replacer(groups []string, a slog.Attr) slog.Attr {
    fmt.Println()
    fmt.Println()
    fmt.Println("Key: ", a.Key)
    fmt.Println("Value: ", a.Value)
    fmt.Println("Full: ", a)
    fmt.Println()
    fmt.Println()

    return a

	//    switch a.Key {
	//    	case slog.TimeKey:
	// 	return slog.Time(string(TimeKey), a.Value.Resolve().Time())
	//
	// case slog.MessageKey:
	// 	return slog.String(string(MessageKey), a.Value.String())
	//
	// case slog.SourceKey:
	// 	src := a.Value.Any().(*slog.Source)
	// 	if src.Function == "" {
	// 		return slog.Attr{}
	// 	}
	// 	base := path.Base(src.File)
	// 	return slog.String(
	// 		string(CallerKey),
	// 		fmt.Sprintf("%s:%s:%d", base, src.Function, src.Line),
	// 	)
	//
	// case slog.LevelKey:
	// 	l := a.Value.Any().(slog.Level)
	// 	return slog.String(string(LevelKey), levels[l])
	//
	// default:
	// 	return a
	//    }
}

func defaultLogger() {
    logger := slog.Default()
    logger.Info("Default logger ")

    // Log with key value pair
    logger.Info("Logger with key value pair",)
}

func setupJSONDefaultLogger() {
    logger := slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
        AddSource: true,
        Level: slog.LevelDebug,
        ReplaceAttr: replacer,
    }))

    logger.Info("JSON log", "key1", "value1")
}

func main() {
    slog.Info("Hello world")
    slog.Debug("Hello world")
    slog.Warn("Hello world")
    slog.Error("Hello world")


    defaultLogger()

    setupJSONDefaultLogger()
}
