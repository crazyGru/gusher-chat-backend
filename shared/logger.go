package shared

import (
	"log"
	"os"
)

var ErrorLogger *log.Logger

func init() {
	ErrorLogger = log.New(os.Stderr, "[err] ", log.Ldate|log.Ltime|log.Llongfile|log.Lmsgprefix)
}
