// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.12.4
// source: api/proto/v1/parkingservice.proto

package v1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_api_proto_v1_parkingservice_proto protoreflect.FileDescriptor

var file_api_proto_v1_parkingservice_proto_rawDesc = []byte{
	0x0a, 0x21, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x70,
	0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x02, 0x76, 0x31, 0x1a, 0x1a, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0xd2, 0x02, 0x0a, 0x0e, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x41, 0x64, 0x64, 0x6e, 0x65, 0x77,
	0x73, 0x6c, 0x6f, 0x74, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x6e, 0x65, 0x77,
	0x73, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x0a, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x15, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x73, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c,
	0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0a,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x0f, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0c, 0x2e, 0x76, 0x31,
	0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37, 0x0a, 0x0b, 0x44, 0x69, 0x73,
	0x70, 0x6c, 0x61, 0x79, 0x73, 0x6c, 0x6f, 0x74, 0x12, 0x0f, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x6c,
	0x6f, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x69, 0x73, 0x70, 0x6c, 0x61, 0x79, 0x73, 0x6c, 0x6f, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x2b, 0x0a, 0x0e, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x12, 0x0b, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x70, 0x61, 0x72, 0x61,
	0x6d, 0x1a, 0x0c, 0x2e, 0x76, 0x31, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x47, 0x0a, 0x0e, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72,
	0x79, 0x12, 0x19, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x61, 0x72, 0x6b, 0x69, 0x6e, 0x67, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x3b, 0x76, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_api_proto_v1_parkingservice_proto_goTypes = []interface{}{
	(*AddnewslotRequest)(nil),      // 0: v1.AddnewslotRequest
	(*UpdateslotRequest)(nil),      // 1: v1.UpdateslotRequest
	(*SlotRequest)(nil),            // 2: v1.SlotRequest
	(*Noparam)(nil),                // 3: v1.Noparam
	(*ParkingHistoryRequest)(nil),  // 4: v1.ParkingHistoryRequest
	(*Response)(nil),               // 5: v1.Response
	(*DisplayslotResponse)(nil),    // 6: v1.DisplayslotResponse
	(*ParkingDetailsResponse)(nil), // 7: v1.ParkingDetailsResponse
}
var file_api_proto_v1_parkingservice_proto_depIdxs = []int32{
	0, // 0: v1.ParkingService.Addnewslot:input_type -> v1.AddnewslotRequest
	1, // 1: v1.ParkingService.Updateslot:input_type -> v1.UpdateslotRequest
	2, // 2: v1.ParkingService.Deleteslot:input_type -> v1.SlotRequest
	2, // 3: v1.ParkingService.Displayslot:input_type -> v1.SlotRequest
	3, // 4: v1.ParkingService.ParkingDetails:input_type -> v1.Noparam
	4, // 5: v1.ParkingService.ParkingHistory:input_type -> v1.ParkingHistoryRequest
	5, // 6: v1.ParkingService.Addnewslot:output_type -> v1.Response
	5, // 7: v1.ParkingService.Updateslot:output_type -> v1.Response
	5, // 8: v1.ParkingService.Deleteslot:output_type -> v1.Response
	6, // 9: v1.ParkingService.Displayslot:output_type -> v1.DisplayslotResponse
	5, // 10: v1.ParkingService.ParkingDetails:output_type -> v1.Response
	7, // 11: v1.ParkingService.ParkingHistory:output_type -> v1.ParkingDetailsResponse
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_proto_v1_parkingservice_proto_init() }
func file_api_proto_v1_parkingservice_proto_init() {
	if File_api_proto_v1_parkingservice_proto != nil {
		return
	}
	file_api_proto_v1_service_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_proto_v1_parkingservice_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_proto_v1_parkingservice_proto_goTypes,
		DependencyIndexes: file_api_proto_v1_parkingservice_proto_depIdxs,
	}.Build()
	File_api_proto_v1_parkingservice_proto = out.File
	file_api_proto_v1_parkingservice_proto_rawDesc = nil
	file_api_proto_v1_parkingservice_proto_goTypes = nil
	file_api_proto_v1_parkingservice_proto_depIdxs = nil
}