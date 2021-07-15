package repository

import "github.com/rprajapati0067/golang-interview-test-main/api/domain"

var (
	m  = make(map[int]domain.Subscriber, 1)
	id = 1
)

func InMemoryDb(sub domain.Subscriber, id int) map[int]domain.Subscriber {

	if sub.Name() != "" && sub.Url() != "" {

		subs := domain.Subscriber{}
		subs.SetId(id)
		subs.SetName(sub.Name())
		subs.SetUrl(sub.Url())
		m = make(map[int]domain.Subscriber, id)
		m[id] = subs

	}

	return m
}
