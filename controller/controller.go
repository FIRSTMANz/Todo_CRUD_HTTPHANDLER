package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"projectfirsty/service"
	"strconv"
)

func GetListData(w http.ResponseWriter, r *http.Request) {

	sv, err := service.NewTodoService()
	if err != nil {
		panic(err)
	}

	datas, err := sv.GetListData()
	if err != nil {
		log.Fatalf("%s", err)
	}

	jsonData, err := json.Marshal(datas)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}

func Insert(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	is_comp := r.URL.Query().Get("is_comp")
	b1, _ := strconv.ParseBool(is_comp)
	sv, err := service.NewTodoService()
	if err != nil {
		panic(err)
	}

	data, err := sv.Insert(title, b1)
	if err != nil {
		log.Fatalf("%s", err)
	}

	fmt.Fprintf(w, "%v", data)

}

func Update(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	is_comp := r.URL.Query().Get("is_comp")
	list_id := r.URL.Query().Get("list_id")
	b1, _ := strconv.ParseBool(is_comp)
	i, _ := strconv.Atoi(list_id)
	sv, err := service.NewTodoService()
	if err != nil {
		panic(err)
	}
	data, err := sv.Update(title, b1, i)
	if err != nil {
		log.Fatalf("%s", err)
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)

}
func Delete(w http.ResponseWriter, r *http.Request) {
	list_id := r.URL.Query().Get("list_id")
	i, _ := strconv.Atoi(list_id)
	sv, err := service.NewTodoService()
	if err != nil {
		panic(err)
	}
	data, err := sv.Delete(i)
	if err != nil {
		log.Fatalf("%s", err)
	}
	fmt.Fprintf(w, "%v", data)

}
func Show(w http.ResponseWriter, r *http.Request) {
	list_id := r.URL.Query().Get("list_id")
	i, _ := strconv.Atoi(list_id)
	sv, err := service.NewTodoService()
	if err != nil {
		panic(err)
	}
	data, err := sv.Show(i)
	if err != nil {
		log.Fatalf("%s", err)
	}

	if data != nil {
		jsonData, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	} else {
		data := map[string]string{"Status": "ไม่พบข้อมูล"}
		jsonData, err := json.Marshal(data)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonData)
	}

}
