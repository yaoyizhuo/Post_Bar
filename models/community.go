package models

type Community struct {
	Id   int    `json:"community_id" db:"community_id"`
	Name string `json:"community_name" db:"community_name"`
}

type CommunityDetail struct {
	Id           int    `json:"community_id" db:"community_id"`
	Name         string `json:"community_name" db:"community_name"`
	Introduction string `json:"introduction" db:"introduction"`
	CreateTime   string `json:"create_time" db:"create_time"`
}
