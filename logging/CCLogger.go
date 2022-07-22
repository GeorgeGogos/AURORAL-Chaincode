package logging

import (
	"fmt"
	"log"
	"time"

	"github.com/s7techlab/cckit/router"
)

type ccLogger struct {
	*log.Logger
	txContext router.Context
}

var CCLoggerInstance *ccLogger

func InitCCLogger() {
	logger := &ccLogger{
		txContext: nil,
	}
	logger.Logger = log.New(logger, "", 0)
	CCLoggerInstance = logger
}

func SetContextMiddlewareFunc() router.MiddlewareFunc {
	return func(next router.HandlerFunc, pos ...int) router.HandlerFunc {
		return func(c router.Context) (interface{}, error) {
			CCLoggerInstance.txContext = c
			return next(c)
		}
	}
}

func (ccl ccLogger) Write(bytes []byte) (int, error) {
	client, _ := ccl.txContext.Client()
	mspID, _ := client.GetMSPID()
	clientCert, _ := client.GetX509Certificate()
	functionName := ccl.txContext.Path()
	txTimestamp, _ := ccl.txContext.Time()
	return fmt.Printf("%s txID=%s txTimestamp=%s client=(%s, %s) function=%s(): %s",
		time.Now().UTC().Format(time.RFC3339),
		ccl.txContext.Stub().GetTxID(),
		txTimestamp.UTC().Format(time.RFC3339),
		mspID, clientCert.Subject.CommonName, functionName, string(bytes))
}
