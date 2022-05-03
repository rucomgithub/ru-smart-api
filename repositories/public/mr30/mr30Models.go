package mr30_repositories

import "github.com/jmoiron/sqlx"

type (
	mr30RepoDB struct {
		oracle_db *sqlx.DB
	}

	Mr30Repo struct {
		ID                    string `db:"ID"`
		COURSE_SEMESTER       string `db:"COURSE_SEMESTER"`
		COURSE_YEAR           string `db:"COURSE_YEAR"`
		COURSE_NO             string `db:"COURSE_NO"`
		COURSE_CREDIT         string `db:"COURSE_CREDIT"`
		COURSE_PR             string `db:"COURSE_PR"`
		COURSE_COMMENT        string `db:"COURSE_COMMENT"`
		COURSE_STUDY_DATETIME string `db:"COURSE_STUDY_DATETIME"`
		COURSE_ROOM           string `db:"COURSE_ROOM"`
		COURSE_INSTRUCTOR     string `db:"COURSE_INSTRUCTOR"`
		COURSE_EXAMDATE       string `db:"COURSE_EXAMDATE"`
	}

	Mr30RepoInterface interface {
		GetMr30(course_year, course_semester string) (*[]Mr30Repo, error)
	}

)

func NewMr30Repo(oracle_db *sqlx.DB) Mr30RepoInterface {
	return &mr30RepoDB{ oracle_db : oracle_db}
}
