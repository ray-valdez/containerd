//go:build linux
// +build linux

/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package containerdShim

import (
	"context"
	"fmt"
	"os"

	"github.com/containerd/containerd/log"
	"github.com/containerd/containerd/pkg/shutdown"
	"github.com/sirupsen/logrus"

	"github.com/containerd/containerd/errdefs"
	"github.com/containerd/containerd/runtime/v2/runc/manager"
	"github.com/containerd/containerd/runtime/v2/runc/task"
	"github.com/containerd/containerd/runtime/v2/shim"
	shimapi "github.com/containerd/containerd/runtime/v2/task"
	taskAPI "github.com/containerd/containerd/runtime/v2/task"
	ptypes "github.com/gogo/protobuf/types"
)

var (
	// check to make sure the *service implements the GRPC API
	_ = (taskAPI.TaskService)(&service{})
)

type shimTaskManager struct {
	shimapi.TaskService
	id      string
	manager shim.Manager
}

// shimLog is logger for shim package
var shimLog = logrus.WithFields(logrus.Fields{
	"source": "containerd-split-shim-v2",
	"name":   "containerd-shim-v2",
})

func (stm *shimTaskManager) Cleanup(ctx context.Context) (*shimapi.DeleteResponse, error) {
	fmt.Println("Cleanup: RV")
	ss, err := stm.manager.Stop(ctx, stm.id)
	if err != nil {
		return nil, err
	}
	return &shimapi.DeleteResponse{
		Pid:        uint32(ss.Pid),
		ExitStatus: uint32(ss.ExitStatus),
		ExitedAt:   ss.ExitedAt,
	}, nil
}

func (stm *shimTaskManager) StartShim(ctx context.Context, opts shim.StartOpts) (string, error) {
	return stm.manager.Start(ctx, opts.ID, opts)
}

// New returns a new shim service
func New(ctx context.Context, id string, publisher shim.Publisher, fn func()) (shim.Shim, error) {
	sd, ok := ctx.(shutdown.Service)
	if !ok {
		ctx, sd = shutdown.WithShutdown(ctx)
		sd.RegisterCallback(func(context.Context) error {
			fn()
			return nil
		})
	}
	ts, err := task.NewTaskService(ctx, publisher, sd)
	if err != nil {
		return nil, err
	}

	return &shimTaskManager{
		TaskService: ts,
		id:          id,
		manager:     manager.NewShimManager("split"),
	}, nil
	return &service{}, nil
}

type service struct {
}

// StartShim is a binary call that executes a new shim returning the address
func (s *service) StartShim(ctx context.Context, opts shim.StartOpts) (string, error) {
	//logrus.Infof("Starting shim with options: %+v", opts)
	return "", nil
}

// Cleanup is a binary call that cleans up any resources used by the shim when the service crashes
func (s *service) Cleanup(ctx context.Context) (*taskAPI.DeleteResponse, error) {
	fmt.Println("Cleanup : RV")
	return nil, errdefs.ErrNotImplemented
}

// Create a new container
func (s *service) Create(ctx context.Context, r *taskAPI.CreateTaskRequest) (_ *taskAPI.CreateTaskResponse, err error) {
	fmt.Println("Create : RV")
	return nil, errdefs.ErrNotImplemented
}

// Start the primary user process inside the container
func (s *service) Start(ctx context.Context, r *taskAPI.StartRequest) (*taskAPI.StartResponse, error) {
	fmt.Println("Start : RV")
	log.G(ctx).Infof("Start: RVV")
	return nil, errdefs.ErrNotImplemented
}

// Delete a process or container
func (s *service) Delete(ctx context.Context, r *taskAPI.DeleteRequest) (*taskAPI.DeleteResponse, error) {
	log.G(ctx).Infof("Delete: RVV")
	fmt.Println("Delete : RV")
	return nil, errdefs.ErrNotImplemented
}

// Exec an additional process inside the container
func (s *service) Exec(ctx context.Context, r *taskAPI.ExecProcessRequest) (*ptypes.Empty, error) {
	fmt.Println("Exec : RV")
	return nil, errdefs.ErrNotImplemented
}

// ResizePty of a process
func (s *service) ResizePty(ctx context.Context, r *taskAPI.ResizePtyRequest) (*ptypes.Empty, error) {
	fmt.Println("Resize : RV")
	return nil, errdefs.ErrNotImplemented
}

// State returns runtime state of a process
func (s *service) State(ctx context.Context, r *taskAPI.StateRequest) (*taskAPI.StateResponse, error) {
	fmt.Println("State : RV")
	return nil, errdefs.ErrNotImplemented
}

// Pause the container
func (s *service) Pause(ctx context.Context, r *taskAPI.PauseRequest) (*ptypes.Empty, error) {
	fmt.Println("Pause : RV")
	return nil, errdefs.ErrNotImplemented
}

// Resume the container
func (s *service) Resume(ctx context.Context, r *taskAPI.ResumeRequest) (*ptypes.Empty, error) {
	fmt.Println("Resume : RV")
	return nil, errdefs.ErrNotImplemented
}

// Kill a process
func (s *service) Kill(ctx context.Context, r *taskAPI.KillRequest) (*ptypes.Empty, error) {
	log.G(ctx).Infof("Kill: RVV")
	fmt.Println("Kill: RV")
	return nil, errdefs.ErrNotImplemented
}

// Pids returns all pids inside the container
func (s *service) Pids(ctx context.Context, r *taskAPI.PidsRequest) (*taskAPI.PidsResponse, error) {
	fmt.Println("Pids : RV")
	return nil, errdefs.ErrNotImplemented
}

// CloseIO of a process
func (s *service) CloseIO(ctx context.Context, r *taskAPI.CloseIORequest) (*ptypes.Empty, error) {
	fmt.Println("CloseIO : RV")
	return nil, errdefs.ErrNotImplemented
}

// Checkpoint the container
func (s *service) Checkpoint(ctx context.Context, r *taskAPI.CheckpointTaskRequest) (*ptypes.Empty, error) {
	fmt.Println("Checkpoint : RV")
	return nil, errdefs.ErrNotImplemented
}

// Connect returns shim information of the underlying service
func (s *service) Connect(ctx context.Context, r *taskAPI.ConnectRequest) (*taskAPI.ConnectResponse, error) {
	fmt.Println("Connect : RV")
	return nil, errdefs.ErrNotImplemented
}

// Shutdown is called after the underlying resources of the shim are cleaned up and the service can be stopped
func (s *service) Shutdown(ctx context.Context, r *taskAPI.ShutdownRequest) (*ptypes.Empty, error) {
	fmt.Println("Shutdown : RV")
	os.Exit(0)
	return &ptypes.Empty{}, nil
}

// Stats returns container level system stats for a container and its processes
func (s *service) Stats(ctx context.Context, r *taskAPI.StatsRequest) (*taskAPI.StatsResponse, error) {
	fmt.Println("Stats : RV")
	return nil, errdefs.ErrNotImplemented
}

// Update the live container
func (s *service) Update(ctx context.Context, r *taskAPI.UpdateTaskRequest) (*ptypes.Empty, error) {
	fmt.Println("Update : RV")
	return nil, errdefs.ErrNotImplemented
}

// Wait for a process to exit
func (s *service) Wait(ctx context.Context, r *taskAPI.WaitRequest) (*taskAPI.WaitResponse, error) {
	fmt.Println("Wait : RV")
	return nil, errdefs.ErrNotImplemented
}
