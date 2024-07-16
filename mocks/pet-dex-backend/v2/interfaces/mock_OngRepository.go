// Code generated by mockery v2.43.2. DO NOT EDIT.

package interfaces

import (
	entity "pet-dex-backend/v2/entity"
	dto "pet-dex-backend/v2/entity/dto"

	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// MockOngRepository is an autogenerated mock type for the OngRepository type
type MockOngRepository struct {
	mock.Mock
}

type MockOngRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *MockOngRepository) EXPECT() *MockOngRepository_Expecter {
	return &MockOngRepository_Expecter{mock: &_m.Mock}
}

// FindByID provides a mock function with given fields: ID
func (_m *MockOngRepository) FindByID(ID uuid.UUID) (*entity.Ong, error) {
	ret := _m.Called(ID)

	if len(ret) == 0 {
		panic("no return value specified for FindByID")
	}

	var r0 *entity.Ong
	var r1 error
	if rf, ok := ret.Get(0).(func(uuid.UUID) (*entity.Ong, error)); ok {
		return rf(ID)
	}
	if rf, ok := ret.Get(0).(func(uuid.UUID) *entity.Ong); ok {
		r0 = rf(ID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*entity.Ong)
		}
	}

	if rf, ok := ret.Get(1).(func(uuid.UUID) error); ok {
		r1 = rf(ID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOngRepository_FindByID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindByID'
type MockOngRepository_FindByID_Call struct {
	*mock.Call
}

// FindByID is a helper method to define mock.On call
//   - ID uuid.UUID
func (_e *MockOngRepository_Expecter) FindByID(ID interface{}) *MockOngRepository_FindByID_Call {
	return &MockOngRepository_FindByID_Call{Call: _e.mock.On("FindByID", ID)}
}

func (_c *MockOngRepository_FindByID_Call) Run(run func(ID uuid.UUID)) *MockOngRepository_FindByID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID))
	})
	return _c
}

func (_c *MockOngRepository_FindByID_Call) Return(_a0 *entity.Ong, _a1 error) *MockOngRepository_FindByID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockOngRepository_FindByID_Call) RunAndReturn(run func(uuid.UUID) (*entity.Ong, error)) *MockOngRepository_FindByID_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: limit, offset, sortBy, order
func (_m *MockOngRepository) List(limit int, offset int, sortBy string, order string) ([]*dto.OngListMapper, error) {
	ret := _m.Called(limit, offset, sortBy, order)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 []*dto.OngListMapper
	var r1 error
	if rf, ok := ret.Get(0).(func(int, int, string, string) ([]*dto.OngListMapper, error)); ok {
		return rf(limit, offset, sortBy, order)
	}
	if rf, ok := ret.Get(0).(func(int, int, string, string) []*dto.OngListMapper); ok {
		r0 = rf(limit, offset, sortBy, order)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*dto.OngListMapper)
		}
	}

	if rf, ok := ret.Get(1).(func(int, int, string, string) error); ok {
		r1 = rf(limit, offset, sortBy, order)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockOngRepository_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockOngRepository_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - limit int
//   - offset int
//   - sortBy string
//   - order string
func (_e *MockOngRepository_Expecter) List(limit interface{}, offset interface{}, sortBy interface{}, order interface{}) *MockOngRepository_List_Call {
	return &MockOngRepository_List_Call{Call: _e.mock.On("List", limit, offset, sortBy, order)}
}

func (_c *MockOngRepository_List_Call) Run(run func(limit int, offset int, sortBy string, order string)) *MockOngRepository_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int), args[1].(int), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockOngRepository_List_Call) Return(ongs []*dto.OngListMapper, err error) *MockOngRepository_List_Call {
	_c.Call.Return(ongs, err)
	return _c
}

func (_c *MockOngRepository_List_Call) RunAndReturn(run func(int, int, string, string) ([]*dto.OngListMapper, error)) *MockOngRepository_List_Call {
	_c.Call.Return(run)
	return _c
}

// Save provides a mock function with given fields: ong
func (_m *MockOngRepository) Save(ong *entity.Ong) error {
	ret := _m.Called(ong)

	if len(ret) == 0 {
		panic("no return value specified for Save")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*entity.Ong) error); ok {
		r0 = rf(ong)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOngRepository_Save_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Save'
type MockOngRepository_Save_Call struct {
	*mock.Call
}

// Save is a helper method to define mock.On call
//   - ong *entity.Ong
func (_e *MockOngRepository_Expecter) Save(ong interface{}) *MockOngRepository_Save_Call {
	return &MockOngRepository_Save_Call{Call: _e.mock.On("Save", ong)}
}

func (_c *MockOngRepository_Save_Call) Run(run func(ong *entity.Ong)) *MockOngRepository_Save_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*entity.Ong))
	})
	return _c
}

func (_c *MockOngRepository_Save_Call) Return(_a0 error) *MockOngRepository_Save_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOngRepository_Save_Call) RunAndReturn(run func(*entity.Ong) error) *MockOngRepository_Save_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: id, ong
func (_m *MockOngRepository) Update(id uuid.UUID, ong entity.Ong) error {
	ret := _m.Called(id, ong)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uuid.UUID, entity.Ong) error); ok {
		r0 = rf(id, ong)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockOngRepository_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockOngRepository_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - id uuid.UUID
//   - ong entity.Ong
func (_e *MockOngRepository_Expecter) Update(id interface{}, ong interface{}) *MockOngRepository_Update_Call {
	return &MockOngRepository_Update_Call{Call: _e.mock.On("Update", id, ong)}
}

func (_c *MockOngRepository_Update_Call) Run(run func(id uuid.UUID, ong entity.Ong)) *MockOngRepository_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(uuid.UUID), args[1].(entity.Ong))
	})
	return _c
}

func (_c *MockOngRepository_Update_Call) Return(_a0 error) *MockOngRepository_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockOngRepository_Update_Call) RunAndReturn(run func(uuid.UUID, entity.Ong) error) *MockOngRepository_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockOngRepository creates a new instance of MockOngRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockOngRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockOngRepository {
	mock := &MockOngRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
