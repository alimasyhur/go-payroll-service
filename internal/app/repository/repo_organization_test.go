package repository_test

import (
	"context"
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/weanan/weanan-service/internal/app/entity"
	"github.com/weanan/weanan-service/internal/app/repository"
	mockPkg "github.com/weanan/weanan-service/internal/pkg/mocks"
)

func TestNewOrganizationRepository(t *testing.T) {
	db, _ := mockPkg.MockGorm()

	t.Run("Should Panic When DB Conn Is Nil", func(t *testing.T) {
		assert.Panics(t, func() {
			repository.NewOrganizationRepository(nil)
		})
	})

	t.Run("Should Not Panic", func(t *testing.T) {
		assert.NotPanics(t, func() {
			repository.NewOrganizationRepository(db)
		})
	})
}

func TestGetOrganizationById(t *testing.T) {
	id := 100
	name := "weanan"

	db, mock := mockPkg.MockGorm()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `organizations` WHERE `organizations`.`id` = ? ORDER BY `organizations`.`id` DESC LIMIT 1")).
		WithArgs(id).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name))

	t.Run("Should Success Query", func(t *testing.T) {
		repo := repository.NewOrganizationRepository(db)
		resp, err := repo.GetById(context.Background(), uint(id))

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, entity.Organization{ID: uint(id), Name: name}, resp)
	})

	t.Run("Should Error Query", func(t *testing.T) {
		repo := repository.NewOrganizationRepository(db)
		resp, err := repo.GetById(context.Background(), uint(200))

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})
}

func TestGetAllOrganization(t *testing.T) {
	id := 100
	name := "weanan"

	db, mock := mockPkg.MockGorm()

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `organizations`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(id, name))

	t.Run("Should Success Query", func(t *testing.T) {
		repo := repository.NewOrganizationRepository(db)
		resp, err := repo.GetAll(context.Background())

		assert.Nil(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, []entity.Organization{{ID: uint(id), Name: name}}, resp)
	})

	t.Run("Should Error Query", func(t *testing.T) {
		repo := repository.NewOrganizationRepository(db)
		resp, err := repo.GetAll(context.Background())

		assert.NotNil(t, err)
		assert.Empty(t, resp)
	})
}

func TestCreateOrganization(t *testing.T) {
	name := "weanan"
	location := "Sukoharjo"

	db, mock := mockPkg.MockGorm()

	t.Run("Should Success Create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `organizations`").
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		data := entity.Organization{Name: name, Location: location}
		repo := repository.NewOrganizationRepository(db)
		err := repo.Create(context.Background(), &data)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), data.ID)
		assert.Equal(t, name, data.Name)
		assert.Equal(t, location, data.Location)
	})

	t.Run("Should Error Create", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO `organizations`").
			WillReturnError(errors.New("insert data error"))
		mock.ExpectRollback()

		data := entity.Organization{Name: name, Location: location}
		repo := repository.NewOrganizationRepository(db)
		err := repo.Create(context.Background(), &data)

		assert.NotNil(t, err)
	})
}

func TestUpdateOrganization(t *testing.T) {
	id := 1
	name := "weanan"
	location := "Sukoharjo"

	db, mock := mockPkg.MockGorm()

	t.Run("Should Success Update", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `organizations`").
			WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		data := entity.Organization{ID: uint(id), Name: name, Location: location}
		repo := repository.NewOrganizationRepository(db)
		err := repo.Update(context.Background(), &data)

		assert.Nil(t, err)
		assert.Equal(t, uint(1), data.ID)
		assert.Equal(t, name, data.Name)
		assert.Equal(t, location, data.Location)
	})

	t.Run("Should Error Update", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `organizations`").
			WillReturnError(errors.New("update data error"))
		mock.ExpectRollback()

		data := entity.Organization{ID: uint(id), Name: name, Location: location}
		repo := repository.NewOrganizationRepository(db)
		err := repo.Update(context.Background(), &data)

		assert.NotNil(t, err)
	})
}

func TestDeleteOrganization(t *testing.T) {
	id := 1
	name := "weanan"
	location := "Sukoharjo"

	db, mock := mockPkg.MockGorm()

	t.Run("Should Success Delete", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `organizations` WHERE `organizations`.`id` = ?").
			WithArgs(id).
			WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		data := entity.Organization{ID: uint(id), Name: name, Location: location}
		repo := repository.NewOrganizationRepository(db)
		err := repo.Delete(context.Background(), &data)

		assert.Nil(t, err)
	})

	t.Run("Should Error Delete", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec("DELETE FROM `organizations` WHERE `organizations`.`id` = ?").
			WithArgs(id).
			WillReturnError(errors.New("delete data error"))
		mock.ExpectRollback()

		data := entity.Organization{ID: uint(id), Name: name, Location: location}
		repo := repository.NewOrganizationRepository(db)
		err := repo.Delete(context.Background(), &data)

		assert.NotNil(t, err)
	})
}
