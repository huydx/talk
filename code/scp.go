func main() {
	replicas := []host{
		"host1",
		"host2",
		"host3"
	}

	for _, r := range replicas {
		if err := scp.CopyPath("/data", "/data", session); err != nil {
			// retry until success
		}
	}
}
