package timestamp

import (
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// TimeToProtoTimestamp converts from time.Time to Protobuf Timestamp
func TimeToProtoTimestamp(time time.Time) *timestamp.Timestamp {
	timestampProto, _ := ptypes.TimestampProto(time)
	return timestampProto
}

// ProtoProtoTimestampNow gets the current Protobuf Timestamp
func ProtoTimestampNow() *timestamp.Timestamp {
	return ptypes.TimestampNow()
}
