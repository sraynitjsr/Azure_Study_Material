package main

import (
	"context"
	"fmt"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/compute/mgmt/compute"
	"github.com/Azure/go-autorest/autorest/azure/auth"
)

func main() {
	// Create an Azure authorizer from environment variables or Azure Managed Service Identity
	authorizer, err := auth.NewAuthorizerFromEnvironment()
	if err != nil {
		fmt.Println("Failed to create authorizer:", err)
		return
	}

	// Create a compute client using the authorizer
	vmClient := compute.NewVirtualMachinesClient("Actual Subscription ID")
	vmClient.Authorizer = authorizer

	// List the virtual machines in the subscription
	vms, err := vmClient.List(context.Background(), "Actual Resource Group Name")
	if err != nil {
		fmt.Println("Failed to list virtual machines:", err)
		return
	}

	// Print the virtual machine names
	for _, vm := range vms.Values() {
		fmt.Println(*vm.Name)
	}
}
