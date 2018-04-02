package asciitable

import (
	"bytes"
	"errors"
)

const (
	JustifyCenter = 1
	JustifyLeft   = 2
	JustifyRight  = 3
)

type cell struct {
	data       string
	dataLength int
}

type column struct {
	header cell
	data   []cell
	width  int
}

type Table struct {
	headerJustification int
	dataJustification   int
	leftPad             int
	rightPad            int
	tableWidth          int
	title               string
	ascii               bytes.Buffer
	headers             []cell
	rows                map[int][]cell
	columns             []column
}

func newCell(data string) cell {
	return cell{
		data:       data,
		dataLength: len(data),
	}
}

func newColumn() *column {
	return &column{
		data: make([]cell, 0),
	}
}

func New() *Table {
	return &Table{
		headerJustification: JustifyCenter,
		headers:             make([]cell, 0),
		rows:                make(map[int][]cell),
	}
}

func (t *Table) SetCellPadding(leftPad, rightPad int) {
	t.leftPad = leftPad
	t.rightPad = rightPad
}

func (t *Table) SetDataJustification(justification int) {
	t.dataJustification = justification
}

func (t *Table) SetHeaderJustification(justification int) {
	t.headerJustification = justification
}

func (t *Table) SetHeaders(headers []string) {
	for _, str := range headers {
		t.headers = append(t.headers, newCell(str))
	}
}

func (t *Table) SetTitle(title string) {
	t.title = title
}

func (t *Table) AddRow(rowdata []string) error {
	l := len(rowdata)
	if l != len(t.headers) {
		return errors.New("row length does not match the header length ")
	}
	var row = make([]cell, l)

	for i, str := range rowdata {
		row[i] = newCell(str)
	}
	t.rows[len(t.rows)] = row
	return nil
}

func (t *Table) String() string {
	t.ascii = bytes.Buffer{}
	t.ascii.WriteString("\n")
	t.createColumns()
	t.calcTableWidth()
	t.writeTitle()
	t.writeHeaders()
	t.writeRows()
	return t.ascii.String()
}
