package resource

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/swordbeta/reddit-frontpage-analyzer-go/src/domain"
	"github.com/swordbeta/reddit-frontpage-analyzer-go/src/util"
	"github.com/swordbeta/reddit-frontpage-analyzer-go/src/util/database"
)

func TestMain(m *testing.M) {
	ret := m.Run()
	database.TearDown()
	os.Exit(ret)
}

func Test_GetPosts(t *testing.T) {
	util.InitConfig()
	db := database.InitDatabase()
	defer db.Close()
	database.SavePost(&domain.Post{
		ID:          "firstPost",
		DateCreated: 1451606400, // 2016-01-01
		PostHint:    "image",
		Tags: []domain.Tag{
			domain.Tag{
				Name:       "Quite sure",
				Confidence: 0.95,
			},
			domain.Tag{
				Name:       "Not quite sure",
				Confidence: 0.45,
			},
		},
	}, db)
	database.SavePost(&domain.Post{
		ID:          "secondPost",
		DateCreated: 1451692800, // 2016-01-02
		PostHint:    "image",
		Tags: []domain.Tag{
			domain.Tag{
				Name:       "Quite sure",
				Confidence: 0.95,
			},
			domain.Tag{
				Name:       "Not quite sure",
				Confidence: 0.45,
			},
		},
	}, db)
	r, _ := http.NewRequest("GET", "/api/v1/posts", nil)
	w := httptest.NewRecorder()
	GetPosts(w, r)
	if w.Code != http.StatusBadRequest {
		t.Errorf("Excepted 400 bad request, got %v.", w.Code)
	}
	r, _ = http.NewRequest("GET", "/api/v1/posts?date=2016-01-01", nil)
	w = httptest.NewRecorder()
	GetPosts(w, r)
	if w.Code != http.StatusOK {
		t.Errorf("Expected 200 OK, got %v.", w.Code)
	}
	contentType := w.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected application/json, got %v.", contentType)
	}
	decoder := json.NewDecoder(w.Body)
	posts := []domain.Post{}
	err := decoder.Decode(&posts)
	if err != nil {
		t.Error(err)
	}
	if len(posts) != 1 || posts[0].ID != "firstPost" {
		t.Errorf(
			"Expected 1 post with ID 'firstPost', got %v posts.",
			len(posts),
		)
	} else if len(posts[0].Tags) != 2 {
		t.Errorf(
			"Expected post with ID 'firstPost' to have 2 tags, got %v.",
			len(posts[0].Tags),
		)
	}
}
