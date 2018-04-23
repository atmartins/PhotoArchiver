package archive

// ArchStats provices statistics about the main archival operation.
type ArchStats struct {
	NumFound   int
	NumError   int
	NumSuccess int
	NumDupes   int
	NumUnique  int
}
