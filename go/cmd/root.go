/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 * 
 *   http://www.apache.org/licenses/LICENSE-2.0
 * 
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package aoc2017

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var RootCmd *cobra.Command

func init() {
	RootCmd = &cobra.Command{
		Use:   "aoc2017",
		Short: "Applications solving Advent of Code 2017 problems in Go-Lang",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Usage()
		},
		Version: "1.0.0",
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
