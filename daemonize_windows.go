package daemonize

import (
	"fmt"
	"os"
	"path/filepath"

	"bitbucket.org/kardianos/osext"
	log "github.com/cihub/seelog"
)

func Daemonize() (logger log.LoggerInterface) {

	path, _ := osext.ExecutableFolder()
	os.Chdir(path)

	cfgfile := fmt.Sprintf("./%s.console.xml", filepath.Base(os.Args[0]))
	see_log_cfg_file_name = cfgfile

	var err error

	if logger, err = log.LoggerFromConfigAsFile(cfgfile); err != nil {
		logger, _ = log.LoggerFromConfigAsString(def_console_log_cfg)
	}

	if err = log.ReplaceLogger(logger); err != nil {
		fmt.Println("Replace SeeLog failed:", err)
		os.Exit(200)
	}

	return
}
