package wordwrap

// WrapperFunc takes a given input string, and returns some wrapped output. The wrapping may
// be altered by currying the wrapper function.
type WrapperFunc func(string) string

// Wrapper creates a curried wrapper function (see WrapperFunc) with the given options applied to
// it. Create a WrapperFunc and store it is a variable, then re-use it elsewhere.
//
//  limit      - The maximum number of characters for a line.
//  breakWords - Whether or not to break long words onto new lines.
//
// Example usage:
//
//  wrapper := wordwrap.Wrapper(10, false)
//  wrapped := wrapper("This string would be split onto several new lines")
func Wrapper(limit uint, breakWords bool) WrapperFunc {
	return func(input string) string {
		return ""
	}
}

// Indent a string with the given prefix at the start of each line.
func Indent(input string, prefix string) string {
	// @todo: implement
	return ""
}