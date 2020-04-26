package packetdef

type msgType int

const (
	REQUEST = 1
	RESPONSE = 2
)

type Header struct {
	Magic string
	Version string
	IsLittleEndian bool
	MsgType msgType
	BodySizeInBytes int
}