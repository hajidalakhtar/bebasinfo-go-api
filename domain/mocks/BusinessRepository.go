package mocks

import (
	"62teknologi-senior-backend-test-muhammad-hajid-al-akhtar/domain"
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type BusinessRepository struct {
	Mock mock.Mock
}

func NewBusinessRepository(mock mock.Mock) domain.BusinessRepository {
	return &BusinessRepository{Mock: mock}
}

func (_m BusinessRepository) Find(ctx context.Context, term string, sortBy string, limit int, offset int, openAt string) ([]domain.Business, error) {
	ret := _m.Mock.Called(ctx, term, sortBy, limit, offset, openAt)

	var r0 []domain.Business
	var r1 error

	if rf, ok := ret.Get(0).(func(context.Context, string, string, int, int, string) ([]domain.Business, error)); ok {
		r0, r1 = rf(ctx, term, sortBy, limit, offset, openAt)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.Business)
		}
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(error)
		}
	}

	return r0, r1
}

func (_m BusinessRepository) Update(ctx context.Context, bs *domain.Business, id uuid.UUID) error {
	ret := _m.Mock.Called(ctx, bs, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Business, uuid.UUID) error); ok {
		r0 = rf(ctx, bs, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

func (_m BusinessRepository) Delete(ctx context.Context, id uuid.UUID) error {
	ret := _m.Mock.Called(ctx, id)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
//func (_m BusinessRepository) GetUsers(ctx context.Context, cursor string, num int64) (res []domain.User, nextCursor string, err error) {
//	ret := _m.Mock.Called(ctx, cursor, num)
//
//	var r0 []domain.User
//	if rf, ok := ret.Get(0).(func(context.Context, string, int64) []domain.User); ok {
//		r0 = rf(ctx, cursor, num)
//	} else {
//		if ret.Get(0) != nil {
//			r0 = ret.Get(0).([]domain.User)
//		}
//	}
//
//	var r1 string
//	if rf, ok := ret.Get(1).(func(context.Context, string, int64) string); ok {
//		r1 = rf(ctx, cursor, num)
//	} else {
//		r1 = ret.Get(1).(string)
//	}
//
//	var r2 error
//	if rf, ok := ret.Get(2).(func(context.Context, string, int64) error); ok {
//		r2 = rf(ctx, cursor, num)
//	} else {
//		r2 = ret.Error(2)
//	}
//
//	return r0, r1, r2
//
//}
//
//func (_m BusinessRepository) UpdateUser(ctx context.Context, u *domain.User, id uuid.UUID) error {
//	ret := _m.Mock.Called(ctx, u, id)
//
//	var r0 error
//	if rf, ok := ret.Get(0).(func(context.Context, *domain.User, uuid.UUID) error); ok {
//		r0 = rf(ctx, u, id)
//	} else {
//		r0 = ret.Error(0)
//	}
//
//	return r0
//}

func (_m BusinessRepository) Store(ctx context.Context, u *domain.Business) error {
	ret := _m.Mock.Called(ctx, u)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *domain.Business) error); ok {
		r0 = rf(ctx, u)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

//
//func (_m BusinessRepository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
//	ret := _m.Mock.Called(ctx, id)
//
//	var r0 error
//	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) error); ok {
//		r0 = rf(ctx, id)
//	} else {
//		r0 = ret.Error(0)
//	}
//
//	return r0
//
//}
