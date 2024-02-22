module github.com/utopiagio/gioui/gio-x

go 1.21.4

require (
	git.sr.ht/~jackmordaunt/go-toast v1.0.0
	git.wow.st/gmp/jni v0.0.0-20210610011705-34026c7e22d0
	github.com/andybalholm/stroke v0.0.0-20221221101821-bd29b49d73f0
	github.com/esiqveland/notify v0.11.0
	github.com/godbus/dbus/v5 v5.0.6
	github.com/utopiagio/gioui/gio v0.0.0
	github.com/yuin/goldmark v1.4.13
	golang.org/x/exp v0.0.0-20240213143201-ec583247a57a
	golang.org/x/exp/shiny v0.0.0-20231226003508-02704c960a9b
	golang.org/x/image v0.15.0
	golang.org/x/sys v0.15.0
	golang.org/x/text v0.14.0
)

require (
	gioui.org/cpu v0.0.0-20210817075930-8d6a761490d2 // indirect
	gioui.org/shader v1.0.8 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-text/typesetting v0.1.0 // indirect
)

replace github.com/utopiagio/gioui/gio => ../gio
