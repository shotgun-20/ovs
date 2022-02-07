package ovs

import (
	"os/exec"
	"strconv"
	"strings"
)

type OvsDumpSourceCLI struct{}

func ovsCtlRun(params ...string) ([]string, error) {
	cmd := exec.Command("ovs-ofctl", params...)
	out, err := cmd.Output()
	outString := string(out)
	//if error was occured we return
	if err != nil {
		return nil, err
	}
	//if command was succesfull we further parse the output

	lines := strings.Split(outString, "\n")
	//skip the first and last lines, since it is just a response header and an empty line
	lines = lines[1:(len(lines) - 1)]
	return lines, nil
}

func (o OvsDumpSourceCLI) TunDumpFlows(ip string, port int) ([]string, error) {
	return ovsCtlRun("dump-flows", "br-tun")
}

func (o OvsDumpSourceCLI) TunDumpPorts(ip string, port int) ([]string, error) {
	return ovsCtlRun("dump-ports", "br-tun")
}

func (o OvsDumpSourceCLI) ExDumpFlows(ip string, port int) ([]string, error) {
	return ovsCtlRun("dump-flows", "br-ex")
}

func (o OvsDumpSourceCLI) ExDumpPorts(ip string, port int) ([]string, error) {
	return ovsCtlRun("dump-ports", "br-ex")
}

func (o OvsDumpSourceCLI) IntDumpFlows(ip string, port int) ([]string, error) {
	return ovsCtlRun("dump-flows", "br-int")
}

func (o OvsDumpSourceCLI) IntDumpPorts(ip string, port int) ([]string, error) {
	return ovsCtlRun("dump-ports", "br-int")
}

func (o OvsDumpSourceCLI) DumpGroups(ip string, port int) ([]string, error) {
	return ovsCtlRun("dump-groups", "tcp:"+ip+":"+strconv.Itoa(port))
}

func (o OvsDumpSourceCLI) DumpGroupStats(ip string, port int) ([]string, error) {
	return ovsCtlRun("dump-group-stats", "tcp:"+ip+":"+strconv.Itoa(port))
}
