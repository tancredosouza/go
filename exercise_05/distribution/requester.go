package distribution

import "../marshaller"
import "../infrastructure"
import "../packetdef"
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
	reqHeader := packetdef.RequestHeader{
		Context: "",
		RequestId: 4242,
		ExpectsResponse: true,
		RemoteObjectKey: remoteObjectKey,
		Operation: operation,
	}

	reqBody := packetdef.RequestBody{
		Data: param,
	}

	header := packetdef.Header{
		Magic: "IF711",
		Version: "1.0",
		MsgType: constants.REQUEST_TYPE,
	}
	body := packetdef.Body{RequestHeader: reqHeader, RequestBody: reqBody}

	packet := packetdef.Packet{Header: header, Body: body}

	// send from CRH
	serializedPacket := crh.SendAndReceive(m.Marshall(packet))

	// receive serializedPacket
	resPacket := m.Unmarshall(serializedPacket)

	return resPacket.Body.ResponseBody.Data
}