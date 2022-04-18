package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"

	"github.com/ingkiller/hackernews/graph/generated"
	"github.com/ingkiller/hackernews/graph/model"
	"github.com/ingkiller/hackernews/internal/album"
	"github.com/ingkiller/hackernews/internal/auth"
	"github.com/ingkiller/hackernews/internal/comment"
	"github.com/ingkiller/hackernews/internal/photo"
	"github.com/ingkiller/hackernews/internal/post"
	"github.com/ingkiller/hackernews/internal/todo"
	"github.com/ingkiller/hackernews/internal/user"
)

func (r *mutationResolver) ToggleTask(ctx context.Context, taskID int, userID int) (*model.Task, error) {
	var task = todo.ToggleTask(taskID, userID)
	var result = &model.Task{
		ID:        task.Id,
		UserID:    task.UserId,
		Title:     task.Title,
		Completed: task.Completed,
	}
	return result, nil
}

func (r *mutationResolver) CreateTask(ctx context.Context, title string, userID int) (*model.Task, error) {
	var newList = todo.CreateTask(title, userID)
	var result = &model.Task{
		ID:        newList.Id,
		UserID:    newList.UserId,
		Title:     newList.Title,
		Completed: newList.Completed,
	}
	return result, nil
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteTask(ctx context.Context, taskID int, userID int) (bool, error) {
	todo.DeleteTask(taskID, userID)
	return true, nil
}

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*model.UserData, error) {
	var user user.User
	user.Username = username
	user.Password = password

	result, err := user.Authenticate()
	if err != nil {
		return nil, errors.New("user or pass incorrect")
	}

	return &model.UserData{Token: result.Token,
		User: &model.User{
			ID:       &result.User.Id,
			Username: username,
			Name:     result.User.Name,
		}}, nil
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

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Albums(ctx context.Context) ([]*model.Album, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Photos(ctx context.Context) ([]*model.Photo, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	return user.GetAll(), nil
}

func (r *queryResolver) GetAlbumsByUserID(ctx context.Context, userID int) ([]*model.Album, error) {
	var result []*model.Album
	var albums []album.Album
	albums = album.GetAlbumByUserId(userID)
	for _, album := range albums {
		result = append(result, &model.Album{ID: album.Id,
			UserID:         album.UserId,
			Title:          album.Title,
			NumberOfPhotos: album.NumberOfPhotos,
		})
	}
	return result, nil
}

func (r *queryResolver) GetPhotosByAlbumID(ctx context.Context, albumID int) ([]*model.Photo, error) {
	var result []*model.Photo
	var photos []photo.Photo
	photos = photo.GetPhotosByAlbumId(albumID)
	for _, photo := range photos {
		result = append(result, &model.Photo{ID: photo.Id,
			AlbumID:      photo.AlbumId,
			Title:        photo.Title,
			URL:          photo.Url,
			ThumbnailURL: photo.ThumbnailUrl,
		})
	}
	return result, nil
}

func (r *queryResolver) GetCommentByPostID(ctx context.Context, postID int) ([]*model.Comment, error) {
	var result []*model.Comment
	var comments []comment.Comment
	comments = comment.GetCommentsByPost(postID)
	for _, comment := range comments {
		result = append(result, &model.Comment{ID: comment.Id,
			PostID: postID,
			Name:   comment.Name,
			Body:   comment.Body,
			Email:  comment.Email,
		})
	}

	return result, nil
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetTodoByUserID(ctx context.Context, userID int) ([]*model.Task, error) {
	var result []*model.Task
	var list []todo.Task

	user := auth.ForContext(ctx)

	if user == nil {
		fmt.Print("access denied GetTodoByUserID v%:")
		//	return result, fmt.Errorf("access denied bad bad")
	}

	list = todo.GetListByUserId(userID)
	for _, task := range list {
		result = append(result, &model.Task{ID: task.Id,
			UserID:    task.UserId,
			Title:     task.Title,
			Completed: task.Completed,
		})
	}
	return result, nil
}

func (r *queryResolver) GetPostByRange(ctx context.Context, start int, long int) ([]*model.Post, error) {
	var result []*model.Post
	var posts []post.Post
	posts = post.GetPostByRange(start, long)
	for _, post := range posts {
		fmt.Print("post.CreatedDatev%:", post.CreatedDate)
		result = append(result, &model.Post{ID: post.Id,
			Title:           post.Title,
			Body:            post.Body,
			NumberOfComment: post.NumberOfComment,
			CreatedDate:     post.CreatedDate.String(),
			User:            &model.User{Name: post.User.Name, Username: post.User.Username, Website: post.User.Website},
		})
	}
	return result, nil
}

func (r *queryResolver) GetPostsByUserID(ctx context.Context, userID int, start int, long int) ([]*model.Post, error) {
	return post.GetPostsByUserId(userID, start, long), nil
}

func (r *queryResolver) GetUserDataByUsername(ctx context.Context, tokenStr string) (*model.User, error) {
	result, err := user.GetUserDataByUsername(tokenStr)
	if err != nil {
		return &model.User{}, err
	}

	return &model.User{ID: &result.Id,
		Username: result.Username,
		Name:     result.Name,
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
