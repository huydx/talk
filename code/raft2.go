type Cache struct {
	raft *raft.Raft
	m    map[int]int
}
