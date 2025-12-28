package domain

type Artist struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	OtherNames []string `json:"other_names"`
	GroupName  string   `json:"group_name"`
	Url        string   `json:"url"`
	PostCount  int      `json:"post_count"`
}
