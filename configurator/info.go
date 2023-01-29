package configurator

import (
	"fmt"
	"os"
	"runtime"
)

var Info InfoT

func initInfo(port string) {
	Info.OS = runtime.GOOS
	Info.Arch = runtime.GOARCH
	if _, err := os.Lstat("/.dockerenv"); err != nil && os.IsNotExist(err) {
		Info.Container = "outside"
	} else {
		Info.Container = "inside"
	}
	Info.Hostname, _ = os.Hostname()
	taddr := getOutboundIP()
	if taddr == nil {
		Info.Address = "internal network is not available"
	} else {
		Info.Address = taddr.String() + port
	}
}

func printInfo() {
	fmt.Printf("v%s\nAddress > %s\nHostname > %s\nContainter > %s\nOS > %s/%s\nDebug > %v\nPath to config > %s\n",
		getVersion(), Info.Address, Info.Hostname,
		Info.Container, Info.OS, Info.Arch, *Debug, *PathToConfig)
}
