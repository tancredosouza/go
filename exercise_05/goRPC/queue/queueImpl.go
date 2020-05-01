package queue

import (
	"fmt"
	"strconv"
)

type QueueRPC struct{
	data []int
}

func (q *QueueRPC) RemoveElement(v *int, reply *string) error {
	if (len(q.data) > 0) {
		q.data = q.data[1:]
		*reply = "Operation successful"
	} else {
		*reply = "Invalid operation. Queue is empty!"
	}
	return nil
}

func (q *QueueRPC) InsertElement(v *int, reply *string) error {
	q.data = append(q.data, *v)
	*reply = "Operation successful"

	return nil
}

func (q *QueueRPC) GetSize(v *int, reply *string) error {
	*reply = strconv.Itoa(len(q.data))

	return nil
}

func (q *QueueRPC) GetFirstElement(v *int, reply *string) error {
	if (len(q.data) > 0) {
		*reply = fmt.Sprintf("Front is: %d", q.data[0])
	} else {
		*reply = "Invalid operation. Queue is empty!"
	}

	return nil
}