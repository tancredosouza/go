package distribution

import (
	"../constants"
	"../infrastructure"
	"fmt"
	"strconv"
)
import "../marshaller"
import "../protocol"

type Invoker struct{
	HostIp string
	HostPort int
}

var stack []float64
var queue []float64

func (i Invoker) Invoke() {
	srh := infrastructure.ServerRequestHandler{
		ServerHost: i.HostIp,
		ServerPort: i.HostPort,
	}

	srh.StartListening()
	for {
		receivedData := srh.Receive()

		processedData := i.demuxAndProcess(receivedData)

		srh.Send(processedData)
	}

	srh.StopListening()
}

func (Invoker) demuxAndProcess(data []byte) []byte {
	m := marshaller.Marshaller{}
	p := m.Unmarshall(data)

	ans := make([]interface{}, 1)
	id := p.Body.RequestHeader.RemoteObjectKey

	// choose operation
	op := p.Body.RequestHeader.Operation
	var v float64 = 0.0

	if (len(p.Body.RequestBody.Data) != 0) {
		v = p.Body.RequestBody.Data[0].(float64)
	}

	var res string = ""
	var statusCode int
	if id == constants.STACK_ID {
		res = onStackPerform(op, v)
		statusCode = constants.OK_STATUS
	} else if id == constants.QUEUE_ID {
		res = onQueuePerform(op, v)
		statusCode = constants.OK_STATUS
	} else {
		res = "Invalid object ID"
		statusCode = constants.NOT_FOUND_STATUS
	}

	ans[0] = res

	respHeader := protocol.ResponseHeader{
		RequestId: p.Body.RequestHeader.RequestId, Status: statusCode,
	}

	respBody := protocol.ResponseBody{
		Data: ans,
	}

	header := protocol.Header{
		Magic:   "IF711",
		Version: "1.0",
		MsgType: 2,
	}

	response := protocol.Packet{Header: header, Body: protocol.Body{ResponseHeader: respHeader, ResponseBody: respBody}}

	return m.Marshall(response)
}

func onStackPerform(operation string, v float64) string {
	var ans string
	switch operation {
	case "pop":
		if (len(stack) > 0) {
			stack = stack[:len(stack)-1]
			ans = "Operation successful"
		} else {
			ans = "Invalid operation. Stack is empty!"
		}
		break
	case "push":
		stack = append(stack, v)
		ans = "Operation successful"
		break
	case "top":
		if (len(stack) > 0) {
			ans = fmt.Sprintf("Top is: %f", stack[len(stack)-1])
		} else {
			ans = "Invalid operation. Stack is empty!"
		}
		break
	case "size":
		ans = "Length is: " + strconv.Itoa(len(stack))
		break
	default:
		ans = "Invalid operation."
	}

	return ans
}

func onQueuePerform(operation string, v float64) string {
	var ans string
	switch operation {
	case "pop":
		if (len(queue) > 0) {
			queue = queue[1:]
			ans = "Operation successful"
		} else {
			ans = "Invalid operation. Queue is empty!"
		}
		break
	case "push":
		queue = append(queue, v)
		ans = "Operation successful"
		break
	case "front":
		if (len(queue) > 0) {
			ans = fmt.Sprintf("Front is: %f", queue[0])
		} else {
			ans = "Invalid operation. Queue is empty!"
		}
		break
	case "size":
		ans = "Length is: " + strconv.Itoa(len(queue))
		break
	default:
		ans = "Invalid operation."
	}

	return ans
}