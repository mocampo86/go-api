package product

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	. "github.com/ahmetb/go-linq/v3"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetAll(ctx context.Context) (products []DAOProduct, err error) {
	args := mock.Called(ctx)
	result := args.Get(0)

	if result != nil {
		return result.([]DAOProduct), args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (mock *MockRepository) Find(ctx context.Context, filter ProductFilterRequest) (products []DAOProduct, err error) {
	args := mock.Called(ctx, filter)
	result := args.Get(0)

	if result != nil {
		var filteredProducts []DAOProduct

		//VER SEBA
		From(result.([]DAOProduct)).Where(func(c interface{}) bool {
			return (filter.ProductId <= 0 || c.(DAOProduct).ProductId == int(filter.ProductId)) &&
				(filter.ProductName == "" || c.(DAOProduct).Name == filter.ProductName)
		}).Select(func(c interface{}) interface{} {
			return c.(DAOProduct)
		}).ToSlice(&filteredProducts)

		return filteredProducts, args.Error(1)
	} else {
		return nil, args.Error(1)
	}
}

func (mock *MockRepository) Insert(ctx context.Context, product DAOProduct) error {
	args := mock.Called(ctx, product)
	return args.Error(0)
}

func (mock *MockRepository) Update(ctx context.Context, product DAOProduct) error {
	args := mock.Called(ctx, product)
	return args.Error(0)
}

func (mock *MockRepository) Remove(ctx context.Context, id int64) error {
	args := mock.Called(ctx, id)
	return args.Error(0)
}

//Test: GetAll
func TestGetAll_WhenSuccess(t *testing.T) {
	mockRepository := new(MockRepository)

	products := []DAOProduct{{}}
	ctx := context.Background()
	mockRepository.On("GetAll", ctx).Return(products, nil)

	testService := NewService(mockRepository)

	result, error := testService.GetAll(ctx)

	//Data Assertion
	assert.Nil(t, error)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result))
}

func TestGetAll_WhenRepositoryError(t *testing.T) {
	mockRepository := new(MockRepository)

	ctx := context.Background()
	repoError := errors.New("Repo Err")
	mockRepository.On("GetAll", ctx).Return(nil, repoError)

	testService := NewService(mockRepository)
	result, servError := testService.GetAll(ctx)

	//Data Assertion
	assert.Nil(t, result)
	assert.NotNil(t, servError)
}

func TestGetAll_WhenEmptyResult(t *testing.T) {
	mockRepository := new(MockRepository)

	products := []DAOProduct{}
	ctx := context.Background()
	mockRepository.On("GetAll", ctx).Return(products, nil)

	testService := NewService(mockRepository)

	result, error := testService.GetAll(ctx)

	//Data Assertion
	assert.Nil(t, error)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result))
}

//Test: Find
func TestFind_WhenSuccess(t *testing.T) {
	mockRepository := new(MockRepository)

	products := []DAOProduct{{ProductId: 1}, {ProductId: 2}}
	ctx := context.Background()
	filter := ProductFilterRequest{ProductId: 1}
	mockRepository.On("Find", ctx, filter).Return(products, nil)

	testService := NewService(mockRepository)

	result, error := testService.Find(ctx, filter)

	//Data Assertion
	assert.Nil(t, error)
	assert.NotNil(t, result)
	assert.Equal(t, 1, len(result))
	assert.Equal(t, 1, result[0].ProductId)
}

func TestFind_WhenRepositoryError(t *testing.T) {
	mockRepository := new(MockRepository)

	filter := ProductFilterRequest{ProductId: 1}
	ctx := context.Background()
	repoError := errors.New("Repo Err")
	mockRepository.On("Find", ctx, filter).Return(nil, repoError)

	testService := NewService(mockRepository)
	result, servError := testService.Find(ctx, filter)

	//Data Assertion
	assert.Nil(t, result)
	assert.NotNil(t, servError)
}

func TestFind_WhenEmptyResult(t *testing.T) {
	mockRepository := new(MockRepository)

	products := []DAOProduct{{ProductId: 1}, {ProductId: 2}}
	ctx := context.Background()
	filter := ProductFilterRequest{ProductId: 3}
	mockRepository.On("Find", ctx, filter).Return(products, nil)

	testService := NewService(mockRepository)

	result, error := testService.Find(ctx, filter)

	//Data Assertion
	assert.Nil(t, error)
	assert.NotNil(t, result)
	assert.Equal(t, 0, len(result))
}

//Test: Insert
func TestInsert_WhenSuccess(t *testing.T) {
	mockRepository := new(MockRepository)

	product := DAOProduct{}
	ctx := context.Background()
	mockRepository.On("Insert", ctx, product).Return(nil)

	testService := NewService(mockRepository)

	error := testService.Insert(ctx, Product{})

	//Data Assertion
	assert.Nil(t, error)
}

func TestInsert_WhenRepositoryError(t *testing.T) {
	mockRepository := new(MockRepository)

	product := DAOProduct{}
	ctx := context.Background()
	repoError := errors.New("Repo Err")
	mockRepository.On("Insert", ctx, product).Return(repoError)

	testService := NewService(mockRepository)

	error := testService.Insert(ctx, Product{})

	//Data Assertion
	assert.NotNil(t, error)
}

//Test: Update
func TestUpdate_WhenSuccess(t *testing.T) {
	mockRepository := new(MockRepository)

	product := DAOProduct{}
	ctx := context.Background()
	mockRepository.On("Update", ctx, product).Return(nil)

	testService := NewService(mockRepository)

	error := testService.Update(ctx, Product{})

	//Data Assertion
	assert.Nil(t, error)
}

func TestUpdate_WhenRepositoryError(t *testing.T) {
	mockRepository := new(MockRepository)

	product := DAOProduct{}
	ctx := context.Background()
	repoError := errors.New("Repo Err")
	mockRepository.On("Update", ctx, product).Return(repoError)

	testService := NewService(mockRepository)

	error := testService.Update(ctx, Product{})

	//Data Assertion
	assert.NotNil(t, error)
}

//Test: Remove
func TestRemove_WhenSuccess(t *testing.T) {
	mockRepository := new(MockRepository)

	idToRemove := int64(1)
	ctx := context.Background()
	mockRepository.On("Remove", ctx, idToRemove).Return(nil)

	testService := NewService(mockRepository)

	error := testService.Remove(ctx, int64(idToRemove))

	//Data Assertion
	assert.Nil(t, error)
}

func TestRemove_WhenRepositoryError(t *testing.T) {
	mockRepository := new(MockRepository)

	idToRemove := int64(1)
	ctx := context.Background()
	repoError := errors.New("Repo Err")
	mockRepository.On("Remove", ctx, idToRemove).Return(repoError)

	testService := NewService(mockRepository)

	error := testService.Remove(ctx, int64(idToRemove))

	//Data Assertion
	assert.NotNil(t, error)
}
