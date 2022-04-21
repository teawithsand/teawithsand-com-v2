package webapp

import (
	_ "expvar"
	"log"
	"net"
	"net/http"

	"github.com/teawithsand/webpage/util"
)

// Runs Pprof HTTP server on 6060 port.
// Note: for now it ignores initialization errors.
func runpprof() {
	go func() {
		log.Println("PProf handler exitted with error", http.ListenAndServe("0.0.0.0:6060", nil))
	}()
}

func runHttp(dic util.DI) (err error) {
	runpprof()

	config := dic.Get(ConfigDI).(*Config)
	server := dic.Get(HTTPServerDI).(*http.Server)
	l, err := net.Listen("tcp", config.ListenAddress)
	if err != nil {
		return
	}

	log.Printf("Serving HTTP started at %s\n", config.ListenAddress)
	err = server.Serve(l)
	return
}

func runHTTPDev(dic util.DI) (err error) {
	runpprof()

	config := dic.Get(ConfigDI).(*Config)
	server := dic.Get(HTTPServerDI).(*http.Server)
	l, err := net.Listen("tcp", config.ListenAddress)
	if err != nil {
		return
	}

	log.Printf("Serving HTTP started at %s\n", config.ListenAddress)
	err = server.Serve(l)
	return
}
