package tables

import (
	"fmt"
	model_event "newsfeeder/httpd/models/event"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

func SetAttendance(c *gin.Context) {
	db := PostSQLConfig()
	event := model_event.PostEvent{}
	err := c.Bind(&event)
	if err != nil {
		fmt.Println(err)
		c.JSON(200, gin.H{
			"message": err,
		})
		return
	}
	stid := []int{}
	rows, err := db.Query("select id from attendance ORDER BY id DESC LIMIT 1")
	repos := []model_event.Event{}
	index := 1
	for rows.Next() {
		p := model_event.Event{}
		err = rows.Scan(&p.ID)
		if err != nil {
			fmt.Println(err)
		}
		repos = append(repos, p)
	}
	if repos != nil && len(repos) > 0 {
		index = repos[0].ID + 1
		fmt.Println(index, repos)
	}
	for _, n := range event.StudentId {
		id, _ := strconv.Atoi(n)
		stid = append(stid, id)
	}
	stime := time.Date(event.StartTime["year"], time.Month(event.StartTime["month"]),
		event.StartTime["day"], event.StartTime["hour"], event.StartTime["minute"], 0, 0, time.UTC)
	etime := time.Date(event.EndTime["year"], time.Month(event.EndTime["month"]),
		event.EndTime["day"], event.EndTime["hour"], event.EndTime["minute"], 0, 0, time.UTC)

	fmt.Println(event.StartTime, stime, etime)
	mid, _ := strconv.Atoi(event.MasterId)
	lid, _ := strconv.Atoi(event.LectionId)
	res, err := db.Exec("insert into attendance (id, lectionid, masterid, starttime, endtime, students) values ( $1, $2,$3,$4,$5,$6)",
		index, lid, mid, stime, etime, pq.Array(stid))
	if err != nil {
		fmt.Println("create error:", err)
	}
	c.JSON(200, gin.H{
		"message": "Succes",
		"user":    &res,
	})
	defer db.Close()

}

func GetAllEvents(c *gin.Context) {
	db := PostSQLConfig()
	query := "select  attend.id, lection.name as lection, string_agg( DISTINCT master.surname || ' ' || master.name || ' ' || master.fathername,'') as master, " +
		"starttime, endtime, attend.students ,  array_agg(DISTINCT s.surname || ' ' || s.name || ' ' || s.fathername) as students, " +
		"array_agg(DISTINCT f.surname || ' ' || f.name || ' ' || f.fathername) as attend, " +
		"array_agg(DISTINCT k.surname || ' ' || k.name || ' ' || k.fathername) as apsent " +
		"FROM attendance attend " +
		"INNER JOIN lection ON attend.lectionid = lection.id " +
		"INNER JOIN master ON attend.masterid = master.id " +
		"LEFT JOIN     students as s ON s.id = ANY(attend.students) " +
		"LEFT JOIN     students as f ON f.id = ANY(attend.attend) " +
		"LEFT JOIN     students as k ON k.id = ANY(attend.upsent) " +
		"GROUP BY attend.id,master.name,starttime,endtime,lection.name;"
	events := []model_event.GetEvent{}
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	for rows.Next() {
		event := model_event.GetEvent{}
		var c pq.Int64Array
		var s pq.StringArray
		var attend pq.StringArray
		var apsent pq.StringArray
		err := rows.Scan(&event.ID, &event.Lection, &event.Master, &event.StartTime,
			&event.EndTime, &c, &s, &attend, &apsent)
		event.StudentId = pq.Int64Array(c)
		event.Students = pq.StringArray(s)
		event.Attend = pq.StringArray(attend)
		event.Apsent = pq.StringArray(apsent)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(event.Students)
		events = append(events, event)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{
			"message": "Succes",
			"user":    &events,
		})
	}
	defer db.Close()
}

