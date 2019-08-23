type Pool struct {
	list []interface{} // offset known to runtime
	mu   Mutex         // guards list
}
func runtime_registerPool(*Pool)
func (p *Pool) Put(x interface{}) {
	...	
	p.mu.Lock()
	p.list = append(p.list, x)
	p.mu.Unlock()
}
func (p *Pool) Get() interface{} {
	...
	p.mu.Lock()
	var x interface{}
	x = p.list[n-1]
	p.list = p.list[:n-1]
	p.mu.Unlock()
	...	
	return x
}
