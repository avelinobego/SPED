package logger

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
)

type logChan chan interface{}
type doneChan chan bool

var Info, Warning, Error logChan

var done doneChan
var ctx context.Context
var cancel context.CancelFunc

var logInfo, logWarning, logError *log.Logger

var InfoFile, ErrorFile, WarningFile io.WriteCloser

var flags int = log.Ldate | log.Ltime | log.Lmsgprefix

func StartLogEvents() {
	done = make(doneChan, 1)
	Info = make(logChan, 1024)
	Error = make(logChan, 1024)
	Warning = make(logChan, 1024)

	if InfoFile == nil {
		InfoFile = os.Stdout
	}
	logInfo = log.New(InfoFile, "Info: ", flags)

	if WarningFile == nil {
		WarningFile = os.Stdout
	}
	logWarning = log.New(WarningFile, "Warning: ", flags)

	if ErrorFile == nil {
		ErrorFile = os.Stderr
	}
	logError = log.New(ErrorFile, "Error: ", flags)

	ctx, cancel = context.WithCancel(context.Background())
	go start()
}

func Done() {
	cancel()
	<-done
	close(InfoFile)
	close(WarningFile)
	close(ErrorFile)
}

func close(w any) {
	if c, ok := w.(io.Closer); ok {
		c.Close()
	}
}

func start() {
	fmt.Println("comecando...")
	for {
		select {
		case text := <-Info:
			logInfo.Println(text)
		case text := <-Warning:
			logWarning.Println(text)
		case text := <-Error:
			logError.Println(text)
		case <-ctx.Done():
			if len(Error) > 0 || len(Warning) > 0 || len(Info) > 0 {
				continue
			}
			close(InfoFile)
			close(WarningFile)
			close(ErrorFile)
			fmt.Println("Terminando processamento de log.")
			done <- true
			return
		}
	}
}
