package main

import (
	"database/sql"
	"log"
	"os"
	"context"
	"github.com/amitader/Gator-RSS/internal/config"
	"github.com/amitader/Gator-RSS/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		log.Fatalf("error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	allCommands := commands{
		registeredCommands : make(map[string]func(*state, command) error),
	}
	allCommands.register("login", handlerLogin)
	allCommands.register("register", handlerRegister)
	allCommands.register("reset", handlerReset)
	allCommands.register("users", handleUsers)
	allCommands.register("agg", handlerAgg)
	allCommands.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	allCommands.register("feeds", handlerFeeds)
	allCommands.register("follow", middlewareLoggedIn(handlerFollow))
	allCommands.register("following", middlewareLoggedIn(handlerFollowing))
	allCommands.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	allCommands.register("browse", middlewareLoggedIn(handlerBrowse))

	userArgs := os.Args
	if len(userArgs) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}

	userCommand := command{
		Name : userArgs[1],
		Args: userArgs[2:],
	}
	err = allCommands.run(s, userCommand)
	if err != nil {
		log.Fatal(err)
	}

}

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return err
		}
		return handler(s, cmd, user)
	}
}


