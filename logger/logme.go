//Author Mohammad Naser Abbasanadi
//Creating Date 2018-11-20
// logme.go is to log all issues

package logger

import (
	"flag"
	"io"
	"log"
	"os"

	"GolangOrdering/config"
)

// var
var (
	Log *log.Logger
)

//init will be fired automaticly in this package
func init() {

	cnf := config.GetConfigInstance()

	var logpath = cnf.LOGPATH
	var filepath = logpath + "/info.log"
	_ = os.Mkdir(logpath, os.ModePerm)

	flag.Parse()
	var file, err1 = os.Create(filepath)

	if err1 != nil {
		panic(err1)
	}
	Log = log.New(file, "", log.LstdFlags|log.Lshortfile)
	mw := io.MultiWriter(os.Stdout, file)
	Log.SetOutput(mw)
	Log.Println("LogFile1 : " + filepath)
}
