package repository

import (
	"context"
	"fmt"
	"pzn_07_golang_database_mysql/database"
	"pzn_07_golang_database_mysql/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	comment := entity.Comment{
		Email: "repository@test.com",
		Comment: "Test repository",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindById(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindById(ctx, 25)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestCommentFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(database.GetConnection())

	ctx := context.Background()

	result, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}

	for _, comment := range result {
		fmt.Println(comment)
	}
}
