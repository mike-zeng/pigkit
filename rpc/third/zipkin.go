package third

import (
	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/reporter/http"
	"log"
	"sync"
)

type zipkinTracer struct {
	tracer *zipkin.Tracer
}

var tracer *zipkinTracer
var onceForZipkin sync.Once

func NewZipkinTracer()*zipkinTracer {
	onceForZipkin.Do(func() {
		zipTracer, err := zipkin.NewTracer(http.NewReporter(""))
		if err != nil {
			log.Fatalln(err)
		}
		tracer = &zipkinTracer{tracer: zipTracer}
	})
	return tracer
}