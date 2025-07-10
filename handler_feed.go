package main
import (
	"fmt"
	"context"
	"time"
	"github.com/amitader/Gator-RSS/internal/database"
	"github.com/google/uuid"
)
func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.Args) != 2 {
		return fmt.Errorf("usage: %s <name> <url>", cmd.Name)
	}
	name := cmd.Args[0]
	url := cmd.Args[1]
	feed, err := s.db.CreateFeed(context.Background(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		Name: name, 
		Url: url,
	})
	if err != nil {
		return fmt.Errorf("couldn't create a feed: %w", err)
	}
	_, err = s.db.InsertFeedFollow(context.Background(),database.InsertFeedFollowParams{
		ID: uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		return fmt.Errorf("couldn't add a feed follow: %w", err)
	}
	return nil
}

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("couldn't get feeds: %w", err)
	}
	for _, feed := range feeds {
		userName, err := s.db.GetUserByID(context.Background(),feed.UserID)
		if err != nil {
			return fmt.Errorf("couldn't find user name by his id: %w", err)
		}	
		fmt.Printf("Feed name: %s\n", feed.Name)
		fmt.Printf("URL: %s\n", feed.Url)
		fmt.Printf("owner of the feed: %s\n", userName)
		fmt.Println("--------------------------")
	}
	return nil
}