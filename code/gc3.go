func poolCleanup() {
[...]
// Drop victim caches from all pools.
for _, p := range oldPools {
   p.victim = nil
   p.victimSize = 0
}

// Move primary cache to victim cache.
for _, p := range allPools {
   p.victim = p.local
   p.victimSize = p.localSize
   p.local = nil
   p.localSize = 0
}

// The pools with non-empty primary caches now have non-empty
// victim caches and no pools have primary caches.
oldPools, allPools = allPools, nil
[...]
}
