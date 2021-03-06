// Copyright © 2018 packet.net

package hardware

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"

	"github.com/spf13/cobra"
	"github.com/tinkerbell/tink/client"
	"github.com/tinkerbell/tink/protos/hardware"
)

// macCmd represents the mac command
var macCmd = &cobra.Command{
	Use:     "mac",
	Short:   "get hardware by any associated mac",
	Example: "tink hardware mac 00:00:00:00:00:01 00:00:00:00:00:02",
	Args: func(_ *cobra.Command, args []string) error {
		for _, arg := range args {
			if _, err := net.ParseMAC(arg); err != nil {
				return fmt.Errorf("invalid mac: %s", arg)
			}
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		for _, mac := range args {
			hw, err := client.HardwareClient.ByMAC(context.Background(), &hardware.GetRequest{Mac: mac})
			if err != nil {
				log.Fatal(err)
			}
			b, err := json.Marshal(hw)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(string(b))
		}
	},
}

func init() {
	SubCommands = append(SubCommands, macCmd)
}
