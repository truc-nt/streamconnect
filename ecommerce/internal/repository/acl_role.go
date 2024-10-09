package repository

import (
	"ecommerce/internal/database"
	"ecommerce/internal/database/gen/model"
	"ecommerce/internal/database/gen/table"

	"github.com/go-jet/jet/v2/postgres"
	"github.com/go-jet/jet/v2/qrm"
)

type IAclRoleRepository interface {
	IBaseRepository[model.ACLRole]
}

type AclRoleRepository struct {
	BaseRepository[model.ACLRole]
}

func NewAclRoleRepository(database *database.PostgresqlDatabase) IAclRoleRepository {
	repo := &AclRoleRepository{}
	repo.Database = database
	return repo
}

func (r *AclRoleRepository) CreateOne(db qrm.Queryable, columnList postgres.ColumnList, data model.ACLRole) (*model.ACLRole, error) {
	stmt := table.ACLRole.INSERT(columnList).MODEL(data).RETURNING(table.ACLRole.AllColumns)
	return r.insertOne(db, stmt)
}

func (r *AclRoleRepository) CreateMany(db qrm.Queryable, columnList postgres.ColumnList, data []*model.ACLRole) ([]*model.ACLRole, error) {
	stmt := table.ACLRole.INSERT(columnList).MODELS(data).RETURNING(table.ACLRole.AllColumns)
	return r.insertMany(db, stmt)
}

func (r *AclRoleRepository) UpdateById(db qrm.Queryable, columnList postgres.ColumnList, data model.ACLRole) (*model.ACLRole, error) {
	stmt := table.ACLRole.UPDATE(columnList).MODEL(data).WHERE(table.ACLRole.IDACLRole.EQ(postgres.Int(int64(data.IDACLRole)))).RETURNING(table.ACLRole.AllColumns)
	return r.update(db, stmt)
}

func (r *AclRoleRepository) GetById(db qrm.Queryable, id int64) (*model.ACLRole, error) {
	stmt := table.ACLRole.SELECT(table.ACLRole.AllColumns).WHERE(table.ACLRole.IDACLRole.EQ(postgres.Int(int64(id))))

	var data model.ACLRole
	err := stmt.Query(db, &data)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
