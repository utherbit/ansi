package ansi

import (
	"strconv"
)

// Cursor Управление курсором
var Cursor = cursor{
	MakeCursorUpOne: "\u001BM",

	CursorSavePositionDEC: "\u001B7",
	CursorLastPositionDEC: "\u001B8",

	CursorSavePositionSCO: "\u001B[s",
	CursorLastPositionSCO: "\u001B[u",
}

type cursor struct {
	MakeCursorUpOne       string // перемещает курсор на одну строку вверх, прокручивая при необходимости
	CursorSavePositionDEC string // сохранить позицию курсора (DEC)
	CursorLastPositionDEC string // восстанавливает курсор в последнюю сохраненную позицию (DEC)
	CursorSavePositionSCO string // сохранить позицию курсора (SCO)
	CursorLastPositionSCO string // восстанавливает курсор в последнюю сохраненную позицию (SCO)
}

// MakeCursor перемещает курсор на строку {{line}}, столбец {{column}}
func (cursor) MakeCursor(line, column int) string {
	return "\u001B[" + strconv.Itoa(line) + ";" + strconv.Itoa(column) + "H"
	//return "\u001B[" + strconv.Itoa(line) + ";" + strconv.Itoa(column) + "f"
}

// MakeCursorUp перемещает курсор вверх на # строки
func (cursor) MakeCursorUp(up int) string {
	return "\u001B[" + strconv.Itoa(up) + "A"
}

// MakeCursorDown перемещает курсор вниз на # строки
func (cursor) MakeCursorDown(down int) string {
	return "\u001B[" + strconv.Itoa(down) + "B"
}

// MakeCursorRight перемещает курсор вправо # столбцов
func (cursor) MakeCursorRight(right int) string {
	return "\u001B[" + strconv.Itoa(right) + "C"
}

// MakeCursorLeft перемещает курсор влево # столбцов
func (cursor) MakeCursorLeft(left int) string {
	return "\u001B[" + strconv.Itoa(left) + "D"
}

// MakeCursorStartDown перемещает курсор в начало следующей строки, на # строки вниз
func (cursor) MakeCursorStartDown(down int) string {
	return "\u001B[" + strconv.Itoa(down) + "E"
}

// MakeCursorStartUp перемещает курсор в начало следующей строки, на # строки вниз
func (cursor) MakeCursorStartUp(up int) string {
	return "\u001B[" + strconv.Itoa(up) + "F"
}

// MakeCursorColumn перемещает курсор в столбец #
func (cursor) MakeCursorColumn(column int) string {
	return "\u001B[" + strconv.Itoa(column) + "G"
}

// Deprecated:
// GetCursorPosition запросить позицию курсора (сообщается как ESC[#;#R)
func (cursor) GetCursorPosition() {
	print("\u001B[6n")
}
