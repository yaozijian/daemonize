package daemonize

import log "github.com/cihub/seelog"

const (
	def_console_log_cfg = `<seelog minlevel="info">
		<outputs formatid="detail">
			<console/>
		</outputs>
		<formats>
			<format id="common" format="%Msg%n" />
			<format id="detail" format="[%File:%Line][%Date(2006-01-02 15:04:05)] %Msg%n" />
		</formats>
	</seelog>`

	def_file_log_cfg = `<seelog minlevel="info">
		<outputs formatid="detail">
			<rollingfile filename="%s.log" type="size" maxsize="1024768" maxrolls="10"/>
		</outputs>
		<formats>
			<format id="common" format="%Msg%n" />
			<format id="detail" format="[%File:%Line][%Date(2006-01-02 15:04:05)] %Msg%n" />
		</formats>
	</seelog>`
)

var (
	see_log_cfg_file_name string
)

func ReloadSeeLogConfig() (err error) {
	var logger log.LoggerInterface
	if logger, err = log.LoggerFromConfigAsFile(see_log_cfg_file_name); err == nil {
		log.ReplaceLogger(logger)
		log.Criticalf("\n\n已经从文件%s中重新加载日志配置\n\n", see_log_cfg_file_name)
	}
	return
}
