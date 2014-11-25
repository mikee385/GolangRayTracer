package table

import "fmt"

type Table struct {
	data   []interface{}
	width  int
	height int
}

func New(width int, height int) Table {
	return Table{
		data:   make([]interface{}, width*height),
		width:  width,
		height: height,
	}
}

func (table *Table) GetWidth() int {
	return table.width
}

func (table *Table) GetHeight() int {
	return table.height
}

func (table *Table) Set(row int, column int, value interface{}) {
	table.data[table.getIndex(row, column)] = value
}

func (table *Table) Get(row int, column int) interface{} {
	return table.data[table.getIndex(row, column)]
}

func (table *Table) getIndex(row int, column int) int {
	if row >= table.height {
		panic(fmt.Sprintf("Table::get_index: `row` overflow (%v >= %v)", row, table.height))
	}
	if column >= table.width {
		panic(fmt.Sprintf("Table::get_index: `column` overflow (%v >= %v)", column, table.width))
	}

	return row*table.width + column
}
