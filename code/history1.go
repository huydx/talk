// Get gets a resource from the pool. 
// Returns nil on failure. 
func (p *Pool) Get() interface{} 

// Put puts the resource into the pool. 
// Returns false if the pool is full. 
func (p *Pool) Put(v interface{}) bool 

// SetCapacity sets capacity of the pool. 
// The method must not be invoked concurrently with other methods on 
the same pool object. 
// Local capacity refers to a private per-CPU pool. 
// Global capacity refers to a centralized shared pool. 
func (p *Pool) SetCapacity(local, global int) 

// Drain removes from the pool and returns all cached objects. 
// The method must not be invoked concurrently with other methods on 
the same pool object. 
func (p *Pool) Drain() []interface{} 
