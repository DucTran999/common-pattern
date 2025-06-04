package observer

type Subscriber interface {
	UpdateNews(newPost string)
}

type Blogs struct {
	observers []Subscriber
}

func (s *Blogs) Subscribe(o Subscriber) {
	s.observers = append(s.observers, o)
}

func (s *Blogs) AddNews(newPost string) {
	for _, o := range s.observers {
		o.UpdateNews(newPost)
	}
}
