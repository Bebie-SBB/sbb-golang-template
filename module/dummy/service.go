package dummy

import (
	"fmt"
	"strconv"

	"sbb-golang-template/pkg/logger"
	"sbb-golang-template/pkg/models"
)

type DummyServiceInterface interface {
	Init() error
	Create(args *CreateArgs) (*Dummy, error)
	List(searchQuery *SearchQuery) ([]*Dummy, error)
	ListMap() (*map[string]Dummy, error)
	Update(args *UpdateArgs) (*Dummy, error)
	Info(args *ReadArgs) (*Dummy, error)
	Delete(args *DeleteArgs) error
	IsExisting(*ReadArgs) (bool, error)
}

type DummyService models.Service[*dummyServiceRepository]

func NewService() (*DummyService, error) {
	service := &DummyService{}
	return service, service.Init()
}

func (s *DummyService) Init() error {
	s.Log = logger.Get()
	return nil
}

func (s *DummyService) IsExisting(args *ReadArgs) (bool, error) {
	existApp, err := s.Info(args)
	if err != nil {
		return false, nil
	}
	return existApp.Code != "", nil
}

func (s *DummyService) Create(args *CreateArgs) (*Dummy, error) {
	exist, err := s.IsExisting(&ReadArgs{Dummy: args.Dummy})
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("dummy code %s already exist", args.Code)
	}

	repo, err := (&dummyServiceRepository{}).CreateRepository()
	if err != nil {
		return nil, err
	}
	return repo.Create(&args.Dummy)
}

func (s *DummyService) List(searchQuery *SearchQuery) ([]*Dummy, error) {
	limit := 10
	offset := 0

	if searchQuery.Limit != "" {
		num, err := strconv.Atoi(searchQuery.Limit)
		if err != nil {
			return nil, fmt.Errorf("limit is not an integer")
		}
		limit = num
	}

	if searchQuery.Page != "" {
		num, err := strconv.Atoi(searchQuery.Page)
		if err != nil {
			return nil, fmt.Errorf("page is not an integer")
		}
		offset = (num - 1) * limit
	}

	repo, err := (&dummyServiceRepository{}).CreateRepository()
	if err != nil {
		return nil, err
	}
	return repo.List(searchQuery.Keyword, limit, offset)
}

func (s *DummyService) ListMap() (*map[string]Dummy, error) {
	repo, err := (&dummyServiceRepository{}).CreateRepository()
	if err != nil {
		return nil, err
	}
	list, err := repo.ListAll()
	if err != nil {
		return nil, err
	}
	newMap := make(map[string]Dummy)
	for _, item := range *list {
		newMap[item.Code] = item
	}
	return &newMap, nil
}

func (s *DummyService) Update(args *UpdateArgs) (*Dummy, error) {
	exist, err := s.IsExisting(&ReadArgs{Dummy: args.Dummy, AllowNotFoundFlg: false})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("dummy code %s not found", args.Code)
	}

	repo, err := (&dummyServiceRepository{}).CreateRepository()
	if err != nil {
		return nil, err
	}
	if err := repo.Update(&args.Dummy); err != nil {
		return nil, err
	}
	return &args.Dummy, nil
}

func (s *DummyService) Info(args *ReadArgs) (*Dummy, error) {
	repo, err := (&dummyServiceRepository{}).CreateRepository()
	if err != nil {
		return nil, err
	}
	return repo.Get(&args.Dummy)
}

func (s *DummyService) Delete(args *DeleteArgs) error {
	exist, err := s.IsExisting(&ReadArgs{Dummy: args.Dummy})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("dummy code %s not found", args.Code)
	}

	repo, err := (&dummyServiceRepository{}).CreateRepository()
	if err != nil {
		return err
	}
	return repo.Delete(&args.Dummy)
}
