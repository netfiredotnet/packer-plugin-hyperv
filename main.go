package main

import (
	"fmt"
	"os"

	"github.com/hashicorp/packer-plugin-sdk/plugin"

	"github.com/netfiredotnet/packer-plugin-hyperv/builder/hyperv/iso"
	"github.com/netfiredotnet/packer-plugin-hyperv/builder/hyperv/vmcx"
	"github.com/netfiredotnet/packer-plugin-hyperv/version"
)

func main() {
	pps := plugin.NewSet()
	pps.RegisterBuilder("iso", new(iso.Builder))
	pps.RegisterBuilder("vmcx", new(vmcx.Builder))
	pps.SetVersion(version.PluginVersion)
	err := pps.Run()
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
