package main

import (
	"context"
	"github.com/amitader/Gator-RSS/internal/database"
	"fmt"
	"strconv"
)
func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.Args) == 1{
		val, err := strconv.Atoi(cmd.Args[0])
		if err != nil {
			return fmt.Errorf("couldn't convert user input to int: %w", err)
		}
		limit = val
	}
	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{UserID: user.ID, Limit: int32(limit),})
	if err != nil {
		return fmt.Errorf("couldn't get the user posts: %w", err)
	}
	fmt.Printf("Found %d posts for user %s:\n", len(posts), user.Name)
	for _, post := range posts {
		fmt.Printf("%s from %s\n", post.PublishedAt.Time.Format("Mon Jan 2"), post.FeedName)
		fmt.Printf("--- %s ---\n", post.Title)
		fmt.Printf("    %v\n", post.Description.String)
		fmt.Printf("Link: %s\n", post.Url)
		fmt.Println("=====================================")
	}
	return nil
		
}