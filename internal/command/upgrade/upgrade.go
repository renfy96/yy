package upgrade

import (
	"fmt"
	"github.com/renfy96/yy/config"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var UpgradeCmd = &cobra.Command{
	Use:     "upgrade",
	Short:   "Upgrade the yy command.",
	Long:    "Upgrade the yy command.",
	Example: "yy upgrade",
	Run: func(_ *cobra.Command, _ []string) {
		fmt.Printf("go install %s\n", config.YYCmd)
		cmd := exec.Command("go", "install", config.YYCmd)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			log.Fatalf("go install %s error\n", err)
		}
		fmt.Printf("\n🎉 yy upgrade successfully!\n\n")
	},
}
