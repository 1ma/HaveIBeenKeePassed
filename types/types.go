package types

// A struct wrapping the relevant bits of
// data from a KeePass XML Entry element.
type Entry struct {
	Title    string
	Password string
}
