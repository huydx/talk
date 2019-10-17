type Datapoint struct {
	Id      uint64
	Samples []Sample	
}

type Sample struct {
	Value float64
	Timestampt int64
}

type Storage interface {
	Put(datapoints []Datapoint)
	QueryRange(id uint64, fromTs int64, toTs int64) []Datapoint
}
