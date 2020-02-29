package dbops

import "log"

func AddVideoDeletionRecord(vid string) error {
	stmtIns, err := dbConn.Prepare("INSERT  into video_del_rec(video_id) VALUE (?)")
	defer stmtIns.Close()
	if err != nil {
		log.Printf("AddVideoDeletionRecord database error %v\n", err)
		return err
	}
	_, err = stmtIns.Exec(vid)
	if err != nil {
		log.Printf("AddVideoDeletionRecord error %v\n", err)
		return err
	}
	return nil
}

func GetVideoDeletionRecord(count int) ([]string, error) {

	var ids []string
	stmtOut, err := dbConn.Prepare("SELECT video_id from video_del_rec limit ?")
	if err != nil {
		log.Printf("GetVideoDeletionRecord database error %v\n", err)
		return ids, err
	}
	rows, err := stmtOut.Query(count)
	if err != nil {
		log.Printf("GetVideoDeletionRecord  error %v\n", err)
		return ids, err
	}
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return ids, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func DelVideoDeletionRecord(vid string) error  {
	stmtDel,err := dbConn.Prepare("DELETE FROM video_del_rec where video_id=?")
	defer stmtDel.Close()
	if err != nil {
		log.Printf("DeleteVideoDeletionRecord database error %v\n", err)
		return err
	}
	_,err = stmtDel.Exec(vid)
	if err != nil {
		log.Printf("DeleteVideoDeletionRecord error %v\n", err)
		return err
	}
	return nil
}