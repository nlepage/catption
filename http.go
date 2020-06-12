package main

import (
	"net/http"

	"github.com/spf13/cobra"
)

func init() {
	cmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Start catption HTTP server",
	RunE: func(_ *cobra.Command, args []string) error {
		http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte(`
				<!DOCTYPE html>
				<html lang="en">
					<head>
						<meta charset="utf-8">
						<title>Catption</title>
						<link rel="icon" href="data:image/svg+xml,<svg xmlns=%22http://www.w3.org/2000/svg%22 viewBox=%220 0 100 100%22><text y=%22.90em%22 font-size=%2290%22>🐱</text></svg>">
					</head>
					<body>
						<h1>Catption</h1>
						<form>
							<input type="file">
						</form>
					</body>
				</html>
			`))
		})

		return http.ListenAndServe(":8888", nil)
	},
}
