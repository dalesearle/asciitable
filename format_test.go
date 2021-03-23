package asciitable

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	table := New()
	table.SetTitle("Testing")
	table.SetHeaders([]string{"1", "22", "333", "4444", "55555"})
	table.AddRow([]string{"1", "2", "3", "4", "5"})
	table.AddRow([]string{"5", "4", "3", "2", "1"})
	table.SetCellPadding(20, 2)
	table.SetHeaderJustification(JustifyCenter)
	fmt.Println(table.String())
}
