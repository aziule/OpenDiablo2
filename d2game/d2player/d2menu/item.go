package d2menu

import "github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"

type item struct {
	name       string
	isSelected bool
	label      d2ui.Label
	nextScreen screen
}
