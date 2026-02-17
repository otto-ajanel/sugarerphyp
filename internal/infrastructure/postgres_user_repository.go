package infrastructure

import (
	"context"
	"database/sql"
	"sugarerpgo/internal/core"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func (r *PostgresUserRepository) GetUsers(ctx context.Context, tenantID string) ([]core.User, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id_user, username, email, id_tenant FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []core.User
	for rows.Next() {
		var u core.User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.TenantID); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}
