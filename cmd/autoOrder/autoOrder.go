package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"github.com/furkansahinfs/AutoOrder-Backend/pkg/server"
	"github.com/furkansahinfs/AutoOrder-Backend/pkg/version"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var (
	configFileFlag = flag.String("config.file", "config.yml", "Path to the configuration file.")
	logFileFlag    = flag.String("log.file", "autoOrder.log", "Path to the log file.")
	versionFlag    = flag.Bool("version", false, "Show version information.")
	debugFlag      = flag.Bool("debug", false, "Show debug information.")
)

func init() {
	// Parse command-line flags
	flag.Parse()

	// Log settings
	if *debugFlag {
		logrus.SetReportCaller(true)
		logrus.SetLevel(logrus.TraceLevel)
	} else {
		logrus.SetReportCaller(false)
		logrus.SetLevel(logrus.InfoLevel)
	}

	logrus.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	logFile, err := os.OpenFile(*logFileFlag, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.WithError(err).Fatal("Could not open log file")
	}

	logrus.SetOutput(logFile)
}

func main() {

	// Show version information
	if *versionFlag {
		fmt.Fprintln(os.Stdout, version.Print("autoOrder"))
		os.Exit(0)
	}

	// Load configuration file
	data, err := ioutil.ReadFile(*configFileFlag)
	if err != nil {
		logrus.WithError(err).Fatal("Could not load configuration")
	}
	var cfg server.Config
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		logrus.WithError(err).Fatal("Could not load configuration")
	}

	// Create server instance
	instance := server.NewInstance(&cfg)

	// Interrupt handler
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		logrus.Infof("Received %s signal", <-c)
		instance.Shutdown()
	}()

	// Start server
	logrus.Infof("Starting autoOrder %s", version.Info())
	logrus.Infof("Build context %s", version.BuildContext())
	instance.Start()
}
