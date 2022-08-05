PROTO_DST_DIR = pkg/pb
PROTO_SRC_DIR = ${PROTO_DST_DIR}/proto
PROTO_THIRD_PARTY = C:/Users/"Dmitry Frolov"/go/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.16.0/third_party/googleapis

proto:
	protoc -I=${PROTO_THIRD_PARTY} \
	-I=${PROTO_SRC_DIR} --go_out=${PROTO_DST_DIR} ${PROTO_SRC_DIR}/*.proto \
	--go-grpc_out ${PROTO_DST_DIR} --go-grpc_opt paths=source_relative