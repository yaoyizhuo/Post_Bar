package models

import "time"

// 内存对齐概念
// 获取数据存储数据库
type Post struct {
	ID          int64     `json:"id" db:"post_id"`
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityID int64     `json:"community_id" db:"community_id" bind:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" bind:"required"`
	Content     string    `json:"content" db:"content" bind:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// 帖子详情获取
type PostDetail struct {
	AuthorName string `json:"author_name"`
	VoteNum      int64  `json:"vote_num"`
	*Post      `json:"post"`
	*Community `json:"community"`
}
