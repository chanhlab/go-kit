package timestamp

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// TimeToProtoTimestamp converts from time.Time to Protobuf Timestamp
func TimeToProtoTimestamp(time time.Time) *timestamppb.Timestamp {
	return timestamppb.New(time)
}

// ProtoProtoTimestampNow gets the current Protobuf Timestamp
func ProtoTimestampNow() *timestamppb.Timestamp {
	return timestamppb.Now()
}
