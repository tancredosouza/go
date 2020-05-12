package service

type Proxy interface {
	RemoveElement() string
	InsertElement(v int) string
	GetSize() string
	GetFirstElement() string
}

