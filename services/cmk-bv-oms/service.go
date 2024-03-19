package main

import (
	"context"
	"fmt"
	"os"

	"github.com/oracle/oci-go-sdk/v65/common"
	"github.com/oracle/oci-go-sdk/v65/core"
)

func updateVolumeKmsKey() {
	// Configure the OCI SDK
	config, _ := common.NewConfiguration(common.DefaultConfig, common.DefaultConfigFilePath())

	// Read values from .env file
	tenancyOCID := os.Getenv("TENANCY_OCID")
	userOCID := os.Getenv("USER_OCID")
	privateKeyPath := os.Getenv("PRIVATE_KEY_PATH")
	volumeOCID := os.Getenv("VOLUME_OCID")
	kmsKeyOCID := os.Getenv("KMS_KEY_OCID")

	// Create the authentication provider
	authProvider, _ := common.NewFileAuthenticationDetailsProvider(userOCID, privateKeyPath)

	// Create the client with the authentication information
	client, err := core.NewVolumserviceClient(config, authProvider)
	if err != nil {
		fmt.Println("Error creating Volume service client:", err)
		return
	}

	// Update the KMS key of the volume
	request := core.UpdateVolumeRequest{
		VolumeId: common.String(volumeOCID),
		UpdateVolumeDetails: core.UpdateVolumeDetails{
			KmsKeyId: common.String(kmsKeyOCID),
		},
	}

	response, err := client.UpdateVolume(context.Background(), request)
	if err != nil {
		fmt.Println("Error updating volume KMS key:", err)
	} else {
		fmt.Println("Volume KMS key updated successfully. New state:", response.Header.Get("opc-work-request-id"))
	}
}

func main() {
	updateVolumeKmsKey()
}
