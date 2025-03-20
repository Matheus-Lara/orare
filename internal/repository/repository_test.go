package repository

import (
	"testing"

	"github.com/Matheus-Lara/orare/internal/db"
	"github.com/stretchr/testify/assert"
)

type TestEntity struct {
	ID   uint   `gorm:"primaryKey"`
	Name string `gorm:"size:255"`
}

func setupTestDB() {
	db.InitTestDB()
	db.GetConnection().AutoMigrate(&TestEntity{})
	db.GetConnection().Exec("DELETE FROM test_entities")
}

func TestRepository_Create(t *testing.T) {
	setupTestDB()
	repo := NewRepository[TestEntity]()

	entity := &TestEntity{Name: "Test Name"}
	err := repo.Create(entity)

	assert.NoError(t, err)
	assert.NotZero(t, entity.ID)
}

func TestRepository_FindByID(t *testing.T) {
	setupTestDB()
	repo := NewRepository[TestEntity]()

	entity := &TestEntity{Name: "Test Name"}
	repo.Create(entity)

	result, err := repo.FindByID(entity.ID)

	assert.NoError(t, err)
	assert.Equal(t, entity.Name, result.Name)
}

func TestRepository_FindAll(t *testing.T) {
	setupTestDB()
	repo := NewRepository[TestEntity]()

	repo.Create(&TestEntity{Name: "Test 1"})
	repo.Create(&TestEntity{Name: "Test 2"})

	results, err := repo.FindAll()

	assert.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestRepository_Update(t *testing.T) {
	setupTestDB()
	repo := NewRepository[TestEntity]()

	entity := &TestEntity{Name: "Old Name"}
	repo.Create(entity)

	entity.Name = "New Name"
	err := repo.Update(entity)

	assert.NoError(t, err)

	result, _ := repo.FindByID(entity.ID)
	assert.Equal(t, "New Name", result.Name)
}

func TestRepository_Delete(t *testing.T) {
	setupTestDB()
	repo := NewRepository[TestEntity]()

	entity := &TestEntity{Name: "Test Name"}
	repo.Create(entity)

	err := repo.Delete(entity.ID)

	assert.NoError(t, err)

	result, err := repo.FindByID(entity.ID)
	assert.Error(t, err)
	assert.Nil(t, result)
}

func TestRepository_Count(t *testing.T) {
	setupTestDB()
	repo := NewRepository[TestEntity]()

	repo.Create(&TestEntity{Name: "Test 1"})
	repo.Create(&TestEntity{Name: "Test 2"})

	count, err := repo.Count()

	assert.NoError(t, err)
	assert.Equal(t, int64(2), count)
}

func TestRepository_Paginate(t *testing.T) {
	setupTestDB()
	repo := NewRepository[TestEntity]()

	repo.Create(&TestEntity{Name: "Test 1"})
	repo.Create(&TestEntity{Name: "Test 2"})
	repo.Create(&TestEntity{Name: "Test 3"})
	repo.Create(&TestEntity{Name: "Test 4"})

	results, err := repo.Paginate(2, 2)

	assert.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestRepository_FindOneBy(t *testing.T) {
	setupTestDB()
	repo := NewRepository[TestEntity]()

	entity := &TestEntity{Name: "Test Name"}
	secondEntity := &TestEntity{Name: "Second Test Name"}
	repo.Create(entity)
	repo.Create(secondEntity)

	result, err := repo.FindOneBy(map[string]interface{}{"name": "Second Test Name"})

	assert.NoError(t, err)
	assert.Equal(t, uint(2), result.ID)
	assert.Equal(t, secondEntity.Name, result.Name)
}
