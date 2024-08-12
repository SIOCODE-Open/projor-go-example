package main

import "errors"

type UserAccount struct {
    Id string `json:"id"`
    Username string `json:"username"`
    Email string `json:"email"`
    Flags []string `json:"flags"`
}

type IUserAccountRepository interface {
    Create(userAccount *UserAccount) error
    Update(userAccount *UserAccount) error
    Delete(userAccount *UserAccount) error
    GetByID(id string) (*UserAccount, error)
    GetAll() ([]UserAccount, error)
}

type UserAccountRepository struct {
    IUserAccountRepository
    entities map[string]*UserAccount
}

func NewUserAccountRepository() *UserAccountRepository {
    return &UserAccountRepository{
        entities: make(map[string]*UserAccount),
    }
}

func (repo *UserAccountRepository) Create(userAccount *UserAccount) error {
    if _, ok := repo.entities[userAccount.Id]; ok {
        return errors.New("UserAccount already exists")
    }
    repo.entities[userAccount.Id] = userAccount
    return nil
}

func (repo *UserAccountRepository) Update(userAccount *UserAccount) error {
    if _, ok := repo.entities[userAccount.Id]; !ok {
        return errors.New("UserAccount not found")
    }
    repo.entities[userAccount.Id] = userAccount
    return nil
}

func (repo *UserAccountRepository) Delete(userAccount *UserAccount) error {
    if _, ok := repo.entities[userAccount.Id]; !ok {
        return errors.New("UserAccount not found")
    }
    delete(repo.entities, userAccount.Id)
    return nil
}

func (repo *UserAccountRepository) GetByID(id string) (*UserAccount, error) {
    if _, ok := repo.entities[id]; !ok {
        return nil, errors.New("UserAccount not found")
    }
    return repo.entities[id], nil
}

func (repo *UserAccountRepository) GetAll() ([]UserAccount, error) {
    var result []UserAccount
    for _, entity := range repo.entities {
        result = append(result, *entity)
    }
    return result, nil
}