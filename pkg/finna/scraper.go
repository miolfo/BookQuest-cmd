package finna

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
)

type AvailabilityResult struct {
	FinnaId   string
	Available bool
}

func AreBooksAvailable(finnaId []string, progressCallback func(doneCount int)) []AvailabilityResult {
	var res []AvailabilityResult

	//User agent spoofing required for headless chrome to correctly make requests from javascript
	opts := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3830.0 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),
		chromedp.Headless,
	}
	actx, acancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer acancel()
	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()

	for i, id := range finnaId {
		available := IsBookAvailable(id, ctx)
		res = append(res, AvailabilityResult{
			FinnaId:   id,
			Available: available,
		})
		progressCallback(i + 1)
	}
	return res
}

func IsBookAvailable(finnaId string, ctx context.Context) bool {

	var nodes []*cdp.Node
	url := fmt.Sprintf("https://www.finna.fi/Record/%s", finnaId)
	log.Println("Checking availability from " + url)
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`.holdings-title`),
		chromedp.Nodes(`.holdings-details > span`, &nodes, chromedp.AtLeast(0)),
	)
	if err != nil {
		log.Fatal(err)
	}

	isAvailable := false
	for _, n := range nodes {
		var res string
		err2 := chromedp.Run(ctx,
			chromedp.TextContent(n.FullXPath(), &res))
		if err2 != nil {
			log.Println(err2)
			continue
		}
		if strings.Contains(res, "saatavissa") {
			isAvailable = true
		}
	}

	return isAvailable
}
