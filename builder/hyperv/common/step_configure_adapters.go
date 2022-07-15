package common

import (
	"context"
	"fmt"
	"math"

	"github.com/hashicorp/packer-plugin-sdk/multistep"
	packersdk "github.com/hashicorp/packer-plugin-sdk/packer"
)

type StepConfigureAdapters struct {
	PrimaryAdapterIdx uint
	SwitchName        []string
	MaxAdapters       uint
}

func (s *StepConfigureAdapters) Run(ctx context.Context, state multistep.StateBag) multistep.StepAction {
	driver := state.Get("driver").(Driver)
	ui := state.Get("ui").(packersdk.Ui)

	errorMsg := "Error configuring adapters"
	vmName := state.Get("vmName").(string)
	actualMax := uint(math.Max(float64(s.MaxAdapters), float64(s.PrimaryAdapterIdx)+1))

	ui.Say(fmt.Sprintf("Configuring %d adapters...", actualMax))

	if actualMax > 1 {
		err := driver.AddVMNetworkAdapters(vmName, actualMax)
		if err != nil {
			err := fmt.Errorf(errorMsg)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}

	for i, v := range s.SwitchName {
		err := driver.ConnectVirtualMachineNetworkAdapterToSwitch(vmName, uint(i), v)
		if err != nil {
			err := fmt.Errorf(errorMsg)
			state.Put("error", err)
			ui.Error(err.Error())
			return multistep.ActionHalt
		}
	}

	return multistep.ActionContinue
}

func (s *StepConfigureAdapters) Cleanup(state multistep.StateBag) {
	//do nothing
}
