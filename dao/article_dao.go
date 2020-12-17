package dao

import (
	"database/sql"
	"log"
	"time"

	"github.com/chengchaos/go-guest-book/entities"
	"github.com/chengchaos/go-guest-book/helpers"
)

type ArticleDao interface {
	GetArticleById(int) *entities.Article

	SaveArticle(*entities.Article)

	UpdateArticle(*entities.Article)

	DeleteArticleById(int)

	ListArticles(int) []*entities.Article
}

type ArticleDaoImpl struct {
	db *sql.DB
}

func NewArticleDao() ArticleDao {

	impl := &ArticleDaoImpl{
		db: GetSqlDB(),
	}
	return impl
}

func (impl *ArticleDaoImpl) GetArticleById(id int) (art *entities.Article) {

	selectSQL := "SELECT * FROM tb_article a WHERE a.id = ?"
	stmt, err := impl.db.Prepare(selectSQL)
	if err != nil {
		log.Println("GetArticleById => ", err)
		return
	}
	log.Println("stmt => ", stmt, "id => ", id)
	defer stmt.Close()

	row := stmt.QueryRow(id)

	art = &entities.Article{}
	var createTime string

	row.Scan(&art.ID, &art.UserID,
		&art.Title, &art.Content, &createTime)

	log.Println("createTime => ", createTime)

	createAt, err := time.Parse("2006-01-02T15:04:05+08:00", createTime)
	if err != nil {
		log.Fatalln(err)
	}
	art.CreateAt = createAt

	log.Println("art => ", art)

	return art

}

func (impl *ArticleDaoImpl) SaveArticle(art *entities.Article) {

	insertSQL := "INSERT INTO tb_article SET user_id = ?, " +
		"title = ?, content = ?"

	stmt, err := impl.db.Prepare(insertSQL)

	if err != nil {
		log.Println(err)
		return
	}

	defer stmt.Close()

	rest, err := stmt.Exec(art.UserID, art.Title, art.Content)
	if err != nil {
		log.Println(err)
		return
	}
	lastInsertID, err := rest.LastInsertId()

	if err != nil {
		log.Println(err)
		return
	}
	log.Println("lastInsertID  => ", lastInsertID)
	art.ID = int(lastInsertID)

}

func (impl *ArticleDaoImpl) UpdateArticle(art *entities.Article) {

	updateSQL := "UPDATE tb_article SET  " +
		"title = ?, content = ? WHERE id = ?"

	stmt, err := impl.db.Prepare(updateSQL)

	if err != nil {
		log.Println(err)
		return
	}

	defer stmt.Close()

	rest, err := stmt.Exec(art.Title, art.Content, art.ID)
	if err != nil {
		log.Println(err)
		return
	}
	rowsAffected, err := rest.RowsAffected()

	if err != nil {
		log.Println(err)
		return
	}
	log.Println("rowsAffected  => ", rowsAffected)

}

func (impl *ArticleDaoImpl) DeleteArticleById(id int) {

	deleteSQL := "DELETE FROM tb_article WHERE id = ?"

	stmt, err := impl.db.Prepare(deleteSQL)

	if err != nil {
		log.Println(err)
		return
	}

	defer stmt.Close()

	rest, err := stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return
	}
	rowsAffected, err := rest.RowsAffected()

	if err != nil {
		log.Println(err)
		return
	}
	log.Println("rowsAffected  => ", rowsAffected)

}

// ListArticles(int) []*entities.Article
func (impl *ArticleDaoImpl) ListArticles(page int) (articles []*entities.Article) {

	countSQL := "SELECT COUNT(id) FROM tb_article"
	row := impl.db.QueryRow(countSQL)
	var countValue int

	row.Scan(&countValue)

	log.Println("countValue =>", countValue)

	selectSQL := "SELECT id, user_id, title, content, create_by " +
		"FROM tb_article a order by id desc limit ? offset ? "
	if page <= 0 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit

	stmt, err := impl.db.Prepare(selectSQL)
	if err != nil {
		log.Println(err)
		return
	}

	defer stmt.Close()

	articles = make([]*entities.Article, 0)
	rows, err := stmt.Query(limit, offset)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	for rows.Next() {
		art := &entities.Article{}
		var createTime string
		rows.Scan(&art.ID, &art.UserID, &art.Title, &art.Content, &createTime)
		art.CreateAt = helpers.String2Time(createTime)
		articles = append(articles, art)
	}

	return articles

}
