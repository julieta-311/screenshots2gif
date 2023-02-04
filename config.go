package main

import (
	"flag"
	"fmt"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type config struct {
	durationSeconds     int
	fps                 int
	outputDir           string
	imgSaveDir          string
	initialSleepSeconds int
	loop                bool
	screen              int
	timeOutMinutes      int
	widthPixels         int
}

func getConfig() (c config, err error) {
	v := viper.New()

	pflag.IntP("durationSeconds", "d", 3, "how many seconds long the animation should be")
	pflag.IntP("fps", "f", 10, "the frame rate the animation should be made in")
	pflag.IntP("initialSleepSeconds", "s", 5, "the number of seconds to wait before taking the first snapshot")
	pflag.BoolP("loop", "l", true, "if the animation should loop indefinitely")
	pflag.StringP("outputDir", "o", ".", "the absolute path to the directory where the output gif is to be stored")
	pflag.IntP(
		"screen",
		"S",
		0,
		"the number identifying the screen that should be captured, default is the main screen = 0",
	)
	pflag.IntP("timeOutMinutes", "t", 5, "the number of minutes until the app is shut down as a safety measure")
	pflag.IntP(
		"widthPixels",
		"w",
		0,
		"the desired width of the animation in pixels, the image will be scaled"+
			"preserving aspect ratio, the original value will be used if this value is set to 0",
	)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	if err := v.BindPFlags(pflag.CommandLine); err != nil {
		return c, fmt.Errorf("failed to bind flags: %w", err)
	}

	c = config{
		durationSeconds:     v.GetInt(("durationSeconds")),
		fps:                 v.GetInt("fps"),
		initialSleepSeconds: v.GetInt("initialSleepSeconds"),
		loop:                v.GetBool("loop"),
		outputDir:           v.GetString("outputDir"),
		screen:              v.GetInt("screen"),
		timeOutMinutes:      v.GetInt(("timeOutMinutes")),
		widthPixels:         v.GetInt("widthPixels"),
	}

	return c, nil
}
