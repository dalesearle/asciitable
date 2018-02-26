package main

import (
	"bytes"
	"errors"
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

type table struct {
	centerHeaders bool
	centerData    bool
	leftPad       int
	rightPad      int
	tableWidth    int
	title         string
	ascii         bytes.Buffer
	headers       []cell
	rows          map[int][]cell
	columns       []column
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

func New() *table {
	return &table{
		headers: make([]cell, 0),
		rows:    make(map[int][]cell),
	}
}

func (t *table) SetCellPadding(leftPad, rightPad int) {
	t.leftPad = leftPad
	t.rightPad = rightPad
}

func (t *table) SetCenterData() {
	t.centerData = true
}

func (t *table) SetCenterHeaders() {
	t.centerHeaders = true
}

func (t *table) SetHeaders(headers []string) {
	for _, str := range headers {
		t.headers = append(t.headers, newCell(str))
	}
}

func (t *table) SetTitle(title string) {
	t.title = title
}

func (t *table) AddRow(rowdata []string) error {
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
