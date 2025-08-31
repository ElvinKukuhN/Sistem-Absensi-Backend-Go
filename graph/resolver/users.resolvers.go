package resolver

import (
	"Sistem-Absensi-Backend-Go/Services/AuthService"
	"Sistem-Absensi-Backend-Go/graph"
	"Sistem-Absensi-Backend-Go/graph/model"
	"context"
)

// Resolver root (copy struct locally untuk menghindari cyclic)
type Resolver struct{}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) SignUp(ctx context.Context, input model.SignUpRequest) (*model.User, error) {
	// panggil fungsi SignUp kamu
	userResp, err := AuthService.SignUp(input.Name, input.Email, input.Password, input.Role)
	if err != nil {
		return nil, err
	}

	return &model.User{
		ID:        userResp.ID,
		Name:      userResp.Name,
		Email:     userResp.Email,
		Role:      userResp.Role,
		CreatedAt: userResp.CreatedAt,
		UpdatedAt: userResp.UpdatedAt,
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() graph.MutationResolver { return &mutationResolver{r} }

func (r *queryResolver) HealthCheck(ctx context.Context) (string, error) {
	return "Server is healthy!", nil
}

func (r *Resolver) Query() graph.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented: Todos - todos"))
}
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }
type queryResolver struct{ *Resolver }
*/
