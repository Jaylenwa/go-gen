package gen

import (
	"os"
)

func GenREADME(req GenReq) {

	srcFile := "cmd/gen/README.MD"
	dstFile := req.BaseDir + "/README.MD"

	srcData, err := os.ReadFile(srcFile)
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(dstFile, srcData, 0644)
	if err != nil {
		panic(err)
	}
}
