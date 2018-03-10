package cmd

import (
	"github.com/ghodss/yaml"
)

var tuiCurrentBindings map[string]string

// This is a default YAML configuration for the key bindings,
// it can serve as an example as well.
const tuiDefaultBindingsYAML = "  \"`\": " + `KEY_PRECH
  " ": KEY_PLAY
  "\"": KEY_PAUSE
  "*": KEY_VOLUP
  "+": KEY_CHUP
  "-": KEY_CHDOWN
  "/": KEY_VOLDOWN
  "0": KEY_0
  "1": KEY_1
  "2": KEY_2
  "3": KEY_3
  "4": KEY_4
  "5": KEY_5
  "6": KEY_6
  "7": KEY_7
  "8": KEY_8
  "9": KEY_9
  "B": KEY_BLUE
  "G": KEY_GREEN
  "H": KEY_HOME
  "L": KEY_CH_LIST
  "M": KEY_MENU
  "P": KEY_POWER
  "Q": KEY_EXIT
  "R": KEY_RED
  "Y": KEY_YELLOW
  "d": KEY_HDMI
  "g": KEY_GUIDE
  "h": KEY_LEFT
  "i": KEY_INFO
  "j": KEY_DOWN
  "k": KEY_UP
  "l": KEY_RIGHT
  "m": KEY_MUTE
  "p": KEY_STOP
  "s": KEY_SOURCE
  "t": KEY_TV

  "q": TUI_QUIT
`

func tuiLoadKeyBindings(yamlText string) error {
	var b map[string]string

	if err := yaml.Unmarshal([]byte(yamlText), &b); err != nil {
		return err
	}

	tuiCurrentBindings = b
	return nil
}