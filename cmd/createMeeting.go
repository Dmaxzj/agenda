// Copyright © 2018 dengzijie
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var createMeetingArgs struct {
	meetingTitle  string
	startTime     string
	endTime       string
	participators []string
}

// createMeetingCmd represents the createMeeting command
var createMeetingCmd = &cobra.Command{
	Use:   "createMeeting",
	Short: "creat a meeting",
	Long:  "creat a meeting as sponsor",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("createMeeting called")
		fmt.Println("meeting title: ", createMeetingArgs.meetingTitle)
		fmt.Println("start time: ", createMeetingArgs.startTime)
		fmt.Println("end time: ", createMeetingArgs.endTime)
		fmt.Printf("participators: ")
		for num, name := range createMeetingArgs.participators {
			fmt.Printf("%d:%s ", num, name)
		}
	},
}

func init() {
	rootCmd.AddCommand(createMeetingCmd)
	createMeetingCmd.Flags().StringVarP(&(createMeetingArgs.meetingTitle), "title", "t", "", "meeting title")
	createMeetingCmd.Flags().StringVarP(&(createMeetingArgs.startTime), "stime", "s", "", "the time when the meeting begin(mm-dd-hh-mm)")
	createMeetingCmd.Flags().StringVarP(&(createMeetingArgs.endTime), "etime", "e", "", "the end time of the meeting(mm-dd-hh-mm)")
	createMeetingCmd.Flags().StringArrayVarP(&(createMeetingArgs.participators), "particpators", "p", nil, "the participators of the meeting")
}
