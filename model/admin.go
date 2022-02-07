package model

import "go-website/db"

type Admin struct {
	Id       int64  `json:"id" gorm:"column:id"` // 列名为 `id`
	NickName string `json:"nickname" gorm:"column:nickname"`
	Password string `json:"password" gorm:"column:password"`
}

func (admin *Admin) Insert() (err error) {
	result := db.Db.Create(&admin) // 新插入数据的id
	if result.Error != nil {
		err = result.Error
		return
	}
	return
}

func (admin *Admin) Find() (admins []Admin, err error) {
	err = db.Db.Find(&admins).Error
	if err != nil {
		return
	}
	return

}

func (admin *Admin) FindOne(id string) (err error) {
	err = db.Db.Where("id = ?", id).First(admin).Error
	if err != nil {
		return
	}
	return
}

//返回更新后的数据
func (admin *Admin) UpdateOne(nickname, password string) (err error) {
	//todo ? db.Model(&admin).Where || db.Where 区别, 为什么只有其他方法不需要加db.Model()
	err = db.Db.Model(&admin).Where("nickname = ?", nickname).Update("password", password).Error
	if err != nil {
		return
	}
	return
}

//返回更新后的数据
func (admin *Admin) Update(nickname, password string) (err error) {
	err = db.Db.Model(&admin).Where("nickname = ?", nickname).Updates(map[string]interface{}{"password": password}).Error
	if err != nil {
		return
	}
	return
}

//删除
func (admin *Admin) Delete(nickname string) (err error) {
	err = db.Db.Where("nickname = ?", nickname).Delete(&admin).Error
	if err != nil {
		return
	}
	return
}
