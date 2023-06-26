# ANSI: Библиотека для работы с ANSI в Go

**ANSI** - это библиотека на языке Go, предназначенная для удобной работы с ANSI-кодами и функционалом командной строки. Она позволяет задавать формат, цвет фона и цвет текста выводимого текста.

## Установка

Для установки библиотеки ANSI выполните команду:

```shell
go get github.com/utherbit/ansi
```

## Использование

Импортируйте пакет ANSI в свой проект:

```go
import "github.com/utherbit/ansi"
```

### Задание формата, цвета фона и цвета текста

Для задания формата, цвета фона и цвета текста используйте соответствующие функции и константы из библиотеки ANSI. Пример использования:

```go
package main

import "github.com/utherbit/ansi"

func main() {
	print(ansi.Format.Colors.Background.SetRGB(255, 0, 0))
	print(ansi.Format.Colors.Foreground.Standard.Cyan)

	print(ansi.Format.Bold.ON, ansi.Format.Underlined.ON)
	print("Hello, world!")
	print(ansi.Format.Reset)
}

```

Вы также можете использовать другие методы и функции для задания цвета фона и цвета текста:

## Лицензия

Эта библиотека распространяется под лицензией [MIT](LICENSE).

## Вклад

Если вы нашли ошибку или хотите внести свой вклад в развитие библиотеки, пожалуйста, создайте issue или отправьте pull request на [GitHub-странице проекта](https://github.com/utherbit/ansi).

## Автор

Автор библиотеки ANSI - utherbit.
