package domain

type Tag struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	NameMatches string `json:"name_matches"`
	PostCount   int    `json:"post_count"`
	Category    int    `json:"category"` // 0: General, 3: Copyright, 4: Character, etc.
	Order       int    `json:"order"`    // exp: name, date, count, etc.
}
