package daemonize

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"bitbucket.org/kardianos/osext"

	log "github.com/cihub/seelog"

	"github.com/tyranron/daemonigo"
)

const (
	daemonigo_flag   = "daemonigo_flag"
	flag_fore_ground = "run_in_foreground"
	flag_no_log      = "run_without_log"
)

func Daemonize() (logger log.LoggerInterface) {

	path, _ := osext.ExecutableFolder()
	os.Chdir(path)

	name := filepath.Base(os.Args[0])
	daemonigo.PidFile = name + ".pid"

	daemonigo.AppName = name
	daemonigo.AppPath, _ = osext.Executable()

	var start_flag string

	start_process_with_flag := func(flag string) {
		running, process, err := daemonigo.Status()
		if running && process != nil && err == nil {
			fmt.Println(daemonigo.AppName + " is running,please stop it and try again.")
			os.Exit(200)
		} else if flag == flag_fore_ground {
			os.Setenv(daemonigo_flag, flag)
			start_flag = flag
		} else {
			fmt.Printf("Starting %s...", daemonigo.AppName)
			os.Setenv(daemonigo_flag, flag)
			if err := daemonigo.Start(1); err != nil {
				fmt.Println("FAILED")
				fmt.Println("Details:", err.Error())
				os.Exit(300)
			} else {
				fmt.Println("OK")
			}
		}
	}

	daemonigo.SetAction("fg", func() { start_process_with_flag(flag_fore_ground) })
	daemonigo.SetAction("nolog", func() { start_process_with_flag(flag_no_log) })

	if ischild, err := daemonigo.Daemonize(); err != nil {
		fmt.Println("Daemonize failed:", err)
		os.Exit(100)
	} else if ischild || start_flag == flag_fore_ground {
		logger = initLogger(name)
		initRlimit()
	} else {
		os.Exit(0)
	}

	return
}

func initLogger(name string) (logger log.LoggerInterface) {

	switch start_flag := os.Getenv(daemonigo_flag); start_flag {
	case flag_no_log:
		logger, _ = log.LoggerFromConfigAsString(`<seelog minlevel="off"></seelog>`)
	default:
		var cfgfile string
		var err error

		if start_flag == flag_fore_ground {
			cfgfile = fmt.Sprintf("./%s.console.xml", name)
		} else {
			cfgfile = fmt.Sprintf("./%s.file.xml", name)
		}

		if logger, err = log.LoggerFromConfigAsFile(cfgfile); err != nil {
			fmt.Println("Load SeeLog config failed:", err)
			if start_flag == flag_fore_ground {
				logger, _ = log.LoggerFromConfigAsString(def_console_log_cfg)
			} else {
				logger, _ = log.LoggerFromConfigAsString(strings.Replace(def_file_log_cfg, `%s`, name, -1))
			}
		}

		if err = log.ReplaceLogger(logger); err != nil {
			fmt.Println("Replace SeeLog failed:", err)
			os.Exit(200)
		}

		see_log_cfg_file_name = cfgfile
	}

	return
}

/* 注意: 一般root用户才有权限调用Setrlimit */
func initRlimit() {
	var rlim syscall.Rlimit
	rlim.Cur = 50000
	rlim.Max = 50000
	err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim)
	if err != nil {
		fmt.Println("set rlimit error: " + err.Error())
		//os.Exit(1)
	}
}
