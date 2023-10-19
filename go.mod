module github.com/cdvelop/fileinput

go 1.20

require github.com/cdvelop/gotools v0.0.45

require (
	github.com/cdvelop/strings v0.0.2 // indirect
	github.com/cdvelop/timeserver v0.0.5 // indirect
	github.com/cdvelop/timetools v0.0.6 // indirect
)

require (
	github.com/cdvelop/input v0.0.39
	github.com/cdvelop/model v0.0.56
	github.com/cdvelop/object v0.0.15
	github.com/cdvelop/unixid v0.0.7
	golang.org/x/text v0.13.0 // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/cutkey => ../cutkey

replace github.com/cdvelop/strings => ../strings

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/unixid => ../unixid
