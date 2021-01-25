package db

import "log"

type Tag struct {
	ID   int    `db:"tag_id"`
	Name string `db:"name"`
}

func CreateTag(name string) (err error) {
	_, err = DB.Exec("INSERT INTO tag (name) VALUES (?)", name)
	if err != nil {
		log.Println(err)
	}

	return
}

func EditTag(id int, name string) (err error) {
	log.Println(id, name)
	_, err = DB.Exec("UPDATE tag SET name = ? WHERE tag_id = ?", name, id)
	if err != nil {
		log.Println(err)
	}

	return
}

func GetTagById(id int) (Tag, error) {
	var tag Tag
	err := DB.Get(&tag, "SELECT * FROM tag WHERE tag_id = ?", id)
	if err != nil {
		log.Println(err)
	}
	log.Println(tag)

	return tag, err
}

func GetAllTags() ([]Tag, error) {
	tags := []Tag{}
	err := DB.Select(&tags, "SELECT * FROM tag")
	if err != nil {
		log.Println(err)
	}

	return tags, err
}
