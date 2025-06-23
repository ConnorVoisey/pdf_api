# PDF API

This is a simple api that wraps Chrome headless to generate PDFs. The main advantage of this is that it doesn't require you to install Chrome headless directly onto your server. The main disadvantage is that it will not allow the complete customisation options of having the headless browser directly available, however, it should offer more than enough options for most users.

There is an openapi spec that can be viewed at the base route (`/`), alternative view the demo.
A demo is available at [https://pdf.shgrid.dev](https://pdf.shgrid.dev)

There are two routes:

 - `/html`, takes in a html string and convert this into a pdf.
 - `/url`, takes in a url string, and convert this website into a pdf. The pdf api will need to be able to connect to this url.

## Deployment

Deployment is recommended through docker. There is a prebuilt container at `ghcr.io/connorvoisey/pdf_api`. Simply run this and expose port 3000, e.g.

```sh
dk run ghcr.io/connorvoisey/pdf_api:latest -p 3000:3000
```
