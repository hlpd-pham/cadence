// Copyright (c) 2021 Uber Technologies, Inc.
// Portions of the Software are attributed to Copyright (c) 2020 Temporal Technologies Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package dynamodb

import (
	"context"

	"github.com/uber/cadence/common/persistence/nosql/nosqlplugin"
)

func (db *ddb) InsertWorkflowExecutionWithTasks(
	ctx context.Context,
	currentWorkflowRequest *nosqlplugin.CurrentWorkflowWriteRequest,
	execution *nosqlplugin.WorkflowExecutionRequest,
	transferTasks []*nosqlplugin.TransferTask,
	crossClusterTasks []*nosqlplugin.CrossClusterTask,
	replicationTasks []*nosqlplugin.ReplicationTask,
	timerTasks []*nosqlplugin.TimerTask,
	shardCondition *nosqlplugin.ShardCondition,
) error {
	panic("TODO")
}

func (db *ddb) UpdateWorkflowExecutionWithTasks(
	ctx context.Context,
	currentWorkflowRequest *nosqlplugin.CurrentWorkflowWriteRequest,
	mutatedExecution *nosqlplugin.WorkflowExecutionRequest,
	insertedExecution *nosqlplugin.WorkflowExecutionRequest,
	resetExecution *nosqlplugin.WorkflowExecutionRequest,
	transferTasks []*nosqlplugin.TransferTask,
	crossClusterTasks []*nosqlplugin.CrossClusterTask,
	replicationTasks []*nosqlplugin.ReplicationTask,
	timerTasks []*nosqlplugin.TimerTask,
	shardCondition *nosqlplugin.ShardCondition,
) error {
	panic("TODO")
}

func (db *ddb) SelectCurrentWorkflow(ctx context.Context, shardID int, domainID, workflowID string) (*nosqlplugin.CurrentWorkflowRow, error) {
	panic("TODO")
}