package main

import (
	"context"
	"fmt"
	"os"

	// "github.com/chromedp/cdproto/page"
	// "github.com/chromedp/chromedp"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"

	"github.com/connorvoisey/pdf_api/server"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
	}
}

func main() {
	ctx := context.Background()
	if err := server.Run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
	// // create context
	// ctx, cancel := chromedp.NewContext(context.Background())
	// defer cancel()
	//
	// // capture pdf
	// var buf []byte
	// if err := chromedp.Run(ctx, printToPDF(`https://www.google.com/`, &buf)); err != nil {
	// 	log.Fatal(err)
	// }
	//
	// if err := os.WriteFile("sample.pdf", buf, 0o644); err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("wrote sample.pdf")
	//
	//    // Create a new router & API
	// router := chi.NewMux()
	// api := humachi.New(router, huma.DefaultConfig("My API", "1.0.0"))
	//
	// // Register GET /greeting/{name}
	// huma.Register(api, huma.Operation{
	// 	OperationID: "get-greeting",
	// 	Method:      http.MethodGet,
	// 	Path:        "/greeting/{name}",
	// 	Summary:     "Get a greeting",
	// 	Description: "Get a greeting for a person by name.",
	// 	Tags:        []string{"Greetings"},
	// }, func(ctx context.Context, input *struct{
	// 	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	// }) (*GreetingOutput, error) {
	// 	resp := &GreetingOutput{}
	// 	resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
	// 	return resp, nil
	// })
	//
	// // Start the server!
	// http.ListenAndServe("127.0.0.1:8888", router)
}
