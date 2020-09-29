package main

import (
	"os"
	"path"
	"strings"

	aw "github.com/deanishe/awgo"
)

var (
	wf            *aw.Workflow
	gitHubIcon    = &aw.Icon{Value: path.Join("github-logo.png")}
	bitBucketIcon = &aw.Icon{Value: path.Join("bitbucket-logo.png")}
	gitIcon       = &aw.Icon{Value: path.Join("git-logo.png")}
	modKeys       = []aw.ModKey{
		aw.ModShift,
		aw.ModFn,
	}
)

func init() {
	wf = aw.New()
}

func main() {
	wf.Run(func() {
		query := strings.Trim(os.Args[1], " \n")
		repos := os.Args[2:len(os.Args)]
		for _, repo := range repos {
			addNewItem(repo)
		}
		if len(query) > 0 {
			wf.Filter(query)
		}
		wf.WarnEmpty("No matching repository", "Try different query?")
		wf.SendFeedback()
	})
}

func addNewItem(repo string) {
	repoPath := strings.Split(repo, "/")
	it := wf.NewItem(repo).
		Title(excludeDomain(repoPath, true)).
		UID(repo).
		Arg(repo).
		Subtitle(getDomainName(repoPath)).
		Icon(getIcon(repoPath)).
		Valid(true)
	for _, modKey := range modKeys {
		mod := createExtraModItem(repoPath, repo, modKey)
		it.SetModifier(mod)
	}
}

func excludeDomain(repo []string, domain bool) string {
	// full_repo_path: strings.Split("/path/to/github.com/user/full_repo_path", "/")
	var i int
	if domain {
		// return user/full_repo_path
		i = 2
	} else {
		// return github.com/user/full_repo_path
		i = 3
	}
	length := len(repo)
	return strings.Join(repo[length-i:length], "/")
}

func getDomainName(repo_path []string) string {
	// return github.com
	return repo_path[len(repo_path)-3]
}

func createExtraModItem(repo []string, path string, modKey aw.ModKey) *aw.Modifier {
	var arg string
	switch modKey {
	case aw.ModShift:
		// Returns repository url with shift modifier.
		// Users can use it for thing like opening in a browser.
		arg = "https://" + excludeDomain(repo, false) + "/"
	case aw.ModFn:
		// Returns [repo owner]/[repo name] with alt modifier.
		// Users can use it for thing like searching on Google.
		arg = excludeDomain(repo, true)
	}
	mod := &aw.Modifier{Key: modKey}
	return mod.
		Arg(arg).
		Valid(true)
}

func getIcon(repoPath []string) *aw.Icon {
	domain := getDomainName(repoPath)
	switch {
	case strings.Contains(domain, "github"):
		return gitHubIcon
	case strings.Contains(domain, "bitbucket"):
		return bitBucketIcon
	default:
		return gitIcon
	}
}
