package danbooru

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"quantum-exposer/internal/domain"
	"quantum-exposer/internal/usecase"
	"strings"
)

type PostDTO struct {
	ID        int64  `json:"id"`
	CreatedAt string `json:"created_at"`
	FileUrl   string `json:"file_url"`
	TagString string `json:"tag_string"`
	Rating    string `json:"rating"`
	Score     int    `json:"score"`
}

type TagDTO struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	PostCount int    `json:"post_count"`
	Category  int    `json:"category"`
}

type ArtistDTO struct {
	ID         int64    `json:"id"`
	Name       string   `json:"name"`
	OtherNames []string `json:"other_names"`
	GroupName  string   `json:"group_name"`
	Url        string   `json:"url"`
	PostCount  int      `json:"post_count"`
}

type DanbooruAPIRepository struct {
	Client   *http.Client
	BaseURL  string
	Username string
	APIKey   string
}

func NewDanbooruAPIRepository(client *http.Client, baseURL, username, apiKey string) usecase.PostRepository {
	return &DanbooruAPIRepository{
		Client:   client,
		BaseURL:  baseURL,
		Username: username,
		APIKey:   apiKey,
	}
}

func (r *DanbooruAPIRepository) FetchByTags(tags []string) ([]domain.Post, error) {
	return r.FetchPosts(usecase.SearchCriteriaPost{
		Tags:  tags,
		Page:  1,
		Limit: 40,
	})
}

func (r *DanbooruAPIRepository) FetchPosts(criteria usecase.SearchCriteriaPost) ([]domain.Post, error) {
	limit := criteria.Limit
	if limit == 0 || limit > 100 {
		limit = 40
	}

	page := criteria.Page
	if page == 0 {
		page = 1
	}

	requestURL := fmt.Sprintf("%s/posts.json?tags=%s&limit=%d&page=%d&login=%s&api_key=%s",
		r.BaseURL,
		url.QueryEscape(strings.Join(criteria.Tags, " ")),
		criteria.Limit,
		criteria.Page%100,
		r.Username,
		r.APIKey,
	)

	if criteria.Random {
		requestURL += "&random=true"
	}

	fmt.Println("requestURL:", requestURL)

	resp, err := r.Client.Get(requestURL)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)

		return nil, fmt.Errorf("api request failed with status %d: %s", resp.StatusCode, string(body))
	}

	var dtos []PostDTO

	if err := json.NewDecoder(resp.Body).Decode(&dtos); err != nil {
		return nil, fmt.Errorf("failed to decode JSON: %w", err)
	}

	var posts []domain.Post

	for _, dto := range dtos {
		posts = append(posts, r.mapPostDTOToDomain(dto))
	}

	return posts, nil
}

func (r *DanbooruAPIRepository) FetchPostByID(id int64) (domain.Post, error) {
	requestURL := fmt.Sprintf("%s/posts/%d.json?login=%s&api_key=%s",
		r.BaseURL,
		id,
		r.Username,
		r.APIKey,
	)

	fmt.Println("requestURL:", requestURL)

	resp, err := r.Client.Get(requestURL)

	if err != nil {
		return domain.Post{}, fmt.Errorf("failed to execute request: %w", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusNotFound {
		return domain.Post{}, fmt.Errorf("post not found")
	}

	if resp.StatusCode != http.StatusOK {
		return domain.Post{}, fmt.Errorf("API request failed with status %d", resp.StatusCode)
	}

	var dto PostDTO

	if err := json.NewDecoder(resp.Body).Decode(&dto); err != nil {
		return domain.Post{}, fmt.Errorf("failed to decode JSON: %w", err)
	}

	return r.mapPostDTOToDomain(dto), nil
}

func (r *DanbooruAPIRepository) FetchTagsByName(criteria usecase.SearchCriteriaTag) ([]domain.Tag, error) {
	requestURL := fmt.Sprintf("%s/tags.json?search[name]=%s&limit=%d&page=%d&hide_empty=yes&login=%s&api_key=%s",
		r.BaseURL,
		url.QueryEscape(strings.Join(criteria.Names, ",")),
		criteria.Limit,
		criteria.Page,
		r.Username,
		r.APIKey,
	)

	if criteria.Order != "" {
		requestURL += "&order=" + url.QueryEscape(criteria.Order)
	}

	resp, err := r.Client.Get(requestURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var tagDTOs []TagDTO

	if err := json.Unmarshal(body, &tagDTOs); err != nil {
		return nil, err
	}

	var tags []domain.Tag

	for _, dto := range tagDTOs {
		tags = append(tags, r.mapTagDTOToDomain(dto))
	}

	return tags, nil
}

func (r *DanbooruAPIRepository) FetchTagsByNameMatches(criteria usecase.SearchCriteriaTag) ([]domain.Tag, error) {
	requestURL := fmt.Sprintf("%s/tags.json?search[name_matches]=%s*&limit=%d&page=%d&hide_empty=yes&login=%s&api_key=%s",
		r.BaseURL,
		url.QueryEscape(criteria.NamePrefix), // namePrefix will be 'blue_h'
		criteria.Limit,
		criteria.Page,
		r.Username,
		r.APIKey,
	)

	if criteria.Order != "" {
		requestURL += "&order=" + url.QueryEscape(criteria.Order)
	}

	resp, err := r.Client.Get(requestURL)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var tagDTOs []TagDTO

	if err := json.Unmarshal(body, &tagDTOs); err != nil {
		return nil, err
	}

	var tags []domain.Tag

	for _, dto := range tagDTOs {
		tags = append(tags, r.mapTagDTOToDomain(dto))
	}

	return tags, nil
}

func (r *DanbooruAPIRepository) FetchArtists(criteria usecase.SearchCriteriaArtist) ([]domain.Artist, error) {
	requestUrl := fmt.Sprintf("%s/artists.json?search[is_banned]=true&search[order]=name&search[query]=%s&limit=%d&page=%d&&login=%s&api_key=%s",
		r.BaseURL,
		criteria.Name,
		criteria.Limit,
		criteria.Page,
		r.Username,
		r.APIKey,
	)

	resp, err := r.Client.Get(requestUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var artistDTOs []ArtistDTO

	if err := json.Unmarshal(body, &artistDTOs); err != nil {
		return nil, err
	}

	var artists []domain.Artist

	for _, dto := range artistDTOs {
		artists = append(artists, r.mapArtistDTOToDomain(dto))
	}

	return artists, nil
}
