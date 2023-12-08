package main

import (
	"os"

	glog "github.com/kataras/golog"
)

var Logger = glog.New()

// configureLocal for local implementation
func configureLocal() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		glog.Fatal(err)
	}
	glog.SetLevelOutput("info", file)
}

func main() {
	configureLocal()
	glog.SetLevel("debug")
	glog.Println("This is a raw message, no levels, no colors.")
	glog.Info("This is an info message, with colors (if the output is terminal)")
	glog.Warn("This is a warning message")
	glog.Error("This is an error message")
	glog.Debug("This is a debug message")
	glog.Fatal(`Fatal will exit no matter what,
    but it will also print the log message if logger's Level is >=FatalLevel`)
}
