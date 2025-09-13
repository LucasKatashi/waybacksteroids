package ui

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func PrintIPblockMessage(threads int) {
	fmt.Printf("\n\n[%s] IP BLOCKED DETECTED!\n", color.RedString("BLOCK"))
	fmt.Printf("[%s] More than 10 connection refused errors detected.\n", color.RedString("BLOCK"))
	fmt.Printf("[%s] Suggestions:\n", color.YellowString("HINT"))
	fmt.Printf("  • Use a VPN or proxy\n")
	fmt.Printf("  • Wait a few minutes before trying again\n")
	fmt.Printf("  • Reduce the number of threads (current: %d)\n", threads)
	os.Exit(1)
}
