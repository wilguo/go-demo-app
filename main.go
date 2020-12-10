package main

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gorilla/mux/otelmux"
	"go.opentelemetry.io/contrib/propagators/aws/xray"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp"

	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

var tracer = otel.Tracer("demo-app")

func main() {

	initTracer()

	// Create a new HTTP router to handle incoming client requests
	r := mux.NewRouter()

	r.Use(otelmux.Middleware("my-server"))

	// When client makes GET request to /hello-world
	// handler() will execute
	r.HandleFunc("/hello-world", handler).Methods(http.MethodGet)

	// Start the server and listen on localhost:8080
	http.ListenAndServe(":8080", r)

}

// Function for handling the /hello-world endpoint
func handler(w http.ResponseWriter, r *http.Request) {

	// Set the header content-type and return hello world
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("hello world")

}

func initTracer() {

	// Create new OTLP Exporter struct
	exporter, err := otlp.NewExporter(
		otlp.WithInsecure(),
		otlp.WithAddress("localhost:30080"),
	)
	if err != nil {
		// Handle error here...
	}

	// AlwaysSample() returns a Sampler that samples every trace.
	// Be careful about using this sampler in a production application with
	// significant traffic: a new trace will be started and exported for every request.
	cfg := sdktrace.Config{
		DefaultSampler: sdktrace.AlwaysSample(),
	}

	// A custom ID Generator to generate traceIDs that conform to
	// AWS X-Ray traceID format
	idg := xray.NewIDGenerator()

	// Create a new TraceProvider struct passing in the config, the exporter
	// and the ID Generator we want to use for our tracing
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithConfig(cfg),
		sdktrace.WithSyncer(exporter),
		sdktrace.WithIDGenerator(idg),
	)

	// Set the traceprovider and the propagator we want to use
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(xray.Propagator{})
}
