package app_typhelper

import (
	"context"
	"fmt"
	app_client "sampleapp/backend/.goku/generated/client"
	svcauth_entsecret_typ "sampleapp/backend/.goku/generated/service/auth/entity/secret/typ"
	svcauth_typ "sampleapp/backend/.goku/generated/service/auth/typ"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
	svccore_entjobapplicant_typ "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/typ"
	svccore_entsecretdecryptable_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_decryptable/typ"
	svccore_entsecretkey_typ "sampleapp/backend/.goku/generated/service/core/entity/secret_key/typ"
	svccore_enttask_typ "sampleapp/backend/.goku/generated/service/core/entity/task/typ"
	svccore_enttaskrun_typ "sampleapp/backend/.goku/generated/service/core/entity/task_run/typ"
	svccore_typ "sampleapp/backend/.goku/generated/service/core/typ"
	svcuser_entteam_typ "sampleapp/backend/.goku/generated/service/user/entity/team/typ"
	svcuser_entuser_typ "sampleapp/backend/.goku/generated/service/user/entity/user/typ"
	app_typ "sampleapp/backend/.goku/generated/typ"
)

func InvokeMethod(ctx context.Context, c app_client.Client, methodName app_typ.MethodName, _req interface{}) (interface{}, error) {
	switch methodName {
	case app_typ.MethodName_Core_SecretDecryptable_ActionDecrypt:
		req, ok := _req.(svccore_entsecretdecryptable_typ.StandardEntityRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [StandardEntityRequest]")
		}
		return c.Core().SecretDecryptable().ActionDecrypt(ctx, req)
	case app_typ.MethodName_Core_Task_ActionRun:
		req, ok := _req.(svccore_enttask_typ.StandardEntityRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [StandardEntityRequest]")
		}
		return c.Core().Task().ActionRun(ctx, req)
	case app_typ.MethodName_Core_TaskRun_ActionRun:
		req, ok := _req.(svccore_enttaskrun_typ.StandardEntityRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [StandardEntityRequest]")
		}
		return c.Core().TaskRun().ActionRun(ctx, req)
	case app_typ.MethodName_Core_JobApplicant_Add:
		req, ok := _req.(svccore_entjobapplicant_typ.AddRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [AddRequest]")
		}
		return c.Core().JobApplicant().Add(ctx, req)
	case app_typ.MethodName_User_Team_Add:
		req, ok := _req.(svcuser_entteam_typ.AddRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [AddRequest]")
		}
		return c.User().Team().Add(ctx, req)
	case app_typ.MethodName_Auth_Secret_Add:
		req, ok := _req.(svcauth_entsecret_typ.AddRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [AddRequest]")
		}
		return c.Auth().Secret().Add(ctx, req)
	case app_typ.MethodName_Core_TaskRun_Add:
		req, ok := _req.(svccore_enttaskrun_typ.AddRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [AddRequest]")
		}
		return c.Core().TaskRun().Add(ctx, req)
	case app_typ.MethodName_Core_File_Add:
		req, ok := _req.(svccore_entfile_typ.AddRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [AddRequest]")
		}
		return c.Core().File().Add(ctx, req)
	case app_typ.MethodName_Core_SecretDecryptable_Add:
		req, ok := _req.(svccore_entsecretdecryptable_typ.AddRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [AddRequest]")
		}
		return c.Core().SecretDecryptable().Add(ctx, req)
	case app_typ.MethodName_Core_SecretKey_Add:
		req, ok := _req.(svccore_entsecretkey_typ.AddRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [AddRequest]")
		}
		return c.Core().SecretKey().Add(ctx, req)
	case app_typ.MethodName_User_User_Add:
		req, ok := _req.(svcuser_entuser_typ.AddRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [AddRequest]")
		}
		return c.User().User().Add(ctx, req)
	case app_typ.MethodName_Core_Task_Add:
		req, ok := _req.(svccore_enttask_typ.AddRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [AddRequest]")
		}
		return c.Core().Task().Add(ctx, req)
	case app_typ.MethodName_Auth_AuthenticateToken:
		req, ok := _req.(svcauth_typ.AuthenticateTokenRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [AuthenticateTokenRequest]")
		}
		return c.Auth().AuthenticateToken(ctx, req)
	case app_typ.MethodName_Core_SecretDecryptable_Get:
		req, ok := _req.(svccore_entsecretdecryptable_typ.GetRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [GetRequest]")
		}
		return c.Core().SecretDecryptable().Get(ctx, req)
	case app_typ.MethodName_Core_Task_Get:
		req, ok := _req.(svccore_enttask_typ.GetRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [GetRequest]")
		}
		return c.Core().Task().Get(ctx, req)
	case app_typ.MethodName_Core_TaskRun_Get:
		req, ok := _req.(svccore_enttaskrun_typ.GetRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [GetRequest]")
		}
		return c.Core().TaskRun().Get(ctx, req)
	case app_typ.MethodName_User_Team_Get:
		req, ok := _req.(svcuser_entteam_typ.GetRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [GetRequest]")
		}
		return c.User().Team().Get(ctx, req)
	case app_typ.MethodName_Core_JobApplicant_Get:
		req, ok := _req.(svccore_entjobapplicant_typ.GetRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [GetRequest]")
		}
		return c.Core().JobApplicant().Get(ctx, req)
	case app_typ.MethodName_Core_SecretKey_Get:
		req, ok := _req.(svccore_entsecretkey_typ.GetRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [GetRequest]")
		}
		return c.Core().SecretKey().Get(ctx, req)
	case app_typ.MethodName_User_User_Get:
		req, ok := _req.(svcuser_entuser_typ.GetRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [GetRequest]")
		}
		return c.User().User().Get(ctx, req)
	case app_typ.MethodName_Core_File_Get:
		req, ok := _req.(svccore_entfile_typ.GetRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [GetRequest]")
		}
		return c.Core().File().Get(ctx, req)
	case app_typ.MethodName_Auth_Secret_Get:
		req, ok := _req.(svcauth_entsecret_typ.GetRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [GetRequest]")
		}
		return c.Auth().Secret().Get(ctx, req)
	case app_typ.MethodName_Core_SecretDecryptable_HookCreatePre:
		req, ok := _req.(svccore_entsecretdecryptable_typ.SecretDecryptable)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [SecretDecryptable]")
		}
		return c.Core().SecretDecryptable().HookCreatePre(ctx, req)
	case app_typ.MethodName_Core_TaskRun_HookCreatePre:
		req, ok := _req.(svccore_enttaskrun_typ.TaskRun)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [TaskRun]")
		}
		return c.Core().TaskRun().HookCreatePre(ctx, req)
	case app_typ.MethodName_Core_File_HookCreatePre:
		req, ok := _req.(svccore_entfile_typ.File)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [File]")
		}
		return c.Core().File().HookCreatePre(ctx, req)
	case app_typ.MethodName_Core_SecretKey_HookCreatePre:
		req, ok := _req.(svccore_entsecretkey_typ.SecretKey)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [SecretKey]")
		}
		return c.Core().SecretKey().HookCreatePre(ctx, req)
	case app_typ.MethodName_Core_Task_HookSavePre:
		req, ok := _req.(svccore_enttask_typ.Task)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [Task]")
		}
		return c.Core().Task().HookSavePre(ctx, req)
	case app_typ.MethodName_User_User_List:
		req, ok := _req.(svcuser_entuser_typ.ListRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [ListRequest]")
		}
		return c.User().User().List(ctx, req)
	case app_typ.MethodName_Core_File_List:
		req, ok := _req.(svccore_entfile_typ.ListRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [ListRequest]")
		}
		return c.Core().File().List(ctx, req)
	case app_typ.MethodName_Core_SecretKey_List:
		req, ok := _req.(svccore_entsecretkey_typ.ListRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [ListRequest]")
		}
		return c.Core().SecretKey().List(ctx, req)
	case app_typ.MethodName_Core_TaskRun_List:
		req, ok := _req.(svccore_enttaskrun_typ.ListRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [ListRequest]")
		}
		return c.Core().TaskRun().List(ctx, req)
	case app_typ.MethodName_Core_JobApplicant_List:
		req, ok := _req.(svccore_entjobapplicant_typ.ListRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [ListRequest]")
		}
		return c.Core().JobApplicant().List(ctx, req)
	case app_typ.MethodName_Core_Task_List:
		req, ok := _req.(svccore_enttask_typ.ListRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [ListRequest]")
		}
		return c.Core().Task().List(ctx, req)
	case app_typ.MethodName_Core_SecretDecryptable_List:
		req, ok := _req.(svccore_entsecretdecryptable_typ.ListRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [ListRequest]")
		}
		return c.Core().SecretDecryptable().List(ctx, req)
	case app_typ.MethodName_User_Team_List:
		req, ok := _req.(svcuser_entteam_typ.ListRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [ListRequest]")
		}
		return c.User().Team().List(ctx, req)
	case app_typ.MethodName_Auth_Secret_List:
		req, ok := _req.(svcauth_entsecret_typ.ListRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [ListRequest]")
		}
		return c.Auth().Secret().List(ctx, req)
	case app_typ.MethodName_Auth_LoginUser:
		req, ok := _req.(svcauth_typ.LoginRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [LoginRequest]")
		}
		return c.Auth().LoginUser(ctx, req)
	case app_typ.MethodName_Core_File_QueryByText:
		req, ok := _req.(svccore_entfile_typ.QueryByTextRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [QueryByTextRequest]")
		}
		return c.Core().File().QueryByText(ctx, req)
	case app_typ.MethodName_Core_TaskRun_QueryByText:
		req, ok := _req.(svccore_enttaskrun_typ.QueryByTextRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [QueryByTextRequest]")
		}
		return c.Core().TaskRun().QueryByText(ctx, req)
	case app_typ.MethodName_Core_SecretDecryptable_QueryByText:
		req, ok := _req.(svccore_entsecretdecryptable_typ.QueryByTextRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [QueryByTextRequest]")
		}
		return c.Core().SecretDecryptable().QueryByText(ctx, req)
	case app_typ.MethodName_Auth_Secret_QueryByText:
		req, ok := _req.(svcauth_entsecret_typ.QueryByTextRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [QueryByTextRequest]")
		}
		return c.Auth().Secret().QueryByText(ctx, req)
	case app_typ.MethodName_Core_Task_QueryByText:
		req, ok := _req.(svccore_enttask_typ.QueryByTextRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [QueryByTextRequest]")
		}
		return c.Core().Task().QueryByText(ctx, req)
	case app_typ.MethodName_Core_SecretKey_QueryByText:
		req, ok := _req.(svccore_entsecretkey_typ.QueryByTextRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [QueryByTextRequest]")
		}
		return c.Core().SecretKey().QueryByText(ctx, req)
	case app_typ.MethodName_User_User_QueryByText:
		req, ok := _req.(svcuser_entuser_typ.QueryByTextRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [QueryByTextRequest]")
		}
		return c.User().User().QueryByText(ctx, req)
	case app_typ.MethodName_User_Team_QueryByText:
		req, ok := _req.(svcuser_entteam_typ.QueryByTextRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [QueryByTextRequest]")
		}
		return c.User().Team().QueryByText(ctx, req)
	case app_typ.MethodName_Core_JobApplicant_QueryByText:
		req, ok := _req.(svccore_entjobapplicant_typ.QueryByTextRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [QueryByTextRequest]")
		}
		return c.Core().JobApplicant().QueryByText(ctx, req)
	case app_typ.MethodName_Auth_RegisterUser:
		req, ok := _req.(svcauth_typ.RegisterUserRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [RegisterUserRequest]")
		}
		return c.Auth().RegisterUser(ctx, req)
	case app_typ.MethodName_Core_SecretDecryptabeAdd:
		req, ok := _req.(svccore_typ.SecretDecryptableAddRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [SecretDecryptableAddRequest]")
		}
		return c.Core().SecretDecryptabeAdd(ctx, req)
	case app_typ.MethodName_Core_Task_Update:
		req, ok := _req.(svccore_enttask_typ.UpdateRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [UpdateRequest]")
		}
		return c.Core().Task().Update(ctx, req)
	case app_typ.MethodName_Core_File_Update:
		req, ok := _req.(svccore_entfile_typ.UpdateRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [UpdateRequest]")
		}
		return c.Core().File().Update(ctx, req)
	case app_typ.MethodName_Core_JobApplicant_Update:
		req, ok := _req.(svccore_entjobapplicant_typ.UpdateRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [UpdateRequest]")
		}
		return c.Core().JobApplicant().Update(ctx, req)
	case app_typ.MethodName_Core_SecretKey_Update:
		req, ok := _req.(svccore_entsecretkey_typ.UpdateRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [UpdateRequest]")
		}
		return c.Core().SecretKey().Update(ctx, req)
	case app_typ.MethodName_User_Team_Update:
		req, ok := _req.(svcuser_entteam_typ.UpdateRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [UpdateRequest]")
		}
		return c.User().Team().Update(ctx, req)
	case app_typ.MethodName_Core_TaskRun_Update:
		req, ok := _req.(svccore_enttaskrun_typ.UpdateRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [UpdateRequest]")
		}
		return c.Core().TaskRun().Update(ctx, req)
	case app_typ.MethodName_User_User_Update:
		req, ok := _req.(svcuser_entuser_typ.UpdateRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [UpdateRequest]")
		}
		return c.User().User().Update(ctx, req)
	case app_typ.MethodName_Core_SecretDecryptable_Update:
		req, ok := _req.(svccore_entsecretdecryptable_typ.UpdateRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [UpdateRequest]")
		}
		return c.Core().SecretDecryptable().Update(ctx, req)
	case app_typ.MethodName_Auth_Secret_Update:
		req, ok := _req.(svcauth_entsecret_typ.UpdateRequest)
		if !ok {
			return nil, fmt.Errorf("Invalid request type: could not cast the request to type [UpdateRequest]")
		}
		return c.Auth().Secret().Update(ctx, req)

	default:
		return nil, fmt.Errorf("Empty or invalid method name")
	}
}

