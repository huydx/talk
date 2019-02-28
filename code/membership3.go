package main


var nodes []*memberlist.Node
http.HandleFunc("/deploy", func(writer http.ResponseWriter, request *http.Request) {
	for _, n := range nodes {
		list.SendBestEffort(n, []byte("deploy!"))
	}
})
go func() {
	http.ListenAndServe(":9999", nil)
}()
for {
	nodes = list.Members()
	time.Sleep(1 * time.Second)
}

