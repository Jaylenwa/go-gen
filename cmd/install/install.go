package install

import (
	"gencode/util"
	"runtime"

	"github.com/gogf/gf/v2/os/gproc"
)

func Install() {
	goroot := runtime.GOROOT()

	src := ""
	dst := ""

	switch runtime.GOOS {
	case "windows":
		src = "./gencode.exe"
		dst = `C:\Program Files` + src
		if goroot != "" && len(goroot) > 0 {
			dst = goroot + "/bin" + src
		}
		_, _ = util.Copy(src, dst)

	default:
		src = "gencode"
		dst = `/usr/local/bin/` + src
		if goroot != "" && len(goroot) > 0 {
			dst = goroot + "/bin/" + src
		}
		_, _ = util.Copy(src, dst)
		_, _ = gproc.ShellExec("chmod -R 755 " + dst)
	}

	return
}
