package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/Paulo-Lopes-Estevao/go_graphql_postegres/app"
	"github.com/Paulo-Lopes-Estevao/go_graphql_postegres/graph/generated"
	"github.com/Paulo-Lopes-Estevao/go_graphql_postegres/graph/model"
	"github.com/Paulo-Lopes-Estevao/go_graphql_postegres/utils"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {

	db := utils.ConnectDB()

	user := model.NewUser{
		Name:  input.Name,
		Email: input.Email,
		Idade: input.Idade,
	}

	repo := app.UserRespositoryDb{Db: db}
	result, err := repo.Insert(&user)

	if err != nil {
		panic(fmt.Errorf("not implemented"))
	}

	fmt.Println(result)

}

func (r *queryResolver) User(ctx context.Context) ([]*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
