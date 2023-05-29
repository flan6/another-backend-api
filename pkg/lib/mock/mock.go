package mock

import "github.com/golang/mock/gomock"

type matchedBy struct {
	matcher func(x interface{}) bool
}

func MatchedBy(fn func(matcher interface{}) bool) gomock.Matcher {
	return &matchedBy{matcher: fn}
}

func (m *matchedBy) Matches(x interface{}) bool {
	return m.matcher(x)
}

func (m *matchedBy) String() string {
	return "value not expected"
}
