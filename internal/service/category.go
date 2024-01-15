package service

import (
	"GeovaneCavalcante/grpc/internal/database"
	"GeovaneCavalcante/grpc/internal/pb"
	"context"
	"io"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.CreateCategory(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	cr := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return cr, nil
}

func (c *CategoryService) ListCategory(ctx context.Context, in *pb.Blank) (*pb.CategoryList, error) {
	var categories []*pb.Category
	categoryList, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	for _, category := range categoryList {
		cr := &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		}
		categories = append(categories, cr)
	}
	return &pb.CategoryList{Categories: categories}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.CategoryGetRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.FindByID(in.Id)
	if err != nil {
		return nil, err
	}
	cr := &pb.Category{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}

	return cr, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryList{}
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}
		cc, err := c.CategoryDB.CreateCategory(category.Name, category.Description)
		if err != nil {
			return err
		}
		cr := &pb.Category{
			Id:          cc.ID,
			Name:        cc.Name,
			Description: cc.Description,
		}
		categories.Categories = append(categories.Categories, cr)
	}
}

func (c *CategoryService) CreateCategoryStreamBidirectional(stream pb.CategoryService_CreateCategoryStreamBidirectionalServer) error {
	for {
		category, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		cc, err := c.CategoryDB.CreateCategory(category.Name, category.Description)
		if err != nil {
			return err
		}
		cr := &pb.Category{
			Id:          cc.ID,
			Name:        cc.Name,
			Description: cc.Description,
		}
		if err := stream.Send(cr); err != nil {
			return err
		}
	}
}
