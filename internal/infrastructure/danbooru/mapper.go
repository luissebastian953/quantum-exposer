package danbooru

import (
	"quantum-exposer/internal/domain"
	"strings"
)

func (r *DanbooruAPIRepository) mapPostDTOToDomain(dto PostDTO) domain.Post {
	tags := strings.Fields(dto.TagString)

	return domain.Post{
		ID:      dto.ID,
		Tags:    tags,
		Rating:  dto.Rating,
		Score:   dto.Score,
		FileURL: dto.FileUrl,
		// add more fields as necessary
	}
}

func (r *DanbooruAPIRepository) mapTagDTOToDomain(dto TagDTO) domain.Tag {
	return domain.Tag{
		ID:        dto.ID,
		Name:      dto.Name,
		PostCount: dto.PostCount,
		Category:  dto.Category,
	}
}

func (r *DanbooruAPIRepository) mapArtistDTOToDomain(dto ArtistDTO) domain.Artist {
	return domain.Artist{
		ID:         dto.ID,
		Name:       dto.Name,
		OtherNames: dto.OtherNames,
		GroupName:  dto.GroupName,
		Url:        dto.Url,
		PostCount:  dto.PostCount,
	}
}
