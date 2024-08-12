package main

import "errors"

type Product struct {
    Id string `json:"id"`
    Name string `json:"name"`
    Price float32 `json:"price"`
    Description string `json:"description"`
}

type IProductRepository interface {
    Create(product *Product) error
    Update(product *Product) error
    Delete(product *Product) error
    GetByID(id string) (*Product, error)
    GetAll() ([]Product, error)
}

type ProductRepository struct {
    IProductRepository
    entities map[string]*Product
}

func NewProductRepository() *ProductRepository {
    return &ProductRepository{
        entities: make(map[string]*Product),
    }
}

func (repo *ProductRepository) Create(product *Product) error {
    if _, ok := repo.entities[product.Id]; ok {
        return errors.New("Product already exists")
    }
    repo.entities[product.Id] = product
    return nil
}

func (repo *ProductRepository) Update(product *Product) error {
    if _, ok := repo.entities[product.Id]; !ok {
        return errors.New("Product not found")
    }
    repo.entities[product.Id] = product
    return nil
}

func (repo *ProductRepository) Delete(product *Product) error {
    if _, ok := repo.entities[product.Id]; !ok {
        return errors.New("Product not found")
    }
    delete(repo.entities, product.Id)
    return nil
}

func (repo *ProductRepository) GetByID(id string) (*Product, error) {
    if _, ok := repo.entities[id]; !ok {
        return nil, errors.New("Product not found")
    }
    return repo.entities[id], nil
}

func (repo *ProductRepository) GetAll() ([]Product, error) {
    var result []Product
    for _, entity := range repo.entities {
        result = append(result, *entity)
    }
    return result, nil
}