// Copyright (c) 2021 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package completion

// BashAutoComplete represents the script used
// to enable automatic completion for the
// Bash (https://www.gnu.org/software/bash/) Unix shell.
const BashAutoComplete = `#! /bin/bash
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

  complete -o bashdefault -o default -o nospace -F _cli_bash_autocomplete vela
`
