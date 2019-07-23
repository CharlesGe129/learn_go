package data

import (
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(Add("https://github.com/EDDYCJY"))
	}()

	http.ListenAndServe("0.0.0.0:6060", nil)
}