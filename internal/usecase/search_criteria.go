package usecase

// SearchCriteriaPost bundles all necessary parameters for searching posts.
type SearchCriteriaPost struct {
	Tags   []string
	Page   int
	Limit  int // Limit 100
	Random bool
}

type SearchCriteriaTag struct {
	NamePrefix string
	Names      []string
	Category   int
	Limit      int
	Page       int
	Order      string
}

type SearchCriteriaArtist struct {
	Name  string
	Limit int
	Page  int
}
