module example/hello

go 1.15

require (
	github.com/golang/protobuf v1.5.2
	github.com/mike-zeng/pigkit/rpc v1.0.0
)
replace github.com/mike-zeng/pigkit/rpc v1.0.0 => ../../../../rpc
