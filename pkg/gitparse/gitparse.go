
	"github.com/trufflesecurity/trufflehog/v3/pkg/common"
	"github.com/trufflesecurity/trufflehog/v3/pkg/context"
func RepoPath(ctx context.Context, source string, head string) (chan Commit, error) {
		defer common.Recover(ctx)
	if len(line) >= 6 && bytes.Equal(line[:3], []byte("---")) {
	if len(line) >= 6 && bytes.Equal(line[:3], []byte("+++")) {