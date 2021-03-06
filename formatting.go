package asciitable

const (
	borderblcorner   rune = 9562 //9495
	borderbrcorner   rune = 9565 //9499
	borderhorizontal rune = 9552 //9473
	bordertlcorner   rune = 9556 //9487
	bordertrcorner   rune = 9559 //9491
	bordervertical   rune = 9553 //9475
	space            rune = 32
	line             rune = 9472
	vertical         rune = 9474
	leftrowtee       rune = 9553 //9567 // 9504
	boldlefttee      rune = 9568 //9507
	rightrowtee      rune = 9553 //9570 //9512
	boldrighttee     rune = 9571 //9515
	upperboldtee     rune = 9574 //9523
	bottomtee        rune = 9552 //9575 // 9527
	cross            rune = 9532
	headercross      rune = 9577 // 9543
)

// TODO: what if the title length is > header width?
func (t *Table) createColumns() {
	var cellWidth int
	var columns = make([]*column, 0)
	var colWidth int
	var fmtOverhead = t.leftPad + t.rightPad
	var workingCell cell
	var workingColumn *column

	for i, header := range t.headers {
		workingColumn = newColumn()
		workingColumn.header = header
		colWidth = header.dataLength + fmtOverhead

		for j := 0; j < len(t.rows); j++ {
			workingCell = t.rows[j][i]
			cellWidth = workingCell.dataLength + fmtOverhead

			workingColumn.data = append(workingColumn.data, workingCell)
			if colWidth < cellWidth {
				colWidth = cellWidth
			}
		}
		workingColumn.width = colWidth
		columns = append(columns, workingColumn)
	}
	t.columns = columns
	t.adjustForTitle()
}

func (t *Table) adjustForTitle() {
	nc := len(t.columns)
	colWidth := nc -1 // adjust for internal vertical rules

	titleWidth := t.title.dataLength + t.leftPad + t.rightPad
	for _,c := range t.columns {
		colWidth += c.width
	}
	if colWidth < titleWidth {
		t.tableWidth = titleWidth
		adjust := titleWidth - colWidth
		for adjust > 0 {
			for _,c := range t.columns {
				if adjust > 0 {
					c.width = c.width + 1
					adjust--
				}
			}
		}

	} else {
		t.tableWidth = colWidth
	}
}

func (t *Table) writeTitle() {
	if t.title.data != "" {
		t.writeRule(bordertlcorner, borderhorizontal, borderhorizontal, bordertrcorner)
		t.writeRunes(bordervertical, 1)
		t.writeCenterJustifiedCell(t.title, t.tableWidth)
		t.writeRunes(bordervertical, 1)
		t.writeRunes('\n', 1)
	}
}

func (t *Table) writeHeaders() {
	var leftCorner = boldlefttee
	var rightCorner = boldrighttee
	if t.title.data == "" {
		leftCorner = bordertlcorner
		rightCorner = bordertrcorner
	}
	t.writeRule(leftCorner, borderhorizontal, upperboldtee, rightCorner)
	t.ascii.WriteRune(bordervertical)
	for _, col := range t.columns {
		if t.headerJustification == JustifyCenter {
			t.writeCenterJustifiedCell(col.header, col.width)
		} else {
			t.writeCell(col.header, col.width)
		}
		t.ascii.WriteRune(bordervertical)
	}
	t.writeRunes('\n', 1)
	t.writeRule(boldlefttee, borderhorizontal, headercross, boldrighttee)
}

func (t *Table) writeRows() {
	var numRows = len(t.rows)
	var numCols = len(t.columns)

	for rowIndex := 0; rowIndex < numRows; rowIndex++ {
		t.writeRunes(bordervertical, 1)
		for colIndex, col := range t.columns {
			cell := col.data[rowIndex]
			if t.dataJustification == JustifyCenter {
				t.writeCenterJustifiedCell(cell, col.width)
			} else {
				t.writeCell(cell, col.width)
			}
			if colIndex < numCols-1 {
				t.writeRunes(vertical, 1)
			} else {
				t.writeRunes(bordervertical, 1)
			}
		}
		t.writeRunes('\n', 1)
		if rowIndex+1 < numRows {
			t.writeRule(leftrowtee, line, cross, rightrowtee)
		}
	}
	t.writeRule(borderblcorner, borderhorizontal, bottomtee, borderbrcorner)
}

func (t *Table) writeRule(start, fill, joint, end rune) {
	var lastJoint = len(t.columns) - 1

	t.writeRunes(start, 1)
	for i, col := range t.columns {
		t.writeRunes(fill, col.width)
		if i < lastJoint {
			t.writeRunes(joint, 1)
		}
	}
	t.writeRunes(end, 1)
	t.ascii.WriteRune('\n')
}

func (t *Table) writeCell(c cell, colWidth int) {
	leftPad, rightPad := t.leftJustify(c, colWidth)
	t.writeRunes(space, leftPad)
	t.ascii.WriteString(c.data)
	t.writeRunes(space, rightPad)
}

func (t *Table) writeCenterJustifiedCell(c cell, colWidth int) {
	leftPad, rightPad := t.centerJustify(c, colWidth)
	t.writeRunes(space, leftPad)
	t.ascii.WriteString(c.data)
	t.writeRunes(space, rightPad)
}

func (t *Table) centerJustify(c cell, columnWidth int) (pre, post int) {
	padding := columnWidth - c.dataLength
	rightPad := padding / 2
	leftPad := rightPad + padding%2
	return leftPad, rightPad
}

func (t *Table) leftJustify(c cell, columnWidth int) (left, right int) {
	leftPad := columnWidth - (c.dataLength + t.rightPad)
	return leftPad, t.rightPad
}

func (t *Table) writeRunes(r rune, writes int) {
	for i := 0; i < writes; i++ {
		t.ascii.WriteRune(r)
	}
}
