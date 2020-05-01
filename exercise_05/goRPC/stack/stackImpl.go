package stack

import (
	"fmt"
	"strconv"
)

type StackRPC struct{
	data []int
}

func (s *StackRPC) RemoveElement(v *int, reply *string) error {
	if (len(s.data) > 0) {
		s.data = s.data[:len(s.data)-1]
		*reply = "Operation successful"
	} else {
		*reply = "Invalid operation. Stack is empty!"
	}

	return nil
}

func (s *StackRPC) InsertElement(v *int, reply *string) error {
	s.data = append(s.data, *v)
	*reply = "Operation successful"

	return nil
}

func (s *StackRPC) GetSize(v *int, reply *string) error {
	*reply = strconv.Itoa(len(s.data))

	return nil
}

func (s *StackRPC) GetFirstElement(v *int, reply *string) error {
	if (len(s.data) > 0) {
		*reply = fmt.Sprintf("Top is: %d", s.data[len(s.data)-1])
	} else {
		*reply = "Invalid operation. Stack is empty!"
	}

	return nil
}