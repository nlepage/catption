package main

import (
	"fmt"
	"html/template"
	"image/jpeg"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	catption "github.com/nlepage/catption/lib"
)

func init() {
	cmd.AddCommand(httpCmd)
}

var httpHandler http.Handler

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start catption HTTP server",
	PreRun: func(_ *cobra.Command, args []string) {
		r := mux.NewRouter()

		r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
			tmpl, err := template.New("index").Parse(`
				<!DOCTYPE html>
				<html lang="en">
					<head>
						<meta charset="utf-8">
						<title>Catption</title>
						<link rel="icon" href="data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.90em%22 font-size=%2290%22>🐱</text></svg>">
					</head>
					<body>
						<h1>Catption</h1>

						<form id="catption-form">
							<p>
								<input name="file" type="file" required>
							</p>
							<p>
								<label>Top <input name="top" type="text"></label>
							</p>
							<p>
								<label>Bottom <input name="bottom" type="text"></label>
							</p>
							<p>
								<label>Font size <input name="font-size" type="number" placeholder="{{ .DefaultFontSize }}"></label>
							</p>
							<p>
								<label>Margin <input name="margin" type="number" placeholder="{{ .DefaultMargin }}"></label>
							</p>
							<p>
								<label>Image size <input name="size" type="number" placeholder="{{ .DefaultSize }}"></label>
							</p>
							<p>
								<button type="submit">Catption</button>
							</p>
						</form>

						<p id="catption"></p>

						<script>
							const form = document.querySelector('#catption-form')
							form.addEventListener('submit', catption)

							async function catption(e) {
								e.preventDefault()

								const body = new FormData(form)

								const res = await fetch("catption", { method: 'POST', body })
								if (!res.ok) throw Error(res.status + ' ' + res.statusText)

								const src = URL.createObjectURL(await res.blob())

								const container = document.querySelector('#catption')
								while (container.lastChild) container.removeChild(container.lastChild)
								
								const img = document.createElement('img')
								img.src = src
								container.appendChild(img)
							}
						</script>
					</body>
				</html>
			`)
			if err != nil {
				panic(err)
			}

			if err = tmpl.Execute(res, map[string]interface{}{
				"DefaultFontSize": catption.DefaultFontSize,
				"DefaultMargin":   catption.DefaultMargin,
				"DefaultSize":     catption.DefaultSize,
			}); err != nil {
				panic(err)
			}
		})

		r.HandleFunc("/catption", func(res http.ResponseWriter, req *http.Request) {
			fmt.Println("/catption 1")

			file, info, err := req.FormFile("file")
			if err != nil {
				panic(err)
			}

			fmt.Println("/catption 2")

			cType := info.Header.Get("content-type")
			var cat *catption.Catption

			switch cType {
			case "image/jpeg":
				fmt.Println("/catption 3")
				cat, err = catption.ReadJPG(file)
			default:
				fmt.Println("/catption 4")
				res.WriteHeader(400)
				return
			}

			if err != nil {
				panic(err)
			}

			fmt.Println("/catption 5")

			cat.Top = req.FormValue("top")
			cat.Bottom = req.FormValue("bottom")
			if sFontSize := req.FormValue("font-size"); sFontSize != "" {
				cat.FontSize, _ = strconv.ParseFloat(sFontSize, 64)
			}
			if sMargin := req.FormValue("margin"); sMargin != "" {
				cat.Margin, _ = strconv.ParseFloat(sMargin, 64)
			}
			if sSize := req.FormValue("size"); sSize != "" {
				cat.Size, _ = strconv.ParseFloat(sSize, 64)
			}

			fmt.Println("/catption 6")

			img, err := cat.Image()
			if err != nil {
				panic(err)
			}

			fmt.Println("/catption 7")

			res.Header().Add("content-type", cType)
			if err = jpeg.Encode(res, img, nil); err != nil {
				panic(err)
			}

			fmt.Println("/catption 8")
		})

		httpHandler = r
	},
}
