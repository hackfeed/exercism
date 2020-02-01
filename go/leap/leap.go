// Package leap offers functions to work with leap years.
package leap

// IsLeapYear is checking whether year is leap or not.
func IsLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}