func GetNewRequestTypeForMethod(ctx context.Context, methodName app_typ.MethodName) (interface{}, error) {
	switch methodName {
	case app_typ.MethodName_Core_SecretDecryptable_ActionDecrypt:
		return &svccore_entsecretdecryptable_typ.StandardEntityRequest{}, nil
	case app_typ.MethodName_Core_Task_ActionRun:
		return &svccore_enttask_typ.StandardEntityRequest{}, nil
	case app_typ.MethodName_Core_TaskRun_ActionRun:
		return &svccore_enttaskrun_typ.StandardEntityRequest{}, nil
	case app_typ.MethodName_Core_JobApplicant_Add:
		return &svccore_entjobapplicant_typ.AddRequest{}, nil
	case app_typ.MethodName_User_Team_Add:
		return &svcuser_entteam_typ.AddRequest{}, nil
	case app_typ.MethodName_Auth_Secret_Add:
		return &svcauth_entsecret_typ.AddRequest{}, nil
	case app_typ.MethodName_Core_TaskRun_Add:
		return &svccore_enttaskrun_typ.AddRequest{}, nil
	case app_typ.MethodName_Core_File_Add:
		return &svccore_entfile_typ.AddRequest{}, nil
	case app_typ.MethodName_Core_SecretDecryptable_Add:
		return &svccore_entsecretdecryptable_typ.AddRequest{}, nil
	case app_typ.MethodName_Core_SecretKey_Add:
		return &svccore_entsecretkey_typ.AddRequest{}, nil
	case app_typ.MethodName_User_User_Add:
		return &svcuser_entuser_typ.AddRequest{}, nil
	case app_typ.MethodName_Core_Task_Add:
		return &svccore_enttask_typ.AddRequest{}, nil
	case app_typ.MethodName_Auth_AuthenticateToken:
		return &svcauth_typ.AuthenticateTokenRequest{}, nil
	case app_typ.MethodName_Core_SecretDecryptable_Get:
		return &svccore_entsecretdecryptable_typ.GetRequest{}, nil
	case app_typ.MethodName_Core_Task_Get:
		return &svccore_enttask_typ.GetRequest{}, nil
	case app_typ.MethodName_Core_TaskRun_Get:
		return &svccore_enttaskrun_typ.GetRequest{}, nil
	case app_typ.MethodName_User_Team_Get:
		return &svcuser_entteam_typ.GetRequest{}, nil
	case app_typ.MethodName_Core_JobApplicant_Get:
		return &svccore_entjobapplicant_typ.GetRequest{}, nil
	case app_typ.MethodName_Core_SecretKey_Get:
		return &svccore_entsecretkey_typ.GetRequest{}, nil
	case app_typ.MethodName_User_User_Get:
		return &svcuser_entuser_typ.GetRequest{}, nil
	case app_typ.MethodName_Core_File_Get:
		return &svccore_entfile_typ.GetRequest{}, nil
	case app_typ.MethodName_Auth_Secret_Get:
		return &svcauth_entsecret_typ.GetRequest{}, nil
	case app_typ.MethodName_Core_SecretDecryptable_HookCreatePre:
		return &svccore_entsecretdecryptable_typ.SecretDecryptable{}, nil
	case app_typ.MethodName_Core_TaskRun_HookCreatePre:
		return &svccore_enttaskrun_typ.TaskRun{}, nil
	case app_typ.MethodName_Core_File_HookCreatePre:
		return &svccore_entfile_typ.File{}, nil
	case app_typ.MethodName_Core_SecretKey_HookCreatePre:
		return &svccore_entsecretkey_typ.SecretKey{}, nil
	case app_typ.MethodName_Core_Task_HookSavePre:
		return &svccore_enttask_typ.Task{}, nil
	case app_typ.MethodName_User_User_List:
		return &svcuser_entuser_typ.ListRequest{}, nil
	case app_typ.MethodName_Core_File_List:
		return &svccore_entfile_typ.ListRequest{}, nil
	case app_typ.MethodName_Core_SecretKey_List:
		return &svccore_entsecretkey_typ.ListRequest{}, nil
	case app_typ.MethodName_Core_TaskRun_List:
		return &svccore_enttaskrun_typ.ListRequest{}, nil
	case app_typ.MethodName_Core_JobApplicant_List:
		return &svccore_entjobapplicant_typ.ListRequest{}, nil
	case app_typ.MethodName_Core_Task_List:
		return &svccore_enttask_typ.ListRequest{}, nil
	case app_typ.MethodName_Core_SecretDecryptable_List:
		return &svccore_entsecretdecryptable_typ.ListRequest{}, nil
	case app_typ.MethodName_User_Team_List:
		return &svcuser_entteam_typ.ListRequest{}, nil
	case app_typ.MethodName_Auth_Secret_List:
		return &svcauth_entsecret_typ.ListRequest{}, nil
	case app_typ.MethodName_Auth_LoginUser:
		return &svcauth_typ.LoginRequest{}, nil
	case app_typ.MethodName_Core_File_QueryByText:
		return &svccore_entfile_typ.QueryByTextRequest{}, nil
	case app_typ.MethodName_Core_TaskRun_QueryByText:
		return &svccore_enttaskrun_typ.QueryByTextRequest{}, nil
	case app_typ.MethodName_Core_SecretDecryptable_QueryByText:
		return &svccore_entsecretdecryptable_typ.QueryByTextRequest{}, nil
	case app_typ.MethodName_Auth_Secret_QueryByText:
		return &svcauth_entsecret_typ.QueryByTextRequest{}, nil
	case app_typ.MethodName_Core_Task_QueryByText:
		return &svccore_enttask_typ.QueryByTextRequest{}, nil
	case app_typ.MethodName_Core_SecretKey_QueryByText:
		return &svccore_entsecretkey_typ.QueryByTextRequest{}, nil
	case app_typ.MethodName_User_User_QueryByText:
		return &svcuser_entuser_typ.QueryByTextRequest{}, nil
	case app_typ.MethodName_User_Team_QueryByText:
		return &svcuser_entteam_typ.QueryByTextRequest{}, nil
	case app_typ.MethodName_Core_JobApplicant_QueryByText:
		return &svccore_entjobapplicant_typ.QueryByTextRequest{}, nil
	case app_typ.MethodName_Auth_RegisterUser:
		return &svcauth_typ.RegisterUserRequest{}, nil
	case app_typ.MethodName_Core_SecretDecryptabeAdd:
		return &svccore_typ.SecretDecryptableAddRequest{}, nil
	case app_typ.MethodName_Core_Task_Update:
		return &svccore_enttask_typ.UpdateRequest{}, nil
	case app_typ.MethodName_Core_File_Update:
		return &svccore_entfile_typ.UpdateRequest{}, nil
	case app_typ.MethodName_Core_JobApplicant_Update:
		return &svccore_entjobapplicant_typ.UpdateRequest{}, nil
	case app_typ.MethodName_Core_SecretKey_Update:
		return &svccore_entsecretkey_typ.UpdateRequest{}, nil
	case app_typ.MethodName_User_Team_Update:
		return &svcuser_entteam_typ.UpdateRequest{}, nil
	case app_typ.MethodName_Core_TaskRun_Update:
		return &svccore_enttaskrun_typ.UpdateRequest{}, nil
	case app_typ.MethodName_User_User_Update:
		return &svcuser_entuser_typ.UpdateRequest{}, nil
	case app_typ.MethodName_Core_SecretDecryptable_Update:
		return &svccore_entsecretdecryptable_typ.UpdateRequest{}, nil
	case app_typ.MethodName_Auth_Secret_Update:
		return &svcauth_entsecret_typ.UpdateRequest{}, nil

	default:
		return nil, fmt.Errorf("Empty or invalid method name")
	}
}
