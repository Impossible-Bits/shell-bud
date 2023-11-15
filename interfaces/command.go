package interfaces

type BashCommand struct {
	Name        string
	Description string
	Command     string
	Pinned      bool
}
