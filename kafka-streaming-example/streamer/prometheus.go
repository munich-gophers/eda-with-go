package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Define Prometheus metrics
var (
	processedMessages = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "processed_messages_total",
			Help: "Total number of processed messages",
		},
	)
	errorMessages = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "error_messages_total",
			Help: "Total number of failed message processing attempts",
		},
	)
	messageProcessingTime = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Name:    "message_processing_duration_seconds",
			Help:    "Histogram of message processing times",
			Buckets: prometheus.DefBuckets,
		},
	)
)

func init() {
	// Register metrics
	prometheus.MustRegister(processedMessages, errorMessages, messageProcessingTime)
}

func startPrometheusServer() {
	http.Handle("/metrics", promhttp.Handler())
	log.Println("Starting Prometheus metrics server at :9091")
	go func() {
		if err := http.ListenAndServe(":9091", nil); err != nil {
			log.Fatalf("Error starting Prometheus server: %v", err)
		}
	}()
}
