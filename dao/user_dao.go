package dao

import (
	"database/sql"
	"log"

	"github.com/chengchaos/go-guest-book/entities"
)

type UserDao interface {
	GetUserByName(string) *entities.User
}

type UserDaoImpl struct {
	db *sql.DB
}

func NewUserDao() UserDao {
	impl := &UserDaoImpl{
		db: GetSqlDB(),
	}
	return impl
}

func (impl *UserDaoImpl) GetUserByName(name string) (user *entities.User) {

	selectSQL := "SELECT * FROM tb_user a WHERE a.username = ?"
	stmt, err := impl.db.Prepare(selectSQL)
	if err != nil {
		log.Println("Prepare =>", err)
		return
	}
	defer stmt.Close()

	row := stmt.QueryRow(name)
	user = &entities.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.State)
	return

}
