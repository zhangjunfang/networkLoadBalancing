package common

type Regedit interface {
	Add(name string, t interface{})
	Remove(name string) interface{}
	Update(name string, t interface{})
	Get(name string) interface{}
}
