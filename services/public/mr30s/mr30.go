package mr30s

import (
	"context"
	"encoding/json"
	"time"
)

var ctx = context.Background()

func (mr30 *mr30Services) GetMr30(course_year, course_semester string) (*Mr30Response, error) {

	mr30Response := Mr30Response{
		COURSE_YEAR:     "",
		COURSE_SEMESTER: "",
		RECORD:          []mr30Record{},
	}

	key := "mr30::" + course_year + "/" + course_semester
	mr30Cache, err := mr30.redis_cache.Get(ctx, key).Result()
	if err == nil {
		_ = json.Unmarshal([]byte(mr30Cache), &mr30Response)
		return &mr30Response, nil
	}

	mr30Repo, err := mr30.mr30Repo.GetMr30(course_year, course_semester)
	if err != nil {
		return nil, err
	}

	mr30Rec := []mr30Record{}
	for _, c := range *mr30Repo {

		mr30Rec = append(mr30Rec, mr30Record{
			ID:                   c.ID,
			COURSE_YEAR:          c.COURSE_YEAR,
			COURSE_SEMESTER:      c.COURSE_SEMESTER,
			COURSE_NO:            c.COURSE_NO,
			COURSE_METHOD:        c.COURSE_METHOD,
			COURSE_METHOD_NUMBER: c.COURSE_METHOD_NUMBER,
			DAY_CODE:             c.DAY_CODE,
			TIME_CODE:            c.TIME_CODE,
			ROOM_GROUP:           c.ROOM_GROUP,
			INSTR_GROUP:          c.INSTR_GROUP,
			COURSE_CREDIT:        c.COURSE_CREDIT,
			COURSE_METHOD_DETAIL: c.COURSE_METHOD_DETAIL,
			DAY_NAME_S:           c.DAY_NAME_S,
			TIME_PERIOD:          c.TIME_PERIOD,
			COURSE_ROOM:          c.COURSE_ROOM,
			COURSE_INSTRUCTOR:    c.COURSE_INSTRUCTOR,
			SHOW_RU30:            c.SHOW_RU30,
			COURSE_PR:            c.COURSE_PR,
			COURSE_COMMENT:       c.COURSE_COMMENT,
			COURSE_EXAMDATE:      c.COURSE_EXAMDATE,
		})
	}

	mr30Response = Mr30Response{
		COURSE_YEAR:     course_year,
		COURSE_SEMESTER: course_semester,
		RECORD:          mr30Rec,
	}

	if len(mr30Rec) != 0 {
		mr30JSON, _ := json.Marshal(&mr30Response)
		timeNow := time.Now()
		redisCacheMr30 := time.Unix(timeNow.Add(time.Hour*1).Unix(), 0)
		_ = mr30.redis_cache.Set(ctx, key, mr30JSON, redisCacheMr30.Sub(timeNow)).Err()
	}

	return &mr30Response, nil
}
 