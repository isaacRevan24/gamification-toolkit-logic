package model

type ControllerError struct {
	err    error
	status int
}
