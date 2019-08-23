// pool.go
func init() {
   runtime_registerPoolCleanup(poolCleanup)
}

func clearpools() {
	 poolcleanup()
}

func poolCleanup() {
	// assign pool to nil, size to 0
}

// mgc.go
func gcStart(trigger gcTrigger) {
   [...]
   // clearpools before we start the GC
   clearpools()
	 [...]
 }

