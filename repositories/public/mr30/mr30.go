package mr30_repositories

func (r *mr30RepoDB) GetMr30(course_year, course_semester string) (*[]Mr30Repo, error) {

	mr30_info := []Mr30Repo{}
	query := "SELECT ID,COURSE_SEMESTER,COURSE_YEAR,COURSE_NO,COURSE_CREDIT,COURSE_PR,COURSE_COMMENT,COURSE_STUDY_DATETIME,COURSE_ROOM,COURSE_INSTRUCTOR,COURSE_EXAMDATE FROM SS000.RU30_SMART WHERE COURSE_YEAR = :param1 AND COURSE_SEMESTER = :param2"

	err := r.oracle_db.Select(&mr30_info, query, course_year, course_semester)
	if err != nil {
		return nil, err
	}

	return &mr30_info, nil
}
