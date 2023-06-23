package main

import "testRegexp/ansi"

func main() {
	print(ansi.Format.Colors.Foreground.Standard.Cyan, ansi.Format.Bold.ON)
	print("Hello world")
	print(ansi.Format.Reset)
}
