package managers

type Channel struct {
	ID int
	Message chan byte
}

type ChannelManager struct {
	chans map[string]Channel
}
func NewChannelManager () *ChannelManager {
	return &ChannelManager{
		chans: make(map[string]Channel),
	}
}

func (cm *ChannelManager) CreateChanell(id int) *Channel{
	return &Channel{
		ID: id,
		Message: make(chan byte),
	} 
}




