package asciitable_test

import (
	"fmt"
	"github.com/dalesearle/asciitable"
	"testing"
)

func TestMultiColLongTitle(t *testing.T) {
	table := asciitable.New()
	table.SetTitle("Testing on a very long title with multi cols")
	table.SetHeaders([]string{"1", "22", "333", "4444", "55555"})
	table.AddRow([]string{"1", "2", "3", "4", "5"})
	table.AddRow([]string{"5", "4", "3", "2", "1"})
	table.SetCellPadding(1, 1)
	table.SetHeaderJustification(asciitable.JustifyCenter)
	fmt.Println(table.String())
}

func TestTitleSingleColLongTitle(t *testing.T) {
	table := asciitable.New()
	table.SetTitle("0123456789")
	table.SetHeaders([]string{"1"})
	table.SetCellPadding(1, 1)
	table.SetHeaderJustification(asciitable.JustifyCenter)
	fmt.Println(table.String())
}

func TestTitleSingleColLongHeader(t *testing.T) {
	table := asciitable.New()
	table.SetTitle("short")
	table.SetHeaders([]string{"longer than title"})
	table.SetCellPadding(1, 1)
	table.SetHeaderJustification(asciitable.JustifyCenter)
	fmt.Println(table.String())
}

func TestShortTitleShortHeaderLongData(t *testing.T) {
	table := asciitable.New()
	table.SetTitle("short")
	table.SetHeaders([]string{"short", "short"})
	table.SetCellPadding(1, 1)
	table.AddRow([]string{"Something Longer", "Something even longer"})
	table.SetHeaderJustification(asciitable.JustifyCenter)
	fmt.Println(table.String())
}