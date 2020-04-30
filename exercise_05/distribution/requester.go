package distribution

import "../marshaller"
import "../infrastructure"
import "../protocol"
import "../constants"

type Requester struct {}

func (Requester) Invoke(serverHost string, serverPort int, remoteObjectKey int, operation string, param []interface{}) []interface{} {
	// create marshaller
	m := marshaller.Marshaller{}

	// declare CRH
	crh := infrastructure.ClientRequestHandler{
		ServerHost: serverHost,
		ServerPort: serverPort,
	}

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
	serializedPacket := crh.SendAndReceive(m.Marshall(packet))

	// receive serializedPacket
	resPacket := m.Unmarshall(serializedPacket)

	return resPacket.Body.ResponseBody.Data
}