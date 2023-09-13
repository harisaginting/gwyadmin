package components

import (
	"time"
)

//TimeSinceOrganizationEpoch returns an epoch time which describes how long its been between organizations epoch
//and when an object was created.
func TimeSinceOrganizationEpoch(orgEpoch int64) int64 {
	now := time.Now()
	nanos := now.UnixNano()
	millis := nanos / 100000

	timeSince := millis - orgEpoch
	return timeSince
}
