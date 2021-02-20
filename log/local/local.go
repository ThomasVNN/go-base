package local


import (
	"github.com/ThomasVNN/go-base/log"
	"github.com/ThomasVNN/go-base/log/level"
	"os"
)

// NewLogger ...
func NewLogger(filepath string) log.Logger {
	var logger log.Logger

	//file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Error opening log file: %v\n", err)
	//	os.Exit(1)
	//}

	//if local.Getenv("ENVIRONMENT") == "dev" {
	logger = log.NewLogfmtLogger(os.Stdout)
	//} else {
	//	logger = log.NewJSONLogger(log.NewSyncWriter(file))
	//}
	logger = log.NewSyncLogger(logger)
	logger = log.With(logger,
		"service", "account",
		"hostname", "staging-1",
		"session", "1ce3f6v",
		"time:", log.DefaultTimestampUTC,
		"caller", log.DefaultCaller,
	)
	_ = level.Info(logger)
	return logger
}
