package ui

import (
	"fmt"

	"github.com/fatih/color"
)

func PrintBanner(silent bool) {
	if !silent {
		art := `                     
               _           _   
 _ _ _ ___ _ _| |_ ___ ___| |_   
| | | | .'| | | . | .'|  _| '_|  by: @LucasKatashi 🤠
|_____|__,|_  |___|__,|___|_,_|     _   _
          |___| ___| |_ ___ ___ ___|_|_| |___                                 
               |_ -|  _| -_|  _| . | | . |_ -|
               |___|_| |___|_| |___|_|___|___|
`

		blue := color.New(color.FgBlue)
		fmt.Printf("%s\n\n", blue.Sprint(art))
	}
}
