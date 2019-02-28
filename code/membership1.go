func main() {
	var (
		seedNode = os.Getenv("SEED")
	)
	flag.Parse()
	conf := memberlist.DefaultLocalConfig()
	list, err := memberlist.Create(conf)
	conf.Delegate = &delegate{}
	if err != nil {
		panic("Failed to create memberlist: " + err.Error())
	}
	_, err = list.Join([]string{seedNode})
	if err != nil {
		panic("Failed to join cluster: " + err.Error())
	}
	for {
		time.Sleep(time.Second)
		for _, member := range list.Members() {
			fmt.Printf("Member: %s %s:%d\n", member.Name, member.Addr, member.Port)
		}
	}
}
