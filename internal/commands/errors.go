package commands

const (
	errDumpDirDoesntExist = strErr("Target dump directory doesn't exist")
	errFeedDirDoesntExist = strErr("Target feed directory doesn't exist")
)

type strErr string

func (e strErr) Error() string {
	return string(e)
}
