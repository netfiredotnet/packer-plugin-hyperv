package common

import (
	"context"
	"fmt"
	"strconv"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepConfigureVlan struct {
	VlanId       []string
	SwitchVlanId string
}

func (s *StepConfigureVlan) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packersdk.Ui)

	errorMsg := "Error configuring vlan: %s"
	vmName := state.Get("vmName").(string)
	switchName := state.Get("SwitchName").(string)
	vlanId := s.VlanId
	switchVlanId := s.SwitchVlanId

	ui.Say("Configuring vlan...")

	if switchVlanId != "" {
		err := driver.SetNetworkAdapterVlanId(switchName, vlanId[0])
		if err != nil {
			err := fmt.Errorf(errorMsg, err)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}

	if len(vlanId) > 0 {
		for i, v := range vlanId {
			err := driver.SetVirtualMachineVlanId(vmName, strconv.Itoa(i), v)
			if err != nil {
				err := fmt.Errorf(errorMsg, err)
				state.Put("error", err)
				ui.Error(err.Error())
				return multistep.ActionHalt
			}
		}
	}

	return multistep.ActionContinue
}

func (s *StepConfigureVlan) Cleanup(state multistep.StateBag) {
	//do nothing
}
