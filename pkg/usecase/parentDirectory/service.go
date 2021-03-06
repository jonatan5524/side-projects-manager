package usecase

import (
	"github.com/jonatan5524/side-projects-manager/pkg/model"
	repository "github.com/jonatan5524/side-projects-manager/pkg/repository/parentDirectory"
)

type ParentDirectoryService struct {
	repository repository.ParentDirectoryRepository
}

func NewParentDirectoryService(repo repository.ParentDirectoryRepository) ParentDirectoryUsecase {
	return &ParentDirectoryService{repo}
}

func (service ParentDirectoryService) Put(parentDir model.ParentDirectory) (model.ParentDirectory, error) {
	id, err := service.repository.Put(parentDir)

	if err != nil {
		return model.NilParentDirectory, err
	}

	parentDir.Id = id

	return parentDir, nil
}

func (service ParentDirectoryService) GetAll() ([]*model.ParentDirectory, error) {
	return service.repository.GetAll()
}

func (service ParentDirectoryService) DeleteByPath(path string) error {
	return service.repository.DeleteByPath(path)
}

func (service ParentDirectoryService) Delete(dir model.ParentDirectory) error {
	return service.repository.Delete(dir)
}
