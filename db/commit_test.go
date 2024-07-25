package db

import (
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestConvertToCommit(t *testing.T) {
	inaccurateConversionError := "failed to commit to CommitResponse successfully"
	response := db.CommitResponse{
		Hash: "commitresponsehash",
		Commit: db.CommitData{
			Message: "commitresponsemessage",
			Author: db.Author{
				Name:  "commitresponseauthorname",
				Email: "commitresponseauthoremail",
				Date:  "commitresponseauthordate",
			},
		},
	}
	commit := response.ConvertToCommit()

	if (reflect.TypeOf(*commit)) != reflect.TypeOf(db.Commit{}) {
		t.Fatal("failed to convert CommitResponse to Commit")
	}

	if response.Hash != commit.Hash || response.Commit.Message != commit.Message || response.Commit.Author.Name != commit.AuthorName || response.Commit.Author.Email != commit.AuthorEmail {
		t.Fatal(inaccurateConversionError)
	}

	if response.Commit.Author.Date != commit.Date {
		t.Fatal(inaccurateConversionError)
	}
}

func TestConvertToCommitResponse(t *testing.T) {
	inaccurateConversionError := "failed to commit to CommitResponse successfully"
	commit := db.Commit{
		Hash:        "hash",
		Message:     "message",
		AuthorName:  "authorname",
		AuthorEmail: "authoremail",
		Date:        "date",
	}
	response := commit.ConvertToCommitResponse()

	if reflect.TypeOf(*response) != reflect.TypeOf(db.CommitResponse{}) {
		t.Fatal("failed to convert Commit to CommitResponse")
	}

	if (*response).Hash != commit.Hash || (*response).Commit.Message != commit.Message || (*response).Commit.Author.Name != commit.AuthorName || (*response).Commit.Author.Email != commit.AuthorEmail {
		t.Fatal(inaccurateConversionError)
	}

	if (*response).Commit.Author.Date != commit.Date {
		t.Fatal(inaccurateConversionError)
	}
}
