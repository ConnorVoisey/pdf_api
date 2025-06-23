package server

import (
	"context"

	"github.com/chromedp/chromedp"
	"github.com/rs/zerolog/log"
)

type UrlPdfReq struct {
	Body struct {
		Url         string       `json:"url" example:"https://shgrid.dev" doc:"Url of the website that will be converted into a pdf"`
		PageOptions *PageOptions `json:"page_options" required:"false" doc:"Configuration for the printing to PDF"`
	}
}

func urlToPdf(ctx context.Context, input *UrlPdfReq) (*ByteRes, error) {
	log.Trace().Any("inputBody", input.Body).Msg("/html")

	resp := &ByteRes{}
	resp.ContentType = "application/pdf"

	ctx, cancel := chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate(input.Body.Url),
		chromedp.WaitReady("body", chromedp.ByQuery),
		chromedp.ActionFunc(func(ctx context.Context) error {
			buf, err := printPdf(ctx, input.Body.PageOptions)
			if err != nil {
				return err
			}
			resp.Body = buf
			return nil
		}),
	); err != nil {
		log.Err(err)
		return nil, err
	}
	return resp, nil
}
