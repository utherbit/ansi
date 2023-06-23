package ansi

import "strconv"

// Format форматирование строки
var Format = format{
	Reset:         "\u001B[0m",
	Bold:          formatSwitch{ON: "\u001B[1m", OFF: "\u001B[22m"},
	Dim:           formatSwitch{ON: "\u001B[2m", OFF: "\u001B[22m"},
	Italic:        formatSwitch{ON: "\u001B[3m", OFF: "\u001B[23m"},
	Underlined:    formatSwitch{ON: "\u001B[4m", OFF: "\u001B[24m"},
	Flashing:      formatSwitch{ON: "\u001B[5m", OFF: "\u001B[25m"},
	Reversible:    formatSwitch{ON: "\u001B[7m", OFF: "\u001B[27m"},
	Invisible:     formatSwitch{ON: "\u001B[8m", OFF: "\u001B[28m"},
	Strikethrough: formatSwitch{ON: "\u001B[9m", OFF: "\u001B[29m"},

	Colors: color{
		Background: commandColors{
			Standard: colorsStandard{
				Black:   "\u001B[40m",
				Red:     "\u001B[41m",
				Green:   "\u001B[42m",
				Yellow:  "\u001B[43m",
				Blue:    "\u001B[44m",
				Purple:  "\u001B[45m",
				Cyan:    "\u001B[46m",
				White:   "\u001B[47m",
				Default: "\u001B[49m",
			},

			commandCODE: "\u001B[48;5;",
			commandRGB:  "\u001B[48;2;",
		},
		Foreground: commandColors{
			Standard: colorsStandard{
				Black:   "\u001B[30m",
				Red:     "\u001B[31m",
				Green:   "\u001B[32m",
				Yellow:  "\u001B[33m",
				Blue:    "\u001B[34m",
				Purple:  "\u001B[35m",
				Cyan:    "\u001B[36m",
				White:   "\u001B[37m",
				Default: "\u001B[39m",
			},

			commandCODE: "\u001B[38;5;",
			commandRGB:  "\u001B[38;2;",
		},
	},
}

type format struct {
	Reset string // сбросить все режимы (стили и цвета)

	Bold          formatSwitch // полужирный режим.
	Dim           formatSwitch // тусклый/слабый режим.
	Italic        formatSwitch // курсивный режим.
	Underlined    formatSwitch // режим подчеркивания.
	Flashing      formatSwitch // режим мигания
	Reversible    formatSwitch // инверсный/реверсивный режим
	Invisible     formatSwitch // скрытый/невидимый режим
	Strikethrough formatSwitch // режим зачеркивания.

	Colors color
}

type formatSwitch struct {
	ON  string
	OFF string
}

type color struct {
	Background commandColors
	Foreground commandColors
}
type commandColors struct {
	Standard colorsStandard

	commandCODE string
	commandRGB  string
}

type colorsStandard struct {
	Black   string //	30	40	Черный
	Red     string //	31	41	Красный
	Green   string //	32	42	Зеленый
	Yellow  string //	33	43	Желтый
	Blue    string //	34	44	Синий
	Purple  string //	35	45	Пурпурный
	Cyan    string //	36	46	Голубой
	White   string //	37	47	Белый
	Default string //	39	49	По умолчанию
}

// SetCode code цвета от 0 до 255
// Таблица начинается с исходных 16 цветов (0-15).
// Исходные 216 цветов (16-231) или сформированы значением RGB 3 бит на канал, смещенным на 16, упакованным в одно значение.
// Последние 24 цвета (232-255) представляют собой оттенки серого, начиная с оттенка немного светлее черного и заканчивая оттенком немного темнее белого.
func (c commandColors) SetCode(code int) string {
	return c.commandCODE + strconv.Itoa(code) + "m"

}
func (c commandColors) SetRGB(r, g, b int) string {
	return c.commandRGB + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
}
