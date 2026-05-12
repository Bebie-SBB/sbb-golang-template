package template

import (
	"fmt"
	"strconv"

	"sbb-golang-template/pkg/logger"
	"sbb-golang-template/pkg/models"
)

type TemplateServiceInterface interface {
	Init() error
	Create(args *CreateArgs) (*Template, error)
	List(searchQuery *SearchQuery) ([]*Template, error)
	Update(args *UpdateArgs) (*Template, error)
	Info(args *ReadArgs) (*Template, error)
	Delete(args *DeleteArgs) error
	IsExisting(*ReadArgs) (bool, error)
}

type TemplateService models.Service[*templateServiceRepository]

func NewService() (*TemplateService, error) {
	s := &TemplateService{}
	return s, s.Init()
}

func (s *TemplateService) Init() error {
	s.Log = logger.Get()
	return nil
}

func (s *TemplateService) IsExisting(args *ReadArgs) (bool, error) {
	t, err := s.Info(args)
	if err != nil {
		return false, nil
	}
	return t.Code != "", nil
}

func (s *TemplateService) Create(args *CreateArgs) (*Template, error) {
	exist, err := s.IsExisting(&ReadArgs{Template: args.Template})
	if err != nil {
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("template code %s already exist", args.Code)
	}
	repo, err := (&templateServiceRepository{}).CreateRepository()
	if err != nil {
		return nil, err
	}
	return repo.Create(&args.Template)
}

func (s *TemplateService) List(searchQuery *SearchQuery) ([]*Template, error) {
	limit, offset := 10, 0
	if searchQuery.Limit != "" {
		n, err := strconv.Atoi(searchQuery.Limit)
		if err != nil {
			return nil, fmt.Errorf("limit is not an integer")
		}
		limit = n
	}
	if searchQuery.Page != "" {
		n, err := strconv.Atoi(searchQuery.Page)
		if err != nil {
			return nil, fmt.Errorf("page is not an integer")
		}
		offset = (n - 1) * limit
	}
	repo, err := (&templateServiceRepository{}).CreateRepository()
	if err != nil {
		return nil, err
	}
	return repo.List(searchQuery.Keyword, limit, offset)
}

func (s *TemplateService) Update(args *UpdateArgs) (*Template, error) {
	exist, err := s.IsExisting(&ReadArgs{Template: args.Template})
	if err != nil {
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("template code %s not found", args.Code)
	}
	repo, err := (&templateServiceRepository{}).CreateRepository()
	if err != nil {
		return nil, err
	}
	if err := repo.Update(&args.Template); err != nil {
		return nil, err
	}
	return &args.Template, nil
}

func (s *TemplateService) Info(args *ReadArgs) (*Template, error) {
	repo, err := (&templateServiceRepository{}).CreateRepository()
	if err != nil {
		return nil, err
	}
	return repo.Get(&args.Template)
}

func (s *TemplateService) Delete(args *DeleteArgs) error {
	exist, err := s.IsExisting(&ReadArgs{Template: args.Template})
	if err != nil {
		return err
	}
	if !exist {
		return fmt.Errorf("template code %s not found", args.Code)
	}
	repo, err := (&templateServiceRepository{}).CreateRepository()
	if err != nil {
		return err
	}
	return repo.Delete(&args.Template)
}
