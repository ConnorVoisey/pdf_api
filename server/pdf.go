package server

import (
	"context"

	"github.com/chromedp/cdproto/page"
)

type (
	PageOptions struct {
		Html                    string   `json:"html" example:"<h1>Hello</h1>" doc:"Body of the PDF"`
		PrintBackground         *bool    `json:"print_background" required:"false" default:"true" doc:"Whether to print background graphics"`
		DisplayHeaderFooter     *bool    `json:"display_header_footer" required:"false" default:"false" doc:"Whether to display header and footer. If no template is provided, but this field is true, it will use the default chrome header and footer."`
		HeaderTemplate          *string  `json:"header_template" required:"false" doc:"HTML template for the PDF header"`
		FooterTemplate          *string  `json:"footer_template" required:"false" doc:"HTML template for the PDF footer"`
		GenerateDocumentOutline *bool    `json:"generate_document_outline" required:"false" default:"false" doc:"Generate a document outline (bookmarks)"`
		GenerateTaggedPDF       *bool    `json:"generate_tagged_pdf" required:"false" default:"false" doc:"Generate a tagged (accessible) PDF"`
		Landscape               *bool    `json:"landscape" required:"false" default:"false" doc:"Paper orientation: landscape if true, portrait if false"`
		MarginTop               *float64 `json:"margin_top" required:"false" default:"0" doc:"Top margin in inches"`
		MarginBottom            *float64 `json:"margin_bottom" required:"false" default:"0" doc:"Bottom margin in inches"`
		MarginLeft              *float64 `json:"margin_left" required:"false" default:"0" doc:"Left margin in inches"`
		MarginRight             *float64 `json:"margin_right" required:"false" default:"0" doc:"Right margin in inches"`
		PageRanges              *string  `json:"page_ranges" required:"false" doc:"Paper page ranges to print, e.g. '1-5, 8, 11-13'"`
		PaperHeight             *float64 `json:"paper_height" required:"false" doc:"Paper height in inches"`
		PaperWidth              *float64 `json:"paper_width" required:"false" doc:"Paper width in inches"`
		PreferCSSPageSize       *bool    `json:"prefer_css_page_size" required:"false" doc:"Whether to prefer page size defined by CSS"`
		Scale                   *float64 `json:"scale" required:"false" default:"1" doc:"Scale of the webpage rendering (1 = 100%)"`
	}
)

func printPdf(ctx context.Context, opts *PageOptions) ([]byte, error) {
	builder := page.PrintToPDF()
	if opts != nil {
		if opts.PrintBackground != nil {
			builder = builder.WithPrintBackground(*opts.PrintBackground)
		}
		if opts.DisplayHeaderFooter != nil {
			builder = builder.WithDisplayHeaderFooter(*opts.DisplayHeaderFooter)
		}
		if opts.HeaderTemplate != nil {
			builder = builder.WithHeaderTemplate(*opts.HeaderTemplate)
		}
		if opts.FooterTemplate != nil {
			builder = builder.WithFooterTemplate(*opts.FooterTemplate)
		}
		if opts.GenerateDocumentOutline != nil {
			builder = builder.WithGenerateDocumentOutline(*opts.GenerateDocumentOutline)
		}
		if opts.GenerateTaggedPDF != nil {
			builder = builder.WithGenerateTaggedPDF(*opts.GenerateTaggedPDF)
		}
		if opts.Landscape != nil {
			builder = builder.WithLandscape(*opts.Landscape)
		}
		if opts.MarginTop != nil {
			builder = builder.WithMarginTop(*opts.MarginTop)
		}
		if opts.MarginBottom != nil {
			builder = builder.WithMarginBottom(*opts.MarginBottom)
		}
		if opts.MarginLeft != nil {
			builder = builder.WithMarginLeft(*opts.MarginLeft)
		}
		if opts.MarginRight != nil {
			builder = builder.WithMarginRight(*opts.MarginRight)
		}
		if opts.PageRanges != nil {
			builder = builder.WithPageRanges(*opts.PageRanges)
		}
		if opts.PaperHeight != nil {
			builder = builder.WithPaperHeight(*opts.PaperHeight)
		}
		if opts.PaperWidth != nil {
			builder = builder.WithPaperWidth(*opts.PaperWidth)
		}
		if opts.PreferCSSPageSize != nil {
			builder = builder.WithPreferCSSPageSize(*opts.PreferCSSPageSize)
		}
		if opts.Scale != nil {
			builder = builder.WithScale(*opts.Scale)
		}
	}
	buf, _, err := builder.
		Do(ctx)
	if err != nil {
		return nil, err
	}
	return buf, nil
}
