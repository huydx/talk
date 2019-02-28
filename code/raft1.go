conf := raft.DefaultConfig()
conf.LocalID = raft.ServerID(id) // <-- id could be host name
addr, err := net.ResolveTCPAddr("tcp", ":11111")
if err != nil {
	panic(err)
}
transport, err := raft.NewTCPTransport(":11111", addr, 3, 10*time.Second, os.Stderr)
if err != nil {
	panic(err)
}
snapshots, err := raft.NewFileSnapshotStore("/tmp", 1, os.Stderr) 
if err != nil {
	panic(err)
}
var logStore = raft.NewInmemStore()
var stableStore = raft.NewInmemStore()
c := &Cache{}
ra, err := raft.NewRaft(conf, c, logStore, stableStore, snapshots, transport)
if err != nil {
	panic(err)
}
c.raft = ra
return c
