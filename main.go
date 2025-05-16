package main

import (
	"context"
	"fmt"

	"singleton_or_no/config"
	"singleton_or_no/repository"
)

func main() {
	ctx := context.Background()

	fmt.Println("🧪 Testing SINGLETON")
	singletonDB := config.GormConnectDB()
	repoSingleton := repository.NewSingletonRepo(singletonDB)
	for i := 0; i < 5; i++ {
		repoSingleton.All(ctx)
	}

	fmt.Println("\n\n🧪 Testing NON-SINGLETON")
	for i := 0; i < 5; i++ {
		repo := repository.NewNonSingletonRepo()
		repo.All(ctx)
	}
}
