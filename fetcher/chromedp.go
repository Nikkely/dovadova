package fetcher

import (
	"context"
	"log"
	"time"

	"github.com/chromedp/chromedp"
)

func runChromedp(headless bool) error {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
	)

	if !headless {
		opts = append(opts,
			chromedp.Flag("headless", false),
			// Like in Puppeteer.
			chromedp.Flag("hide-scrollbars", false),
			chromedp.Flag("mute-audio", false),
		)
	}

	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	// also set up a custom logger
	taskCtx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(logf))
	defer cancel()

	ctx, cancel := chromedp.NewContext(
		taskCtx,
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 15*time.Second)
	defer cancel()
	// navigate to a page, wait for an element, click
	var example string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://pkg.go.dev/time`),
		// wait for footer element is visible (ie, page is loaded)
		chromedp.WaitVisible(`body > footer`),
		// find and click "Example" link
		chromedp.Click(`#example-After`, chromedp.NodeVisible),
		// retrieve the text of the textarea
		chromedp.Value(`#example-After textarea`, &example),
		// chromedp.ScrollIntoView()
		chromedp.Sleep(1*time.Second),
	)
	if err != nil {
		return err
	}
	log.Printf("Go's time.After example:\n%s", example)
	return nil
}

func logf(s string, v ...interface{}) {
	log.Printf(s, v...)
}
