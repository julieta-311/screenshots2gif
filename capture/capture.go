package capture

import (
	"context"
	"fmt"
	"image"
	"image/png"
	"log"
	"os"
	"path"
	"time"

	"github.com/kbinani/screenshot"
	"github.com/nfnt/resize"
)

const dateTimeLayout = "20060102_150405"

// Screenshot is the screenshot handler
type ScreenShot struct{}

// captureRect is a wrapper for screenshot.CaptureRect.
func (s *ScreenShot) captureRect(ctx context.Context, rect image.Rectangle) (*image.RGBA, error) {
	return screenshot.CaptureRect(rect)
}

// getDisplayBounds is a wrapper for screenshot.GetDisplayBounds.
func (s *ScreenShot) getDisplayBounds(ctx context.Context, displayIndex int) image.Rectangle {
	return screenshot.GetDisplayBounds(displayIndex)
}

// numActiveDisplays is a wrapper for screenshot.NumActiveDisplays.
func (s *ScreenShot) numActiveDisplays(ctx context.Context) int {
	return screenshot.NumActiveDisplays()
}

func (s *ScreenShot) getScreensShot(
	ctx context.Context,
	screen int,
	frameNumber int,
	saveDir string,
	width uint,
) error {
	bounds := s.getDisplayBounds(ctx, screen)

	capt, err := s.captureRect(ctx, bounds)
	if err != nil {
		return fmt.Errorf("failed to capture screen: %w", err)
	}

	t := time.Now()
	fileName := fmt.Sprintf("%s_%d_%dx%d.png", t.Format(dateTimeLayout), frameNumber, bounds.Dx(), bounds.Dy())
	filePath := path.Join(saveDir, fileName)
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file %q: %w", filePath, err)
	}
	defer file.Close()

	img := image.Image(capt)
	if width > 0 {
		img = resize.Resize(width, 0, img, resize.Lanczos3)
	}

	if err := png.Encode(file, img); err != nil {
		return fmt.Errorf("failed to encode image: %w", err)
	}

	log.Printf("take %05d captured screen #%d: \"%s\"\n", frameNumber, screen, filePath)

	return nil
}

// GetAllScreenshots takes a series of nFrame screenshots with a delay between shots given by delayBetweenShots
// and saves it to saveDir.
func (s *ScreenShot) GetAllScreenshots(
	ctx context.Context,
	screen int,
	saveDir string,
	delayBetweenShots time.Duration,
	nFrames int,
	width uint,
) error {
	nDisp := s.numActiveDisplays(ctx)
	if nDisp == 0 {
		return fmt.Errorf("no screens found")
	}

	if screen > nDisp-1 || screen < 0 {
		return fmt.Errorf("invalid screen number: %d", screen)
	}

	for k := 0; k < nFrames; k++ {
		if err := s.getScreensShot(ctx, screen, k, saveDir, width); err != nil {
			return err
		}
		time.Sleep(delayBetweenShots)
	}
	return nil
}
