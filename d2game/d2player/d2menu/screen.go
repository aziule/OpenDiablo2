package d2menu

import (
	"github.com/OpenDiablo2/OpenDiablo2/d2common/d2resource"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2asset"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2render"
	"github.com/OpenDiablo2/OpenDiablo2/d2core/d2ui"
)

type screen interface {
	Advance(elapsed float64) error
	IsActive() bool
	Load()
	Name() string
	OnSelectDown()
	OnSelectUp()
	OnEnter()
	Render(target d2render.Surface)
	Reset()
	SetActive(active bool)
}

type rootScreen struct {
	isActive         bool
	lSelectionSprite *d2ui.Sprite
	rSelectionSprite *d2ui.Sprite
	totalHeight      int
	items            []*item

	optionsScreen screen
}

func newRootScreen(optionsScreen screen) *rootScreen {
	return &rootScreen{
		optionsScreen: optionsScreen,
	}
}

func (s *rootScreen) Advance(elapsed float64) error {
	s.lSelectionSprite.Advance(elapsed)
	s.rSelectionSprite.Advance(elapsed)
	return nil
}

func (s *rootScreen) IsActive() bool {
	return s.isActive
}

func (s *rootScreen) Load() {
	lAnimation, _ := d2asset.LoadAnimation(d2resource.PentSpin, d2resource.PaletteSky)
	lSelectionSprite, _ := d2ui.LoadSprite(lAnimation)
	lSelectionSprite.PlayForward()
	lSelectionSprite.SetBlend(true)
	s.lSelectionSprite = lSelectionSprite

	rAnimation, _ := d2asset.LoadAnimation(d2resource.PentSpin, d2resource.PaletteSky)
	rSelectionSprite, _ := d2ui.LoadSprite(rAnimation)
	rSelectionSprite.PlayForward()
	rSelectionSprite.SetBlend(true)
	s.rSelectionSprite = rSelectionSprite

	totalHeight := 0

	itemOptions := &item{
		name:       "options",
		isSelected: true,
		label:      d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteMenu0),
		nextScreen: optionsScreen,
	}
	itemOptions.label.SetText("options")
	_, h := itemOptions.label.GetSize()
	totalHeight += h

	itemSave := &item{
		name:       "save",
		isSelected: false,
		label:      d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteMenu0),
	}
	itemSave.label.SetText("save and exit game")
	_, h = itemSave.label.GetSize()
	totalHeight += h

	itemBack := &item{
		name:       "back",
		isSelected: false,
		label:      d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteMenu0),
	}
	itemBack.label.SetText("return to game")
	_, h = itemBack.label.GetSize()
	totalHeight += h

	s.totalHeight = totalHeight
	s.items = []*item{itemOptions, itemSave, itemBack}
}

func (s *rootScreen) Name() string {
	return "root"
}

func (s *rootScreen) Render(target d2render.Surface) {
	if !s.isActive {
		return
	}

	tw, th := target.GetSize()
	startY := (th - s.totalHeight) / 2

	_, sh := s.lSelectionSprite.GetCurrentFrameSize()

	for i, item := range s.items {
		iw, ih := item.label.GetSize()
		ix := (tw - iw) / 2
		iy := startY + i*ih

		item.label.SetPosition(ix, iy)
		item.label.Render(target)

		if item.isSelected {
			s.lSelectionSprite.SetPosition(100, iy+sh)
			s.rSelectionSprite.SetPosition(tw-100, iy+sh)
		}
	}

	s.lSelectionSprite.Render(target)
	s.rSelectionSprite.Render(target)
}

func (s *rootScreen) OnSelectDown() {
	selectNext := false
	for i, item := range s.items {
		if item.isSelected {
			if i == len(s.items)-1 {
				return
			}
			item.isSelected = false
			selectNext = true
			continue
		}

		if selectNext {
			item.isSelected = true
			return
		}
	}
}

func (s *rootScreen) OnSelectUp() {
	selectPrev := false
	for i := len(s.items) - 1; i >= 0; i-- {
		if s.items[i].isSelected {
			if i == 0 {
				return
			}
			s.items[i].isSelected = false
			selectPrev = true
			continue
		}

		if selectPrev {
			s.items[i].isSelected = true
			return
		}
	}
}

func (s *rootScreen) OnEnter() {
	for _, item := range s.items {
		if !item.isSelected {
			continue
		}
		if item.nextScreen != nil {
			s.isActive = false
			item.nextScreen.SetActive(true)
			s.Reset()
			return
		}
	}
}

