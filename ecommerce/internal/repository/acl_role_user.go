package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IAclRoleUserRepository interface {
	IBaseRepository[model.ACLRoleUser]

	GetByParam(db qrm.Queryable, param *model.ACLRoleUser) (*model.ACLRoleUser, error)
}

type AclRoleUserRepository struct {
	BaseRepository[model.ACLRoleUser]
}

func NewAclRoleUserRepository(database *database.PostgresqlDatabase) IAclRoleUserRepository {
	repo := &AclRoleUserRepository{}
	repo.Database = database
	return repo
}

func (r *AclRoleUserRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.ACLRoleUser) (*model.ACLRoleUser, error) {
	stmt := table.ACLRoleUser.INSERT(columnList).MODEL(data).RETURNING(table.ACLRoleUser.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *AclRoleUserRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.ACLRoleUser) ([]*model.ACLRoleUser, error) {
	stmt := table.ACLRoleUser.INSERT(columnList).MODELS(data).RETURNING(table.ACLRoleUser.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *AclRoleUserRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ACLRoleUser) (*model.ACLRoleUser, error) {
	stmt := table.ACLRoleUser.UPDATE(columnList).MODEL(data).WHERE(table.ACLRoleUser.IDACLRoleUser.EQ(postgres.Int(int64(data.IDACLRoleUser)))).RETURNING(table.ACLRoleUser.AllColumns)
	return r.update(db, stmt)
}

func (r *AclRoleUserRepository) GetById(db qrm.Queryable, id int64) (*model.ACLRoleUser, error) {
	stmt := table.ACLRoleUser.SELECT(table.ACLRoleUser.AllColumns).WHERE(table.ACLRoleUser.IDACLRoleUser.EQ(postgres.Int(int64(id))))

	var data model.ACLRoleUser
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}

func (r *AclRoleUserRepository) GetByParam(db qrm.Queryable, param *model.ACLRoleUser) (*model.ACLRoleUser, error) {
	conditions := postgres.Bool(true)

	if param.IDACLRoleUser != 0 {
		conditions = conditions.AND(table.ACLRoleUser.IDACLRoleUser.EQ(postgres.Int(int64(param.IDACLRoleUser))))
	}

	if param.FkACLRole != 0 {
		conditions = conditions.AND(table.ACLRoleUser.FkACLRole.EQ(postgres.Int(int64(param.FkACLRole))))
	}

	if param.FkUser != 0 {
		conditions = conditions.AND(table.ACLRoleUser.FkUser.EQ(postgres.Int(int64(param.FkUser))))
	}

	if param.FkShop != 0 {
		conditions = conditions.AND(table.ACLRoleUser.FkShop.EQ(postgres.Int(int64(param.FkShop))))
	}

	stmt := table.ACLRoleUser.SELECT(table.ACLRoleUser.AllColumns).WHERE(conditions)

	var data model.ACLRoleUser
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
