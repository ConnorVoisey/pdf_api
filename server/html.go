package server

import (
	"context"

	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/chromedp"
	"github.com/rs/zerolog/log"
)

type HtmlPdfReq struct {
	Body struct {
		Html        string       `json:"html" example:"<h1>Hello</h1>" doc:"Body of the PDF"`
		PageOptions *PageOptions `json:"page_options" required:"false" doc:"Configuration for the printing to PDF"`
	}
}

func htmlToPdf(ctx context.Context, input *HtmlPdfReq) (*ByteRes, error) {
	log.Trace().Any("inputBody", input.Body).Msg("/html")

	resp := &ByteRes{}
	resp.ContentType = "application/pdf"

	ctx, cancel := chromedp.NewContext(ctx, chromedp.WithLogf(log.Printf))
	defer cancel()

	if err := chromedp.Run(ctx,
		chromedp.Navigate("about:blank"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			frameTree, err := page.GetFrameTree().Do(ctx)
			if err != nil {
				return err
			}

			return page.SetDocumentContent(frameTree.Frame.ID, input.Body.Html).Do(ctx)
		}),
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
