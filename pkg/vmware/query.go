package vmware

import (
	"net/http"
)

// SearchQuery is query struct for search function
type SearchQuery struct {
	FilterHeader map[string][]string
}

func (s *SearchQuery) AddFilter(param string, values []string) {
	m := s.FilterHeader

	_, ok := m[param]
	if !ok {
		m[param] = values
	} else {
		m[param] = append(m[param], values...)
	}

	s.FilterHeader = m
}

func ToFilter(param string, values []string) map[string][]string {
	return map[string][]string{
		param: values,
	}
}

// NewSearchQueryDatacenters create datacenter filter SearchQuery
func NewSearchQueryDatacenters(datacenters []string) *SearchQuery {
	return &SearchQuery{
		FilterHeader: ToFilter("datacenters", datacenters),
	}
}

// NewSearchQueryHosts create host filter SearchQuery
func NewSearchQueryHosts(hosts []string) *SearchQuery {
	return &SearchQuery{
		FilterHeader: ToFilter("hosts", hosts),
	}
}

// AddSearchQuery add url parameter by SearchQuery
func AddSearchQuery(req *http.Request, query *SearchQuery) *http.Request {
	if query == nil {
		return req
	}

	q := req.URL.Query()
	if len(query.FilterHeader) != 0 {
		for key, values := range query.FilterHeader {
			for _, v := range values {
				q.Add(key, v)
			}
		}
	}

	req.URL.RawQuery = q.Encode()
	return req
}
