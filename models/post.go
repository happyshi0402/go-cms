package models

import "errors"

type Post struct {
	Model
	Id        int       `json:"id"        form:"id"        gorm:"default:''"`
	PostCode  string    `json:"post_code" form:"post_code" gorm:"default:''"`
	PostName  string    `json:"post_name" form:"post_name" gorm:"default:''"`
	PostSort  int       `json:"post_sort" form:"post_sort" gorm:"default:''"`
	Status    string    `json:"status"    form:"status"    gorm:"default:''"`
	CreateBy  string    `json:"create_by" form:"create_by" gorm:"default:''"`
	CreatedAt int       `json:"created_at"form:"created_at"gorm:"default:''"`
	UpdateBy  string    `json:"update_by" form:"update_by" gorm:"default:''"`
	UpdatedAt int       `json:"updated_at"form:"updated_at"gorm:"default:''"`
	Remark    string    `json:"remark"    form:"remark"    gorm:"default:''"`
	
}


func NewPost() (post *Post) {
	return &Post{}
}

func (m *Post) Pagination(offset, limit int, key string) (res []Post, count int) {
	query := Db
	if key != "" {
		query = query.Where("name like ?", "%"+key+"%")
	}
	query.Offset(offset).Limit(limit).Order("id desc").Find(&res)
	query.Model(Post{}).Count(&count)
	return
}

func (m *Post) Create() (newAttr Post, err error) {

    tx := Db.Begin()
	err = tx.Create(m).Error
	
	if err != nil{
       tx.Rollback()
	}else {
		tx.Commit()
	}

	newAttr = *m
	return
}

func (m *Post) Update() (newAttr Post, err error) {
    tx := Db.Begin()
	if m.Id > 0 {
		err = tx.Where("id=?", m.Id).Save(m).Error
	} else {
		err = errors.New("id参数错误")
	}
    if err != nil{
       tx.Rollback()
	}else {
		tx.Commit()
	}
	newAttr = *m
	return
}

func (m *Post) Delete() (err error) {
    tx := Db.Begin()
	if m.Id > 0 {
		err = tx.Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
    if err != nil{
       tx.Rollback()
	}else {
		tx.Commit()
	}
	return
}

func (m *Post) DelBatch(ids []int) (err error) {
    tx := Db.Begin()
	if len(ids) > 0 {
		err = tx.Where("id in (?)", ids).Delete(m).Error
	} else {
		err = errors.New("id参数错误")
	}
    if err != nil{
       tx.Rollback()
	}else {
		tx.Commit()
	}
	return
}

func (m *Post) FindById(id int) (post Post, err error) {
	err = Db.Where("id=?", id).First(&post).Error
	return
}

