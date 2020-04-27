package distribution

import "../requestHandlers"
import "../common"
import "../packetdef"

type invoker struct {}

var stack []int

func (i invoker) exec() {
	srh := new(requestHandlers.ServerRequestHandler)

	for {
		receivedData := srh.Receive()

		// process
		processedData := process(receivedData)

		srh.Send(processedData)
	}
}

func process(data []byte) []byte {
	m := common.Marshaller{}
	p := m.Unmarshall(data)

	ans := make([]interface{}, 1)
	id := p.Body.RequestHeader.RemoteObjectId

	// choose operation
	op := p.Body.RequestHeader.Operation
	if (id == 1) {
		switch op {
		case "pop": // pop from stack
			stack = stack[:len(stack)-1]
			ans[0] = "Operation successful"
		case "push":
			stack = append(stack, p.Body.RequestBody.Data[0].(int))
			ans[0] = "Operation successful"
		case "top": // get top
			ans[0] = stack[len(stack)-1]
		case "size": // get stack size
			ans[0] = len(stack)
		default:
			// TODO: send error message
		}
	} else if (id == 2) {
		// TODO: do something on queue
	} else {
		// TODO: send error message
	}

	//assembly packet
	respHeader := packetdef.ResponseHeader{
		"", p.Body.RequestHeader.RequestId, 200,
	}

	respBody := packetdef.ResponseBody{
		ans,
	}
	header := packetdef.Header{
		Magic: "IF711",
		Version: "1.0",
		MsgType: 2,
	}

	response := packetdef.Packet{header, packetdef.Body{ResponseHeader:respHeader, ResponseBody:respBody}}

	// return answer
	return m.Marshall(response)
}