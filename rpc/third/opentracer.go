package third

import (
	"context"
	"github.com/mike-zeng/pigkit/rpc/codec"
	log2 "github.com/mike-zeng/pigkit/rpc/log"

	"github.com/mike-zeng/pigkit/rpc/config"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	jaegerConfig "github.com/uber/jaeger-client-go/config"
	"log"
	"sync"
)

type tracerClient struct {

}

var client *tracerClient
var jaegerInit sync.Once

func GetTracerClient() *tracerClient {
	jaegerInit.Do(func() {
		pigConf := config.GetConfig().PigServer
		cfg := &jaegerConfig.Configuration{
			Sampler : &jaegerConfig.SamplerConfig{
				Type : "const",  // Fixed sampling
				Param : 1,       // 1= full sampling, 0= no sampling
			},
			Reporter : &jaegerConfig.ReporterConfig{
				LogSpans: true,
				CollectorEndpoint: pigConf.JaegerConf.CollectorEndpoint,
			},
			ServiceName : pigConf.ServiceName,
		}
		newTracer, _, err := cfg.NewTracer()
		if err != nil {
			log.Fatalln(err)
		}
		opentracing.SetGlobalTracer(newTracer)
		client = &tracerClient{}
	})
	return client
}


func (cli *tracerClient) StartSpanClient(ctx context.Context,methodName string)(context.Context,opentracing.Span){
	globalTracer := opentracing.GlobalTracer()

	rootSpan := opentracing.SpanFromContext(ctx)
	var rootCtx opentracing.SpanContext
	if rootSpan != nil {
		rootCtx = rootSpan.Context()
	}

	clientSpan := globalTracer.StartSpan(
		methodName,
		opentracing.ChildOf(rootCtx),
		ext.SpanKindRPCClient,
	)
	mdCarrier := &opentracing.TextMapCarrier{}
	if err := globalTracer.Inject(clientSpan.Context(), opentracing.TextMap, mdCarrier); err != nil {
		log2.DefaultLog.ERROR("tracer inject err:",err)
		return nil,nil
	}
	metadata := codec.GetMetadataFromCtx(ctx)
	codec.WithCarrier(metadata,mdCarrier)
	newCtx := codec.WithMetadata(ctx, metadata)
	return newCtx,clientSpan
}

func (cli *tracerClient) StartSpanServer(ctx context.Context,methodName string)  (context.Context,opentracing.Span){
	// 拿到上游传递过来的东西
	globalTracer := opentracing.GlobalTracer()
	metaData := codec.GetMetadataFromCtx(ctx)
	mdCarrier := codec.GetCarrier(metaData)
	spanContext, err := globalTracer.Extract(opentracing.TextMap, mdCarrier)
	if err != nil{
		log2.DefaultLog.ERROR("extract inject err:",err)
		return nil,nil
	}
	span := globalTracer.StartSpan(methodName, ext.RPCServerOption(spanContext), ext.SpanKindRPCServer)

	// span set to ctx
	withSpanCtx := opentracing.ContextWithSpan(ctx, span)
	return withSpanCtx,span
}