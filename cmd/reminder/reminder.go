package reminder

import "io"

func New() *Reminder {
	return &Reminder{}
}

func Handler(body io.ReadCloser) error {
	return nil
}
