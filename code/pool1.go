type Pool struct {
		New func() interface{}
}

func (p *Pool) Put(x interface{}) {}

func (p *Pool) Get() interface{} {}
