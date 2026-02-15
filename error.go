// Copyright © 2026 Timothy E. Peoples

package constr

// Error is an error that's a string -- and can thus be a const value.
type Error string

// Error implements error for Error.
func (e Error) Error() string {
	return string(e)
}
