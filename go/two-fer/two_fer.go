// Package twofer allows you to share somethin cool with your friends.
package twofer

// ShareWith function let you share with someone named name.
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
