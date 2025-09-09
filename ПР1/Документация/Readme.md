# Практическая работа 1

## Описание

Кроссплатформенное консольное приложение-калькулятор, реализованное на языке Go. Поддерживаются версии для Linux, macOS и Windows.

---

## Системные требования
- Go 1.18+
- Операционная система: Linux, macOS, Windows

---

## Сборка и запуск

### Готовые исполняемые файлы
В папке находятся готовые бинарные файлы для каждой платформы:
- **Linux:** `Calculator_linux`
- **macOS:** `Calculator_macOS`
- **Windows:** `Calculator_windows.exe`

#### Запуск
- **Linux/macOS:**
	```sh
	./Calculator_linux      # для Linux
	./Calculator_macOS      # для macOS
	```
	Если файл не запускается, дайте права на исполнение:
	```sh
	chmod +x Calculator_linux
	chmod +x Calculator_macOS
	```
- **Windows:**
	Двойной клик по `Calculator_windows.exe` или запуск через командную строку:
	```bat
	Calculator_windows.exe
	```

---

### Сборка из исходников

1. Установите Go: https://go.dev/dl/
2. Перейдите в папку с нужным исходником:
	 - `Calculator_linux.go` — для Linux
	 - `Calculator_macOS.go` — для macOS
	 - `Calculator_windows.go` — для Windows
3. Выполните команду:
	 ```sh
	 go build -o Calculator_<ваша_платформа> Calculator_<ваша_платформа>.go
	 ```
	 Пример для Linux:
	 ```sh
	 go build -o Calculator_linux Calculator_linux.go
	 ```

---

## Использование

После запуска калькулятора следуйте инструкциям в консоли. Введите выражение для вычисления и нажмите Enter.

---

## Техническая документация
- [Техническое задание](technical_task.pdf)
- [Описание работы](Readme.md)

---

## Авторы
- Команда ForestBeavers, ИКБО-66-23