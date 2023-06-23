package ansi

// Clear Функции стирания
var Clear = clear{
	ClearScreenCur2End:   "\u001B[0J",
	ClearScreenStart2Cur: "\u001B[1J",
	ClearScreen:          "\u001B[2J",
	ClearSaveStr:         "\u001B[3J",
	ClearLineCur2End:     "\u001B[0K",
	ClearLineStart2Cur:   "\u001B[1K",
	ClearLine:            "\u001B[2K",
}

type clear struct {
	ClearScreenCur2End   string // стереть от курсора до конца экрана
	ClearScreenStart2Cur string // стереть от курсора до начала экрана
	ClearScreen          string // стереть весь экран
	ClearSaveStr         string // стереть сохраненные строки
	ClearLineCur2End     string // стереть от курсора до конца строки
	ClearLineStart2Cur   string // стереть начало строки до курсора
	ClearLine            string // стереть всю строку
}

const (
// ClearDisplay         = "\u001B[J"  Стереть на дисплее (аналогично ESC[0J)
// Clear6               = "\u001B[K"  стереть в строке (то же, что и ESC[0K)
)
