package fetcher

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

const (
	listEndpoint = `https://shadowverse-evolve.com/cardlist/cardsearch/?card_name=&class%5B0%5D=all&expansion_name=&cost%5B0%5D=all&card_kind%5B0%5D=all&rare%5B0%5D=all&power_from=&power_to=&hp_from=&hp_to=&type=&ability=&keyword=&view=image`
)

func GetCardList(headless bool) ([]string, error) {
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
	ctx, cancel := chromedp.NewContext(allocCtx, chromedp.WithLogf(logf))
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	var resultNumRaw string
	if err := chromedp.Run(ctx,
		chromedp.Navigate(listEndpoint),
		chromedp.WaitVisible(`span.num`),
		chromedp.Text(`span.num`, &resultNumRaw),
	); err != nil {
		return nil, err
	}

	resultNum, err := strconv.Atoi(resultNumRaw)
	if err != nil {
		return nil, err
	}
	as, err := loopRun(ctx, resultNum)
	if err != nil {
		return nil, err
	}

	links := make([]string, 0, len(as))
	for _, a := range as {
		if href, ok := a["href"]; !ok {
			log.Printf("no href attr found. unxepected <a>")
		} else {
			links = append(links, "https://shadowverse-evolve.com"+href)
		}
	}
	return links, nil
}

func logf(s string, v ...any) {
	log.Printf(s, v...)
}

func loopRun(ctx context.Context, num int) ([]map[string]string, error) {
	if err := chromedp.Run(ctx,
		chromedp.Navigate(listEndpoint),
	); err != nil {
		return nil, err
	}

	// ctx has been set timeout, so it is not infinite loop
	for {
		var as []map[string]string
		if err := chromedp.Run(ctx,
			chromedp.ScrollIntoView(`footer`),
			// HACK: need a little above the footer to load
			chromedp.ActionFunc(func(ctx context.Context) error {
				_, exp, err := runtime.Evaluate(`window.scrollBy(0,-10);`).Do(ctx)
				if err != nil {
					return err
				}
				if exp != nil {
					return exp
				}
				return nil
			}),
			chromedp.Sleep(1*time.Second),
			chromedp.AttributesAll(`ul.cardlist-Result_List > li > a`, &as),
		); err != nil {
			return nil, err
		}

		if num <= len(as) {
			return as, nil
		}
	}
}
