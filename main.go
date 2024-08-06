package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Subjects []Grade `json:"subjects"`
}

type Grade struct {
	Subject string  `json:"subject"`
	Score   float64 `json:"score"`
}

type Class struct {
	Name     string    `json:"name"`
	Students []Student `json:"students"`
}

var class Class

func init() {
	class = Class{
		Name: "10-А",
		Students: []Student{
			{ID: 1, Name: "Іван Іванов", Subjects: []Grade{{Subject: "Математика", Score: 85}, {Subject: "Історія", Score: 90}}},
			{ID: 2, Name: "Марія Петренко", Subjects: []Grade{{Subject: "Математика", Score: 95}, {Subject: "Історія", Score: 85}}},
			{ID: 3, Name: "Петро Сидоров", Subjects: []Grade{{Subject: "Математика", Score: 75}, {Subject: "Історія", Score: 80}}},
			{ID: 4, Name: "Оксана Коваленко", Subjects: []Grade{{Subject: "Математика", Score: 92}, {Subject: "Історія", Score: 88}}},
			{ID: 5, Name: "Андрій Мельник", Subjects: []Grade{{Subject: "Математика", Score: 66}, {Subject: "Історія", Score: 70}}},
			{ID: 6, Name: "Ірина Дмитренко", Subjects: []Grade{{Subject: "Математика", Score: 89}, {Subject: "Історія", Score: 92}}},
			{ID: 7, Name: "Володимир Романов", Subjects: []Grade{{Subject: "Математика", Score: 73}, {Subject: "Історія", Score: 75}}},
			{ID: 8, Name: "Олена Кравченко", Subjects: []Grade{{Subject: "Математика", Score: 91}, {Subject: "Історія", Score: 87}}},
			{ID: 9, Name: "Максим Левченко", Subjects: []Grade{{Subject: "Математика", Score: 68}, {Subject: "Історія", Score: 74}}},
			{ID: 10, Name: "Катерина Вовк", Subjects: []Grade{{Subject: "Математика", Score: 85}, {Subject: "Історія", Score: 90}}},
			{ID: 11, Name: "Дмитро Поліщук", Subjects: []Grade{{Subject: "Математика", Score: 77}, {Subject: "Історія", Score: 82}}},
			{ID: 12, Name: "Аліна Кучеренко", Subjects: []Grade{{Subject: "Математика", Score: 93}, {Subject: "Історія", Score: 89}}},
			{ID: 13, Name: "Роман Савченко", Subjects: []Grade{{Subject: "Математика", Score: 64}, {Subject: "Історія", Score: 71}}},
			{ID: 14, Name: "Юлія Павленко", Subjects: []Grade{{Subject: "Математика", Score: 88}, {Subject: "Історія", Score: 84}}},
			{ID: 15, Name: "Олег Жуков", Subjects: []Grade{{Subject: "Математика", Score: 79}, {Subject: "Історія", Score: 81}}},
		},
	}
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/class", getClassInfo)
	mux.HandleFunc("/student/", getStudentInfo)

	log.Println("Сервер працює на порту 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func getClassInfo(w http.ResponseWriter, r *http.Request) {

	if !isTeacher(r) {
		http.Error(w, "Заборонено (дозволено для вчителя)", http.StatusForbidden)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(class)
}

func getStudentInfo(w http.ResponseWriter, r *http.Request) {
	if !isTeacher(r) {
		http.Error(w, "Заборонено (дозволено для вчителя)", http.StatusForbidden)
		return
	}

	idStr := r.URL.Path[len("/student/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Неправильний ID учня", http.StatusBadRequest)
		return
	}

	for _, student := range class.Students {
		if student.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(student)
			return
		}
	}

	http.Error(w, "Учня не знайдено", http.StatusNotFound)
}

func isTeacher(r *http.Request) bool {
	username, _, ok := r.BasicAuth()
	if !ok {
		return false
	}

	return username == "teacher"
}
