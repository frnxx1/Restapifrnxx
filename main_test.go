package main

import (
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	
)



func TestJson(t *testing.T) {
	router := GetRouter()

	request, _ := http.NewRequest("GET", "/albums", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	if w.Code != 200 {
		t.Fatal("epic fail")
	}

	log.Println(w.Body.String())

}

func TestId(t *testing.T) {
	router := GetRouter()

	for _, i := range albums {
		request, _ := http.NewRequest("GET", "/albums/"+i.ID, nil)
		w := httptest.NewRecorder()
		router.ServeHTTP(w, request)

		if w.Code != 200 {
			t.Fatal("epic fail")
		}

		log.Println(w.Body.String())
	}

}

func TestDelete(t *testing.T) {
	router := GetRouter()
	delId := "2"
	request, _ := http.NewRequest("DELETE", "/albums/"+delId, nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	if w.Code != 204 {
		t.Fatal("epic fail")
	}

	log.Println(w.Body.String())

}


func TestUpdate(t *testing.T) {
	router := GetRouter()

	for _, i := range albums {
		request, _ := http.NewRequest("PUT", "/albums/"+i.ID, strings.NewReader(`{"title": "loop"}`))
		w := httptest.NewRecorder()
		router.ServeHTTP(w, request)

		if w.Code != 200 {
			t.Fatal("epic fail")
		}

		log.Println(w.Body.String())
	}

}

func TestAdd(t *testing.T) {
	router := GetRouter()

	request, _ := http.NewRequest("POST", "/albums", strings.NewReader(`{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}`))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, request)

	if w.Code != 201 {
		t.Fatal("epic fail")
	}

	log.Println(w.Body.String())

}
