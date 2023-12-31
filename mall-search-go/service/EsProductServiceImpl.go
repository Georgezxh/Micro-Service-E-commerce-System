package service

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall-search-go/config"
	"mall-search-go/model"
	"mall-search-go/repository"
	"mall-search-go/store"
)

type EsProductServiceImpl struct {
	prouductDao store.EsproductDao
	elasticRepo repository.EsProductRepository
}

func NewEsProductServiceImpl() EsProductService {
	//dsn := "root:root@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return &EsProductServiceImpl{prouductDao: store.NewEsProductDao(db), elasticRepo: repository.Repo}
}

func (s *EsProductServiceImpl) ImportAll() (int, error) {
	esProductList, err := s.prouductDao.GetAllProductList(nil)
	if err != nil {
		return 0, err
	}
	num, err := s.elasticRepo.SaveAll(esProductList)
	if err != nil {
		return 0, err
	}
	return num, nil
}

func (s *EsProductServiceImpl) Create(id int64) (*model.EsProduct, error) {
	product, err := s.prouductDao.GetAllProductList(&id)
	if err != nil {
		return nil, err
	}

	return s.elasticRepo.Save(&product[0])
}

func (s *EsProductServiceImpl) Delete(id int64) error {
	return s.elasticRepo.Delete(id)
}

func (s *EsProductServiceImpl) DeleteBatch(ids []int64) error {
	if len(ids) == 0 {
		return nil
	}
	return s.elasticRepo.DeletaBatch(ids)
}

func (s *EsProductServiceImpl) SearchByNameOrSubTitleOrKeywords(keyword string, pageNum, pageSize int) (model.Page, error) {
	return s.elasticRepo.Search(keyword, pageNum, pageSize)
}

func (s *EsProductServiceImpl) SearchByProductCategoryId(keyword string, brandId *int64, productCategoryId *int64, pageNum int, pageSize int, sort int) (model.Page, error) {
	return s.elasticRepo.SearchById(keyword, brandId, productCategoryId, pageNum, pageSize, sort)
}

func (s *EsProductServiceImpl) Recommend(id int64, pageNum int, pageSize int) (model.Page, error) {
	var result model.Page

	product, err := s.prouductDao.GetAllProductList(&id)
	if err != nil {
		return result, err
	}

	return s.elasticRepo.Recommend(id, product[0], pageNum, pageSize)
}

func (s *EsProductServiceImpl) SearchRelated(keyword string) (model.EsProductRelatedInfo, error) {
	return s.elasticRepo.SearchRelated(keyword)
}
