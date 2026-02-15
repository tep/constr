// Copyright © 2026 Timothy E. Peoples

// Package constr provides a single type (Error) which is a string that
// implements error, and... since it's a string, values of type Error may
// be declared `const`.  Therefor, instead of:
//
//	var ErrBadThing = errors.New("something bad happened")
//
// ...you can declare:
//
//	const ErrBadThing = constr.Error("something bad happened")
//
// ...safe in the knowledge that nobody's ever going to futz with it.
//
// Yes, I know, this is absurdly pedantic... but, be honest... you've wanted
// to do it this way from the beginning, haven't you?  I know I have! In fact,
// I've written this same bit of code hundreds of times over the years but
// finally decided it needed a home of its own.
//
// NB. For those interested, the package name is a portmanteau of the
// words 'constant' and 'string' -- a mnemonic or sorts.
package constr

// Error is an error that's a string -- and can thus be a const value.
type Error string

// Sentinel is a convenience alias for Error -- since, in this context, there
// are times when the term "sentinel" is a more semantically accurate
// description of what's being declared.
type Sentinel = Error

// Error implements error for Error.
func (e Error) Error() string {
	return string(e)
}
