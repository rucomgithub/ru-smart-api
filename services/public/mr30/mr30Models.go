package mr30_services

import (
	mr30_repositories "RU-Smart-Workspace/ru-smart-api/repositories/public/mr30"

	"github.com/go-redis/redis/v8"
)

type (
	
	mr30Services struct {
		mr30Repo    mr30_repositories.Mr30RepoInterface
		redis_cache *redis.Client
	}

	Mr30Response struct {
		COURSE_YEAR           string `json:"course_year"`
		COURSE_SEMESTER       string `json:"course_semester"`
		RECORD				  []mr30Record
	}

	mr30Record struct {
		ID                    string `json:"id"`
		COURSE_SEMESTER       string `json:"course_semester"`
		COURSE_YEAR           string `json:"course_year"`
		COURSE_NO             string `json:"course_no"`
		COURSE_CREDIT         string `json:"course_credit"`
		COURSE_PR             string `json:"course_pr"`
		COURSE_COMMENT        string `json:"course_comment"`
		COURSE_STUDY_DATETIME string `json:"course_study_datetime"`
		COURSE_ROOM           string `json:"course_room"`
		COURSE_INSTRUCTOR     string `json:"course_instructor"`
		COURSE_EXAMDATE       string `json:"course_examdate"`
	}

	Mr30Request struct {
		Course_year     string `json:"course_year"`
		Course_semester string `json:"course_semester"`
	}

	Mr30ServiceInterface interface {
		GetMr30(course_year, course_semester string) (*Mr30Response, error)
	}
)

func NewMr30Services(mr30Repo mr30_repositories.Mr30RepoInterface, redis_cache *redis.Client) Mr30ServiceInterface {
	return &mr30Services{
		mr30Repo:    mr30Repo,
		redis_cache: redis_cache,
	}
}



