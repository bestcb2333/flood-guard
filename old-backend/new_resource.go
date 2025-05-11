package main

import "gorm.io/gorm"

type ListResourceDTO struct {
	Type   string `form:"type"`
	Region uint   `form:"region"`
	ListDTO
}

var NewGetResources = ListHandler[Resource](
	new(ListResourceDTO),
	func(query *gorm.DB, u *User, r *ListResourceDTO) *gorm.DB {
		query = query.Preload("Region", Select("id", "name"))
		if r.Type != "" {
			query = query.Where("type = ?", r.Type)
		}
		if r.Region != 0 {
			query = query.Where("region_id = ?", r.Region)
		}
		return query
	},
)

var NewAddResource = AddHandler(
	new(ResourceDTO),
	func(data *Resource, u *User, dto *ResourceDTO) *Resource {
		data.ResourceDTO = *dto
		return data
	},
)

var DeleteResource = DeleteHandler[Resource]()
