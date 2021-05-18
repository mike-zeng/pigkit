package codec

import (
	"context"
	"encoding/json"
	"github.com/opentracing/opentracing-go"
)

type Metadata map[string][]byte

type MetadataStruct struct {

}

func GetMetadataFromCtx(ctx context.Context) Metadata {
	if md, ok := ctx.Value(MetadataStruct{}).(Metadata); ok {
		return md
	}
	md := make(map[string][]byte)
	WithMetadata(ctx, md)
	return md
}

func WithMetadata(ctx context.Context, metadata map[string][]byte) context.Context{
	return context.WithValue(ctx, MetadataStruct{}, Metadata(metadata))
}

func WithCarrier(metadata map[string][]byte,carrier *opentracing.TextMapCarrier)  {
	if metadata == nil {
		metadata = make(map[string][]byte)
	}
	if marshal, err := json.Marshal(carrier);err==nil {
		metadata["carrier"] = marshal
	}
}

func GetCarrier(metadata map[string][]byte)*opentracing.TextMapCarrier  {
	if metadata == nil {
		return nil
	}
	unmarshal := metadata["carrier"]
	if unmarshal == nil {
		return nil
	}
	var carrier = opentracing.TextMapCarrier{}
	err := json.Unmarshal(unmarshal, &carrier)
	if err != nil {
		return nil
	}
	return &carrier
}