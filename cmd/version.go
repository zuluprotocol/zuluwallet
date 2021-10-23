package cmd

import (
	"code.vegaprotocol.io/vegawallet/cmd/printer"
	"code.vegaprotocol.io/vegawallet/version"
	vgjson "code.vegaprotocol.io/shared/libs/json"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Show the version of the Vega wallet",
	Long:  "Show the version of the Vega wallet",
	RunE:  runVersion,
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

func runVersion(_ *cobra.Command, _ []string) error {
	if rootArgs.output == "human" {
		p := printer.NewHumanPrinter()
		p.Text("Version:").Jump().WarningText(version.Version).NJump(2)
		p.Text("Git hash:").Jump().WarningText(version.VersionHash).NJump(2)
	} else if rootArgs.output == "json" {
		return printVersionJson()
	}
	return nil
}

func printVersionJson() error {
	return vgjson.Print(struct {
		Version string `json:"version"`
		GitHash string `json:"gitHash"`
	}{
		Version: version.Version,
		GitHash: version.VersionHash,
	})
}
