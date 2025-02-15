// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: BUSL-1.1

// Actions are executable commands that can be run on an allocation within
// the context of a task. They are left open-ended enough to be applied to
// other Nomad concepts like Nodes in the future.

package structs

import "slices"

type Action struct {
	Name    string
	Command string
	Args    []string
}

type JobAction struct {
	Action
	TaskName      string
	TaskGroupName string
}

type ActionListResponse struct {
	Actions []*JobAction
	QueryMeta
}

func (a *Action) Copy() *Action {
	if a == nil {
		return nil
	}
	na := new(Action)
	*na = *a
	na.Args = slices.Clone(a.Args)
	return na
}

func (a *Action) Equal(o *Action) bool {
	if a == o {
		return true
	}
	if a == nil || o == nil {
		return false
	}
	return a.Name == o.Name &&
		a.Command == o.Command &&
		slices.Equal(a.Args, o.Args)
}
