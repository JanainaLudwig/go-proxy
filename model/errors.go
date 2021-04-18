package model

import "strconv"

type ModelError struct {
	Status int
	Message string
}

func (m ModelError) Error() string {
	return "[ERROR " + strconv.Itoa(m.Status) + "] " + m.Message
}
