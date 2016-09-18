// Copyright (c) 2015, Stefano Brilli
// All rights reserved.
// This source code is released under the the 3 Clause BSD License.
// Read LICENSE at https://github.com/cybercase/gospace/blob/master/LICENSE

// Generates a script that set and updates environment variables recommended by
// https://golang.org/doc/code.html. See README file for example of use.

package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	destFile, err := os.Create(path.Join(workingDir, "activate"))
	if err != nil {
		panic(err)
	}
	defer destFile.Close() // TODO: consider deletion of this line

	_, err = fmt.Fprintf(destFile, template)
	if err != nil {
		panic(err)
	}

	fmt.Println(`You're ready to go.

- Type 'source activate' jump into your new workspace
- Type 'deactivate' to step out`)
}

// Original: https://github.com/pypa/virtualenv/tree/develop/virtualenv_embedded
const template = `# This file must be used with "source bin/activate" *from bash*
# you cannot run it directly

deactivate () {
    # reset old environment variables
    if [ -n "$_OLD_WORKSPACE_PATH" ] ; then
        PATH="$_OLD_WORKSPACE_PATH"
        export PATH
        unset _OLD_WORKSPACE_PATH
    fi

    if [ -n "$_OLD_GOPATH" ] ; then
        GOPATH="$_OLD_GOPATH"
        export GOPATH
        unset _OLD_GOPATH
    fi

    # This should detect bash and zsh, which have a hash command that must
    # be called to get it to forget past commands.  Without forgetting
    # past commands the $PATH changes we made may not be respected
    if [ -n "$BASH" -o -n "$ZSH_VERSION" ] ; then
        hash -r 2>/dev/null
    fi

    if [ -n "$_OLD_WORKSPACE_PS1" ] ; then
        PS1="$_OLD_WORKSPACE_PS1"
        export PS1
        unset _OLD_WORKSPACE_PS1
    fi

    unset WORKSPACE
    if [ ! "$1" = "nondestructive" ] ; then
    # Self destruct!
        unset -f deactivate
    fi
}

# unset irrelevant variables
deactivate nondestructive

WORKSPACE="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
export WORKSPACE

_OLD_GOPATH="$GOPATH"
GOPATH=$WORKSPACE
export GOPATH

_OLD_WORKSPACE_PATH="$PATH"
PATH="$WORKSPACE/bin:$PATH"
export PATH

if [ -z "$WORKSPACE_DISABLE_PROMPT" ] ; then
    _OLD_WORKSPACE_PS1="$PS1"
    if [ "x" != x ] ; then
        PS1="$PS1"
    else
        PS1="(` + "`" + `basename \"$WORKSPACE\"` + "`" + `)$PS1"
    fi
    export PS1
fi

# This should detect bash and zsh, which have a hash command that must
# be called to get it to forget past commands.  Without forgetting
# past commands the $PATH changes we made may not be respected
if [ -n "$BASH" -o -n "$ZSH_VERSION" ] ; then
    hash -r 2>/dev/null
fi
`
