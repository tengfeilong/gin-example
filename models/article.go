package models

type Article struct {
	Model
	TagID         int    `json:"tag_id" gorm:"index"`
	Tag           Tag    `json:"tag"`
	CoverImageUrl string `json:"cover_image_url"`
	Title         string `json:"title"`
	Desc          string `json:"desc"`
	Content       string `json:"content"`
	CreatedBy     string `json:"created_by"`
	ModifiedBy    string `json:"modified_by"`
	State         int    `json:"state"`
}

func ExistArticleByID(id int) bool {
	var article Article
	db.Select("id").Where("id = ?", id).First(&article)
	if article.ID > 0 {
		return true
	}
	return false
}
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)
	return
}

func GetArticleList(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

func EditArticle(id int, data interface{}) bool {
	tx := db.Begin()
	if err := tx.Model(&Article{}).Where("id = ?", id).Updates(data).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func AddArticle(data map[string]interface{}) bool {
	tx := db.Begin()
	err := tx.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	}).Error
	if err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}
