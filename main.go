package main

import (
	"flag"
	"io"
	"os"

	"github.com/diamondburned/gotk4-adwaita/pkg/adw"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	log "github.com/sirupsen/logrus"

	"github.com/sumnerevans/sublime-music-next/resources"
	"github.com/sumnerevans/sublime-music-next/ui"
)

func main() {
	logLevelStr := flag.String("loglevel", "debug", "the log level")
	logFilename := flag.String("logfile", "", "the log file to use (defaults to '' meaning no log file)")
	flag.Parse()

	// Configure logging
	if *logFilename != "" {
		logFile, err := os.OpenFile(*logFilename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
		if err == nil {
			mw := io.MultiWriter(os.Stdout, logFile)
			log.SetOutput(mw)
		} else {
			log.Errorf("Failed to open logging file; using default stderr: %s", err)
		}
	}
	log.SetFormatter(&log.TextFormatter{})
	log.SetLevel(log.InfoLevel)
	log.SetReportCaller(true)
	logLevel, err := log.ParseLevel(*logLevelStr)
	if err == nil {
		log.SetLevel(logLevel)
	} else {
		log.Errorf("Invalid loglevel '%s'. Using default 'debug'.", logLevel)
	}

	// Load the resources
	resources.LoadResources()

	app := &SublimeMusic{
		Application: adw.NewApplication("app.sublimemusic.SublimeMusicNext", gio.ApplicationFlagsNone),
	}
	app.ConnectActivate(app.activate)
	ui.SetupActions(app.Application)

	// Exit with whatever code the app exits with
	os.Exit(app.Run(os.Args))
}

type SublimeMusic struct {
	*adw.Application

	window *ui.RootWindow
}

func (s *SublimeMusic) activate() {
	s.window = ui.CreateRootWindow(&s.Application.Application)

	// audio := gtk.NewMediaFileForFilename("/home/sumner/tmp/ohea.mp3")
	// audio.Play()

	s.window.Show()
}
