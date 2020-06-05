package locks

import "sync"

var ConnectionInfoLock sync.Mutex
var OutgoingLock sync.Mutex