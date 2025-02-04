package storage

type Storage interface {
	Create(n *Notify) error
	Save(n *Notify) error
	Delete(n *Notify) error
	GetNotify(channelId string) (*Notify, error)
}

type Notify struct {
	Author   string
	Schedule string
	Title    string
}
