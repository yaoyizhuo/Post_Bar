package mysql

import (
	"bulebell/models"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"strings"
)

// CreatePost 添加数据
func CreatePost(p *models.Post) (err error) {
	sqlStr := "insert into post(post_id,author_id,community_id,title,content) values(?,?,?,?,?)"
	_, err = db.Exec(sqlStr, p.ID, p.AuthorID, p.CommunityID, p.Title, p.Content)

	return
}

// GetPostDetail 根据 id 查询数据
func GetPostByID(id int64) (postDetail *models.Post, err error) {
	postDetail = new(models.Post)
	sqlStr := `select post_id,author_id,community_id,status,title,content,create_time from post where post_id=?`

	if err := db.Get(postDetail, sqlStr, id); err != nil {
		if err == sql.ErrNoRows {
			err = ErrInvalidID
		}
	}
	return postDetail, err
}

// GetPostList	获取帖子列表
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	posts = make([]*models.Post, 0, 2)

	sqlStr := `select post_id,author_id,community_id,status,title,content from post limit ?,?`
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}

// GetPostListByIDs 根据id列表查询帖子数据
func GetPostListByIDs(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id,author_id,community_id,status,title,content,create_time 
			from post where post_id in(?) order by find_in_set(post_id)`

	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
