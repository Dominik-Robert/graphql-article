package book_api

import (
	"context"
	"errors"
) // THIS CODE IS A STARTING POINT ONLY. IT WILL NOT BE UPDATED WITH SCHEMA CHANGES.

type Resolver struct {
	MyBooks []*Book
}

func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateBook(ctx context.Context, input NewBook) (*Book, error) {
	newBook := &Book{
		Isbn: input.Isbn,
		Name: input.Isbn,
		Author: &Author{
			Name: input.Authorname,
		},
	}

	r.MyBooks = append(r.MyBooks, newBook)

	return newBook, nil
}

func (r *mutationResolver) DeleteBook(ctx context.Context, isbn string) (*Book, error) {
	for i, book := range r.MyBooks {
		if book.Isbn == isbn {
			r.MyBooks = append(r.MyBooks[:i], r.MyBooks[i+1:]...)
			return &Book{Isbn: isbn}, nil
		}
	}
	return nil, errors.New("No isbn found for deletion")
}

func (r *mutationResolver) UpdateBook(ctx context.Context, input UpdateBook) (*Book, error) {
	updatedBook := &Book{
		Isbn: input.Isbn,
		Name: input.Isbn,
		Author: &Author{
			Name: *input.Authorname,
		},
	}

	for i, book := range r.MyBooks {
		if book.Isbn == input.Isbn {
			r.MyBooks[i] = updatedBook
			return updatedBook, nil
		}
	}

	return nil, errors.New("No isbn found for deletion")
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Books(ctx context.Context) ([]*Book, error) {
	return r.MyBooks, nil
}

func (r *queryResolver) Book(ctx context.Context, isbn string) (*Book, error) {
	for _, book := range r.MyBooks {
		if book.Isbn == isbn {
			return book, nil
		}
	}
	return nil, errors.New("Found no Book with this ISBN")
}
