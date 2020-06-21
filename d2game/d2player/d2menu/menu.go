package d2menu

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
)

type Menu interface {
	Load()
	Toggle()
	Advance(elapsed float64) error
	Render(target d2render.Surface)
	OnSelectDown()
	OnSelectUp()
	OnEnter()
}

type IngameMenu struct {
	isOpen bool

	screens []screen
}

func NewMenu() *IngameMenu {
	optionsScreen := newOptionsScreen()
	rootScreen := newRootScreen(optionsScreen)
	return &IngameMenu{
		screens: []screen{
			rootScreen,
		},
	}
}

func (m *IngameMenu) Load() {
	for _, screen := range m.screens {
		screen.Load()
	}
}

func (m *IngameMenu) Advance(elapsed float64) error {
	if !m.isOpen {
		return nil
	}

	for _, screen := range m.screens {
		screen.Advance(elapsed)
	}

	return nil
}

func (m *IngameMenu) Render(target d2render.Surface) {
	if !m.isOpen {
		return
	}

	for _, screen := range m.screens {
		screen.Render(target)
	}
}

func (m *IngameMenu) Toggle() {
	m.isOpen = !m.isOpen

	if m.isOpen {
		m.reset()
	}
}

func (m *IngameMenu) OnSelectDown() {
	for _, screen := range m.screens {
		if !screen.IsActive() {
			continue
		}
		screen.OnSelectDown()
	}
}

func (m *IngameMenu) OnSelectUp() {
	for _, screen := range m.screens {
		if !screen.IsActive() {
			continue
		}
		screen.OnSelectUp()
	}
}

func (m *IngameMenu) OnEnter() {
	for _, screen := range m.screens {
		if !screen.IsActive() {
			continue
		}
		screen.OnEnter()
	}
}

func (m *IngameMenu) reset() {
	for i, screen := range m.screens {
		screen.Reset()

		if i == 0 {
			screen.SetActive(true)
			continue
		}

		screen.SetActive(false)
	}
}
