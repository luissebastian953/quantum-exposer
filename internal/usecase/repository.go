package usecase

import "quantum-exposer/internal/domain"

// PostRepository is the Port (interface) that the Use Case depends on.
// The inner layer (usecase) is defined by this contract.
type PostRepository interface {
	FetchByTags(tags []string) ([]domain.Post, error)
	FetchTagsByName(criteria SearchCriteriaTag) ([]domain.Tag, error)
	FetchTagsByNameMatches(criteria SearchCriteriaTag) ([]domain.Tag, error)
	FetchPosts(criteria SearchCriteriaPost) ([]domain.Post, error)
	FetchPostByID(id int64) (domain.Post, error)
	FetchArtists(criteria SearchCriteriaArtist) ([]domain.Artist, error)
}
