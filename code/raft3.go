// FSM provides an interface that can be implemented by
// clients to make use of the replicated log.
type FSM interface {
	Apply(*Log) interface{}

	Snapshot() (FSMSnapshot, error)

	Restore(io.ReadCloser) error
}
