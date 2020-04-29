package distribution

import "../common"
import "../infrastructure"
import "../packetdef"

type Requester struct {}

func (Requester) Invoke(serverHost string, serverPort int, remoteObjectKey int, operation string, param []interface{}) string {
	// create marshaller
	m := common.Marshaller{}

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
		MsgType: 11232, //TODO create request type
	}
	body := packetdef.Body{RequestHeader: reqHeader, RequestBody: reqBody}

	packet := packetdef.Packet{Header: header, Body: body}

	// send from CRH
	data := crh.SendAndReceive(m.Marshall(packet))

	// receive data
	resPacket := m.Unmarshall(data)

	return resPacket.Body.ResponseBody.Data[0].(string)
}