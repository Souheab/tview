package tview

import "github.com/gdamore/tcell/v2"

func (l *List) CustomInsertItem(index int, mainText string, selectedStyle *tcell.Style) *List {
	item := &listItem{
		MainText:      mainText,
		SecondaryText: "",
		Shortcut:      0,
		Selected:      nil,
		SelectedStyle: selectedStyle,
	}

	// Shift index to range.
	if index < 0 {
		index = len(l.items) + index + 1
	}
	if index < 0 {
		index = 0
	} else if index > len(l.items) {
		index = len(l.items)
	}

	// Shift current item.
	if l.currentItem < len(l.items) && l.currentItem >= index {
		l.currentItem++
	}

	// Insert item (make space for the new item, then shift and insert).
	l.items = append(l.items, nil)
	if index < len(l.items)-1 { // -1 because l.items has already grown by one item.
		copy(l.items[index+1:], l.items[index:])
	}
	l.items[index] = item

	// Fire a "change" event for the first item in the list.
	if len(l.items) == 1 && l.changed != nil {
		item := l.items[0]
		l.changed(0, item.MainText, item.SecondaryText, item.Shortcut)
	}

	return l
}
