proto1 :
			 protoc -I . --go_out ./pkg/ --go_opt=paths=source_relative --go-grpc_out ./pkg/ --go-grpc_opt=paths=source_relative api/proto/v1/*.proto

proto2 :
			 protoc -I . --grpc-gateway_out ./pkg/ --grpc-gateway_opt=logtostderr=true,paths=source_relative api/proto/v1/*.proto