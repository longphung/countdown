package tests

import (
	"github.com/golang/mock/gomock"
	"github.com/longphung/countdown/countdown"
	"github.com/longphung/countdown/countdown/mocks"
	"github.com/longphung/countdown/countdown/models"
	"reflect"
	"testing"
	"time"
)

func TestServices_GetAllCountdowns(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockRepository := mocks.NewMockRepository(mockCtl)
	service := countdown.NewService(mockRepository)
	mockRepository.EXPECT().GetAllCountdowns().Return(nil, nil).Times(1)

	_, err := service.GetAllCountdowns()
	if err != nil {
		t.Errorf("Expected nil, got %s", err)
	}
}

func TestServices_GetCountdown(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockRepository := mocks.NewMockRepository(mockCtl)
	service := countdown.NewService(mockRepository)

	t.Run("Get valid countdown", func(t *testing.T) {
		mockRepository.EXPECT().GetCountdown("1").Return(&models.Countdown{}, nil).Times(1)
		countdown, err := service.GetCountdown("1")
		if reflect.TypeOf(countdown) != reflect.TypeOf(&models.Countdown{}) {
			t.Errorf("Expected %s, got %s", reflect.TypeOf(&models.Countdown{}), reflect.TypeOf(countdown))
		}
		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}
	})

	t.Run("ErrInvalidId on get invalid countdown id", func(t *testing.T) {
		mockRepository.EXPECT().GetCountdown("-1").Return(nil, countdown.ErrInvalidId).Times(1)
		_, err := service.GetCountdown("-1")
		if err != countdown.ErrInvalidId {
			t.Errorf("Expected %s, got %s", countdown.ErrInvalidId, err)
		}
	})

	t.Run("ErrNotFound on get not found countdown id", func(t *testing.T) {
		mockRepository.EXPECT().GetCountdown("10").Return(nil, countdown.ErrNotFound).Times(1)
		_, err := service.GetCountdown("10")
		if err != countdown.ErrNotFound {
			t.Errorf("Expected %s, got %s", countdown.ErrNotFound, err)
		}
	})
}

func TestServices_CreateCountdown(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockRepository := mocks.NewMockRepository(mockCtl)
	service := countdown.NewService(mockRepository)

	t.Run("Create valid countdown", func(t *testing.T) {
		mockValidCountdown := models.Countdown{
			Name:    "title",
			DueDate: time.Now(),
		}
		mockRepository.EXPECT().CreateCountdown(mockValidCountdown).Return(int64(1), nil).Times(1)
		id, err := service.CreateCountdown(mockValidCountdown)
		if id != 1 {
			t.Errorf("Expected %d, got %d", 1, id)
		}
		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}
	})

	t.Run("Missing due date", func(t *testing.T) {
		mockInvalidCountdown := models.Countdown{
			Name: "title",
		}
		mockRepository.EXPECT().CreateCountdown(mockInvalidCountdown).Return(int64(0), countdown.ErrNoDueDate).Times(1)
		_, err := service.CreateCountdown(mockInvalidCountdown)
		if err != countdown.ErrNoDueDate {
			t.Errorf("Expected %s, got %s", countdown.ErrNoDueDate, err)
		}
	})
}

func TestServices_UpdateCountdown(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockRepository := mocks.NewMockRepository(mockCtl)
	service := countdown.NewService(mockRepository)

	t.Run("Update valid countdown", func(t *testing.T) {
		mockValidCountdown := models.Countdown{
			Name:    "title",
			DueDate: time.Now(),
		}
		mockRepository.EXPECT().UpdateCountdown("1", mockValidCountdown).Return(&mockValidCountdown, int64(1), nil).Times(1)
		countdown, id, err := service.UpdateCountdown("1", mockValidCountdown)
		if reflect.TypeOf(countdown) != reflect.TypeOf(&models.Countdown{}) {
			t.Errorf("Expected %s, got %s", reflect.TypeOf(&models.Countdown{}), reflect.TypeOf(countdown))
		}
		if id != 1 {
			t.Errorf("Expected %d, got %d", 1, id)
		}
		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}
	})

	t.Run("ErrNotFound on update not found countdown id", func(t *testing.T) {
		mockValidCountdown := models.Countdown{
			Name:    "title",
			DueDate: time.Now(),
		}
		mockRepository.EXPECT().UpdateCountdown("10", mockValidCountdown).Return(nil, int64(0), countdown.ErrNotFound).Times(1)
		_, id, err := service.UpdateCountdown("10", mockValidCountdown)
		if id != 0 {
			t.Errorf("Expected %d, got %d", 0, id)
		}
		if err != countdown.ErrNotFound {
			t.Errorf("Expected %s, got %s", countdown.ErrNotFound, err)
		}
	})
}

func TestServices_DeleteCountdown(t *testing.T) {
	mockCtl := gomock.NewController(t)
	defer mockCtl.Finish()

	mockRepository := mocks.NewMockRepository(mockCtl)
	service := countdown.NewService(mockRepository)

	t.Run("Delete valid countdown", func(t *testing.T) {
		mockRepository.EXPECT().DeleteCountdown("1").Return(nil).Times(1)
		err := service.DeleteCountdown("1")
		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}
	})

	t.Run("ErrNotFound on delete not found countdown id", func(t *testing.T) {
		mockRepository.EXPECT().DeleteCountdown("10").Return(countdown.ErrNotFound).Times(1)
		err := service.DeleteCountdown("10")
		if err != countdown.ErrNotFound {
			t.Errorf("Expected %s, got %s", countdown.ErrNotFound, err)
		}
	})
}
