package main

import (
	"flag"
	log "github.com/sirupsen/logrus"
	"github.com/zkqiang/kuaizhi-feeder/enum"
	"github.com/zkqiang/kuaizhi-feeder/scheduler"
	"os"
	"time"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	var srcStr string
	flag.StringVar(&srcStr, "src", "", "feed source")
	flag.Parse()
	if srcStr == "" {
		log.Fatal("source is empty")
	}
	source := enum.GetSource(srcStr)
	if source == -1 {
		log.Fatalf("source %s is error", srcStr)
	}
	for {
		scheduler.Feed(source)
		time.Sleep(5 * time.Second)
	}
}
