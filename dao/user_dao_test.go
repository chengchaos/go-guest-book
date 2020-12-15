package dao

import "testing"

func Test_GetUserByName(t *testing.T) {

	userDao := NewUserDao()
	user := userDao.GetUserByName("admin")
	t.Log("ID => ", user.ID)
	t.Log("Password => ", user.Password)
}
