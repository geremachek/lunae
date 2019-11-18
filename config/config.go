/*

+--------------------------+
| Lune Configuration File. |
+--------------------------+

*/

package config

import "github.com/gdamore/tcell"

var (
	XBuff int = 0 // Blank space on sides
	YBuffTop int = 0 // Blank space on top
	YBuffBottom int = 2 // Blank space on bottom (remember to make room for the bar!)

	// Key bindings (arrow keys will always work for the movement functions)

	KeyRefresh rune = 'f'
	KeyQuit rune = 'q'
	KeyDelete rune = 'd'
	KeyGroupDelete rune = 'D'
	KeyCopy rune = 'C'
	KeyMove rune = 'M'
	KeySelect rune = ' '
	KeyDotToggle rune = '.'
	KeyLeft rune = 'h'
	KeyDown rune = 'j'
	KeyUp rune = 'k'
	KeyRight rune = 'l'

	BarBg tcell.Color = tcell.GetColor("#ffffff")
	BarFg tcell.Color = tcell.GetColor("#000000")
	BarDiv string = " "
	BarStyle = map[string]tcell.Style{
		// Make sure to order your elements!
		// To get the output of a command/script (using dash), encolse the script in brackets
		// To access the currently selected file use the "@" symbol, i.e. 1[ls -l @]

		"1cwd": tcell.StyleDefault.Background(BarBg).Bold(true),
		"2size": tcell.StyleDefault.Background(BarBg),
		"3mode": tcell.StyleDefault.Background(BarBg),
	}

	FileOpen = map[string]string {
		// Key is the file type, formatted like "jpg"
		// Value is the program,t/g. "t" for terminal, "g" for gui

		"*": "vi,t", // vi, a terminal program will open all files.
	}

	FileColors = map[string]tcell.Style {
		"[dir]": tcell.StyleDefault.Foreground(tcell.GetColor("#508cbe")).Bold(true),
	}

	Bindings = map[rune][]string {
		'1': []string{"cd", "~"}, // "cd" into the home directory when the user presses 1
	}

	Selected []string
)