package commands

import (
	"github.com/urfave/cli"
)

var REST_Commands = []cli.Command{
	{
		Name:   "SET",
		Usage:  "set a key value",
		Action: getAction,
	},

	{
		Name:   "GET",
		Usage:  "get a key value",
		Action: postAction,
	},
	{
		Name:   "QPUSH",
		Usage:  "create a queue",
		Action: postAction,
	},
	{
		Name:   "QPOP",
		Usage:  "pop a key value",
		Action: postAction,
	},
}
