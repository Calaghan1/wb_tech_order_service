package server

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
		html := `<!DOCTYPE html>
		<html lang="en-US">
		<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1">
		<title>Today's Date</title>
		<script>
		let d = new Date();
		alert("Today's date is " + d);
		</script>
		</head>
		<body>
		</body>
		</html>`
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprint(w, html)
}

func StartServer() {
	http.HandleFunc("/", handler)
}