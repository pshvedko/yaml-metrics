package main

import (
	"context"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/pshvedko/yaml-metrics/collector"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	r := prometheus.NewRegistry()
	for t, f := range map[string]func(name string) (*collector.Collector, error){
		"yaml": collector.NewCollectorYaml,
		"json": collector.NewCollectorJson,
	} {
		c, err := f("metrics." + t)
		if err != nil {
			log.Fatal(err)
		}
		c.Map("currencies", "", "value", prometheus.GaugeValue, []string{"name"}, prometheus.Labels{"from": t})
		err = r.Register(c)
		if err != nil {
			log.Fatal(err)
		}
	}
	h := mux.NewRouter()
	h.Handle("/metrics", handlers.LoggingHandler(os.Stdout, promhttp.HandlerFor(r, promhttp.HandlerOpts{})))
	s := http.Server{
		Addr:    ":8080",
		Handler: h,
	}
	errs := make(chan error)
	defer close(errs)
	go func() {
		errs <- s.ListenAndServe()
	}()
	done := make(chan os.Signal)
	defer close(done)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-done:
			err := s.Shutdown(context.TODO())
			if err != nil {
				log.Fatal(err)
			}
		case err := <-errs:
			if err != nil && err != http.ErrServerClosed {
				log.Fatal(err)
			}
			signal.Stop(done)
			return
		}
	}
}
