package LoadBalancerGo

import (
	"log"
	"net/http"
	"net/http/httputil"
	"io/ioutil"
	"sync"
	"net/url"
	"encoding/json"
	"os"
	"fmt"
)

var mu sync.Mutex
var idx int = 0

type LoadBalancer struct {
	Name			string
	Port			string
	EndPoints []EndPoint
}

func New(name string, port string) (*LoadBalancer) {

	lb := LoadBalancer{ Name: name, Port: port }
	lb.loadServerEndPoints()
	return &lb
}

func (lb *LoadBalancer)RoudRobin(w http.ResponseWriter, r *http.Request) {

	maxLen := len(lb.EndPoints)

	mu.Lock()

	if currentEndPoint := lb.EndPoints[idx%maxLen]; currentEndPoint.IsAlive() == false {
		idx++
	}

	targetURL, err := url.Parse(lb.EndPoints[idx%maxLen].URL)

	if err != nil {
			log.Fatal(err.Error())
	}

	idx++
	mu.Unlock()

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)
	reverseProxy.ServeHTTP(w, r)
}

func (lb *LoadBalancer) loadServerEndPoints() {

	path, err := os.Getwd()

	if err != nil {
    log.Println(err)
	}

	data, err := ioutil.ReadFile(path + "/LoadBalancerGo/" + "/server_config.json")

	if err != nil {
		log.Fatal(err.Error())
	}
	json.Unmarshal(data, lb)
}


func (lb *LoadBalancer) Serve() {
	
	http.HandleFunc("/", lb.RoudRobin)
	http.ListenAndServe(":" + lb.Port, nil)
}

func (lb *LoadBalancer) ServeBackend(name string, port string) {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Server name: %v\n ", name)
		fmt.Fprintf(w, "Response header: %v\n", r.Header)
	}))
	http.ListenAndServe(":" + port, mux)
}