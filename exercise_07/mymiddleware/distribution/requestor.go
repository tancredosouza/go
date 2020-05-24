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

	log.Println("initializing connection")
	r.Crh.Initialize()
	log.Println("initializing connection")
}

func (r *Requestor) Invoke(serverHost string, serverPort int, remoteObjectKey int, operation string, param []interface{}) {
	// create marshaller
	m := marshaller.Marshaller{}

	packet := assemblePacket(remoteObjectKey, operation, param)
	log.Println("Sending -> ", param)
	go r.Crh.Send(m.Marshall(packet))
}

func (r *Requestor) ResultCallback() []interface{} {
	for {
		log.Println("Waiting to receive")
		log.Println(r.Receive())
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