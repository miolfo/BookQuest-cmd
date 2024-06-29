package util

import (
	"context"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"log"
	"strings"
)

type ping struct {
	Message string
}

type availabilityResponse struct {
	//Possible types are 'ORDERED' | 'UNAVAILABLE' | 'AVAILABLE' | 'WAITING' | 'IN_TRANSIT' | 'UNKNOWN'
	Statuses []string
}

type AvailabilityResult struct {
	FinnaId   string
	Available bool
}

func AreBooksAvailable(finnaId []string) []AvailabilityResult {
	var res []AvailabilityResult
	//User agent spoofing required for headless chrome to correctly make requests from javascript
	/*opts := []chromedp.ExecAllocatorOption{
		chromedp.UserAgent("Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3830.0 Safari/537.36"),
		chromedp.WindowSize(1920, 1080),
		chromedp.Headless,
	}
	actx, acancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer acancel()
	ctx, cancel := chromedp.NewContext(actx)
	defer cancel()*/
	for _, id := range finnaId {
		res = append(res, AvailabilityResult{
			FinnaId:   id,
			Available: false,
		})
	}
	return res
}

func IsBookAvailable(finnaId string, ctx context.Context) bool {

	// run task list
	var title string
	var html string
	var nodes []*cdp.Node
	var text string
	err := chromedp.Run(ctx,
		chromedp.Navigate(`https://www.finna.fi/Record/helmet.2527262`),
		chromedp.Title(&title),
		chromedp.WaitVisible(`.holdings-title`),
		//chromedp.InnerHTML(`.holdings-title`, &html, chromedp.NodeVisible),
		chromedp.Nodes(`.holdings-details > span`, &nodes),
		//chromedp.Nodes(`.holdings-details`, childNodes)
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(strings.TrimSpace(title))
	log.Println(strings.TrimSpace(html))
	log.Println(strings.TrimSpace(text))

	for _, n := range nodes {
		log.Println(n.FullXPath())
		var res string
		err2 := chromedp.Run(ctx,
			chromedp.TextContent(n.FullXPath(), &res))
		if err2 != nil {
			log.Println(err2)
			continue
		}
		log.Println(res)
	}

	return false
}
