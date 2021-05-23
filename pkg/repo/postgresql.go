package repo

import (
	"database/sql"
	"dateservice/pkg/models"
	_ "github.com/lib/pq"
	"log"
	"time"
)

type Repository interface {
	AddUser(name string) error
	AddComment(videoId int, userId int, text string, date string) error
	AddLike(videoId int, userId int, date string) error
	GetLikesByDate(currentDate time.Time) ([]models.Likes, error)
	GetCommentsByDate(currentDate time.Time) ([]models.Comments, error)
	GetActivityByDate(currentDate time.Time) ([]models.User, error)
}
type repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) (Repository, error) {
	return &repo{
		db: db,
	}, nil
}
func (r *repo) AddUser(name string) error {
	_, err := r.db.Exec("insert into users (name) values ($1)", name)
	if err != nil {
		log.Println(err)
	}
	log.Println("add new user to database")
	return nil
}
func (r *repo) AddComment(videoId int, userId int, text string, date string) error {
	_, err := r.db.Exec("insert into comments (video_id, text, user_id, date) values ($1, $2, $3, $4)", videoId, text, userId, date)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("add new comment to database")
	return nil
}
func (r *repo) AddLike(videoId int, userId int, date string) error {
	_, err := r.db.Exec("insert into likes (video_id, user_id, date) values ($1, $2, $3)", videoId, userId, date)
	if err != nil {
		log.Println(err)
	}
	log.Println("add new like to database")
	return nil
}
func (r *repo) GetLikesByDate(currentDate time.Time) ([]models.Likes, error) {
	rows, err := r.db.Query("SELECT * FROM likes WHERE date = $1", currentDate.Format("01-02-2006"))
	if err != nil {
		log.Println("query likes error")
		return nil, err
	}
	defer rows.Close()
	var likesArr []models.Likes

	for rows.Next() {
		l := models.Likes{}
		err = rows.Scan(&l.ID, &l.VideoID, &l.UserID, &l.Date)
		if err != nil {
			log.Println("Scan likes error", err)
			return nil, err
		}
		likesArr = append(likesArr, l)

	}
	return likesArr, nil
}
func (r *repo) GetCommentsByDate(currentDate time.Time) ([]models.Comments, error) {
	rows, err := r.db.Query("SELECT * FROM comments WHERE date = $1", currentDate.Format("01-02-2006"))
	if err != nil {
		log.Println("query comments error")
		return nil, err
	}
	defer rows.Close()
	var commentsArr []models.Comments

	for rows.Next() {
		c := models.Comments{}
		err = rows.Scan(&c.ID, &c.VideoID, &c.Text, &c.UserID, &c.Date)
		if err != nil {
			log.Println("Scan comments error", err)
			return nil, err
		}
		commentsArr = append(commentsArr, c)

	}
	return commentsArr, nil
}
func (r *repo) GetActivityByDate(currentDate time.Time) ([]models.User, error) {
	comments, err := r.GetCommentsByDate(currentDate)
	likes, err := r.GetLikesByDate(currentDate)
	if err != nil {
		return nil, err
	}
	rows, err := r.db.Query("SELECT * FROM users")
	if err != nil {
		log.Println("query activiry error")
		return nil, err
	}
	defer rows.Close()
	var users []models.User

	for rows.Next() {
		u := models.User{}
		err = rows.Scan(&u.ID, &u.Name)
		if err != nil {
			log.Println("Scan comments error", err)
			return nil, err
		}
		users = append(users, u)
	}
	for i := 0; i < len(users); i++ {
		for j := 0; j < len(comments); j++ {
			if users[i].ID == comments[j].UserID {
				users[i].Comments = comments
			}
		}
		for j := 0; j < len(likes); j++ {
			if users[i].ID == likes[j].UserID {
				users[i].Likes = likes
			}
		}
	}
	return users, nil
}
