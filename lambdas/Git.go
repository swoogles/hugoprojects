package main

type SmallCommit struct {
	Message string
}

type CommitList struct {
	Commits []GitHubCommit
}

type GitHubCommit struct {
	Sha string
	Commit SmallCommit
}

