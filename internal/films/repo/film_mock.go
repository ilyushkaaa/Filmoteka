// Code generated by MockGen. DO NOT EDIT.
// Source: film.go

// Package repo is a generated GoMock package.
package repo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/ilyushkaaa/Filmoteka/internal/films/entity"
)

// MockFilmRepo is a mock of FilmRepo interface.
type MockFilmRepo struct {
	ctrl     *gomock.Controller
	recorder *MockFilmRepoMockRecorder
}

// MockFilmRepoMockRecorder is the mock recorder for MockFilmRepo.
type MockFilmRepoMockRecorder struct {
	mock *MockFilmRepo
}

// NewMockFilmRepo creates a new mock instance.
func NewMockFilmRepo(ctrl *gomock.Controller) *MockFilmRepo {
	mock := &MockFilmRepo{ctrl: ctrl}
	mock.recorder = &MockFilmRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFilmRepo) EXPECT() *MockFilmRepoMockRecorder {
	return m.recorder
}

// AddFilm mocks base method.
func (m *MockFilmRepo) AddFilm(film entity.Film, actorIDs []uint64) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFilm", film, actorIDs)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddFilm indicates an expected call of AddFilm.
func (mr *MockFilmRepoMockRecorder) AddFilm(film, actorIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFilm", reflect.TypeOf((*MockFilmRepo)(nil).AddFilm), film, actorIDs)
}

// DeleteFilm mocks base method.
func (m *MockFilmRepo) DeleteFilm(ID uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteFilm", ID)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DeleteFilm indicates an expected call of DeleteFilm.
func (mr *MockFilmRepoMockRecorder) DeleteFilm(ID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteFilm", reflect.TypeOf((*MockFilmRepo)(nil).DeleteFilm), ID)
}

// GetFilmByID mocks base method.
func (m *MockFilmRepo) GetFilmByID(filmID uint64) (*entity.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmByID", filmID)
	ret0, _ := ret[0].(*entity.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmByID indicates an expected call of GetFilmByID.
func (mr *MockFilmRepoMockRecorder) GetFilmByID(filmID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmByID", reflect.TypeOf((*MockFilmRepo)(nil).GetFilmByID), filmID)
}

// GetFilms mocks base method.
func (m *MockFilmRepo) GetFilms(sortParam string) ([]entity.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilms", sortParam)
	ret0, _ := ret[0].([]entity.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilms indicates an expected call of GetFilms.
func (mr *MockFilmRepoMockRecorder) GetFilms(sortParam interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilms", reflect.TypeOf((*MockFilmRepo)(nil).GetFilms), sortParam)
}

// GetFilmsBySearch mocks base method.
func (m *MockFilmRepo) GetFilmsBySearch(searchStr string) ([]entity.Film, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFilmsBySearch", searchStr)
	ret0, _ := ret[0].([]entity.Film)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFilmsBySearch indicates an expected call of GetFilmsBySearch.
func (mr *MockFilmRepoMockRecorder) GetFilmsBySearch(searchStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFilmsBySearch", reflect.TypeOf((*MockFilmRepo)(nil).GetFilmsBySearch), searchStr)
}

// UpdateFilm mocks base method.
func (m *MockFilmRepo) UpdateFilm(film entity.Film, actorIDs []uint64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateFilm", film, actorIDs)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateFilm indicates an expected call of UpdateFilm.
func (mr *MockFilmRepoMockRecorder) UpdateFilm(film, actorIDs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateFilm", reflect.TypeOf((*MockFilmRepo)(nil).UpdateFilm), film, actorIDs)
}
