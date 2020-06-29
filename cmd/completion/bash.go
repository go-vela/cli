package completion

import (
	"bytes"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var BashCmd = cli.Command{
	Name:        "bash",
	Usage:       "Use this command to enable auto completion in bash",
	Description: "Bash shell auto completion for vela",
	Action:      executeBash,
	CustomHelpTemplate: fmt.Sprintf(`%s
1.  To enable auto completion for current bash session. Make sure bash version is 4+:
    source <(vela completion bash)
2.  To permanently enable bash auto completion for vela, visit https://go-vela.github.io/docs/cli/autocompletion/bash/
`, cli.CommandHelpTemplate),
}

// To generate bash completion script which can be sourced to enable vela auto completion in bash shell
func executeBash(_ *cli.Context) error {
	buf := new(bytes.Buffer)

	// urfave cli bash auto completion script tailor made for vela
	buf.WriteString(`#! /bin/bash
		_cli_bash_autocomplete() {
  			if [[ "${COMP_WORDS[0]}" != "source" ]]; then
    			local cur opts base
    			COMPREPLY=()
    			cur="${COMP_WORDS[COMP_CWORD]}"
    			
				if [[ "$cur" == "-"* ]]; then
      				opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} ${cur} --generate-bash-completion )
    			else
      				opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
    			fi
    			
				COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
    			return 0
  			fi
		}
		complete -o bashdefault -o default -o nospace -F _cli_bash_autocomplete vela`,
	)

	_, err := os.Stdout.Write(buf.Bytes())

	if err != nil {
		return err
	}

	return nil
}
