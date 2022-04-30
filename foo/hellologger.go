package foo

import "github.com/ethereum/go-ethereum/log"

func BuildLogger() log.Logger {
	logger := log.New("test", 123)
	logger.SetHandler(log.StdoutHandler)
	return logger
}
