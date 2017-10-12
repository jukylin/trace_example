package components

import (
	"github.com/astaxie/beego"
	"github.com/uber/jaeger-client-go/config"
	"github.com/opentracing/opentracing-go"
	"time"
)

var Trace opentracing.Tracer

func InitTrace() opentracing.Tracer {
	cfg := config.Configuration{
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            false,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:"0.0.0.0:5775",
		},
	}
	tracer, _, err := cfg.New("trace_example")
	if err != nil {
		beego.Error("cannot initialize Jaeger Tracer", err.Error())
	}

	Trace = tracer
	return tracer
}



