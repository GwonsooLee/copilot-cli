// Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"github.com/aws/amazon-ecs-cli-v2/cmd/archer/template"
	"github.com/aws/amazon-ecs-cli-v2/internal/pkg/cli/group"
	"github.com/spf13/cobra"
)

// BuildAppCmd is the top level command for applications.
func BuildAppCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "app",
		Short: "Application commands",
		Long: `Command for working with applications.
An application represents an Amazon ECS service or task.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			persistProjectName()
		},
	}

	cmd.AddCommand(BuildAppInitCmd())
	cmd.SetUsageTemplate(template.Usage)
	cmd.Annotations = map[string]string{
		"group": group.Develop,
	}
	return cmd
}