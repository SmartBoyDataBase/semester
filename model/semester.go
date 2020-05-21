package model

import (
	"fmt"
	"sbdb-semester/infrastructure"
)

type Semester struct {
	Id    uint64 `json:"id"`
	Name  string `json:"name"`
	Start string `json:"start"`
	End   string `json:"end"`
}

func Get(id uint64) (Semester, error) {
	result := Semester{
		Id: id,
	}
	row := infrastructure.DB.QueryRow(`
	SELECT name, lower(date_range), upper(date_range)
	FROM semester;
	`)
	err := row.Scan(&result.Name, &result.Start, &result.End)
	return result, err
}

func Create(semester Semester) (Semester, error) {
	fmt.Println(semester)
	row := infrastructure.DB.QueryRow(`
	INSERT INTO semester(name, date_range)
	VALUES ($1, [$2, $3))
	RETURNING id;`, semester.Name, semester.Start[:10], semester.End[:10])
	err := row.Scan(&semester.Id)
	return semester, err
}

func All() ([]Semester, error) {
	rows, err := infrastructure.DB.Query(`
	SELECT id,name,lower(date_range), upper(date_range)
	FROM semester;
	`)
	if err != nil {
		return nil, err
	}
	var semesters []Semester
	for rows.Next() {
		var semester Semester
		err := rows.Scan(&semester.Id, &semester.Name, &semester.Start, &semester.End)
		if err != nil {
			return semesters, err
		}
		semesters = append(semesters, semester)
	}
	return semesters, nil
}
