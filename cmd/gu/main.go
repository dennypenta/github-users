package main

import (
	"context"
	"fmt"
	"github.com/google/go-github/github"
	"github.com/joho/godotenv"
	"net/http"
	"sync"
	"time"
	"log"
	"os"
	"strconv"
)

const (
	defaultLimit = 5
	defaultDebugMode = false
)

var (
	debug = false
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	requestLimit := os.Getenv("MAX_PARALLEL_REQUESTS")
	limit, err := strconv.Atoi(requestLimit)
	if err != nil {
		limit = defaultLimit
	}

	if os.Getenv("DEBUG") == "" {
		debug = defaultDebugMode
	} else {
		debug = true
	}

	if debug {
		fmt.Println("max parallel requests: %s", requestLimit)
	}


	client := github.NewClient(http.DefaultClient)
	wg := &sync.WaitGroup{}

	var inProgress chan int64
	if limit <= 1 {
		inProgress = make(chan int64, 1)
	} else {
		inProgress = make(chan int64, limit)
	}

	for _, id := range []int64{18236918, 4324516, 9} {
		inProgress <- id
		wg.Add(1)
		fmt.Println("time: %s", time.Now())
		fmt.Println("start request for user #%s", id)
		go func(id int64) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*5000)
			defer println()
			defer cancel()
			defer wg.Done()
			defer func(){
				<- inProgress
				fmt.Println("finish request for user #%s", id)
			}()

			u, _, err := client.Users.GetByID(ctx, id)
			if err != nil {
				if err == context.DeadlineExceeded {
					fmt.Println("time is out")
				} else {
					fmt.Println(err.Error())
				}
				return
			}

			fmt.Println("it's a user #%d", *u.ID)
			fmt.Println("name: %s", *u.Name)
			fmt.Println("avatar url: %s", *u.AvatarURL)
			fmt.Println("public repos: %d", *u.PublicRepos)
		}(id)
	}

	wg.Wait()
}
