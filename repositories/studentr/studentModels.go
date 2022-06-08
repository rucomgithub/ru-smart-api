package studentr

import (
	"github.com/jmoiron/sqlx"
)

type (
	studentRepoDB struct {
		oracle_db *sqlx.DB
	}

	StudentProfileRepo struct {
		STD_CODE             string `db:"STD_CODE"`
		NAME_THAI            string `db:"NAME_THAI"`
		NAME_ENG             string `db:"NAME_ENG"`
		BIRTH_DATE           string `db:"BIRTH_DATE"`
		STD_STATUS_DESC_THAI string `db:"STD_STATUS_DESC_THAI"`
		CITIZEN_ID           string `db:"CITIZEN_ID"`
		REGIONAL_NAME_THAI   string `db:"REGIONAL_NAME_THAI"`
		STD_TYPE_DESC_THAI   string `db:"STD_TYPE_DESC_THAI"`
		FACULTY_NAME_THAI    string `db:"FACULTY_NAME_THAI"`
		MAJOR_NAME_THAI      string `db:"MAJOR_NAME_THAI"`
		WAIVED_NO            string `db:"WAIVED_NO"`
		WAIVED_PAID          string `db:"WAIVED_PAID"`
		WAIVED_TOTAL_CREDIT  int    `db:"WAIVED_TOTAL_CREDIT"`
		CHK_CERT_NAME_THAI   string `db:"CHK_CERT_NAME_THAI"`
		PENAL_NAME_THAI      string `db:"PENAL_NAME_THAI"`
		MOBILE_TELEPHONE     string `db:"MOBILE_TELEPHONE"`
		EMAIL_ADDRESS        string `db:"EMAIL_ADDRESS"`
	}

	PrepareTokenRepo struct {
		STD_CODE string `db:"STD_CODE"`
		STATUS   int `db:"STATUS"`
	}

	StudentRepoInterface interface {
		GetStudentProfile(studentCode string) (*StudentProfileRepo, error)
		Authentication(studentCode string) (*PrepareTokenRepo, error)
	}

)

func NewStudentRepo(oracle_db *sqlx.DB) StudentRepoInterface {
	return &studentRepoDB{ oracle_db: oracle_db }
}