func GetAllEventsForMaster(c *gin.Context) {
	db := PostSQLConfig()
	s := c.Request.URL.Query().Get("master")
	rows, err := db.Query("select  attend.id, lection.name as lection, string_agg( DISTINCT master.surname || ' ' || master.name || ' ' || master.fathername,'') as master, "+
		"starttime, endtime, attend.students ,  array_agg(DISTINCT s.surname || ' ' || s.name || ' ' || s.fathername) as students, "+
		"array_agg(DISTINCT f.surname || ' ' || f.name || ' ' || f.fathername) as attend, "+
		"array_agg(DISTINCT k.surname || ' ' || k.name || ' ' || k.fathername) as apsent "+
		"FROM attendance attend "+
		"INNER JOIN lection ON attend.lectionid = lection.id "+
		"INNER JOIN master ON attend.masterid = master.id "+
		"LEFT JOIN     students as s ON s.id = ANY(attend.students) "+
		"LEFT JOIN     students as f ON f.id = ANY(attend.attend) "+
		"LEFT JOIN     students as k ON k.id = ANY(attend.upsent) "+
		"where attend.masterid = $1"+
		"GROUP BY attend.id,master.name,starttime,endtime,lection.name;", s)
	events := []model_event.GetEvent{}
	// rows, err := db.Query(query)
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	for rows.Next() {
		event := model_event.GetEvent{}
		var c pq.Int64Array
		var s pq.StringArray
		var attend pq.StringArray
		var apsent pq.StringArray
		err := rows.Scan(&event.ID, &event.Lection, &event.Master, &event.StartTime,
			&event.EndTime, &c, &s, &attend, &apsent)
		event.StudentId = pq.Int64Array(c)
		event.Students = pq.StringArray(s)
		event.Attend = pq.StringArray(attend)
		event.Apsent = pq.StringArray(apsent)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(event.Students)
		events = append(events, event)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{
			"message": "Succes",
			"user":    &events,
		})
	}
	defer db.Close()
}

func GetAllEventsForStudent(c *gin.Context) {
	db := PostSQLConfig()
	s := c.Request.URL.Query().Get("student")
	rows, err := db.Query("select  attend.id, lection.name as lection, string_agg( DISTINCT master.surname || ' ' || master.name || ' ' || master.fathername,'') as master, "+
		"starttime, endtime, attend.students ,  array_agg(DISTINCT s.surname || ' ' || s.name || ' ' || s.fathername) as students, "+
		"array_agg(DISTINCT f.surname || ' ' || f.name || ' ' || f.fathername) as attend, "+
		"array_agg(DISTINCT k.surname || ' ' || k.name || ' ' || k.fathername) as apsent "+
		"FROM attendance attend "+
		"INNER JOIN lection ON attend.lectionid = lection.id "+
		"INNER JOIN master ON attend.masterid = master.id "+
		"LEFT JOIN     students as s ON s.id = ANY(attend.students) "+
		"LEFT JOIN     students as f ON f.id = ANY(attend.attend) "+
		"LEFT JOIN     students as k ON k.id = ANY(attend.upsent) "+
		"where s.id = $1"+
		"GROUP BY attend.id,master.name,starttime,endtime,lection.name;", s)
	events := []model_event.GetEvent{}
	// rows, err := db.Query(query)
	if err != nil {
		fmt.Println("ERROR IS1", err)
		createUserTable()
	}
	defer rows.Close()
	for rows.Next() {
		event := model_event.GetEvent{}
		var c pq.Int64Array
		var s pq.StringArray
		var attend pq.StringArray
		var apsent pq.StringArray
		err := rows.Scan(&event.ID, &event.Lection, &event.Master, &event.StartTime,
			&event.EndTime, &c, &s, &attend, &apsent)
		event.StudentId = pq.Int64Array(c)
		event.Students = pq.StringArray(s)
		event.Attend = pq.StringArray(attend)
		event.Apsent = pq.StringArray(apsent)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(event.Students)
		events = append(events, event)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	} else {
		c.JSON(200, gin.H{
			"message": "Succes",
			"user":    &events,
		})
	}
	defer db.Close()
}
