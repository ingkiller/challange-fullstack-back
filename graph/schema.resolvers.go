package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ingkiller/hackernews/graph/generated"
	"github.com/ingkiller/hackernews/graph/model"
	"github.com/ingkiller/hackernews/internal/comment"
	"github.com/ingkiller/hackernews/internal/post"
	"github.com/ingkiller/hackernews/internal/story"
)

func (r *mutationResolver) ToggleTask(ctx context.Context, input model.Task) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) GetCommentByPostIDMutation(ctx context.Context, input model.PostID) ([]*model.Comment, error) {
	var result []*model.Comment
	var comments []comment.Comment
	comments = comment.GetCommentsByPost(input.PostID)
	for _, comment := range comments {
		result = append(result, &model.Comment{ID: comment.Id,
			PostID: input.PostID,
			Name:   comment.Name,
			Body:   comment.Body,
			Email:  comment.Email,
		})
	}

	return result, nil
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Stories(ctx context.Context) ([]*model.Story, error) {
	var resultLinks []*model.Story
	var dbLinks []story.Story
	dbLinks = story.GetAll()

	for _, story := range dbLinks {
		resultLinks = append(resultLinks, &model.Story{ID: story.Id,
			Title:       story.Title,
			By:          story.By,
			Descendants: story.Descendants,
			Kids:        story.Kids,
			Score:       story.Score,
			Time:        story.Time,
			Type:        story.Type,
			URL:         story.URL,
		})
	}
	return resultLinks, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	var result []*model.Post
	var dbLinks []post.Post
	dbLinks = post.GetAll()

	for _, post := range dbLinks {

		result = append(result, &model.Post{ID: post.Id,
			Title:           post.Title,
			Body:            post.Body,
			NumberOfComment: post.NumberOfComment,
			User:            &model.User{Name: post.User.Name, Username: post.User.Username, Website: post.User.Website},
		})
	}

	return result, nil
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todo(ctx context.Context) ([]*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Albums(ctx context.Context) ([]*model.Album, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Photos(ctx context.Context) ([]*model.Photo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetCommentByPostID(ctx context.Context, input model.PostID) ([]*model.Comment, error) {
	var result []*model.Comment
	var comments []comment.Comment
	comments = comment.GetCommentsByPost(input.PostID)
	for _, comment := range comments {
		result = append(result, &model.Comment{ID: comment.Id,
			PostID: input.PostID,
			Name:   comment.Name,
			Body:   comment.Body,
			Email:  comment.Email,
		})
	}

	return result, nil
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
