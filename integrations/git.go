package integrations

import (
	"path"

	"github.com/chrisledet/rebasebot/git"
	"github.com/chrisledet/rebasebot/github"
)

// var witty_witty []string;

// var witty_fails []string;

// func WittyCharacterInit(){
// 	witty_witty = []string{"apple",
// 		"banana",
// 		"kiwi"}

// 		witty_fails = []string{"Done!\n Probably you heard this before:\n I did it, but I did not enjoy it...",

			
// 		}
// }


// Ties the git operations together to perform a branch rebase
func GitRebase(pr *github.PullRequest) error {
	
	filepath := git.GetRepositoryFilePath(pr.Head.Repository.FullName)
	remoteRepositoryURL := git.GenerateCloneURL(pr.Head.Repository.FullName)

	if !git.Exists(filepath) {
		if _, err := git.Clone(remoteRepositoryURL); err != nil {
			pr.PostComment("I could not pull " + pr.Head.Repository.FullName + " from GitHub.")
			return err
		}
	}

	if err := git.Fetch(filepath); err != nil {
		git.Prune(filepath)
		pr.PostComment("I could not fetch the latest changes from GitHub. Please try again in a few minutes.")
		return err
	}

	if err := git.Checkout(filepath, pr.Head.Ref); err != nil {
		pr.PostComment("I could not checkout " + pr.Head.Ref + " locally.")
		return err
	}

	if err := git.Reset(filepath, path.Join("origin", pr.Head.Ref)); err != nil {
		pr.PostComment("I could not checkout " + pr.Head.Ref + " locally.")
		return err
	}

	if err := git.Config(filepath, "user.name", git.GetName()); err != nil {
		pr.PostComment("I could run git config for user.name on the server.")
		return err
	}

	if err := git.Config(filepath, "user.email", git.GetEmail()); err != nil {
		pr.PostComment("I could run git config for user.email on the server.")
		return err
	}

	if err := git.Rebase(filepath, path.Join("origin", pr.Base.Ref)); err != nil {
		pr.PostComment("I could not rebase " + pr.Head.Ref + " with " + pr.Base.Ref + ". There are conflicts.")
		return err
	}

	if err := git.Push(filepath, pr.Head.Ref); err != nil {
		pr.PostComment("I could not push the changes to " + pr.Base.Ref + "." + "Next time use your fingers for more than just picking your nose.")
		return err
	}

	pr.PostComment("Rebase done!")
	return nil
}


// Ties the git operations together to perform a branch rebase
func GitMerge(pr *github.PullRequest, message string) error {

	filepath := git.GetRepositoryFilePath(pr.Head.Repository.FullName)
	remoteRepositoryURL := git.GenerateCloneURL(pr.Head.Repository.FullName)

	if !git.Exists(filepath) {
		if _, err := git.Clone(remoteRepositoryURL); err != nil {
			pr.PostComment("I could not pull " + pr.Head.Repository.FullName + " from GitHub.")
			return err
		}
	}

	if err := git.Fetch(filepath); err != nil {
		git.Prune(filepath)
		pr.PostComment("I could not fetch the latest changes from GitHub. Please try again in a few minutes.")
		return err
	}

	if err := git.Checkout(filepath, pr.Head.Ref); err != nil {
		pr.PostComment("I could not checkout " + pr.Head.Ref + " locally.")
		return err
	}

	if err := git.Reset(filepath, path.Join("origin", pr.Head.Ref)); err != nil {
		pr.PostComment("I could not checkout " + pr.Head.Ref + " locally.")
		return err
	}

	if err := git.Config(filepath, "user.name", git.GetName()); err != nil {
		pr.PostComment("I could run git config for user.name on the server.")
		return err
	}

	if err := git.Config(filepath, "user.email", git.GetEmail()); err != nil {
		pr.PostComment("I could run git config for user.email on the server.")
		return err
	}

	if err := git.Rebase(filepath, path.Join("origin", pr.Base.Ref)); err != nil {
		pr.PostComment("I could not rebase " + pr.Head.Ref + " with " + pr.Base.Ref + ". There are conflicts.")
		return err
	}

	if err := git.Push(filepath, pr.Head.Ref); err != nil {
		pr.PostComment("I could not push the changes to " + pr.Base.Ref + ".")
		return err
	}

	if err := git.Checkout(filepath, pr.Base.Ref); err != nil {
		pr.PostComment("I could not checkout " + pr.Base.Ref + " locally.")
		return err
	}

	if err := git.Fetch(filepath); err != nil {
		git.Prune(filepath)
		pr.PostComment("I could not fetch the latest changes from GitHub. Please try again in a few minutes.")
		return err
	}

	if err := git.Reset(filepath, path.Join("origin", pr.Base.Ref)); err != nil {
		pr.PostComment("I could not checkout " + pr.Base.Ref + " locally.")
		return err
	}

	if err := git.Merge(filepath, pr.Head.Ref, message); err != nil {

		pr.PostComment("I could not merge " + pr.Head.Ref + " into " + pr.Base.Ref + "." + "\nNext time use your fingers for more than just picking your nose.")
		return err
	}

	if err := git.Push(filepath, pr.Base.Ref); err != nil {
		pr.PostComment("I could not push the changes to " + pr.Base.Ref + ".")
		return err
	}

	pr.PostComment("I just merged " + pr.Head.Ref + " into " + pr.Base.Ref+ "\nProbably you heard this before:\nI did it, but I did not enjoy it...")
	return nil
}
