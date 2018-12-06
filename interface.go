package lumbarcheck

// Checker interface to check if given dsn is valid
type Checker interface {
	Check(dsn string) error
}
