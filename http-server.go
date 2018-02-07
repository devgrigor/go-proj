package main 

import ("net/http"; "io")

func handle(res http.ResponseWriter, req *http.Request) {
	res.Header().Set(
		"Content-type",
		"text/html",
	)
	io.WriteString(
		res,
		`<doctype html>
<html>
	<head>
		<title>This is go server</title>
	</head>
	<body>
		This is sparta
	</body>
</html>`,

	)
}

func main() {
	http.Handle("/sub/",
		http.StripPrefix(
			"/sub/",
			http.FileServer(http.Dir("sub")),
		))
	http.HandleFunc("/", handle)
	http.ListenAndServe(":8000", nil)
}