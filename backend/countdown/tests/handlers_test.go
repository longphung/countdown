package tests

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/longphung/countdown/countdown"
	"github.com/longphung/countdown/countdown/mocks"
	"github.com/longphung/countdown/countdown/models"
	"github.com/longphung/countdown/utils"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var testCountdowns = []models.Countdown{
	{
		CommonModelFields: utils.CommonModelFields{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt(sql.NullTime{
				Valid: false,
			}),
		},
		Name:    "Test Countdown 1",
		DueDate: time.Now(),
	},
	{
		CommonModelFields: utils.CommonModelFields{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt(sql.NullTime{
				Valid: false,
			}),
		},
		Name:    "Test Countdown 2",
		DueDate: time.Now().Add(time.Hour * 24),
	},
	{
		CommonModelFields: utils.CommonModelFields{
			ID:        1,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			DeletedAt: gorm.DeletedAt(sql.NullTime{
				Valid: false,
			}),
		},
		Name:    "Test Countdown 3",
		DueDate: time.Now().Add(time.Hour * 48),
	},
}

func setup(t *testing.T) (*gin.Engine, *mocks.MockServices, *countdown.Handler, func(t *testing.T)) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	mockCtl := gomock.NewController(t)

	mockServices := mocks.NewMockServices(mockCtl)
	handler := countdown.NewHandler(mockServices)
	return r, mockServices, handler, func(t *testing.T) {
		mockCtl.Finish()
	}
}

func TestHandler_GetAllCountdowns(t *testing.T) {
	r, mockServices, handler, teardown := setup(t)
	defer teardown(t)
	r.GET("/countdowns", handler.GetAllCountdowns)

	t.Run("Success", func(t *testing.T) {
		mockServices.EXPECT().GetAllCountdowns().Return(testCountdowns, nil).Times(1)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/countdowns", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
		testCoutndownBytes, err := json.Marshal(testCountdowns)
		if err != nil {
			t.Errorf("Failed to marshal testCountdowns: %s", err)
		}
		res := bytes.Compare(w.Body.Bytes(), testCoutndownBytes)
		if res != 0 {
			t.Errorf("Expected %s, got %s", testCoutndownBytes, w.Body.Bytes())
		}
	})

	t.Run("Error", func(t *testing.T) {
		dummyError := errors.New("test error")
		mockServices.EXPECT().GetAllCountdowns().Return(nil, dummyError).Times(1)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/countdowns", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusInternalServerError {
			t.Errorf("Expected status code %d, got %d", http.StatusInternalServerError, w.Code)
		}

		resultCountdownBytes, err := json.Marshal(gin.H{
			"error": dummyError.Error(),
		})
		if err != nil {
			t.Errorf("Failed to marshal testCountdowns: %s", err)
		}
		res := bytes.Compare(w.Body.Bytes(), resultCountdownBytes)
		if res != 0 {
			t.Errorf("Expected %s, got %s", resultCountdownBytes, w.Body.Bytes())
		}
	})
}

func TestHandler_GetCountdown(t *testing.T) {
	r, mockServices, handler, teardown := setup(t)
	defer teardown(t)
	r.GET("/countdown/:id", handler.GetCountdown)

	t.Run("Success", func(t *testing.T) {
		mockServices.EXPECT().GetCountdown("1").Return(&testCountdowns[0], nil).Times(1)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/countdown/1", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusOK {
			t.Errorf("Expected status code %d, got %d", http.StatusOK, w.Code)
		}
		countdownBytes, err := json.Marshal(testCountdowns[0])
		if err != nil {
			t.Errorf("Failed to marshal testCountdowns: %s", err)
		}
		res := bytes.Compare(w.Body.Bytes(), countdownBytes)
		if res != 0 {
			t.Errorf("Expected %s, got %s", countdownBytes, w.Body.Bytes())
		}
	})

	t.Run("Invalid ID", func(t *testing.T) {
		// Random string id
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/countdown/invalid", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
		// Negative number id
		mockServices.EXPECT().GetCountdown("-10").Return(nil, countdown.ErrInvalidId).Times(1)

		w = httptest.NewRecorder()
		req, _ = http.NewRequest("GET", "/countdown/-10", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})

	t.Run("Not found", func(t *testing.T) {
		mockServices.EXPECT().GetCountdown("10").Return(nil, countdown.ErrNotFound).Times(1)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/countdown/10", nil)
		r.ServeHTTP(w, req)

		if w.Code != http.StatusNotFound {
			t.Errorf("Expected status code %d, got %d", http.StatusNotFound, w.Code)
		}
	})
}

func TestHandlers_CreateCountdown(t *testing.T) {
	r, mockServices, handler, teardown := setup(t)
	defer teardown(t)
	r.POST("/countdown", handler.CreateCountdown)

	t.Run("Success", func(t *testing.T) {
		dueDateTime, err := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
		if err != nil {
			t.Errorf("Failed to parse dueDateTime: %s", err)
		}
		testData := models.Countdown{
			Name:    "test",
			DueDate: dueDateTime,
		}
		mockServices.EXPECT().CreateCountdown(testData).Return(int64(1), nil).Times(1)

		w := httptest.NewRecorder()
		countdownBytes, err := json.Marshal(gin.H{
			"name":    "test",
			"dueDate": "2006-01-02T15:04:05Z",
		})
		if err != nil {
			t.Errorf("Failed to marshal testCountdowns: %s", err)
		}
		req, _ := http.NewRequest("POST", "/countdown", bytes.NewBuffer(countdownBytes))
		r.ServeHTTP(w, req)

		if w.Code != http.StatusCreated {
			t.Errorf("Expected status code %d, got %d", http.StatusCreated, w.Code)
		}

		resultBytes, err := json.Marshal(gin.H{
			"id": 1,
		})
		if err != nil {
			t.Errorf("Failed to marshal testCountdowns: %s", err)
		}
		res := bytes.Compare(w.Body.Bytes(), resultBytes)
		if res != 0 {
			t.Errorf("Expected %s, got %s", resultBytes, w.Body.Bytes())
		}
	})

	t.Run("Invalid date", func(t *testing.T) {
		testData, err := json.Marshal(gin.H{
			"name":    "test",
			"dueDate": "invalid",
		})
		if err != nil {
			t.Errorf("Failed to marshal testCountdowns: %s", err)
		}

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("POST", "/countdown", bytes.NewBuffer(testData))
		r.ServeHTTP(w, req)

		if w.Code != http.StatusBadRequest {
			t.Errorf("Expected status code %d, got %d", http.StatusBadRequest, w.Code)
		}
	})
}
