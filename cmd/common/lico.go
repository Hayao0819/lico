package common

import (
	"encoding/base64"
	"os"

	"fmt"

	//p "github.com/Hayao0819/lico/paths"
	"github.com/Hayao0819/lico/utils"
	"github.com/Hayao0819/lico/vars"
	random "github.com/mazen160/go-random"
	//"github.com/spf13/cobra"
)

func HasCorrectRepoDir() bool {
	isDir := utils.IsDir(vars.RepoDir)
	//hasGitDir := utils.Exists(path.Join(vars.RepoDir, ".git"))
	hasListFile := utils.Exists(vars.GetList())
	//cmd.Println(isDir)
	//cmd.Println(hasGitDir)
	return isDir && hasListFile //&& hasGitDir
}

/*
func RunCmd(f func() *cobra.Command, args ...string) error {
	//return f().RunE(f(), args)
	cmd := f()
	cmd.SilenceUsage=true
	cmd.SilenceErrors=true
	cmd.SetArgs(args)
	return cmd.Execute()
}
*/

/*
func formatHomePath(path *p.Path) (*p.Path, error) {
	rtn, err := path.Abs(*homePathBase)
	if ; err != nil {
		return nil, err
	}
	return &rtn, nil
}

func formatRepoPath(path *p.Path) (*p.Path, error) {
	rtn, err := path.Abs(*repoPathBase)
	if err != nil {
		return nil, err
	}
	return &rtn, nil
}
*/

func Lico() string {
	list := []string{
		"44KI44GP6KaL44Gf44KJ44GT44Gu5a2Q44Gq44KT44GL5aSJ44Gq44KT55Sf44GI44Go44KL77yB77yBCg==",
		"5LmF44CF44Gr44GE44GN44Gj44Gf5ber5aWz44Gv44KT44KS44GT44GI44Gg44KB44Gr6JC944Go44Gb44KL44KP44GBCg==",
		"5LuK5pel44Gv44G244G25rys44GR44GX44GL44Gq44GE44KT44KE44GR44Gp5aSn5LiI5aSr77yfCg==",
		"6Z2e5oim6ZeY5ZOh44Gq5LiK44Gr44GI44GS44Gk44Gq44GP5byx44GE44GR44GM5Lq644Gu44Oe44K544K/44O844GM44Km44OB44KS5a6I44KK44CB5pyJ5Yip44Gr5Lqk5riJ44GZ44KL44Gf44KB44Gr44OP44OD44K/44Oq44KS44GL44GR44Gm44GP44KM44KL44Gq44KT44GmLi4u6KaL55u044GX44Gf44KPCg==",
		"44GX44Gz44KM44KL44KE44GkIOOBreOCi+OChOOBpCDkvZPjgYzjgbvjgabjgovjgoTjgaQK",
		"5ber5aWz44Gv44KT44GM5YyW44GL44GV44KM44Gm6I2J44KS6aOf44KA44Go44GT44G/44Gm56yR44GE44Gf44GE5rCX5oyB44Gh44KC44G+44O844G+44O844GC44Gj44Gf44KICg==",
		"552A5bSp44KM5pat5Zu66Zi75q2i44Gu6KGTfua3seWknOOBq+OBr+iEseOBkuOCi+OBrgo=",
	}

	i, _ := random.IntRange(0, len(list))
	dec, _ := base64.StdEncoding.DecodeString(list[i])

	return string(dec)
}

func GetRepoUrl() ([]string, error) {
	if !HasCorrectRepoDir() {
		return []string{}, vars.ErrNoRepoDir
	}

	remoteList, stderr, err := utils.RunCmdAndGet("git", "-C", vars.RepoDir, "remote", "show")
	if err != nil {
		fmt.Fprintln(os.Stderr, stderr)
		return []string{}, err
	}

	urlList := []string{}
	for _, r := range remoteList {
		if utils.IsEmpty(r) {
			continue
		}
		u, stderr, err := utils.RunCmdAndGet("git", "-C", vars.RepoDir, "config", "--get", fmt.Sprintf("remote.%v.url", r))
		if err != nil {
			fmt.Fprintln(os.Stderr, stderr)
			return []string{}, err
		}
		urlList = append(urlList, u[0])
	}
	return urlList, err
}
