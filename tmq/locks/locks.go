package locks

import "sync"

var TopicsLock sync.Mutex
var OutgoingLock sync.Mutex
var SubscribersLock sync.Mutex
var ConnectionInfoLock sync.Mutex