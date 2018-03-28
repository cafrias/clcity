package mock

// Spy helper to spy on method activity.
type Spy struct {
	Invoked bool
	Args    []interface{}
	Ret     []interface{}
}
