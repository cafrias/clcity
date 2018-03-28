package mock

// Spy helper to spy on method activity.
type Spy struct {
	Invoked bool
	Calls   int
	Args    [][]interface{}
	Ret     [][]interface{}
}
