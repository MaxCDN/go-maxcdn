package maxcdn

import (
	"fmt"
)

func ExampleResponse() {
	var data Account
	response, _ := max.Get(&data, Endpoint.Account, nil)
	fmt.Printf("%+v\n", response)
}

func ExampleGeneric() {
	var data Generic
	if _, err := max.Get(&data, Endpoint.Account, nil); err == nil {
		alias := data.Data["alias"].(string)
		name := data.Data["name"].(string)
		fmt.Printf("alias: %s\n", alias)
		fmt.Printf("name:  %s\n", name)
	}
}

func ExampleAccount() {
	var data Account
	if _, err := max.Get(&data, Endpoint.Account, nil); err == nil {
		fmt.Printf("%+v\n", data.Account)
	}
}

func ExampleAccountAddress() {
	var data AccountAddress
	if _, err := max.Get(&data, Endpoint.AccountAddress, nil); err == nil {
		fmt.Printf("%+v\n", data.Address)
	}
}

func ExamplePopularFiles() {
	var data PopularFiles
	if _, err := max.Get(&data, Endpoint.Reports.PopularFiles, nil); err == nil {
		for i, file := range data.PopularFiles {
			fmt.Printf("%2d: %30s=%s, \n", i, file.Uri, file.Hit)
		}
	}
	fmt.Println("----")
	fmt.Printf("    %30s=%s, \n", "summary", data.Summary.Hit)
}

func ExampleStatsSummary() {
	var data StatsSummary
	if _, err := max.Get(&data, Endpoint.Reports.Stats, nil); err == nil {
		fmt.Printf("%+v\n", data.Stats)
	}
}

func ExampleStats() {
	var data Stats
	if _, err := max.Get(&data, Endpoint.Reports.StatsBy("hourly"), nil); err == nil {
		fmt.Printf("%+v\n", data.Stats)
		fmt.Printf("%+v\n", data.Summary)
	}
}
