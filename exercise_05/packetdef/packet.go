package packetdef

type Packet struct {
	Header Header
	Body   Body
}

type Header struct {
	Magic string
	Version string
	IsLittleEndian bool
	MsgType int
	BodySizeInBytes int
}

type Body struct {
	RequestHeader  RequestHeader
	RequestBody    RequestBody
	ResponseHeader ResponseHeader
	ResponseBody   ResponseBody
}

type RequestHeader struct {
	Context         string
	RequestId       int
	ExpectsResponse bool
	RemoteObjectKey int
	Operation       string
}

type RequestBody struct {
	Data []interface{}
}

type ResponseHeader struct {
	Context string
	RequestId int
	Status int
}

type ResponseBody struct {
	Data []interface{}
}