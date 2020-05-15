package distribution

import (
	"errors"
	"github.com/my/repo/mymiddleware/constants"
	"github.com/my/repo/mymiddleware/infrastructure"
	"github.com/my/repo/mymiddleware/marshaller"
	"github.com/my/repo/mymiddleware/protocol"
	"log"
	"net"
)

type Requester struct {
	Crh  infrastructure.ClientRequestHandler
	Conn net.Conn
}

func (r *Requester) Initialize(serverHost string, serverPort int) {
	r.Crh = infrastructure.ClientRequestHandler{
		ServerHost: serverHost,
		ServerPort: serverPort,
	}

	log.Println("stablishing connection")
	r.Conn = r.Crh.StablishConnection()
	log.Println("stablished connection")
}

func (r Requester) Invoke(serverHost string, serverPort int, remoteObjectKey int, operation string, param []interface{}) ([]interface{}, error) {
	// create marshaller
	m := marshaller.Marshaller{}

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

	// send from CRH
	serializedPacket := r.Crh.SendAndReceive(m.Marshall(packet), r.Conn)

	// receive serializedPacket
	resPacket := m.Unmarshall(serializedPacket)
	status := resPacket.Body.ResponseHeader.Status
	if (status != constants.OK_STATUS) {
		return resPacket.Body.ResponseBody.Data, errors.New("")
	}

	return resPacket.Body.ResponseBody.Data, nil
}