package main

import (
	"context"
	"fmt"
	"log"

	gofeatureflag "github.com/open-feature/go-sdk-contrib/providers/go-feature-flag/pkg"
	of "github.com/open-feature/go-sdk/pkg/openfeature"
)

func main() {
	ctx := context.Background()

	provider, err := gofeatureflag.NewProviderWithContext(ctx, gofeatureflag.ProviderOptions{
		Endpoint: "http://localhost:1031",
	})
	if err != nil {
		log.Fatalln("gofeatureflag new provider: %w", err)
	}
	of.SetProvider(provider)

	client := of.NewClient("my-app")
	evaluationCtx := of.NewEvaluationContext(
		"1d1b9238-2591-4a47-94cf-d2bc080892f1",
		map[string]interface{}{
			"firstname": "mirai",
			"lastname":  "taro",
			"role":      "admin",
		})

	adminFlag, _ := client.BooleanValue(ctx, "flag-only-for-admin", false, evaluationCtx)
	if adminFlag {
		fmt.Println("アドミン向け機能ON")
	} else {
		fmt.Println("アドミン向け機能OFF")
	}

}
