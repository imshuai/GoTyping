package model

import (
	"errors"
)

//Category 定义文章分类
type Category struct {
	ID          int64    `json:"id"`
	Name        string   `json:"name" xorm:"unique"`
	Alias       string   `json:"alias" xorm:"unique"`
	Keywords    []string `json:"keywords"`
	Description string   `json:"description"`
	Common
}

//NewCatagory 根据名称和别名新建分类
func NewCatagory(name, alias string) (*Category, error) {
	ca := &Category{
		Name:  name,
		Alias: alias,
	}
	_, err := db.InsertOne(ca)
	if err != nil {
		return &Category{}, err
	}
	return ca, nil
}

//GetCategoryByID 根据给定ID获得分类
func GetCategoryByID(id int64) (*Category, error) {
	ca := new(Category)
	exist, err := db.ID(id).Get(ca)
	if err != nil {
		return &Category{}, err
	}
	if !exist {
		return &Category{}, errors.New("cannot find category by id")
	}
	return ca, nil
}

//GetCategoryByName 根据给定分类名称获得分类
func GetCategoryByName(name string) (*Category, error) {
	ca := new(Category)
	exist, err := db.Where("`name`=?", name).Get(ca)
	if err != nil {
		return &Category{}, err
	}
	if !exist {
		return &Category{}, errors.New("cannot find category by name")
	}
	return ca, nil
}

//GetCategoryByAlias 根据给定分类别名获得分类
func GetCategoryByAlias(alias string) (*Category, error) {
	ca := new(Category)
	exist, err := db.Where("`alias`=?", alias).Get(ca)
	if err != nil {
		return &Category{}, err
	}
	if !exist {
		return &Category{}, errors.New("cannot find category by alias")
	}
	return ca, nil
}

//GetCategories 获取所有分类名称
func GetCategories(currentID int64) (map[string]bool, error) {
	rows, err := db.Asc("id").Rows(new(Category))
	defer rows.Close()
	if err != nil {
		return make(map[string]bool, 0), err
	}
	cats := make(map[string]bool, 0)
	for rows.Next() {
		cat := new(Category)
		rows.Scan(cat)
		cats[cat.Name] = func() bool {
			if cat.ID == currentID {
				return true
			}
			return false
		}()
	}
	return cats, nil
}

//SetKeywords 设置分类的关键词
func (cat *Category) SetKeywords(keywords ...string) (bool, error) {
	cat.Keywords = append(cat.Keywords, keywords...)
	_, err := db.ID(cat.ID).Cols("keywords").Update(cat)
	if err != nil {
		return false, err
	}
	return true, nil
}

//SetDescription 设置分类的详细描述
func (cat *Category) SetDescription(description string) (bool, error) {
	cat.Description = description
	_, err := db.ID(cat.ID).Cols("description").Update(cat)
	if err != nil {
		return false, err
	}
	return true, nil
}

//SetName 设置分类的名称
func (cat *Category) SetName(name string) (bool, error) {
	cat.Name = name
	_, err := db.ID(cat.ID).Cols("name").Update(cat)
	if err != nil {
		return false, err
	}
	return true, nil
}

//SetAlias 设置分类的别名
func (cat *Category) SetAlias(alias string) (bool, error) {
	cat.Alias = alias
	_, err := db.ID(cat.ID).Cols("alias").Update(cat)
	if err != nil {
		return false, err
	}
	return true, nil
}

//Delete 删除当前分类
func (cat *Category) Delete(confirm bool) (bool, error) {
	if confirm {
		_, err := db.ID(cat.ID).Delete(cat)
		if err != nil {
			return false, err
		}
		return true, nil
	}
	return false, errors.New("confirm needed")
}
