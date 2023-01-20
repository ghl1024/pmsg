// Copyright 2022-2023 The pmsg Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/spf13/cobra"

	"github.com/lenye/pmsg/pkg/flags"
)

// feiShuCmd 飞书
var feiShuCmd = &cobra.Command{
	Use:     "feishu",
	Aliases: []string{"fs"},
	Short:   "fei shu",
}

func init() {
	feiShuCmd.PersistentFlags().StringVarP(&userAgent, flags.UserAgent, "a", "", "http user agent")

	feiShuCmd.AddCommand(feiShuBotCmd)
}
