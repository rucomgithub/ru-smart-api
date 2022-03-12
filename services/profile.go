package services

func (s studentServices) GetStudentProfile(studentCode string) (studentProfileResponse *StudentProfileService, err error) {

	student := StudentProfileService{}

	STUDENT_CODE := studentCode

	sp, err := s.studentRepo.GetStudentProfile(STUDENT_CODE)
	if err != nil {
		return studentProfileResponse, err
	}

	student = StudentProfileService{
		STD_CODE:             sp.STD_CODE,
		NAME_THAI:            sp.NAME_THAI,
		NAME_ENG:             sp.NAME_ENG,
		BIRTH_DATE:           sp.BIRTH_DATE,
		STD_STATUS_DESC_THAI: sp.STD_STATUS_DESC_THAI,
		CITIZEN_ID:           sp.CITIZEN_ID,
		REGIONAL_NAME_THAI:   sp.REGIONAL_NAME_THAI,
		STD_TYPE_DESC_THAI:   sp.STD_TYPE_DESC_THAI,
		FACULTY_NAME_THAI:    sp.FACULTY_NAME_THAI,
		MAJOR_NAME_THAI:      sp.MAJOR_NAME_THAI,
		WAIVED_NO:            sp.WAIVED_NO,
		WAIVED_PAID:          sp.WAIVED_PAID,
		WAIVED_TOTAL_CREDIT:  sp.WAIVED_TOTAL_CREDIT,
		CHK_CERT_NAME_THAI:   sp.CHK_CERT_NAME_THAI,
		PENAL_NAME_THAI:      sp.PENAL_NAME_THAI,
		MOBILE_TELEPHONE:     sp.MOBILE_TELEPHONE,
		EMAIL_ADDRESS:        sp.EMAIL_ADDRESS,
	}

	studentProfileResponse = &student

	return studentProfileResponse, nil
}