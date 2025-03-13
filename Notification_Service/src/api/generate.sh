PROTO_DIR=./proto

protoc --go_out=. --go-grpc_out=. --grpc-gateway_out=. --openapiv2_out=. \
  --proto_path=$PROTO_DIR $PROTO_DIR/notification.proto
