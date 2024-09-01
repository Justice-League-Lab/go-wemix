package channel

var Txchan chan struct{}

func Init() {
	Txchan = make(chan struct{})
}
