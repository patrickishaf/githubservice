package services

import "testing"

func TestGetRepoByName(t *testing.T) {
	orgName := "patrickishaf"
	name := "rage"

	repo, err := GetRepoByName(orgName, name)

	if err != nil {
		t.Fatal(err)
	}

	if repo.Name != name {
		t.Fatal("repo name does not match")
	}
}

func TestGetRepoByFullName(t *testing.T) {
	orgName := "patrickishaf"
	name := "rage"

	_, err := GetCommitsInRepo(orgName, name)

	if err != nil {
		t.Fatal(err)
	}
}
