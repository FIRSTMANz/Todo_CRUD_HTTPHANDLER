package service

import (
	"database/sql"
	"fmt"
	"log"
	"projectfirsty/dba"
	"time"

	_ "github.com/lib/pq"
)

type TodoService struct {
	db *sql.DB
}

func NewTodoService() (*TodoService, error) {
	sv := &TodoService{}

	db, err := dba.NewConnection()
	if err != nil {
		return nil, err
	}

	sv.db = db

	return sv, nil
}

func (sv *TodoService) GetListData() ([]ListData, error) {

	stmt, err := sv.db.Prepare("select list_id,title,is_comp,date from list ")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Listitem []ListData
	for rows.Next() {
		var ConvItem = ListData{}
		err = rows.Scan(&ConvItem.List_id, &ConvItem.Title, &ConvItem.Is_comp, &ConvItem.Date)
		if err != nil {
			return nil, err
		}
		Listitem = append(Listitem, ConvItem)

	}

	return Listitem, nil
}

func (sv *TodoService) Insert(title string, is_comp bool) (string, error) {
	now := time.Now()
	secs := now.Unix()

	txn, err := sv.db.Begin()
	if err != nil {
		return "", err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare("INSERT INTO list(title,is_comp,date) VALUES($1,$2,$3)")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, is_comp, secs)
	if err != nil {
		return "", err
	}

	err = txn.Commit()
	if err != nil {
		return "", err
	}

	return "success fully", nil
}

func (sv *TodoService) Update(title string, is_comp bool, list_id int) (*ListData, error) {

	now := time.Now()
	secs := now.Unix()

	txn, err := sv.db.Begin()
	if err != nil {
		return nil, err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare("UPDATE list SET title=$1, is_comp=$2 ,date=$3 WHERE list_id=$4")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, is_comp, secs, list_id)
	if err != nil {
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	returnData := ListData{
		List_id: list_id,
		Title:   title,
		Is_comp: is_comp,
	}

	fmt.Println("success fully!!")
	return &returnData, nil
}

func (sv *TodoService) Delete(list_id int) (string, error) {
	txn, err := sv.db.Begin()
	if err != nil {
		return "", err
	}
	defer txn.Rollback()

	stmt, err := txn.Prepare("DELETE FROM list WHERE list_id=$1")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(list_id)
	if err != nil {
		return "", err
	}

	err = txn.Commit()
	if err != nil {
		return "", err
	}

	log.Println("DELETE")

	return "Delete Complete", nil
}

func (sv *TodoService) Show(list_id int) (*ListData, error) {
	stmt, err := sv.db.Prepare("select list_id,title,is_comp,date from list where list_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(list_id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var Dataget = ListData{}
	for rows.Next() {
		ConvItem := ListData{}
		err = rows.Scan(&ConvItem.List_id, &ConvItem.Title, &ConvItem.Is_comp, &ConvItem.Date)
		if err != nil {
			return nil, err
		}
		Dataget = ConvItem
	}

	if Dataget.List_id == 0 {
		return nil, nil
	} else {
		return &Dataget, nil
	}

}
