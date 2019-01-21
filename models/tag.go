package models

type Tag struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

/*func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now())
	return nil
}

func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now())

	return nil
}*/

func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

func AddTag(name string, state int, createdBy string) bool {
	tx := db.Begin()

	err := tx.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}).Error
	if err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}
func EditTag(id int, data interface{}) bool {
	tx := db.Begin()
	if err := tx.Model(&Tag{}).Where("id = ?", id).Updates(data).Error; err != nil {
		tx.Rollback()
		return false
	}
	tx.Commit()
	return true
}

func DeleteTag(id int) {
	db.Where("id = ?", id).Delete(&Tag{})
}

func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

func ExistTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name = ?", name).First(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}

func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).Find(&tag)
	if tag.ID > 0 {
		return true
	}
	return false
}
