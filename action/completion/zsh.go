// Copyright (c) 2022 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package completion

// ZshAutoComplete represents the script used
// to enable automatic completion for the
// Zsh (https://ohmyz.sh/) Unix shell.
//
// nolint: lll // ignore long line length
const ZshAutoComplete = `#comdef vela
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
  compdef _cli_zsh_autocomplete vela
`
