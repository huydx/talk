func (p *Pool) Get() interface{} 

func (p *Pool) Put(v interface{}) bool 

func (p *Pool) SetCapacity(local, global int) 

func (p *Pool) Drain() []interface{} 
