PROTO_DST_DIR = pkg/pb
PROTO_SRC_DIR = ${PROTO_DST_DIR}/proto

protoc:
	protoc -IC:/protoc/include -I. \
		-IC:\Users\fd239\go\pkg\mod\github.com\grpc-ecosystem\grpc-gateway@v1.16.0\third_party\googleapis \
		-IC:\Users\fd239\go\pkg\mod\github.com\envoyproxy\protoc-gen-validate@v0.6.7 \
		-I=${PROTO_SRC_DIR} --go_out=${PROTO_DST_DIR} ${PROTO_SRC_DIR}/*.proto \
		--go-grpc_out ${PROTO_DST_DIR} --go-grpc_opt paths=source_relative