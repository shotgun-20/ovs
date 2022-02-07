package ovs

//the passive TCP port where OVS entries are listening
//for OpenFlow commands
const OvsDefaultPort int = 6633

type OvsStatReader interface {
	TunFlows(ip string, port int) ([]Flow, error)
	TunPorts(ip string, port int) ([]Port, error)
	ExFlows(ip string, port int) ([]Flow, error)
	ExPorts(ip string, port int) ([]Port, error)
	IntFlows(ip string, port int) ([]Flow, error)
	IntPorts(ip string, port int) ([]Port, error)
//	Groups(ip string, port int) ([]Group, error)
}

var (
	OvsDefaultReader OvsStatReader = OvsDumpReader{}
)
