package main

import (
	"image/jpeg"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	catption "github.com/nlepage/catption/lib"
)

func init() {
	cmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start catption HTTP server",
	RunE: func(_ *cobra.Command, args []string) error {
		r := mux.NewRouter()

		r.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
			if _, err := res.Write([]byte(`
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

								const res = await fetch("/catption", { method: 'POST', body })
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
			`)); err != nil {
				panic(err)
			}
		})

		r.HandleFunc("/catption", func(res http.ResponseWriter, req *http.Request) {
			file, info, err := req.FormFile("file")
			if err != nil {
				panic(err)
			}

			cType := info.Header.Get("content-type")
			var cat *catption.Catption

			switch cType {
			case "image/jpeg":
				cat, err = catption.ReadJPG(file)
			default:
				res.WriteHeader(400)
				return
			}

			if err != nil {
				panic(err)
			}

			cat.Top = req.FormValue("top")
			cat.Bottom = req.FormValue("bottom")

			img, err := cat.Image()
			if err != nil {
				panic(err)
			}

			res.Header().Add("content-type", cType)
			if err = jpeg.Encode(res, img, nil); err != nil {
				panic(err)
			}
		})

		return http.ListenAndServe(":8888", r)
	},
}
