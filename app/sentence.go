package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type Config struct {
	ageService  string
	nameService string
}

var config Config

var (
	httpReqs = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "sentence_requests_total",
			Help: "Number of requests",
		},
		[]string{"type"},
	)
)

func handler(httpReqs *prometheus.CounterVec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := makeRequest(config.nameService)
		age := makeRequest(config.ageService)

		fmt.Fprintf(w, "%s is %s years ", name, age)

		m := httpReqs.WithLabelValues("sentence")
		m.Inc()
	}
}

func makeRequest(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	return string(body)
}

func main() {
	flag.StringVar(&config.ageService, "age-service", "http://127.0.0.1:8080", "age service")
	flag.StringVar(&config.nameService, "name-service", "http://127.0.0.1:8080", "name service")
	flag.Parse()

	prometheus.MustRegister(httpReqs)

	http.HandleFunc("/", handler(httpReqs))
	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
