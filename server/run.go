package server

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/signal"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humachi"
	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	"github.com/danielgtaylor/huma/v2/humacli"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

func Run(ctx context.Context, w io.Writer, args []string) error {
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	var api huma.API
	cli := humacli.New(func(hooks humacli.Hooks, options *Options) {
		err := Init(options)
		if err != nil {
			panic(err)
		}
		router := chi.NewMux()

		router.Use(cors.Handler(cors.Options{
			AllowedOrigins:   []string{"https://*", "http://*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link"},
			AllowCredentials: true,
			MaxAge:           300,
		}))

		config := huma.DefaultConfig("My API", "1.0.0")
		config.DocsPath = "/"
		config.Info.Title = "PDF API"
		config.Info.Description = `Generate PDF's from a rest api.

This is a simple api that wraps Chrome headless to generate PDFs. The main advantage of this is that it doesn't require you to install Chrome headless directly onto your server. The main disadvantage is that it will not allow the complete customisation options of having the headless browser directly available, however, it should offer more than enough options for most users.
Git repo can be found at https://github.com/connorvoisey/pdf_api. This includes more information such as a deployment guide.`
		api = humachi.New(router, config)

		AddRoutes(&api)

		hooks.OnStart(func() {
			log.Info().
				Int("Port", options.Port).
				Msg("Started server")
			err := http.ListenAndServe(fmt.Sprintf(":%d", options.Port), router)
			if err != nil {
				log.Err(err).Msg("Failed to listen and serve")
				panic(err)
			}
		})
	})

	cli.Root().AddCommand(&cobra.Command{
		Use:   "openapi",
		Short: "Print the OpenAPI spec",
		Run: func(cmd *cobra.Command, args []string) {
			b, _ := api.OpenAPI().MarshalJSON()
			fmt.Println(string(b))
		},
	})

	cli.Run()

	return nil
}

type ByteRes struct {
	ContentType string `header:"Content-Type"`
	Body        []byte
}

func AddRoutes(api *huma.API) {
	huma.Get(*api, "/", func(ctx context.Context, _ *struct{}) (*ByteRes, error) {
		return &ByteRes{Body: []byte(`<!doctype html>
<html>
  <head>
    <title>API Reference</title>
    <meta charset="utf-8" />
    <meta
      name="viewport"
      content="width=device-width, initial-scale=1" />
  </head>
  <body>
    <script
      id="api-reference"
      data-url="/openapi.json"></script>
    <script src="https://cdn.jsdelivr.net/npm/@scalar/api-reference"></script>
  </body>
</html>`)}, nil
	})
	huma.Register(*api, huma.Operation{
		OperationID:   "html-to-pdf",
		Method:        http.MethodPost,
		Path:          "/html",
		Summary:       "Html to PDF",
		Description:   "Generate a PDF from html. Header and footer can be specified as html templates that will be applied to each page.",
		Tags:          []string{"PDF"},
		DefaultStatus: http.StatusOK,
	}, htmlToPdf)

	huma.Register(*api, huma.Operation{
		OperationID:   "url-to-pdf",
		Method:        http.MethodPost,
		Path:          "/url-to-pdf",
		Summary:       "URL to PDF",
		Tags:          []string{"PDF"},
		DefaultStatus: http.StatusOK,
	}, urlToPdf)
}
