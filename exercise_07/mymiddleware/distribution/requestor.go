package distribution

import (
	"github.com/my/repo/mymiddleware/constants"
	"github.com/my/repo/mymiddleware/infrastructure"
	"github.com/my/repo/mymiddleware/marshaller"
	"github.com/my/repo/mymiddleware/protocol"
	"log"
)

type Requestor struct {
	Crh infrastructure.ClientRequestHandler
}

func (r *Requestor) Initialize(serverHost string, serverPort int) {
	r.Crh = infrastructure.ClientRequestHandler{
		ServerHost: serverHost,
		ServerPort: serverPort,
	}

	r.Crh.Initialize()
}

func (r *Requestor) Invoke(serverHost string, serverPort int, remoteObjectKey int, operation string, param []interface{}) {
	// create marshaller
	m := marshaller.Marshaller{}

	packet := assemblePacket(remoteObjectKey, operation, param)
	r.Crh.Send(m.Marshall(packet))
}

func (r *Requestor) WaitForResponseAsync() {
	i := 0
	for {
		log.Println(i, "---->", r.Receive())
		i++
	}
}

func (r *Requestor) Receive() []interface{} {
	m := marshaller.Marshaller{}

	serializedPacket := r.Crh.Receive()

	resPacket := m.Unmarshall(serializedPacket)
	status := resPacket.Body.ResponseHeader.Status

	if (status != constants.OK_STATUS) {
		log.Fatal(resPacket.Body.ResponseBody.Data)
	}

	return resPacket.Body.ResponseBody.Data
}

func assemblePacket(remoteObjectKey int, operation string, param []interface{}) protocol.Packet {
	// assemble packet
	reqHeader := protocol.RequestHeader{
		Context: "",
		RequestId: 4242,
		ExpectsResponse: true,
		RemoteObjectKey: remoteObjectKey,
		Operation: operation,
	}

	reqBody := protocol.RequestBody{
		Data: param,
	}

	header := protocol.Header{
		Magic: "IF711",
		Version: "1.0",
		MsgType: constants.REQUEST_TYPE,
	}
	body := protocol.Body{RequestHeader: reqHeader, RequestBody: reqBody}

	packet := protocol.Packet{Header: header, Body: body}

	return packet
}