package main

import (
	"fmt"
	"os"
	"time"

	sdk "github.com/bitwarden/sdk-go/v2"
)

func main() {
	apiURL := os.Getenv("API_URL")
	identityURL := os.Getenv("IDENTITY_URL")
	organizationID := os.Getenv("ORGANIZATION_ID")
	accessToken := os.Getenv("ACCESS_TOKEN")
	stateFile := "/tmp/bw.state"

	client, err := sdk.NewBitwardenClient(&apiURL, &identityURL)
	if err != nil {
		panic("Could not create Bitwarden client")
	}

	// Generators
	request := sdk.PasswordGeneratorRequest{
		AvoidAmbiguous: true,
		Length:         64,
		Lowercase:      true,
		MinLowercase:   new(int64),
		MinNumber:      new(int64),
		MinSpecial:     new(int64),
		MinUppercase:   new(int64),
		Numbers:        true,
		Special:        true,

		Uppercase: true,
	}

	generatedPassword, err := client.Generators().GeneratePassword(request)
	if err != nil {
		panic("Could not generate password")
	}

	fmt.Println("Generators:")
	fmt.Println("  generated password:", *generatedPassword)
	fmt.Println()

	// Authentication
	client.AccessTokenLogin(accessToken, &stateFile)

	// Projects
	fmt.Println("Projects:")
	newProject, err := client.Projects().Create(organizationID, "New Project from Go SDK")
	if err != nil {
		panic("Could not create `New Project from Go SDK`")
	}

	fmt.Println("  create:", *&newProject.Name)

	getProject, err := client.Projects().Get(newProject.ID)
	if err != nil {
		panic("Could not get project")
	}

	fmt.Println("  read:  ", *&getProject.Name)

	updatedProject, err := client.Projects().Update(newProject.ID, organizationID, "New Project from Go SDK - Updated")
	if err != nil {
		panic("Could not update `New Project from Go SDK`")
	}

	fmt.Println("  update:", *&updatedProject.Name)

	listOfProjects, err := client.Projects().List(organizationID)
	if err != nil {
		panic("Could not list projects")
	}

	fmt.Println("  list:  ", *&listOfProjects.Data)
	fmt.Println()

	fmt.Println("  # deletions are deferred until the end...")
	fmt.Println()

	// Secrets
	fmt.Println("Secrets:")

	lastSyncedDate := time.Now() // prepare for syncing secrets...
	secretsSync, err := client.Secrets().Sync(organizationID, nil)
	if err != nil {
		panic("Could not do initial sync of secrets with org")
	}
	fmt.Println("  sync.HasChanges (initial): ", *&secretsSync.HasChanges, "# should say `true`")
	redundantSecretsSync, err := client.Secrets().Sync(organizationID, &lastSyncedDate)
	if err != nil {
		panic("Could not do redundant sync of secrets with org")
	}
	fmt.Println("  sync.HasChanges (redunant):", *&redundantSecretsSync.HasChanges, "# should say `false`")

	newSecret, err := client.Secrets().Create("New Secret from Go SDK", "super secret value", "my note", organizationID, []string{updatedProject.ID})
	if err != nil {
		panic("Could not create `New Secret from Go SDK`")
	}

	fmt.Println("  create:                    ", *&newSecret.Key)

	getSecret, err := client.Secrets().Get(newSecret.ID)
	if err != nil {
		panic("Could not get secret")
	}

	fmt.Println("  read (get):                ", *&getSecret.Key)

	getSecretsByIDs, err := client.Secrets().GetByIDS([]string{newSecret.ID})
	if err != nil {
		panic("Could not get secret with GetByIDs")
	}

	fmt.Println("  read (GetByIDs):           ", *&getSecretsByIDs.Data[len(getSecretsByIDs.Data)-1])
	fmt.Println()

	updatedSecret, err := client.Secrets().Update(newSecret.ID, "New Secret from Go SDK - Updated", "super secret value", "my note", organizationID, []string{updatedProject.ID})
	if err != nil {
		panic("Could not update `New Secret from Go SDK`")
	}

	fmt.Println("  update:                    ", *&updatedSecret.Key)

	listOfSecrets, err := client.Secrets().List(organizationID)
	if err != nil {
		panic("Could not list secrets")
	}

	fmt.Println("  list:                      ", *&listOfSecrets.Data)
	fmt.Println()

	secretsSync, err = client.Secrets().Sync(organizationID, &lastSyncedDate)
	fmt.Println("  sync.HasChanges (final):   ", *&secretsSync.HasChanges, " # should say `true`")

	fmt.Println("  # deletions are deferred until the end...")
	fmt.Println()

	// Deletions
	fmt.Println("Deletions:")
	projectDeletionResponse, err := client.Projects().Delete([]string{newProject.ID})
	if err != nil {
		panic("Could not delete project")
	}

	fmt.Println("  project:", *&projectDeletionResponse)

	secretDeletionResponse, err := client.Secrets().Delete([]string{newSecret.ID})
	if err != nil {
		panic("Could not delete secret")
	}

	fmt.Println("  secret: ", *&secretDeletionResponse)
	fmt.Println()
}
