package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/julieta-311/screenshots2gif/animate"
	"github.com/julieta-311/screenshots2gif/capture"
)

const maximumAllowedFPS = 32

func main() {
	ctx := context.Background()

	cfg, err := getConfig()
	if err != nil {
		log.Fatalf("failed to get config: %v", err)
	}

	timeOut := time.Duration(cfg.timeOutMinutes)*time.Minute + time.Second*time.Duration(cfg.initialSleepSeconds)
	ctx, cancel := context.WithTimeout(ctx, timeOut)

	log.Printf("\U0001F4F7 Starting screenshots2gif \U0001F4F7\n\n")
	log.Printf("Operation will time out if not done after %v minutes.\n", cfg.timeOutMinutes)
	log.Printf("Taking screenshots to generate %d seconds long gif at %d fps.\n", cfg.durationSeconds, cfg.fps)
	log.Printf("Selected screen: %d (main screen = 0).\n", cfg.screen)
	log.Printf("Output directory: %s.\n To cancel hit ctrl + c.", cfg.outputDir)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	defer func() {
		signal.Stop(quit)
		cancel()
	}()

	finished := make(chan bool)
	go func() {
		run(ctx, cfg)
		finished <- true
	}()

	select {
	case <-finished:
		log.Printf("*** Done. ***\n\n")
	case <-quit:
		log.Printf("*** Got cancel signal. Shutting down. ***\n\n")
	case <-ctx.Done():
		log.Printf("*** Operation timed out. ***\n\n")
	}

	log.Println("Goodbye.")
}

func run(ctx context.Context, cfg config) {
	if cfg.initialSleepSeconds > 0 {
		log.Printf("Sleeping for %v seconds before starting.\n", cfg.initialSleepSeconds)
		time.Sleep(time.Duration(cfg.initialSleepSeconds) * time.Second)
	}

	animDelay, err := calculateDelay(cfg.fps)
	if err != nil {
		log.Fatalf("failed to calculate delay between shots: %v", err)
	}

	delayBetweenShots := time.Duration(10) * time.Millisecond * time.Duration(animDelay)
	nFrames := cfg.fps * cfg.durationSeconds

	imgSaveDir, err := os.MkdirTemp("", "*-screenshots")
	if err != nil {
		log.Fatalf("failed to create temp dir: %v", err)
	}
	defer cleanUpDir(ctx, imgSaveDir)

	s := capture.ScreenShot{}
	if err = s.GetAllScreenshots(ctx, cfg.screen, imgSaveDir, delayBetweenShots, nFrames, uint(cfg.widthPixels)); err != nil {
		cleanUpDir(ctx, imgSaveDir) // needed os.Exit does not honor deferred calls
		log.Fatalf("failed to get screenshots: %v", err)
	}

	log.Printf("Creating animation...\n")
	if err = animate.Animate(ctx, imgSaveDir, cfg.outputDir, cfg.loop, cfg.fps); err != nil {
		cleanUpDir(ctx, imgSaveDir) // needed os.Exit does not honor deferred calls
		log.Fatalf("failed to animate: %v", err)
	}

	log.Printf("Animation saved to %s/out.gif.\n", cfg.outputDir)
	log.Printf("Cleaning up temporary files...\n")
}

// calculateDelay works out the delay in 100ths of seconds needed
// to achieve the frame rate given by fps.
func calculateDelay(fps int) (int, error) {
	if fps <= 0 || fps >= maximumAllowedFPS {
		return 0, fmt.Errorf("invalid fps provided, allowed range is 1 to 32")
	}

	const secsPart int = 100
	return secsPart / fps, nil
}

func cleanUpDir(ctx context.Context, dir string) {
	if err := os.RemoveAll(dir); err != nil {
		log.Printf("failed to clean up temp files: %v", err)
	}
}
