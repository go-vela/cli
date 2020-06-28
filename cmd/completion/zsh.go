package completion

import (
	"bytes"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var ZSHCmd = cli.Command{
	Name:        "zsh",
	Usage:       "Use this command to enable auto completion in zsh",
	Description: "zsh shell auto completion for vela",
	Action:      executeZSH,
	CustomHelpTemplate: fmt.Sprintf(`%s
1.  To enable auto completion for current zsh session:
    source <(vela completion zsh)
2.  To permanently enable zsh auto completion for vela, visit https://github.com/go-vela/docs
`, cli.CommandHelpTemplate),
}

func executeZSH(_ *cli.Context) error {
	buf := new(bytes.Buffer)

	// urfave cli zsh auto completion script tailor made for vela
	buf.WriteString(`#comdef vela
		_cli_zsh_autocomplete() {
  			local -a opts
  			local cur
  			cur=${words[-1]}
  			
			if [[ "$cur" == "-"* ]]; then
    			opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} ${cur} --generate-bash-completion)}")
			else
    			opts=("${(@f)$(_CLI_ZSH_AUTOCOMPLETE_HACK=1 ${words[@]:0:#words[@]-1} --generate-bash-completion)}")
  			fi
  			
			if [[ "${opts[1]}" != "" ]]; then
   		  		_describe 'values' opts
  			fi

  			return
		}
		export _CLI_ZSH_AUTOCOMPLETE_HACK=1
		compdef _cli_zsh_autocomplete vela`,
	)

	_, err := os.Stdout.Write(buf.Bytes())
	if err != nil {
		return err
	}

	return nil
}
