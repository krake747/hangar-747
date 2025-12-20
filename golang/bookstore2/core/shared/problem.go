package shared

import "fmt"

type ProblemDetail struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Detail   string `json:"detail"`
	Status   int    `json:"status"`
	Instance string `json:"instance,omitempty"`
}

func (p ProblemDetail) Error() string {
	return fmt.Sprintf("%s: %s", p.Title, p.Detail)
}

func NewProblemDetail(status int, title, detail string) *ProblemDetail {
	return &ProblemDetail{
		Type:   "about:blank",
		Title:  title,
		Detail: detail,
		Status: status,
	}
}
