package timestamp

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestTimeToTimestampProto(t *testing.T) {
	currentTime := time.Now()
	expected := timestamppb.New(currentTime)

	assert.Equal(t, expected, TimeToProtoTimestamp(currentTime))
}

func TestTimestampProtoNow(t *testing.T) {
	assert.NotEmpty(t, ProtoTimestampNow())
}
