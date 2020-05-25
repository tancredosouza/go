package distribution

import (
	"fmt"
	"github.com/my/repo/mymiddleware/constants"
	"github.com/my/repo/mymiddleware/infrastructure"
	"github.com/my/repo/mymiddleware/marshaller"
	"github.com/my/repo/mymiddleware/protocol"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Invoker struct{
	HostIp string
	HostPort int
}

var servants chan *[]float64

func (i Invoker) Invoke() {
	addServants(50)

	srh := infrastructure.ServerRequestHandler{
		ServerHost: i.HostIp,
		ServerPort: i.HostPort,
	}

	srh.Initialize()

	srh.AcceptNewConnection()
	srh.CreateQueues()

	go i.handleNewClientConnection(srh)
}

func addServants(n int) {
	servants = make(chan *[]float64, n)

	for i:=0; i<n;i++{
		servants <- &[]float64{}
	}
}

func (i Invoker) handleNewClientConnection(srh infrastructure.ServerRequestHandler) {
	for {
		receivedData := srh.Receive()

		processedData := i.demuxAndProcess(receivedData)

		srh.Send(processedData)
	}
}

func (Invoker) demuxAndProcess(data []byte) []byte {
	m := marshaller.Marshaller{}
	p := m.Unmarshall(data)

	ans := make([]interface{}, 1)
	id := p.Body.RequestHeader.RemoteObjectKey

	// choose operation
	op := p.Body.RequestHeader.Operation
	var v float64 = 0
	var queueNumber = int(p.Body.RequestBody.Data[0].(float64))

	if (len(p.Body.RequestBody.Data) > 1) {
		v = p.Body.RequestBody.Data[1].(float64)
	}

	var res string = ""
	var statusCode int
	if id == constants.QUEUE_ID {
		res = onQueuePerform(op, v, queueNumber)
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

func onQueuePerform(operation string, v float64, queueNumber int) string {
	<- servants
	acquiredServant := fetchDataForQueue(queueNumber)
	var ans string
	var x []float64
	switch operation {
	case "pop":
		if (len(acquiredServant) > 0) {
			x = (acquiredServant)[1:]
			acquiredServant = x
			ans = "Operation successful"
		} else {
			ans = "Invalid operation. Queue is empty!"
		}
		break
	case "push":
		x = append(acquiredServant, v)
		acquiredServant = x
		ans = "Operation successful"
		break
	case "front":
		if (len(acquiredServant) > 0) {
			ans = fmt.Sprintf("Front is: %f", (acquiredServant)[0])
		} else {
			ans = "Invalid operation. Queue is empty!"
		}
		break
	case "size":
		ans = "Length is: " + strconv.Itoa(len(acquiredServant))
		break
	default:
		ans = "Invalid operation."
	}
	saveDataOnDatabase(&acquiredServant, queueNumber)
	acquiredServant = nil
	servants <- &acquiredServant
	return ans
}

func fetchDataForQueue(queueId int) []float64 {
	sourceFile := fmt.Sprintf("./mymiddleware/database/queue_%d.txt", queueId)

	servant, err := readFile(sourceFile)
	if (err != nil) {
		log.Fatal("Error while fetching data from database ", err)
	}

	return *servant
}

func readFile(fname string) (*[]float64, error) {
	if(!fileExists(fname)) {
		return &[]float64{}, nil
	}

	b, err := ioutil.ReadFile(fname)
	if err != nil { return nil, err }

	lines := strings.Split(string(b), "\n")
	// Assign cap to avoid resize on every append.
	nums := make([]float64, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 { continue }
		// Atoi better suits the job when we know exactly what we're dealing
		// with. Scanf is the more general option.
		n, err := strconv.ParseFloat(l,64)
		if err != nil { return nil, err }
		nums = append(nums, float64(n))
	}

	return &nums, nil
}

func saveDataOnDatabase(values *[]float64, queueId int) {
	destFilepath := fmt.Sprintf("./mymiddleware/database/queue_%d.txt", queueId)

	err := writeFile(values, destFilepath)
	if (err != nil) {
		log.Fatal("Error while saving data to database ", err)
	}
}

func writeFile(values *[]float64, filepath string) error {
	if (!fileExists(filepath)){
		_, err := os.Create(filepath)
		if (err != nil) {
			return err
		}
	}

	f, err := os.OpenFile(filepath, os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	for i:=0; i<len(*values);i++ {
		if _, err = f.WriteString(fmt.Sprintf("%f\n", (*values)[i])); err != nil {
			panic(err)
		}
	}

	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}