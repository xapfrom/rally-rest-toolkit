/**
* Copyright 2014 Comcast Cable Communications Management, LLC
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
* http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
 */

package models

type Reference struct {
	Count         int    `json:",omitempty"`
	Ref           string `json:"_ref,omitempty"`
	Type          string `json:"_type,omitempty"`
	RefObjectName string `json:"_refObjectName,omitempty"`
	RefObjectUUID string `json:"_refObjectUUID,omitempty"`
}

type Defect struct {
	Ref                 string     `json:"_ref,omitempty"`
	CreationDate        string     `json:",omitempty"`
	ObjectID            int        `json:",omitempty"`
	ObjectUUID          string     `json:",omitempty"`
	Subscription        *Reference `json:",omitempty"`
	Workspace           *Reference `json:",omitempty"`
	Changesets          *Reference `json:",omitempty"`
	Requirement         *Reference `json:",omitempty"`
	Description         string     `json:",omitempty"`
	FormattedID         string     `json:",omitempty"`
	Name                string     `json:",omitempty"`
	Notes               string     `json:",omitempty"`
	Owner               *Reference `json:",omitempty"`
	Project             *Reference `json:",omitempty"`
	LastBuild           string     `json:",omitempty"`
	LastRun             string     `json:",omitempty"`
	ScheduleState       string     `json:",omitempty"`
	ScheduleStatePrefix string     `json:",omitempty"`
	Iteration           *Reference `json:",omitempty"`
	State               string     `json:",omitempty"`
	Priority            string     `json:",omitempty"`
	Severity            string     `json:",omitempty"`
	Tasks               *Reference `json:",omitempty"`
	Resolution          string     `json:",omitempty"`
}

type Feature struct {
	Ref         string      `json:"_ref,omitempty"`
	ObjectID    int         `json:",omitempty"`
	Description string      `json:",omitempty"`
	FormattedID string      `json:",omitempty"`
	State       interface{} `json:",omitempty"`
}

type HierarchicalRequirement struct {
	Ref                 string     `json:"_ref,omitempty"`
	Project             *Reference `json:",omitempty"`
	CreationDate        string     `json:",omitempty"`
	ObjectID            int        `json:",omitempty"`
	ObjectUUID          string     `json:",omitempty"`
	Subscription        *Reference `json:",omitempty"`
	Workspace           *Reference `json:",omitempty"`
	Changesets          *Reference `json:",omitempty"`
	Description         string     `json:",omitempty"`
	FormattedID         string     `json:",omitempty"`
	Name                string     `json:",omitempty"`
	LastBuild           string     `json:",omitempty"`
	LastRun             string     `json:",omitempty"`
	ScheduleState       string     `json:",omitempty"`
	ScheduleStatePrefix string     `json:",omitempty"`
	AcceptedDate        string     `json:",omitempty"`
	InProgressDate      string     `json:",omitempty"`
	Tasks               *Reference `json:",omitempty"`
}

type Task struct {
	Ref             string     `json:"_ref,omitempty"`
	CreationDate    string     `json:",omitempty"`
	ObjectID        int        `json:",omitempty"`
	ObjectUUID      string     `json:",omitempty"`
	Subscription    *Reference `json:",omitempty"`
	Workspace       *Reference `json:",omitempty"`
	Changesets      *Reference `json:",omitempty"`
	FormattedID     string     `json:",omitempty"`
	Name            string     `json:",omitempty"`
	Description     string     `json:",omitempty"`
	Actuals         float32    `json:",omitempty"`
	Attachments     *Reference `json:",omitempty"`
	Blocked         bool       `json:",omitempty"`
	BlockedReason   string     `json:",omitempty"`
	DragAndDropRank string     `json:",omitempty"`
	Estimate        float32    `json:",omitempty"`
	Iteration       *Reference `json:",omitempty"`
	Project         *Reference `json:",omitempty"`
	Recycled        bool       `json:",omitempty"`
	Release         *Reference `json:",omitempty"`
	State           string     `json:",omitempty"`
	TaskIndex       int64      `json:",omitempty"`
	TimeSpent       float32    `json:",omitempty"`
	ToDo            float32    `json:",omitempty"`
	WorkProduct     *Reference `json:",omitempty"`
}

type BuildDefinition struct {
	Ref          string     `json:"_ref,omitempty"`
	CreationDate string     `json:",omitempty"`
	ObjectID     int        `json:",omitempty"`
	ObjectUUID   string     `json:",omitempty"`
	Subscription *Reference `json:",omitempty"`
	Workspace    *Reference `json:",omitempty"`
	Builds       *Reference `json:",omitempty"`
	Projects     *Reference `json:",omitempty"`
	Project      *Reference `json:",omitempty"`
	LastBuild    *Reference `json:",omitempty"`
	Name         string     `json:",omitempty"`
	Description  string     `json:",omitempty"`
	LastStatus   string     `json:",omitempty"`
	Uri          string     `json:",omitempty"`
}

type Build struct {
	Ref             string       `json:"_ref,omitempty"`
	CreationDate    string       `json:",omitempty"`
	ObjectID        int          `json:",omitempty"`
	ObjectUUID      string       `json:",omitempty"`
	Subscription    *Reference   `json:",omitempty"`
	Workspace       *Reference   `json:",omitempty"`
	BuildDefinition *Reference   `json:",omitempty"`
	Changesets      []*Reference `json:",omitempty"`
	Number          string       `json:",omitempty"`
	Duration        float32      `json:",omitempty"`
	Start           string       `json:",omitempty"`
	Message         string       `json:",omitempty"`
	Status          string       `json:",omitempty"`
	Uri             string       `json:",omitempty"`
}

type Changeset struct {
	Ref             string     `json:"_ref,omitempty"`
	CreationDate    string     `json:",omitempty"`
	ObjectID        int        `json:",omitempty"`
	ObjectUUID      string     `json:",omitempty"`
	Subscription    *Reference `json:",omitempty"`
	Workspace       *Reference `json:",omitempty"`
	Artifacts       *Reference `json:",omitempty"`
	Author          *Reference `json:",omitempty"`
	Branch          string     `json:",omitempty"`
	Builds          *Reference `json:",omitempty"`
	Changes         *Reference `json:",omitempty"`
	CommitTimestamp string     `json:",omitempty"`
	Message         string     `json:",omitempty"`
	Name            string     `json:",omitempty"`
	Revision        string     `json:",omitempty"`
	SCMRepository   *Reference `json:",omitempty"`
	Uri             string     `json:",omitempty"`
}
