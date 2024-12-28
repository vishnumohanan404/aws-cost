package main

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer"
	"github.com/aws/aws-sdk-go-v2/service/costexplorer/types"
)

func main() {
	// Input arguments (hardcoded for simplicity, replace with flags as needed)
	startDate := "2024-12-01"
	endDate := "2024-12-31"

	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		log.Fatalf("Failed to load AWS config: %v", err)
	}

	// Create Cost Explorer client
	client := costexplorer.NewFromConfig(cfg)

	// Call GetCostAndUsage API with grouping by SERVICE
	result, err := client.GetCostAndUsage(context.Background(), &costexplorer.GetCostAndUsageInput{
		TimePeriod: &types.DateInterval{
			Start: aws.String(startDate),
			End:   aws.String(endDate),
		},
		Granularity: types.GranularityMonthly,  // No daily granularity needed
		Metrics:     []string{"UnblendedCost"}, // Adjust metrics as needed
		GroupBy: []types.GroupDefinition{
			{
				Type: types.GroupDefinitionTypeDimension, // Use the provided constant
				Key:  aws.String("SERVICE"),
			},
		},
	})
	if err != nil {
		log.Fatalf("Failed to retrieve cost data: %v", err)
	}

	// Display results in a table
	fmt.Printf("AWS Cost Breakdown by Service (%s to %s):\n", startDate, endDate)
	fmt.Printf("%-30s %15s\n", "Service", "Cost (USD)")
	fmt.Println(strings.Repeat("-", 45))

	totalCost := 0.0
	for _, group := range result.ResultsByTime[0].Groups {
		service := group.Keys[0] // Use directly as it's already a string
		cost, _ := strconv.ParseFloat(aws.ToString(group.Metrics["UnblendedCost"].Amount), 64)
		fmt.Printf("%-30s %15.2f\n", service, cost)
		totalCost += cost
	}

	fmt.Println(strings.Repeat("-", 45))
	fmt.Printf("%-30s %15.2f\n", "Total", totalCost)
}
