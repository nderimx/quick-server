package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":9090", nil))
}

func handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, `<!doctype html>
<html>
	<head>
		<title>Just a Server</title>
	</head>
	<body>
		<h1>Wabalaba Dub Dub!</h1>
	</body>
	<style>
		h1 {
			font-family: courier;
			color: rgb(200, 90, 8);
			margin: 90px 0px 0px 200px;
		}
	</style>
</html>`)
}
