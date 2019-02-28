func (c *Cache) Apply(l *raft.Log) interface{} {
	var cm command
	if err := json.Unmarshal(l.Data, &c); err != nil {return err}
	switch cm.Op {
	case "set": return c.applySet(cm.Key, cm.Value)
	}
}

func (c *Cache) Snapshot() (raft.FSMSnapshot, error) {
	o := make(map[int]int)
	for k, v := range c.m {
		o[k] = v
	}
	return &snapshot{m: o}, nil
}

func (c *Cache) Restore(rc io.ReadCloser) error {
	o := make(map[int]int)
	if err := json.NewDecoder(rc).Decode(&o); err != nil {return err}
	c.m = o
	return nil
}
