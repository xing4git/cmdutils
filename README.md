lckage cmdutils
    import "github.com/xing4git/cmdutils"

FUNCTIONS

func BashExecute(cmdstr string) (ret string, err error)  
    Execute some commands with bash. The `cmdstr` param could be a single
    command such as `ls ~`, or a pipeline command such as `ps aux | grep
    'dropbox' | grep -v 'grep'`, a complete script file content is also can
    be work. The `cmdstr` will be written to a temp file, and then execute
    that file with bash. The output of your command will be returned by this
    function.

func Run(cmdname string, params ...string)  
    This function is a wrapper for cmd.Run() Pipe stdout, stderr, stdin of
    os to stdout, stderr, stdin of cmd
