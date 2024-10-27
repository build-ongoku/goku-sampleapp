package mthd

import (
	"context"
	"fmt"

	app_client "sampleapp/backend/.goku/generated/client"
	svccore_enttaskrun_dal "sampleapp/backend/.goku/generated/service/core/entity/task_run/dal"
	svccore_enttaskrun_typ "sampleapp/backend/.goku/generated/service/core/entity/task_run/typ"
	svccore_enttaskrun_typhelper "sampleapp/backend/.goku/generated/service/core/entity/task_run/typhelper"
	app_typhelper "sampleapp/backend/.goku/generated/typhelper"

	"github.com/teejays/gokutil/log"
	"github.com/teejays/gokutil/scalars"
)

type T = svccore_enttaskrun_typ.TaskRun

func HookCreatePre(ctx context.Context, c app_client.Client, in T) (T, error) {

	// Default the status to pending
	in.Status = svccore_enttaskrun_typ.Status_Pending

	return in, nil
}

func HookSavePre(ctx context.Context, c app_client.Client, in T) (T, error) {

	return in, nil
}

func ActionRun(ctx context.Context, c app_client.Client, req svccore_enttaskrun_typ.StandardEntityRequest) (svccore_enttaskrun_typ.StandardEntityResponse, error) {
	var ret svccore_enttaskrun_typ.StandardEntityResponse
	var retErr error

	// Change the status of the task run to running
	dalClient, err := svccore_enttaskrun_dal.NewClient(ctx)
	if err != nil {
		return ret, err
	}
	updateReq := svccore_enttaskrun_typ.UpdateRequest{
		Object: svccore_enttaskrun_typ.TaskRun{
			ID:     req.ID,
			Status: svccore_enttaskrun_typ.Status_Running,
		},
		Fields: []svccore_enttaskrun_typ.TaskRunField{
			svccore_enttaskrun_typ.TaskRunField_Status,
		},
	}
	updateResp, err := dalClient.Update(ctx, updateReq)
	if err != nil {
		return ret, fmt.Errorf("Updating task run status to [%s]: %w", svccore_enttaskrun_typ.Status_Running, err)
	}
	taskRun := updateResp.Object

	// Get the task run
	// getReq := svccore_enttaskrun_typ.GetRequest{}
	// getReq.ID = req.ID
	// taskRun, err := c.Core().TaskRun().Get(ctx, getReq)
	// if err != nil {
	// 	return ret, err
	// }

	// Get the task
	task, err := svccore_enttaskrun_typhelper.GetTask(ctx, c.Core().Task(), taskRun)
	if err != nil {
		retErr = fmt.Errorf("Getting task: %w", err)
		return ret, retErr
	}

	// Get the request data
	reqType, err := app_typhelper.GetNewRequestTypeForMethod(ctx, task.Method)
	if err != nil {
		retErr = fmt.Errorf("Getting request type for method: %w", err)
		return ret, retErr
	}

	// If task has a request JSON, unmarshal it into the request type
	// otherwise the request will stay empty
	if !task.MethodRequestData.IsEmpty() {
		err = task.MethodRequestData.LoadTo(&reqType)
		if err != nil {
			retErr = fmt.Errorf("Unmarshalling method request JSON: %w", err)
			return ret, retErr
		}
	}

	// Upon failure, update the task run status to failed
	defer func() {
		if retErr == nil {
			return
		}

		// If there was an error, update the task run status to failed
		updateReq := svccore_enttaskrun_typ.UpdateRequest{
			Object: svccore_enttaskrun_typ.TaskRun{
				ID:     req.ID,
				Status: svccore_enttaskrun_typ.Status_Failed,
			},
			Fields: []svccore_enttaskrun_typ.TaskRunField{
				svccore_enttaskrun_typ.TaskRunField_Status,
			},
		}
		_, err := dalClient.Update(ctx, updateReq)
		if err != nil {
			log.Error(ctx, "Task Run failed. And we failed to update task run status to failed", "originalError", retErr, "updateError", err)
		}

	}()

	mthdResp, err := app_typhelper.InvokeMethod(ctx, c, task.Method, reqType)
	if err != nil {
		retErr = fmt.Errorf("Invoking method: %w", err)
		return ret, retErr
	}

	// Update the task run
	taskRun.MethodResponseData, err = scalars.NewGenericDataFromStruct(mthdResp)
	if err != nil {
		retErr = fmt.Errorf("Marshalling method response: %w", err)
		return ret, retErr
	}

	taskRun.Status = svccore_enttaskrun_typ.Status_Success

	updateReq = svccore_enttaskrun_typ.UpdateRequest{
		Object: taskRun,
		Fields: []svccore_enttaskrun_typ.TaskRunField{
			svccore_enttaskrun_typ.TaskRunField_Status,
			svccore_enttaskrun_typ.TaskRunField_MethodResponseData,
		},
	}
	_, err = dalClient.Update(ctx, updateReq)
	if err != nil {
		return ret, fmt.Errorf("Method Completed succesfully. However, there was an error updating task run status to [%s]: %w", svccore_enttaskrun_typ.Status_Success, err)
	}

	return ret, nil

}
