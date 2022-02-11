package users

import (
	"context"
)

type Repository struct {
}

func (r *Repository) Create(ctx context.Context, payload *User) error {
	return nil
}

func (r *Repository) Get(ctx context.Context, uid int) (interface{}, error) {
	return nil, nil
}

func (r *Repository) Update(ctx context.Context, payload *User) error {
	return nil
}

func (r *Repository) Delete(ctx context.Context, uid int) error {
	return nil
}

// Modify must update only specific user field/fields
// todo: implement me
func (r *Repository) Modify() {}
