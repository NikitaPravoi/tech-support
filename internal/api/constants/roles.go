package constants

type Role int

const (
	SystemAdmin Role = iota + 1
	TechnicalSupport
	Developer
	Reporter
	Viewer
)
