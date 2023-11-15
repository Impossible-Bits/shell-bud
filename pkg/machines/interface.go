package machines

var MACHINE_KEY = "machines"

type Machine struct {
	Name   string
	IP     string
	Status bool
}