func (s *rootScreen) Reset() {
	for i, item := range s.items {
		if i == 0 {
			item.isSelected = true
			continue
		}

		item.isSelected = false
	}
}

func (s *rootScreen) SetActive(active bool) {
	s.isActive = active
}

type optionsScreen struct {
	isActive         bool
	lSelectionSprite *d2ui.Sprite
	rSelectionSprite *d2ui.Sprite
	totalHeight      int
	items            []*item
}

func newOptionsScreen() *optionsScreen {
	return &optionsScreen{}
}

func (s *optionsScreen) Advance(elapsed float64) error {
	s.lSelectionSprite.Advance(elapsed)
	s.rSelectionSprite.Advance(elapsed)
	return nil
}

func (s *optionsScreen) IsActive() bool {
	return s.isActive
}

func (s *optionsScreen) Load() {
	lAnimation, _ := d2asset.LoadAnimation(d2resource.PentSpin, d2resource.PaletteSky)
	lSelectionSprite, _ := d2ui.LoadSprite(lAnimation)
	lSelectionSprite.PlayForward()
	lSelectionSprite.SetBlend(true)
	s.lSelectionSprite = lSelectionSprite

	rAnimation, _ := d2asset.LoadAnimation(d2resource.PentSpin, d2resource.PaletteSky)
	rSelectionSprite, _ := d2ui.LoadSprite(rAnimation)
	rSelectionSprite.PlayForward()
	rSelectionSprite.SetBlend(true)
	s.rSelectionSprite = rSelectionSprite

	totalHeight := 0

	soundOptions := &item{
		isSelected: true,
		label:      d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteMenu0),
	}
	soundOptions.label.SetText("sound options")
	_, h := soundOptions.label.GetSize()
	totalHeight += h

	videoOptions := &item{
		isSelected: false,
		label:      d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteMenu0),
	}
	videoOptions.label.SetText("video options")
	_, h = videoOptions.label.GetSize()
	totalHeight += h

	automapOptions := &item{
		isSelected: false,
		label:      d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteMenu0),
	}
	automapOptions.label.SetText("automap options")
	_, h = automapOptions.label.GetSize()
	totalHeight += h

	configureOptions := &item{
		isSelected: false,
		label:      d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteMenu0),
	}
	configureOptions.label.SetText("configure options")
	_, h = configureOptions.label.GetSize()
	totalHeight += h

	previousMenu := &item{
		isSelected: false,
		label:      d2ui.CreateLabel(d2resource.Font42, d2resource.PaletteMenu0),
	}
	previousMenu.label.SetText("previous menu")
	_, h = previousMenu.label.GetSize()
	totalHeight += h

	s.totalHeight = totalHeight
	s.items = []*item{soundOptions, videoOptions, automapOptions, configureOptions, previousMenu}
}

func (s *optionsScreen) Name() string {
	return "options"
}

func (s *optionsScreen) Render(target d2render.Surface) {
	if !s.isActive {
		return
	}

	tw, th := target.GetSize()
	startY := (th - s.totalHeight) / 2

	_, sh := s.lSelectionSprite.GetCurrentFrameSize()

	for i, item := range s.items {
		iw, ih := item.label.GetSize()
		ix := (tw - iw) / 2
		iy := startY + i*ih

		item.label.SetPosition(ix, iy)
		item.label.Render(target)

		if item.isSelected {
			s.lSelectionSprite.SetPosition(100, iy+sh)
			s.rSelectionSprite.SetPosition(tw-100, iy+sh)
		}
	}

	s.lSelectionSprite.Render(target)
	s.rSelectionSprite.Render(target)
}

func (s *optionsScreen) OnSelectDown() {
	selectNext := false
	for i, item := range s.items {
		if item.isSelected {
			if i == len(s.items)-1 {
				return
			}
			item.isSelected = false
			selectNext = true
			continue
		}

		if selectNext {
			item.isSelected = true
			return
		}
	}
}

func (s *optionsScreen) OnSelectUp() {
	selectPrev := false
	for i := len(s.items) - 1; i >= 0; i-- {
		if s.items[i].isSelected {
			if i == 0 {
				return
			}
			s.items[i].isSelected = false
			selectPrev = true
			continue
		}

		if selectPrev {
			s.items[i].isSelected = true
			return
		}
	}
}

func (s *optionsScreen) OnEnter() {
}

func (s *optionsScreen) Reset() {
	for i, item := range s.items {
		if i == 0 {
			item.isSelected = true
			continue
		}

		item.isSelected = false
	}
}

func (s *optionsScreen) SetActive(active bool) {
	s.isActive = active
}
