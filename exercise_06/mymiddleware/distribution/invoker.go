package distribution

import (
	"fmt"
	"github.com/my/repo/mymiddleware/constants"
	"github.com/my/repo/mymiddleware/infrastructure"
	"github.com/my/repo/mymiddleware/marshaller"
	"github.com/my/repo/mymiddleware/protocol"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Invoker struct{
	HostIp string
	HostPort int
}

var stack []int64

var queueServant []int64

func (i Invoker) Invoke() {
	srh := infrastructure.ServerRequestHandler{
		ServerHost: i.HostIp,
		ServerPort: i.HostPort,
	}

	srh.StartListening()

	for {
		srh.AcceptNewConnection()
		go i.handleNewClientConnection(srh)
	}

	srh.StopListening()
}

func (i Invoker) handleNewClientConnection(srh infrastructure.ServerRequestHandler) {
	for {
		//log.Println("Waiting to receive data from client")
		receivedData, err := srh.Receive()
		if (err != nil) {
			break;
		}

		processedData := i.demuxAndProcess(receivedData)

		//log.Println("Sending data to client")
		srh.Send(processedData)
	}

	srh.CloseConnection()
}

func (Invoker) demuxAndProcess(data []byte) []byte {
	m := marshaller.Marshaller{}
	p := m.Unmarshall(data)

	ans := make([]interface{}, 1)
	id := p.Body.RequestHeader.RemoteObjectKey

	// choose operation
	op := p.Body.RequestHeader.Operation
	var v int64 = 0

	if (len(p.Body.RequestBody.Data) != 0) {
		v = p.Body.RequestBody.Data[0].(int64)
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

func onStackPerform(operation string, v int64) string {
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

func onQueuePerform(operation string, v int64) string {
	var ans string
	switch operation {
	case "pop":
		if (len(queueServant) > 0) {
			queueServant = queueServant[1:]
			ans = "Operation successful"
		} else {
			ans = "Invalid operation. Queue is empty!"
		}
		break
	case "push":
		queueServant = append(queueServant, v)
		ans = "Operation successful"
		break
	case "front":
		if (len(queueServant) > 0) {
			ans = fmt.Sprintf("Front is: %f", queueServant[0])
		} else {
			ans = "Invalid operation. Queue is empty!"
		}
		break
	case "size":
		ans = "Length is: " + strconv.Itoa(len(queueServant))
		break
	default:
		ans = "Invalid operation."
	}

	return ans
}

func fetchDataForQueue(queueId int) {
	sourceFile := fmt.Sprintf("database/queue_%d.txt", queueId)
	queueServant, _ = readFile(sourceFile)
}

// It would be better for such a function to return error, instead of handling
// it on their own.
func readFile(fname string) (nums []int64, err error) {
	b, err := ioutil.ReadFile(fname)
	if err != nil { return nil, err }

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums = make([]int64, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 { continue }
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.Atoi(l)
		if err != nil { return nil, err }
		nums = append(nums, int64(n))
	}

	return nums, nil
}

func saveDataOnDatabase(queueId int) {
	destFilepath := fmt.Sprintf("database/queue_%d.txt", queueId)

	writeFile(destFilepath)
}

func writeFile(filepath string) error {
	outputFile, err := os.Create(filepath)
	if (err != nil) {
		return err
	}

	for i:=0; i<len(queueServant);i++ {
		fmt.Fprintln(outputFile, queueServant[i])
	}

	return nil
}