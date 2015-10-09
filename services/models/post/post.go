package post

import (
	"time"

	"github.com/jmcvetta/neoism"
	"github.com/nathandao/vantaa/core/vantaadb"
)

type Post struct {
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	Created     time.Time `json:"created"`
	PublishTime time.Time `json:"publish_time"`
	Status      string    `json:"status"`
	UserId      int       `json:"user_id"`
}

type PostAdapter struct {
	Title       string    `json:"p.title"`
	Content     string    `json:"p.content"`
	Created     time.Time `json:"p.created"`
	Updated     time.Time `json:"p.updated"`
	PublishTime time.Time `json:"p.publish_time"`
	Status      string    `json:"p.status"`
	UserId      int       `json:"id(u)"`
}

// Save or p.Save creates or update a Post
func (p *Post) Save() (*Post, error) {
	newp, err := CreatePost(p)
	if err != nil {
		return nil, err
	}
	return newp, nil
}

// Transform transform a PostAdapter into a Post struct
func (pa *PostAdapter) Transform() *Post {
	p := Post{
		Title:       pa.Title,
		Content:     pa.Content,
		Created:     pa.Created,
		PublishTime: pa.PublishTime,
		Status:      pa.Status,
		UserId:      pa.UserId,
	}
	return &p
}

// CreatePost creates a new Post
func CreatePost(*Post) (*Post, error) {
	res := []PostAdapter{}
	db := vantaadb.Connect()
	cq := neoism.CypherQuery{
		Statement: `MATCH (u:User)
                WHERE id(u) = {uid}
                CREATE (p:Post {
                  title:{title},
                  content:{content},
                  created:{created},
                  updated:{updated},
                  publish_time:{publish_time},
                  status:{status},
                })
                CREATE
                  (p)-[:written_by]->(u)
                  (u)-[:writes]->(p)
                RETURN
                  id(u),
                  p.title, p.content,
                  p.created,
                  p.updated,
                  p.publish_time,
                  p.status`,
		Parameters: neoism.Props{},
		Result:     &res,
	}
	if err := db.Cypher(&cq); err != nil {
		return nil, err
	}
	pa := &res[0]
	return pa.Transform(), nil
}
