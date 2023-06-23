package ansi

// https://gist.github.com/fnky/458719343aabd01cfb17a3a4f7296797

var Sequences = sequences{
	ESC: "\u001B",
	CSI: "\u009B",
	DCS: "\u0090",
	OSC: "\u009D",
}

type sequences struct {
	ESC string // последовательность, начинающаяся с ESC( \x1B)
	CSI string // Ввод управляющей последовательности: последовательность, начинающаяся с ESC [или CSI ( \x9B)
	DCS string // Строка управления устройством: последовательность, начинающаяся с ESC Pили DCS ( \x90)
	OSC string // Команда операционной системы: последовательность, начинающаяся с ESC ]или OSC ( \x9D)
}
