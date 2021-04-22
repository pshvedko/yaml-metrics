package main

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pshvedko/yaml-metrics/collector"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
)

func serve(addr string, handler *http.ServeMux, signals ...os.Signal) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	server := http.Server{
		Handler: handler,
	}
	errs := make(chan error, 2)
	defer close(errs)
	go func() {
		errs <- server.Serve(listener)
	}()
	done := make(chan os.Signal)
	defer close(done)
	signal.Notify(done, signals...)
	for {
		select {
		case <-done:
			errs <- server.Shutdown(context.TODO())
		case err = <-errs:
			signal.Stop(done)
			return nil
		}
	}
}

func main() {
	yaml, err := collector.NewCollectorYaml("metrics.yaml")
	if err != nil {
		log.Fatal(err)
	}
	yaml.Map("currencies", "", "value", prometheus.GaugeValue, []string{"name"}, nil)
	r := prometheus.NewRegistry()
	err = r.Register(yaml)
	if err != nil {
		log.Fatal(err)
	}
	h := http.NewServeMux()
	h.Handle("/metrics", promhttp.HandlerFor(r, promhttp.HandlerOpts{}))
	err = serve(":8080", h)
	if err != nil {
		log.Fatal(err)
	}
}
