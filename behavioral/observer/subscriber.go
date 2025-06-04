package observer

type subscriber struct {
	name string
	news []string
}

func NewSubscriber(name string) *subscriber {
	return &subscriber{name: name}
}

func (s *subscriber) UpdateNews(news string) {
	println(s.name + " received news: " + news)
	s.news = append(s.news, news)
}

func (s *subscriber) GetNews() []string {
	return s.news
}
