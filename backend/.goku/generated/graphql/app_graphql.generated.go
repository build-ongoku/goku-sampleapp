package app_graphql

import (
	"context"

	"github.com/teejays/gokutil/log"
	"github.com/teejays/gokutil/scalars"

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
	svcuser_typ "sampleapp/backend/.goku/generated/service/user/typ"
	app_typ "sampleapp/backend/.goku/generated/typ"
)

var llog = log.GetLogger().WithHeading("Graphql Resolver")

type Core_File_AddRequestResolver struct {
	svccore_entfile_typ.AddRequest
}

func (r *Core_File_AddRequestResolver) Object() *Core_File_FileInputResolver {
	return &Core_File_FileInputResolver{FileInput: r.AddRequest.Object}
}

type Core_SecretDecryptable_AddRequestResolver struct {
	svccore_entsecretdecryptable_typ.AddRequest
}

func (r *Core_SecretDecryptable_AddRequestResolver) Object() *Core_SecretDecryptable_SecretDecryptableInputResolver {
	return &Core_SecretDecryptable_SecretDecryptableInputResolver{SecretDecryptableInput: r.AddRequest.Object}
}

type Core_JobApplicant_AddRequestResolver struct {
	svccore_entjobapplicant_typ.AddRequest
}

func (r *Core_JobApplicant_AddRequestResolver) Object() *Core_JobApplicant_JobApplicantInputResolver {
	return &Core_JobApplicant_JobApplicantInputResolver{JobApplicantInput: r.AddRequest.Object}
}

type Core_Task_AddRequestResolver struct {
	svccore_enttask_typ.AddRequest
}

func (r *Core_Task_AddRequestResolver) Object() *Core_Task_TaskInputResolver {
	return &Core_Task_TaskInputResolver{TaskInput: r.AddRequest.Object}
}

type Core_SecretKey_AddRequestResolver struct {
	svccore_entsecretkey_typ.AddRequest
}

func (r *Core_SecretKey_AddRequestResolver) Object() *Core_SecretKey_SecretKeyInputResolver {
	return &Core_SecretKey_SecretKeyInputResolver{SecretKeyInput: r.AddRequest.Object}
}

type Auth_Secret_AddRequestResolver struct {
	svcauth_entsecret_typ.AddRequest
}

func (r *Auth_Secret_AddRequestResolver) Object() *Auth_Secret_SecretInputResolver {
	return &Auth_Secret_SecretInputResolver{SecretInput: r.AddRequest.Object}
}

type Core_TaskRun_AddRequestResolver struct {
	svccore_enttaskrun_typ.AddRequest
}

func (r *Core_TaskRun_AddRequestResolver) Object() *Core_TaskRun_TaskRunInputResolver {
	return &Core_TaskRun_TaskRunInputResolver{TaskRunInput: r.AddRequest.Object}
}

type User_User_AddRequestResolver struct {
	svcuser_entuser_typ.AddRequest
}

func (r *User_User_AddRequestResolver) Object() *User_User_UserInputResolver {
	return &User_User_UserInputResolver{UserInput: r.AddRequest.Object}
}

type User_Team_AddRequestResolver struct {
	svcuser_entteam_typ.AddRequest
}

func (r *User_Team_AddRequestResolver) Object() *User_Team_TeamInputResolver {
	return &User_Team_TeamInputResolver{TeamInput: r.AddRequest.Object}
}

type AddressResolver struct {
	app_typ.Address
}

func (r *AddressResolver) Line2() *string {
	return &r.Address.Line2
}

type AddressWithMetaResolver struct {
	app_typ.AddressWithMeta
}

func (r *AddressWithMetaResolver) Line2() *string {
	return &r.AddressWithMeta.Line2
}

type AddressInputResolver struct {
	app_typ.AddressInput
}

func (r *AddressInputResolver) Line2() *string {
	return &r.AddressInput.Line2
}

type AddressFieldConditionResolver struct {
	app_typ.AddressFieldCondition
}

type AddressFilterResolver struct {
	app_typ.AddressFilter
}

func (r *AddressFilterResolver) Country() *CountryConditionResolver {
	return &CountryConditionResolver{CountryCondition: *r.AddressFilter.Country}
}
func (r *AddressFilterResolver) And() []*AddressFilterResolver {
	var ret []*AddressFilterResolver
	for _, elem := range r.AddressFilter.And {
		ret = append(ret, &AddressFilterResolver{AddressFilter: elem})
	}
	return ret
}
func (r *AddressFilterResolver) Or() []*AddressFilterResolver {
	var ret []*AddressFilterResolver
	for _, elem := range r.AddressFilter.Or {
		ret = append(ret, &AddressFilterResolver{AddressFilter: elem})
	}
	return ret
}

type Auth_AuthenticateResponseResolver struct {
	svcauth_typ.AuthenticateResponse
}

type Auth_AuthenticateTokenRequestResolver struct {
	svcauth_typ.AuthenticateTokenRequest
}

type ContactResolver struct {
	app_typ.Contact
}

func (r *ContactResolver) Name() *PersonNameResolver {
	return &PersonNameResolver{PersonName: r.Contact.Name}
}
func (r *ContactResolver) Address() *AddressResolver {
	return &AddressResolver{Address: r.Contact.Address}
}

type ContactWithMetaResolver struct {
	app_typ.ContactWithMeta
}

func (r *ContactWithMetaResolver) Name() *PersonNameWithMetaResolver {
	return &PersonNameWithMetaResolver{PersonNameWithMeta: r.ContactWithMeta.Name}
}
func (r *ContactWithMetaResolver) Address() *AddressWithMetaResolver {
	return &AddressWithMetaResolver{AddressWithMeta: r.ContactWithMeta.Address}
}

type ContactInputResolver struct {
	app_typ.ContactInput
}

func (r *ContactInputResolver) Name() *PersonNameInputResolver {
	return &PersonNameInputResolver{PersonNameInput: r.ContactInput.Name}
}
func (r *ContactInputResolver) Address() *AddressInputResolver {
	return &AddressInputResolver{AddressInput: r.ContactInput.Address}
}

type ContactFieldConditionResolver struct {
	app_typ.ContactFieldCondition
}

type ContactFilterResolver struct {
	app_typ.ContactFilter
}

func (r *ContactFilterResolver) Name() *PersonNameFilterResolver {
	return &PersonNameFilterResolver{PersonNameFilter: *r.ContactFilter.Name}
}
func (r *ContactFilterResolver) Address() *AddressFilterResolver {
	return &AddressFilterResolver{AddressFilter: *r.ContactFilter.Address}
}
func (r *ContactFilterResolver) And() []*ContactFilterResolver {
	var ret []*ContactFilterResolver
	for _, elem := range r.ContactFilter.And {
		ret = append(ret, &ContactFilterResolver{ContactFilter: elem})
	}
	return ret
}
func (r *ContactFilterResolver) Or() []*ContactFilterResolver {
	var ret []*ContactFilterResolver
	for _, elem := range r.ContactFilter.Or {
		ret = append(ret, &ContactFilterResolver{ContactFilter: elem})
	}
	return ret
}

type CountryConditionResolver struct {
	app_typ.CountryCondition
}

type Core_CronExpressionResolver struct {
	svccore_typ.CronExpression
}

type Core_CronExpressionWithMetaResolver struct {
	svccore_typ.CronExpressionWithMeta
}

type Core_CronExpressionInputResolver struct {
	svccore_typ.CronExpressionInput
}

type Core_CronExpressionFieldConditionResolver struct {
	svccore_typ.CronExpressionFieldCondition
}

type Core_CronExpressionFilterResolver struct {
	svccore_typ.CronExpressionFilter
}

func (r *Core_CronExpressionFilterResolver) And() []*Core_CronExpressionFilterResolver {
	var ret []*Core_CronExpressionFilterResolver
	for _, elem := range r.CronExpressionFilter.And {
		ret = append(ret, &Core_CronExpressionFilterResolver{CronExpressionFilter: elem})
	}
	return ret
}
func (r *Core_CronExpressionFilterResolver) Or() []*Core_CronExpressionFilterResolver {
	var ret []*Core_CronExpressionFilterResolver
	for _, elem := range r.CronExpressionFilter.Or {
		ret = append(ret, &Core_CronExpressionFilterResolver{CronExpressionFilter: elem})
	}
	return ret
}

type EntityNameConditionResolver struct {
	app_typ.EntityNameCondition
}

type Auth_EntityNameConditionResolver struct {
	svcauth_typ.EntityNameCondition
}

type User_EntityNameConditionResolver struct {
	svcuser_typ.EntityNameCondition
}

type Core_EntityNameConditionResolver struct {
	svccore_typ.EntityNameCondition
}

type Core_FileResolver struct {
	svccore_entfile_typ.File
}

func (r *Core_FileResolver) SizeBytes() int32 {
	return int32(r.File.SizeBytes)
}

type Core_File_FileInputResolver struct {
	svccore_entfile_typ.FileInput
}

type Core_File_FileFieldConditionResolver struct {
	svccore_entfile_typ.FileFieldCondition
}

type Core_File_FileFilterResolver struct {
	svccore_entfile_typ.FileFilter
}

func (r *Core_File_FileFilterResolver) StorageClient() *Core_File_StorageClientConditionResolver {
	return &Core_File_StorageClientConditionResolver{StorageClientCondition: *r.FileFilter.StorageClient}
}
func (r *Core_File_FileFilterResolver) And() []*Core_File_FileFilterResolver {
	var ret []*Core_File_FileFilterResolver
	for _, elem := range r.FileFilter.And {
		ret = append(ret, &Core_File_FileFilterResolver{FileFilter: elem})
	}
	return ret
}
func (r *Core_File_FileFilterResolver) Or() []*Core_File_FileFilterResolver {
	var ret []*Core_File_FileFilterResolver
	for _, elem := range r.FileFilter.Or {
		ret = append(ret, &Core_File_FileFilterResolver{FileFilter: elem})
	}
	return ret
}

type Core_SecretKey_FormatConditionResolver struct {
	svccore_entsecretkey_typ.FormatCondition
}

type GenderConditionResolver struct {
	app_typ.GenderCondition
}

type Core_Task_GetRequestResolver struct {
	svccore_enttask_typ.GetRequest
}

type User_Team_GetRequestResolver struct {
	svcuser_entteam_typ.GetRequest
}

type Auth_Secret_GetRequestResolver struct {
	svcauth_entsecret_typ.GetRequest
}

type Core_File_GetRequestResolver struct {
	svccore_entfile_typ.GetRequest
}

type User_User_GetRequestResolver struct {
	svcuser_entuser_typ.GetRequest
}

type Core_TaskRun_GetRequestResolver struct {
	svccore_enttaskrun_typ.GetRequest
}

type Core_JobApplicant_GetRequestResolver struct {
	svccore_entjobapplicant_typ.GetRequest
}

type Core_SecretDecryptable_GetRequestResolver struct {
	svccore_entsecretdecryptable_typ.GetRequest
}

type Core_SecretKey_GetRequestResolver struct {
	svccore_entsecretkey_typ.GetRequest
}

type Core_JobApplicantResolver struct {
	svccore_entjobapplicant_typ.JobApplicant
}

type Core_JobApplicant_JobApplicantInputResolver struct {
	svccore_entjobapplicant_typ.JobApplicantInput
}

type Core_JobApplicant_JobApplicantFieldConditionResolver struct {
	svccore_entjobapplicant_typ.JobApplicantFieldCondition
}

type Core_JobApplicant_JobApplicantFilterResolver struct {
	svccore_entjobapplicant_typ.JobApplicantFilter
}

func (r *Core_JobApplicant_JobApplicantFilterResolver) And() []*Core_JobApplicant_JobApplicantFilterResolver {
	var ret []*Core_JobApplicant_JobApplicantFilterResolver
	for _, elem := range r.JobApplicantFilter.And {
		ret = append(ret, &Core_JobApplicant_JobApplicantFilterResolver{JobApplicantFilter: elem})
	}
	return ret
}
func (r *Core_JobApplicant_JobApplicantFilterResolver) Or() []*Core_JobApplicant_JobApplicantFilterResolver {
	var ret []*Core_JobApplicant_JobApplicantFilterResolver
	for _, elem := range r.JobApplicantFilter.Or {
		ret = append(ret, &Core_JobApplicant_JobApplicantFilterResolver{JobApplicantFilter: elem})
	}
	return ret
}

type Core_SecretDecryptable_ListRequestResolver struct {
	svccore_entsecretdecryptable_typ.ListRequest
}

func (r *Core_SecretDecryptable_ListRequestResolver) Filter() *Core_SecretDecryptable_SecretDecryptableFilterResolver {
	return &Core_SecretDecryptable_SecretDecryptableFilterResolver{SecretDecryptableFilter: r.ListRequest.Filter}
}

type Core_Task_ListRequestResolver struct {
	svccore_enttask_typ.ListRequest
}

func (r *Core_Task_ListRequestResolver) Filter() *Core_Task_TaskFilterResolver {
	return &Core_Task_TaskFilterResolver{TaskFilter: r.ListRequest.Filter}
}

type User_Team_ListRequestResolver struct {
	svcuser_entteam_typ.ListRequest
}

func (r *User_Team_ListRequestResolver) Filter() *User_Team_TeamFilterResolver {
	return &User_Team_TeamFilterResolver{TeamFilter: r.ListRequest.Filter}
}

type Core_JobApplicant_ListRequestResolver struct {
	svccore_entjobapplicant_typ.ListRequest
}

func (r *Core_JobApplicant_ListRequestResolver) Filter() *Core_JobApplicant_JobApplicantFilterResolver {
	return &Core_JobApplicant_JobApplicantFilterResolver{JobApplicantFilter: r.ListRequest.Filter}
}

type User_User_ListRequestResolver struct {
	svcuser_entuser_typ.ListRequest
}

func (r *User_User_ListRequestResolver) Filter() *User_User_UserFilterResolver {
	return &User_User_UserFilterResolver{UserFilter: r.ListRequest.Filter}
}

type Core_SecretKey_ListRequestResolver struct {
	svccore_entsecretkey_typ.ListRequest
}

func (r *Core_SecretKey_ListRequestResolver) Filter() *Core_SecretKey_SecretKeyFilterResolver {
	return &Core_SecretKey_SecretKeyFilterResolver{SecretKeyFilter: r.ListRequest.Filter}
}

type Core_TaskRun_ListRequestResolver struct {
	svccore_enttaskrun_typ.ListRequest
}

func (r *Core_TaskRun_ListRequestResolver) Filter() *Core_TaskRun_TaskRunFilterResolver {
	return &Core_TaskRun_TaskRunFilterResolver{TaskRunFilter: r.ListRequest.Filter}
}

type Auth_Secret_ListRequestResolver struct {
	svcauth_entsecret_typ.ListRequest
}

func (r *Auth_Secret_ListRequestResolver) Filter() *Auth_Secret_SecretFilterResolver {
	return &Auth_Secret_SecretFilterResolver{SecretFilter: r.ListRequest.Filter}
}

type Core_File_ListRequestResolver struct {
	svccore_entfile_typ.ListRequest
}

func (r *Core_File_ListRequestResolver) Filter() *Core_File_FileFilterResolver {
	return &Core_File_FileFilterResolver{FileFilter: r.ListRequest.Filter}
}

type Core_JobApplicant_ListResponseResolver struct {
	svccore_entjobapplicant_typ.ListResponse
}

func (r *Core_JobApplicant_ListResponseResolver) Items() []*Core_JobApplicantResolver {
	var ret []*Core_JobApplicantResolver
	for _, elem := range r.ListResponse.Items {
		ret = append(ret, &Core_JobApplicantResolver{JobApplicant: elem})
	}
	return ret
}
func (r *Core_JobApplicant_ListResponseResolver) Count() int32 {
	return int32(r.ListResponse.Count)
}

type User_User_ListResponseResolver struct {
	svcuser_entuser_typ.ListResponse
}

func (r *User_User_ListResponseResolver) Items() []*User_UserResolver {
	var ret []*User_UserResolver
	for _, elem := range r.ListResponse.Items {
		ret = append(ret, &User_UserResolver{User: elem})
	}
	return ret
}
func (r *User_User_ListResponseResolver) Count() int32 {
	return int32(r.ListResponse.Count)
}

type Core_File_ListResponseResolver struct {
	svccore_entfile_typ.ListResponse
}

func (r *Core_File_ListResponseResolver) Items() []*Core_FileResolver {
	var ret []*Core_FileResolver
	for _, elem := range r.ListResponse.Items {
		ret = append(ret, &Core_FileResolver{File: elem})
	}
	return ret
}
func (r *Core_File_ListResponseResolver) Count() int32 {
	return int32(r.ListResponse.Count)
}

type User_Team_ListResponseResolver struct {
	svcuser_entteam_typ.ListResponse
}

func (r *User_Team_ListResponseResolver) Items() []*User_TeamResolver {
	var ret []*User_TeamResolver
	for _, elem := range r.ListResponse.Items {
		ret = append(ret, &User_TeamResolver{Team: elem})
	}
	return ret
}
func (r *User_Team_ListResponseResolver) Count() int32 {
	return int32(r.ListResponse.Count)
}

type Core_TaskRun_ListResponseResolver struct {
	svccore_enttaskrun_typ.ListResponse
}

func (r *Core_TaskRun_ListResponseResolver) Items() []*Core_TaskRunResolver {
	var ret []*Core_TaskRunResolver
	for _, elem := range r.ListResponse.Items {
		ret = append(ret, &Core_TaskRunResolver{TaskRun: elem})
	}
	return ret
}
func (r *Core_TaskRun_ListResponseResolver) Count() int32 {
	return int32(r.ListResponse.Count)
}

type Core_Task_ListResponseResolver struct {
	svccore_enttask_typ.ListResponse
}

func (r *Core_Task_ListResponseResolver) Items() []*Core_TaskResolver {
	var ret []*Core_TaskResolver
	for _, elem := range r.ListResponse.Items {
		ret = append(ret, &Core_TaskResolver{Task: elem})
	}
	return ret
}
func (r *Core_Task_ListResponseResolver) Count() int32 {
	return int32(r.ListResponse.Count)
}

type Auth_Secret_ListResponseResolver struct {
	svcauth_entsecret_typ.ListResponse
}

func (r *Auth_Secret_ListResponseResolver) Items() []*Auth_SecretResolver {
	var ret []*Auth_SecretResolver
	for _, elem := range r.ListResponse.Items {
		ret = append(ret, &Auth_SecretResolver{Secret: elem})
	}
	return ret
}
func (r *Auth_Secret_ListResponseResolver) Count() int32 {
	return int32(r.ListResponse.Count)
}

type Core_SecretKey_ListResponseResolver struct {
	svccore_entsecretkey_typ.ListResponse
}

func (r *Core_SecretKey_ListResponseResolver) Items() []*Core_SecretKeyResolver {
	var ret []*Core_SecretKeyResolver
	for _, elem := range r.ListResponse.Items {
		ret = append(ret, &Core_SecretKeyResolver{SecretKey: elem})
	}
	return ret
}
func (r *Core_SecretKey_ListResponseResolver) Count() int32 {
	return int32(r.ListResponse.Count)
}

type Core_SecretDecryptable_ListResponseResolver struct {
	svccore_entsecretdecryptable_typ.ListResponse
}

func (r *Core_SecretDecryptable_ListResponseResolver) Items() []*Core_SecretDecryptableResolver {
	var ret []*Core_SecretDecryptableResolver
	for _, elem := range r.ListResponse.Items {
		ret = append(ret, &Core_SecretDecryptableResolver{SecretDecryptable: elem})
	}
	return ret
}
func (r *Core_SecretDecryptable_ListResponseResolver) Count() int32 {
	return int32(r.ListResponse.Count)
}

type Auth_LoginRequestResolver struct {
	svcauth_typ.LoginRequest
}

type Core_TaskRun_MethodNameConditionResolver struct {
	svccore_enttaskrun_typ.MethodNameCondition
}

type Core_Task_MethodNameConditionResolver struct {
	svccore_enttask_typ.MethodNameCondition
}

type Auth_Secret_MethodNameConditionResolver struct {
	svcauth_entsecret_typ.MethodNameCondition
}

type MethodNameConditionResolver struct {
	app_typ.MethodNameCondition
}

type Auth_MethodNameConditionResolver struct {
	svcauth_typ.MethodNameCondition
}

type Core_MethodNameConditionResolver struct {
	svccore_typ.MethodNameCondition
}

type User_User_MethodNameConditionResolver struct {
	svcuser_entuser_typ.MethodNameCondition
}

type User_Team_MethodNameConditionResolver struct {
	svcuser_entteam_typ.MethodNameCondition
}

type Core_File_MethodNameConditionResolver struct {
	svccore_entfile_typ.MethodNameCondition
}

type Core_SecretDecryptable_MethodNameConditionResolver struct {
	svccore_entsecretdecryptable_typ.MethodNameCondition
}

type Core_SecretKey_MethodNameConditionResolver struct {
	svccore_entsecretkey_typ.MethodNameCondition
}

type Core_JobApplicant_MethodNameConditionResolver struct {
	svccore_entjobapplicant_typ.MethodNameCondition
}

type PersonNameResolver struct {
	app_typ.PersonName
}

func (r *PersonNameResolver) MiddleInitial() *string {
	return &r.PersonName.MiddleInitial
}

type PersonNameWithMetaResolver struct {
	app_typ.PersonNameWithMeta
}

func (r *PersonNameWithMetaResolver) MiddleInitial() *string {
	return &r.PersonNameWithMeta.MiddleInitial
}

type PersonNameInputResolver struct {
	app_typ.PersonNameInput
}

func (r *PersonNameInputResolver) MiddleInitial() *string {
	return &r.PersonNameInput.MiddleInitial
}

type PersonNameFieldConditionResolver struct {
	app_typ.PersonNameFieldCondition
}

type PersonNameFilterResolver struct {
	app_typ.PersonNameFilter
}

func (r *PersonNameFilterResolver) And() []*PersonNameFilterResolver {
	var ret []*PersonNameFilterResolver
	for _, elem := range r.PersonNameFilter.And {
		ret = append(ret, &PersonNameFilterResolver{PersonNameFilter: elem})
	}
	return ret
}
func (r *PersonNameFilterResolver) Or() []*PersonNameFilterResolver {
	var ret []*PersonNameFilterResolver
	for _, elem := range r.PersonNameFilter.Or {
		ret = append(ret, &PersonNameFilterResolver{PersonNameFilter: elem})
	}
	return ret
}

type PhoneNumberResolver struct {
	app_typ.PhoneNumber
}

func (r *PhoneNumberResolver) CountryCode() int32 {
	return int32(r.PhoneNumber.CountryCode)
}
func (r *PhoneNumberResolver) Extension() *string {
	return &r.PhoneNumber.Extension
}

type PhoneNumberWithMetaResolver struct {
	app_typ.PhoneNumberWithMeta
}

func (r *PhoneNumberWithMetaResolver) CountryCode() int32 {
	return int32(r.PhoneNumberWithMeta.CountryCode)
}
func (r *PhoneNumberWithMetaResolver) Extension() *string {
	return &r.PhoneNumberWithMeta.Extension
}

type PhoneNumberInputResolver struct {
	app_typ.PhoneNumberInput
}

func (r *PhoneNumberInputResolver) CountryCode() int32 {
	return int32(r.PhoneNumberInput.CountryCode)
}
func (r *PhoneNumberInputResolver) Extension() *string {
	return &r.PhoneNumberInput.Extension
}

type PhoneNumberFieldConditionResolver struct {
	app_typ.PhoneNumberFieldCondition
}

type PhoneNumberFilterResolver struct {
	app_typ.PhoneNumberFilter
}

func (r *PhoneNumberFilterResolver) And() []*PhoneNumberFilterResolver {
	var ret []*PhoneNumberFilterResolver
	for _, elem := range r.PhoneNumberFilter.And {
		ret = append(ret, &PhoneNumberFilterResolver{PhoneNumberFilter: elem})
	}
	return ret
}
func (r *PhoneNumberFilterResolver) Or() []*PhoneNumberFilterResolver {
	var ret []*PhoneNumberFilterResolver
	for _, elem := range r.PhoneNumberFilter.Or {
		ret = append(ret, &PhoneNumberFilterResolver{PhoneNumberFilter: elem})
	}
	return ret
}

type Core_TaskRun_QueryByTextRequestResolver struct {
	svccore_enttaskrun_typ.QueryByTextRequest
}

type Core_Task_QueryByTextRequestResolver struct {
	svccore_enttask_typ.QueryByTextRequest
}

type Auth_Secret_QueryByTextRequestResolver struct {
	svcauth_entsecret_typ.QueryByTextRequest
}

type Core_SecretKey_QueryByTextRequestResolver struct {
	svccore_entsecretkey_typ.QueryByTextRequest
}

type User_User_QueryByTextRequestResolver struct {
	svcuser_entuser_typ.QueryByTextRequest
}

type Core_SecretDecryptable_QueryByTextRequestResolver struct {
	svccore_entsecretdecryptable_typ.QueryByTextRequest
}

type User_Team_QueryByTextRequestResolver struct {
	svcuser_entteam_typ.QueryByTextRequest
}

type Core_File_QueryByTextRequestResolver struct {
	svccore_entfile_typ.QueryByTextRequest
}

type Core_JobApplicant_QueryByTextRequestResolver struct {
	svccore_entjobapplicant_typ.QueryByTextRequest
}

type Auth_RegisterUserRequestResolver struct {
	svcauth_typ.RegisterUserRequest
}

func (r *Auth_RegisterUserRequestResolver) Name() *PersonNameInputResolver {
	return &PersonNameInputResolver{PersonNameInput: r.RegisterUserRequest.Name}
}

type Auth_SecretResolver struct {
	svcauth_entsecret_typ.Secret
}

type Auth_Secret_SecretInputResolver struct {
	svcauth_entsecret_typ.SecretInput
}

type Core_SecretDecryptableResolver struct {
	svccore_entsecretdecryptable_typ.SecretDecryptable
}

func (r *Core_SecretDecryptableResolver) RawValue() *string {
	return &r.SecretDecryptable.RawValue
}

type Core_SecretDecryptable_SecretDecryptableInputResolver struct {
	svccore_entsecretdecryptable_typ.SecretDecryptableInput
}

func (r *Core_SecretDecryptable_SecretDecryptableInputResolver) RawValue() *string {
	return &r.SecretDecryptableInput.RawValue
}

type Core_SecretDecryptable_SecretDecryptableActionConditionResolver struct {
	svccore_entsecretdecryptable_typ.SecretDecryptableActionCondition
}

type Core_SecretDecryptableAddRequestResolver struct {
	svccore_typ.SecretDecryptableAddRequest
}

func (r *Core_SecretDecryptableAddRequestResolver) SecretKeyID() *scalars.ID {
	return &r.SecretDecryptableAddRequest.SecretKeyID
}

type Core_SecretDecryptableAddRequestWithMetaResolver struct {
	svccore_typ.SecretDecryptableAddRequestWithMeta
}

func (r *Core_SecretDecryptableAddRequestWithMetaResolver) SecretKeyID() *scalars.ID {
	return &r.SecretDecryptableAddRequestWithMeta.SecretKeyID
}

type Core_SecretDecryptableAddRequestInputResolver struct {
	svccore_typ.SecretDecryptableAddRequestInput
}

func (r *Core_SecretDecryptableAddRequestInputResolver) SecretKeyID() *scalars.ID {
	return &r.SecretDecryptableAddRequestInput.SecretKeyID
}

type Core_SecretDecryptableAddRequestFieldConditionResolver struct {
	svccore_typ.SecretDecryptableAddRequestFieldCondition
}

type Core_SecretDecryptableAddRequestFilterResolver struct {
	svccore_typ.SecretDecryptableAddRequestFilter
}

func (r *Core_SecretDecryptableAddRequestFilterResolver) And() []*Core_SecretDecryptableAddRequestFilterResolver {
	var ret []*Core_SecretDecryptableAddRequestFilterResolver
	for _, elem := range r.SecretDecryptableAddRequestFilter.And {
		ret = append(ret, &Core_SecretDecryptableAddRequestFilterResolver{SecretDecryptableAddRequestFilter: elem})
	}
	return ret
}
func (r *Core_SecretDecryptableAddRequestFilterResolver) Or() []*Core_SecretDecryptableAddRequestFilterResolver {
	var ret []*Core_SecretDecryptableAddRequestFilterResolver
	for _, elem := range r.SecretDecryptableAddRequestFilter.Or {
		ret = append(ret, &Core_SecretDecryptableAddRequestFilterResolver{SecretDecryptableAddRequestFilter: elem})
	}
	return ret
}

type Core_SecretDecryptable_SecretDecryptableFieldConditionResolver struct {
	svccore_entsecretdecryptable_typ.SecretDecryptableFieldCondition
}

type Core_SecretDecryptable_SecretDecryptableFilterResolver struct {
	svccore_entsecretdecryptable_typ.SecretDecryptableFilter
}

func (r *Core_SecretDecryptable_SecretDecryptableFilterResolver) And() []*Core_SecretDecryptable_SecretDecryptableFilterResolver {
	var ret []*Core_SecretDecryptable_SecretDecryptableFilterResolver
	for _, elem := range r.SecretDecryptableFilter.And {
		ret = append(ret, &Core_SecretDecryptable_SecretDecryptableFilterResolver{SecretDecryptableFilter: elem})
	}
	return ret
}
func (r *Core_SecretDecryptable_SecretDecryptableFilterResolver) Or() []*Core_SecretDecryptable_SecretDecryptableFilterResolver {
	var ret []*Core_SecretDecryptable_SecretDecryptableFilterResolver
	for _, elem := range r.SecretDecryptableFilter.Or {
		ret = append(ret, &Core_SecretDecryptable_SecretDecryptableFilterResolver{SecretDecryptableFilter: elem})
	}
	return ret
}

type Auth_Secret_SecretFieldConditionResolver struct {
	svcauth_entsecret_typ.SecretFieldCondition
}

type Auth_Secret_SecretFilterResolver struct {
	svcauth_entsecret_typ.SecretFilter
}

func (r *Auth_Secret_SecretFilterResolver) Type() *Auth_Secret_TypeConditionResolver {
	return &Auth_Secret_TypeConditionResolver{TypeCondition: *r.SecretFilter.Type}
}
func (r *Auth_Secret_SecretFilterResolver) And() []*Auth_Secret_SecretFilterResolver {
	var ret []*Auth_Secret_SecretFilterResolver
	for _, elem := range r.SecretFilter.And {
		ret = append(ret, &Auth_Secret_SecretFilterResolver{SecretFilter: elem})
	}
	return ret
}
func (r *Auth_Secret_SecretFilterResolver) Or() []*Auth_Secret_SecretFilterResolver {
	var ret []*Auth_Secret_SecretFilterResolver
	for _, elem := range r.SecretFilter.Or {
		ret = append(ret, &Auth_Secret_SecretFilterResolver{SecretFilter: elem})
	}
	return ret
}

type Core_SecretKeyResolver struct {
	svccore_entsecretkey_typ.SecretKey
}

type Core_SecretKey_SecretKeyInputResolver struct {
	svccore_entsecretkey_typ.SecretKeyInput
}

type Core_SecretKey_SecretKeyFieldConditionResolver struct {
	svccore_entsecretkey_typ.SecretKeyFieldCondition
}

type Core_SecretKey_SecretKeyFilterResolver struct {
	svccore_entsecretkey_typ.SecretKeyFilter
}

func (r *Core_SecretKey_SecretKeyFilterResolver) Type() *Core_SecretKey_TypeConditionResolver {
	return &Core_SecretKey_TypeConditionResolver{TypeCondition: *r.SecretKeyFilter.Type}
}
func (r *Core_SecretKey_SecretKeyFilterResolver) PrivateKeyFormat() *Core_SecretKey_FormatConditionResolver {
	return &Core_SecretKey_FormatConditionResolver{FormatCondition: *r.SecretKeyFilter.PrivateKeyFormat}
}
func (r *Core_SecretKey_SecretKeyFilterResolver) PublicKeyFormat() *Core_SecretKey_FormatConditionResolver {
	return &Core_SecretKey_FormatConditionResolver{FormatCondition: *r.SecretKeyFilter.PublicKeyFormat}
}
func (r *Core_SecretKey_SecretKeyFilterResolver) And() []*Core_SecretKey_SecretKeyFilterResolver {
	var ret []*Core_SecretKey_SecretKeyFilterResolver
	for _, elem := range r.SecretKeyFilter.And {
		ret = append(ret, &Core_SecretKey_SecretKeyFilterResolver{SecretKeyFilter: elem})
	}
	return ret
}
func (r *Core_SecretKey_SecretKeyFilterResolver) Or() []*Core_SecretKey_SecretKeyFilterResolver {
	var ret []*Core_SecretKey_SecretKeyFilterResolver
	for _, elem := range r.SecretKeyFilter.Or {
		ret = append(ret, &Core_SecretKey_SecretKeyFilterResolver{SecretKeyFilter: elem})
	}
	return ret
}

type ServiceNameConditionResolver struct {
	app_typ.ServiceNameCondition
}

type Core_JobApplicant_StandardEntityRequestResolver struct {
	svccore_entjobapplicant_typ.StandardEntityRequest
}

type Core_TaskRun_StandardEntityRequestResolver struct {
	svccore_enttaskrun_typ.StandardEntityRequest
}

type Core_SecretKey_StandardEntityRequestResolver struct {
	svccore_entsecretkey_typ.StandardEntityRequest
}

type User_Team_StandardEntityRequestResolver struct {
	svcuser_entteam_typ.StandardEntityRequest
}

type Core_SecretDecryptable_StandardEntityRequestResolver struct {
	svccore_entsecretdecryptable_typ.StandardEntityRequest
}

type Core_File_StandardEntityRequestResolver struct {
	svccore_entfile_typ.StandardEntityRequest
}

type Core_Task_StandardEntityRequestResolver struct {
	svccore_enttask_typ.StandardEntityRequest
}

type Auth_Secret_StandardEntityRequestResolver struct {
	svcauth_entsecret_typ.StandardEntityRequest
}

type User_User_StandardEntityRequestResolver struct {
	svcuser_entuser_typ.StandardEntityRequest
}

type Core_TaskRun_StandardEntityResponseResolver struct {
	svccore_enttaskrun_typ.StandardEntityResponse
}

func (r *Core_TaskRun_StandardEntityResponseResolver) Object() *Core_TaskRunResolver {
	return &Core_TaskRunResolver{TaskRun: r.StandardEntityResponse.Object}
}

type Auth_Secret_StandardEntityResponseResolver struct {
	svcauth_entsecret_typ.StandardEntityResponse
}

func (r *Auth_Secret_StandardEntityResponseResolver) Object() *Auth_SecretResolver {
	return &Auth_SecretResolver{Secret: r.StandardEntityResponse.Object}
}

type Core_SecretDecryptable_StandardEntityResponseResolver struct {
	svccore_entsecretdecryptable_typ.StandardEntityResponse
}

func (r *Core_SecretDecryptable_StandardEntityResponseResolver) Object() *Core_SecretDecryptableResolver {
	return &Core_SecretDecryptableResolver{SecretDecryptable: r.StandardEntityResponse.Object}
}

type Core_SecretKey_StandardEntityResponseResolver struct {
	svccore_entsecretkey_typ.StandardEntityResponse
}

func (r *Core_SecretKey_StandardEntityResponseResolver) Object() *Core_SecretKeyResolver {
	return &Core_SecretKeyResolver{SecretKey: r.StandardEntityResponse.Object}
}

type Core_Task_StandardEntityResponseResolver struct {
	svccore_enttask_typ.StandardEntityResponse
}

func (r *Core_Task_StandardEntityResponseResolver) Object() *Core_TaskResolver {
	return &Core_TaskResolver{Task: r.StandardEntityResponse.Object}
}

type Core_File_StandardEntityResponseResolver struct {
	svccore_entfile_typ.StandardEntityResponse
}

func (r *Core_File_StandardEntityResponseResolver) Object() *Core_FileResolver {
	return &Core_FileResolver{File: r.StandardEntityResponse.Object}
}

type User_Team_StandardEntityResponseResolver struct {
	svcuser_entteam_typ.StandardEntityResponse
}

func (r *User_Team_StandardEntityResponseResolver) Object() *User_TeamResolver {
	return &User_TeamResolver{Team: r.StandardEntityResponse.Object}
}

type User_User_StandardEntityResponseResolver struct {
	svcuser_entuser_typ.StandardEntityResponse
}

func (r *User_User_StandardEntityResponseResolver) Object() *User_UserResolver {
	return &User_UserResolver{User: r.StandardEntityResponse.Object}
}

type Core_JobApplicant_StandardEntityResponseResolver struct {
	svccore_entjobapplicant_typ.StandardEntityResponse
}

func (r *Core_JobApplicant_StandardEntityResponseResolver) Object() *Core_JobApplicantResolver {
	return &Core_JobApplicantResolver{JobApplicant: r.StandardEntityResponse.Object}
}

type Core_TaskRun_StatusConditionResolver struct {
	svccore_enttaskrun_typ.StatusCondition
}

type Core_File_StorageClientConditionResolver struct {
	svccore_entfile_typ.StorageClientCondition
}

type Core_TaskResolver struct {
	svccore_enttask_typ.Task
}

type Core_Task_TaskInputResolver struct {
	svccore_enttask_typ.TaskInput
}

type Core_Task_TaskActionConditionResolver struct {
	svccore_enttask_typ.TaskActionCondition
}

type Core_Task_TaskFieldConditionResolver struct {
	svccore_enttask_typ.TaskFieldCondition
}

type Core_Task_TaskFilterResolver struct {
	svccore_enttask_typ.TaskFilter
}

func (r *Core_Task_TaskFilterResolver) Method() *MethodNameConditionResolver {
	return &MethodNameConditionResolver{MethodNameCondition: *r.TaskFilter.Method}
}
func (r *Core_Task_TaskFilterResolver) And() []*Core_Task_TaskFilterResolver {
	var ret []*Core_Task_TaskFilterResolver
	for _, elem := range r.TaskFilter.And {
		ret = append(ret, &Core_Task_TaskFilterResolver{TaskFilter: elem})
	}
	return ret
}
func (r *Core_Task_TaskFilterResolver) Or() []*Core_Task_TaskFilterResolver {
	var ret []*Core_Task_TaskFilterResolver
	for _, elem := range r.TaskFilter.Or {
		ret = append(ret, &Core_Task_TaskFilterResolver{TaskFilter: elem})
	}
	return ret
}

type Core_TaskRunResolver struct {
	svccore_enttaskrun_typ.TaskRun
}

type Core_TaskRun_TaskRunInputResolver struct {
	svccore_enttaskrun_typ.TaskRunInput
}

type Core_TaskRun_TaskRunActionConditionResolver struct {
	svccore_enttaskrun_typ.TaskRunActionCondition
}

type Core_TaskRun_TaskRunFieldConditionResolver struct {
	svccore_enttaskrun_typ.TaskRunFieldCondition
}

type Core_TaskRun_TaskRunFilterResolver struct {
	svccore_enttaskrun_typ.TaskRunFilter
}

func (r *Core_TaskRun_TaskRunFilterResolver) Status() *Core_TaskRun_StatusConditionResolver {
	return &Core_TaskRun_StatusConditionResolver{StatusCondition: *r.TaskRunFilter.Status}
}
func (r *Core_TaskRun_TaskRunFilterResolver) TriggeredBy() *Core_TaskRun_TriggeredByConditionResolver {
	return &Core_TaskRun_TriggeredByConditionResolver{TriggeredByCondition: *r.TaskRunFilter.TriggeredBy}
}
func (r *Core_TaskRun_TaskRunFilterResolver) And() []*Core_TaskRun_TaskRunFilterResolver {
	var ret []*Core_TaskRun_TaskRunFilterResolver
	for _, elem := range r.TaskRunFilter.And {
		ret = append(ret, &Core_TaskRun_TaskRunFilterResolver{TaskRunFilter: elem})
	}
	return ret
}
func (r *Core_TaskRun_TaskRunFilterResolver) Or() []*Core_TaskRun_TaskRunFilterResolver {
	var ret []*Core_TaskRun_TaskRunFilterResolver
	for _, elem := range r.TaskRunFilter.Or {
		ret = append(ret, &Core_TaskRun_TaskRunFilterResolver{TaskRunFilter: elem})
	}
	return ret
}

type User_TeamResolver struct {
	svcuser_entteam_typ.Team
}

type User_Team_TeamInputResolver struct {
	svcuser_entteam_typ.TeamInput
}

type User_Team_TeamFieldConditionResolver struct {
	svcuser_entteam_typ.TeamFieldCondition
}

type User_Team_TeamFilterResolver struct {
	svcuser_entteam_typ.TeamFilter
}

func (r *User_Team_TeamFilterResolver) And() []*User_Team_TeamFilterResolver {
	var ret []*User_Team_TeamFilterResolver
	for _, elem := range r.TeamFilter.And {
		ret = append(ret, &User_Team_TeamFilterResolver{TeamFilter: elem})
	}
	return ret
}
func (r *User_Team_TeamFilterResolver) Or() []*User_Team_TeamFilterResolver {
	var ret []*User_Team_TeamFilterResolver
	for _, elem := range r.TeamFilter.Or {
		ret = append(ret, &User_Team_TeamFilterResolver{TeamFilter: elem})
	}
	return ret
}

type Core_TaskRun_TriggeredByConditionResolver struct {
	svccore_enttaskrun_typ.TriggeredByCondition
}

type Auth_Secret_TypeConditionResolver struct {
	svcauth_entsecret_typ.TypeCondition
}

type Core_SecretKey_TypeConditionResolver struct {
	svccore_entsecretkey_typ.TypeCondition
}

type Core_SecretKey_UpdateRequestResolver struct {
	svccore_entsecretkey_typ.UpdateRequest
}

func (r *Core_SecretKey_UpdateRequestResolver) Object() *Core_SecretKeyResolver {
	return &Core_SecretKeyResolver{SecretKey: r.UpdateRequest.Object}
}

type Core_File_UpdateRequestResolver struct {
	svccore_entfile_typ.UpdateRequest
}

func (r *Core_File_UpdateRequestResolver) Object() *Core_FileResolver {
	return &Core_FileResolver{File: r.UpdateRequest.Object}
}

type Core_TaskRun_UpdateRequestResolver struct {
	svccore_enttaskrun_typ.UpdateRequest
}

func (r *Core_TaskRun_UpdateRequestResolver) Object() *Core_TaskRunResolver {
	return &Core_TaskRunResolver{TaskRun: r.UpdateRequest.Object}
}

type User_Team_UpdateRequestResolver struct {
	svcuser_entteam_typ.UpdateRequest
}

func (r *User_Team_UpdateRequestResolver) Object() *User_TeamResolver {
	return &User_TeamResolver{Team: r.UpdateRequest.Object}
}

type Core_Task_UpdateRequestResolver struct {
	svccore_enttask_typ.UpdateRequest
}

func (r *Core_Task_UpdateRequestResolver) Object() *Core_TaskResolver {
	return &Core_TaskResolver{Task: r.UpdateRequest.Object}
}

type Core_SecretDecryptable_UpdateRequestResolver struct {
	svccore_entsecretdecryptable_typ.UpdateRequest
}

func (r *Core_SecretDecryptable_UpdateRequestResolver) Object() *Core_SecretDecryptableResolver {
	return &Core_SecretDecryptableResolver{SecretDecryptable: r.UpdateRequest.Object}
}

type User_User_UpdateRequestResolver struct {
	svcuser_entuser_typ.UpdateRequest
}

func (r *User_User_UpdateRequestResolver) Object() *User_UserResolver {
	return &User_UserResolver{User: r.UpdateRequest.Object}
}

type Core_JobApplicant_UpdateRequestResolver struct {
	svccore_entjobapplicant_typ.UpdateRequest
}

func (r *Core_JobApplicant_UpdateRequestResolver) Object() *Core_JobApplicantResolver {
	return &Core_JobApplicantResolver{JobApplicant: r.UpdateRequest.Object}
}

type Auth_Secret_UpdateRequestResolver struct {
	svcauth_entsecret_typ.UpdateRequest
}

func (r *Auth_Secret_UpdateRequestResolver) Object() *Auth_SecretResolver {
	return &Auth_SecretResolver{Secret: r.UpdateRequest.Object}
}

type Core_File_UpdateResponseResolver struct {
	svccore_entfile_typ.UpdateResponse
}

func (r *Core_File_UpdateResponseResolver) Object() *Core_FileResolver {
	return &Core_FileResolver{File: r.UpdateResponse.Object}
}

type User_Team_UpdateResponseResolver struct {
	svcuser_entteam_typ.UpdateResponse
}

func (r *User_Team_UpdateResponseResolver) Object() *User_TeamResolver {
	return &User_TeamResolver{Team: r.UpdateResponse.Object}
}

type Core_JobApplicant_UpdateResponseResolver struct {
	svccore_entjobapplicant_typ.UpdateResponse
}

func (r *Core_JobApplicant_UpdateResponseResolver) Object() *Core_JobApplicantResolver {
	return &Core_JobApplicantResolver{JobApplicant: r.UpdateResponse.Object}
}

type Core_Task_UpdateResponseResolver struct {
	svccore_enttask_typ.UpdateResponse
}

func (r *Core_Task_UpdateResponseResolver) Object() *Core_TaskResolver {
	return &Core_TaskResolver{Task: r.UpdateResponse.Object}
}

type User_User_UpdateResponseResolver struct {
	svcuser_entuser_typ.UpdateResponse
}

func (r *User_User_UpdateResponseResolver) Object() *User_UserResolver {
	return &User_UserResolver{User: r.UpdateResponse.Object}
}

type Core_SecretKey_UpdateResponseResolver struct {
	svccore_entsecretkey_typ.UpdateResponse
}

func (r *Core_SecretKey_UpdateResponseResolver) Object() *Core_SecretKeyResolver {
	return &Core_SecretKeyResolver{SecretKey: r.UpdateResponse.Object}
}

type Core_SecretDecryptable_UpdateResponseResolver struct {
	svccore_entsecretdecryptable_typ.UpdateResponse
}

func (r *Core_SecretDecryptable_UpdateResponseResolver) Object() *Core_SecretDecryptableResolver {
	return &Core_SecretDecryptableResolver{SecretDecryptable: r.UpdateResponse.Object}
}

type Auth_Secret_UpdateResponseResolver struct {
	svcauth_entsecret_typ.UpdateResponse
}

func (r *Auth_Secret_UpdateResponseResolver) Object() *Auth_SecretResolver {
	return &Auth_SecretResolver{Secret: r.UpdateResponse.Object}
}

type Core_TaskRun_UpdateResponseResolver struct {
	svccore_enttaskrun_typ.UpdateResponse
}

func (r *Core_TaskRun_UpdateResponseResolver) Object() *Core_TaskRunResolver {
	return &Core_TaskRunResolver{TaskRun: r.UpdateResponse.Object}
}

type User_UserResolver struct {
	svcuser_entuser_typ.User
}

func (r *User_UserResolver) Name() *PersonNameWithMetaResolver {
	return &PersonNameWithMetaResolver{PersonNameWithMeta: r.User.Name}
}
func (r *User_UserResolver) Addresses() []*AddressWithMetaResolver {
	var ret []*AddressWithMetaResolver
	for _, elem := range r.User.Addresses {
		ret = append(ret, &AddressWithMetaResolver{AddressWithMeta: elem})
	}
	return ret
}

type User_User_UserInputResolver struct {
	svcuser_entuser_typ.UserInput
}

func (r *User_User_UserInputResolver) Name() *PersonNameInputResolver {
	return &PersonNameInputResolver{PersonNameInput: r.UserInput.Name}
}
func (r *User_User_UserInputResolver) Addresses() []*AddressInputResolver {
	var ret []*AddressInputResolver
	for _, elem := range r.UserInput.Addresses {
		ret = append(ret, &AddressInputResolver{AddressInput: elem})
	}
	return ret
}

type User_User_UserFieldConditionResolver struct {
	svcuser_entuser_typ.UserFieldCondition
}

type User_User_UserFilterResolver struct {
	svcuser_entuser_typ.UserFilter
}

func (r *User_User_UserFilterResolver) Name() *PersonNameFilterResolver {
	return &PersonNameFilterResolver{PersonNameFilter: *r.UserFilter.Name}
}
func (r *User_User_UserFilterResolver) HavingAddresses() *AddressFilterResolver {
	return &AddressFilterResolver{AddressFilter: *r.UserFilter.HavingAddresses}
}
func (r *User_User_UserFilterResolver) And() []*User_User_UserFilterResolver {
	var ret []*User_User_UserFilterResolver
	for _, elem := range r.UserFilter.And {
		ret = append(ret, &User_User_UserFilterResolver{UserFilter: elem})
	}
	return ret
}
func (r *User_User_UserFilterResolver) Or() []*User_User_UserFilterResolver {
	var ret []*User_User_UserFilterResolver
	for _, elem := range r.UserFilter.Or {
		ret = append(ret, &User_User_UserFilterResolver{UserFilter: elem})
	}
	return ret
}

// Resolver structs and their methods

type Resolver struct {
	QueryResolver
	MutationResolver
}

func NewResolver(getClientFn app_client.ClientGetterFn) *Resolver {
	return &Resolver{}
}

type QueryResolver struct {
	getClientFn app_client.ClientGetterFn
}

type Auth_AuthenticateTokenArgs struct {
	Req *Auth_AuthenticateTokenRequestResolver
}

func (r *QueryResolver) Auth_AuthenticateToken(ctx context.Context, args Auth_AuthenticateTokenArgs) (*Auth_AuthenticateResponseResolver, error) {
	var ret Auth_AuthenticateResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for AuthenticateToken...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Auth().AuthenticateToken(ctx, args.Req.AuthenticateTokenRequest)
	if err != nil {
		return nil, err
	}

	ret.AuthenticateResponse = resp

	return &ret, nil
}

type Core_SecretDecryptable_GetArgs struct {
	Req *Core_SecretDecryptable_GetRequestResolver
}

func (r *QueryResolver) Core_SecretDecryptable_Get(ctx context.Context, args Core_SecretDecryptable_GetArgs) (*Core_SecretDecryptableResolver, error) {
	var ret Core_SecretDecryptableResolver
	var err error

	llog.Info(ctx, "Resolving query for Get...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretDecryptable().Get(ctx, args.Req.GetRequest)
	if err != nil {
		return nil, err
	}

	ret.SecretDecryptable = resp

	return &ret, nil
}

type Core_Task_GetArgs struct {
	Req *Core_Task_GetRequestResolver
}

func (r *QueryResolver) Core_Task_Get(ctx context.Context, args Core_Task_GetArgs) (*Core_TaskResolver, error) {
	var ret Core_TaskResolver
	var err error

	llog.Info(ctx, "Resolving query for Get...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().Task().Get(ctx, args.Req.GetRequest)
	if err != nil {
		return nil, err
	}

	ret.Task = resp

	return &ret, nil
}

type Core_TaskRun_GetArgs struct {
	Req *Core_TaskRun_GetRequestResolver
}

func (r *QueryResolver) Core_TaskRun_Get(ctx context.Context, args Core_TaskRun_GetArgs) (*Core_TaskRunResolver, error) {
	var ret Core_TaskRunResolver
	var err error

	llog.Info(ctx, "Resolving query for Get...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().TaskRun().Get(ctx, args.Req.GetRequest)
	if err != nil {
		return nil, err
	}

	ret.TaskRun = resp

	return &ret, nil
}

type User_Team_GetArgs struct {
	Req *User_Team_GetRequestResolver
}

func (r *QueryResolver) User_Team_Get(ctx context.Context, args User_Team_GetArgs) (*User_TeamResolver, error) {
	var ret User_TeamResolver
	var err error

	llog.Info(ctx, "Resolving query for Get...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.User().Team().Get(ctx, args.Req.GetRequest)
	if err != nil {
		return nil, err
	}

	ret.Team = resp

	return &ret, nil
}

type Core_JobApplicant_GetArgs struct {
	Req *Core_JobApplicant_GetRequestResolver
}

func (r *QueryResolver) Core_JobApplicant_Get(ctx context.Context, args Core_JobApplicant_GetArgs) (*Core_JobApplicantResolver, error) {
	var ret Core_JobApplicantResolver
	var err error

	llog.Info(ctx, "Resolving query for Get...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().JobApplicant().Get(ctx, args.Req.GetRequest)
	if err != nil {
		return nil, err
	}

	ret.JobApplicant = resp

	return &ret, nil
}

type Core_SecretKey_GetArgs struct {
	Req *Core_SecretKey_GetRequestResolver
}

func (r *QueryResolver) Core_SecretKey_Get(ctx context.Context, args Core_SecretKey_GetArgs) (*Core_SecretKeyResolver, error) {
	var ret Core_SecretKeyResolver
	var err error

	llog.Info(ctx, "Resolving query for Get...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretKey().Get(ctx, args.Req.GetRequest)
	if err != nil {
		return nil, err
	}

	ret.SecretKey = resp

	return &ret, nil
}

type User_User_GetArgs struct {
	Req *User_User_GetRequestResolver
}

func (r *QueryResolver) User_User_Get(ctx context.Context, args User_User_GetArgs) (*User_UserResolver, error) {
	var ret User_UserResolver
	var err error

	llog.Info(ctx, "Resolving query for Get...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.User().User().Get(ctx, args.Req.GetRequest)
	if err != nil {
		return nil, err
	}

	ret.User = resp

	return &ret, nil
}

type Core_File_GetArgs struct {
	Req *Core_File_GetRequestResolver
}

func (r *QueryResolver) Core_File_Get(ctx context.Context, args Core_File_GetArgs) (*Core_FileResolver, error) {
	var ret Core_FileResolver
	var err error

	llog.Info(ctx, "Resolving query for Get...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().File().Get(ctx, args.Req.GetRequest)
	if err != nil {
		return nil, err
	}

	ret.File = resp

	return &ret, nil
}

type Auth_Secret_GetArgs struct {
	Req *Auth_Secret_GetRequestResolver
}

func (r *QueryResolver) Auth_Secret_Get(ctx context.Context, args Auth_Secret_GetArgs) (*Auth_SecretResolver, error) {
	var ret Auth_SecretResolver
	var err error

	llog.Info(ctx, "Resolving query for Get...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Auth().Secret().Get(ctx, args.Req.GetRequest)
	if err != nil {
		return nil, err
	}

	ret.Secret = resp

	return &ret, nil
}

type User_User_ListArgs struct {
	Req *User_User_ListRequestResolver
}

func (r *QueryResolver) User_User_List(ctx context.Context, args User_User_ListArgs) (*User_User_ListResponseResolver, error) {
	var ret User_User_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for List...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.User().User().List(ctx, args.Req.ListRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Core_File_ListArgs struct {
	Req *Core_File_ListRequestResolver
}

func (r *QueryResolver) Core_File_List(ctx context.Context, args Core_File_ListArgs) (*Core_File_ListResponseResolver, error) {
	var ret Core_File_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for List...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().File().List(ctx, args.Req.ListRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Core_SecretKey_ListArgs struct {
	Req *Core_SecretKey_ListRequestResolver
}

func (r *QueryResolver) Core_SecretKey_List(ctx context.Context, args Core_SecretKey_ListArgs) (*Core_SecretKey_ListResponseResolver, error) {
	var ret Core_SecretKey_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for List...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretKey().List(ctx, args.Req.ListRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Core_TaskRun_ListArgs struct {
	Req *Core_TaskRun_ListRequestResolver
}

func (r *QueryResolver) Core_TaskRun_List(ctx context.Context, args Core_TaskRun_ListArgs) (*Core_TaskRun_ListResponseResolver, error) {
	var ret Core_TaskRun_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for List...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().TaskRun().List(ctx, args.Req.ListRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Core_JobApplicant_ListArgs struct {
	Req *Core_JobApplicant_ListRequestResolver
}

func (r *QueryResolver) Core_JobApplicant_List(ctx context.Context, args Core_JobApplicant_ListArgs) (*Core_JobApplicant_ListResponseResolver, error) {
	var ret Core_JobApplicant_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for List...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().JobApplicant().List(ctx, args.Req.ListRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Core_Task_ListArgs struct {
	Req *Core_Task_ListRequestResolver
}

func (r *QueryResolver) Core_Task_List(ctx context.Context, args Core_Task_ListArgs) (*Core_Task_ListResponseResolver, error) {
	var ret Core_Task_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for List...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().Task().List(ctx, args.Req.ListRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Core_SecretDecryptable_ListArgs struct {
	Req *Core_SecretDecryptable_ListRequestResolver
}

func (r *QueryResolver) Core_SecretDecryptable_List(ctx context.Context, args Core_SecretDecryptable_ListArgs) (*Core_SecretDecryptable_ListResponseResolver, error) {
	var ret Core_SecretDecryptable_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for List...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretDecryptable().List(ctx, args.Req.ListRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type User_Team_ListArgs struct {
	Req *User_Team_ListRequestResolver
}

func (r *QueryResolver) User_Team_List(ctx context.Context, args User_Team_ListArgs) (*User_Team_ListResponseResolver, error) {
	var ret User_Team_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for List...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.User().Team().List(ctx, args.Req.ListRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Auth_Secret_ListArgs struct {
	Req *Auth_Secret_ListRequestResolver
}

func (r *QueryResolver) Auth_Secret_List(ctx context.Context, args Auth_Secret_ListArgs) (*Auth_Secret_ListResponseResolver, error) {
	var ret Auth_Secret_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for List...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Auth().Secret().List(ctx, args.Req.ListRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Auth_LoginUserArgs struct {
	Req *Auth_LoginRequestResolver
}

func (r *QueryResolver) Auth_LoginUser(ctx context.Context, args Auth_LoginUserArgs) (*Auth_AuthenticateResponseResolver, error) {
	var ret Auth_AuthenticateResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for LoginUser...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Auth().LoginUser(ctx, args.Req.LoginRequest)
	if err != nil {
		return nil, err
	}

	ret.AuthenticateResponse = resp

	return &ret, nil
}

type Core_File_QueryByTextArgs struct {
	Req *Core_File_QueryByTextRequestResolver
}

func (r *QueryResolver) Core_File_QueryByText(ctx context.Context, args Core_File_QueryByTextArgs) (*Core_File_ListResponseResolver, error) {
	var ret Core_File_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for QueryByText...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().File().QueryByText(ctx, args.Req.QueryByTextRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Core_TaskRun_QueryByTextArgs struct {
	Req *Core_TaskRun_QueryByTextRequestResolver
}

func (r *QueryResolver) Core_TaskRun_QueryByText(ctx context.Context, args Core_TaskRun_QueryByTextArgs) (*Core_TaskRun_ListResponseResolver, error) {
	var ret Core_TaskRun_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for QueryByText...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().TaskRun().QueryByText(ctx, args.Req.QueryByTextRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Core_SecretDecryptable_QueryByTextArgs struct {
	Req *Core_SecretDecryptable_QueryByTextRequestResolver
}

func (r *QueryResolver) Core_SecretDecryptable_QueryByText(ctx context.Context, args Core_SecretDecryptable_QueryByTextArgs) (*Core_SecretDecryptable_ListResponseResolver, error) {
	var ret Core_SecretDecryptable_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for QueryByText...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretDecryptable().QueryByText(ctx, args.Req.QueryByTextRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Auth_Secret_QueryByTextArgs struct {
	Req *Auth_Secret_QueryByTextRequestResolver
}

func (r *QueryResolver) Auth_Secret_QueryByText(ctx context.Context, args Auth_Secret_QueryByTextArgs) (*Auth_Secret_ListResponseResolver, error) {
	var ret Auth_Secret_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for QueryByText...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Auth().Secret().QueryByText(ctx, args.Req.QueryByTextRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Core_Task_QueryByTextArgs struct {
	Req *Core_Task_QueryByTextRequestResolver
}

func (r *QueryResolver) Core_Task_QueryByText(ctx context.Context, args Core_Task_QueryByTextArgs) (*Core_Task_ListResponseResolver, error) {
	var ret Core_Task_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for QueryByText...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().Task().QueryByText(ctx, args.Req.QueryByTextRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Core_SecretKey_QueryByTextArgs struct {
	Req *Core_SecretKey_QueryByTextRequestResolver
}

func (r *QueryResolver) Core_SecretKey_QueryByText(ctx context.Context, args Core_SecretKey_QueryByTextArgs) (*Core_SecretKey_ListResponseResolver, error) {
	var ret Core_SecretKey_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for QueryByText...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretKey().QueryByText(ctx, args.Req.QueryByTextRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type User_User_QueryByTextArgs struct {
	Req *User_User_QueryByTextRequestResolver
}

func (r *QueryResolver) User_User_QueryByText(ctx context.Context, args User_User_QueryByTextArgs) (*User_User_ListResponseResolver, error) {
	var ret User_User_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for QueryByText...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.User().User().QueryByText(ctx, args.Req.QueryByTextRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type User_Team_QueryByTextArgs struct {
	Req *User_Team_QueryByTextRequestResolver
}

func (r *QueryResolver) User_Team_QueryByText(ctx context.Context, args User_Team_QueryByTextArgs) (*User_Team_ListResponseResolver, error) {
	var ret User_Team_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for QueryByText...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.User().Team().QueryByText(ctx, args.Req.QueryByTextRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type Core_JobApplicant_QueryByTextArgs struct {
	Req *Core_JobApplicant_QueryByTextRequestResolver
}

func (r *QueryResolver) Core_JobApplicant_QueryByText(ctx context.Context, args Core_JobApplicant_QueryByTextArgs) (*Core_JobApplicant_ListResponseResolver, error) {
	var ret Core_JobApplicant_ListResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for QueryByText...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().JobApplicant().QueryByText(ctx, args.Req.QueryByTextRequest)
	if err != nil {
		return nil, err
	}

	ret.ListResponse = resp

	return &ret, nil
}

type MutationResolver struct {
	getClientFn app_client.ClientGetterFn
}

type Core_SecretDecryptable_ActionDecryptArgs struct {
	Req *Core_SecretDecryptable_StandardEntityRequestResolver
}

func (r *MutationResolver) Core_SecretDecryptable_ActionDecrypt(ctx context.Context, args Core_SecretDecryptable_ActionDecryptArgs) (*Core_SecretDecryptable_StandardEntityResponseResolver, error) {
	var ret Core_SecretDecryptable_StandardEntityResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for ActionDecrypt...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretDecryptable().ActionDecrypt(ctx, args.Req.StandardEntityRequest)
	if err != nil {
		return nil, err
	}

	ret.StandardEntityResponse = resp

	return &ret, nil
}

type Core_Task_ActionRunArgs struct {
	Req *Core_Task_StandardEntityRequestResolver
}

func (r *MutationResolver) Core_Task_ActionRun(ctx context.Context, args Core_Task_ActionRunArgs) (*Core_Task_StandardEntityResponseResolver, error) {
	var ret Core_Task_StandardEntityResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for ActionRun...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().Task().ActionRun(ctx, args.Req.StandardEntityRequest)
	if err != nil {
		return nil, err
	}

	ret.StandardEntityResponse = resp

	return &ret, nil
}

type Core_TaskRun_ActionRunArgs struct {
	Req *Core_TaskRun_StandardEntityRequestResolver
}

func (r *MutationResolver) Core_TaskRun_ActionRun(ctx context.Context, args Core_TaskRun_ActionRunArgs) (*Core_TaskRun_StandardEntityResponseResolver, error) {
	var ret Core_TaskRun_StandardEntityResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for ActionRun...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().TaskRun().ActionRun(ctx, args.Req.StandardEntityRequest)
	if err != nil {
		return nil, err
	}

	ret.StandardEntityResponse = resp

	return &ret, nil
}

type Core_JobApplicant_AddArgs struct {
	Req *Core_JobApplicant_AddRequestResolver
}

func (r *MutationResolver) Core_JobApplicant_Add(ctx context.Context, args Core_JobApplicant_AddArgs) (*Core_JobApplicantResolver, error) {
	var ret Core_JobApplicantResolver
	var err error

	llog.Info(ctx, "Resolving query for Add...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().JobApplicant().Add(ctx, args.Req.AddRequest)
	if err != nil {
		return nil, err
	}

	ret.JobApplicant = resp

	return &ret, nil
}

type User_Team_AddArgs struct {
	Req *User_Team_AddRequestResolver
}

func (r *MutationResolver) User_Team_Add(ctx context.Context, args User_Team_AddArgs) (*User_TeamResolver, error) {
	var ret User_TeamResolver
	var err error

	llog.Info(ctx, "Resolving query for Add...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.User().Team().Add(ctx, args.Req.AddRequest)
	if err != nil {
		return nil, err
	}

	ret.Team = resp

	return &ret, nil
}

type Auth_Secret_AddArgs struct {
	Req *Auth_Secret_AddRequestResolver
}

func (r *MutationResolver) Auth_Secret_Add(ctx context.Context, args Auth_Secret_AddArgs) (*Auth_SecretResolver, error) {
	var ret Auth_SecretResolver
	var err error

	llog.Info(ctx, "Resolving query for Add...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Auth().Secret().Add(ctx, args.Req.AddRequest)
	if err != nil {
		return nil, err
	}

	ret.Secret = resp

	return &ret, nil
}

type Core_TaskRun_AddArgs struct {
	Req *Core_TaskRun_AddRequestResolver
}

func (r *MutationResolver) Core_TaskRun_Add(ctx context.Context, args Core_TaskRun_AddArgs) (*Core_TaskRunResolver, error) {
	var ret Core_TaskRunResolver
	var err error

	llog.Info(ctx, "Resolving query for Add...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().TaskRun().Add(ctx, args.Req.AddRequest)
	if err != nil {
		return nil, err
	}

	ret.TaskRun = resp

	return &ret, nil
}

type Core_File_AddArgs struct {
	Req *Core_File_AddRequestResolver
}

func (r *MutationResolver) Core_File_Add(ctx context.Context, args Core_File_AddArgs) (*Core_FileResolver, error) {
	var ret Core_FileResolver
	var err error

	llog.Info(ctx, "Resolving query for Add...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().File().Add(ctx, args.Req.AddRequest)
	if err != nil {
		return nil, err
	}

	ret.File = resp

	return &ret, nil
}

type Core_SecretDecryptable_AddArgs struct {
	Req *Core_SecretDecryptable_AddRequestResolver
}

func (r *MutationResolver) Core_SecretDecryptable_Add(ctx context.Context, args Core_SecretDecryptable_AddArgs) (*Core_SecretDecryptableResolver, error) {
	var ret Core_SecretDecryptableResolver
	var err error

	llog.Info(ctx, "Resolving query for Add...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretDecryptable().Add(ctx, args.Req.AddRequest)
	if err != nil {
		return nil, err
	}

	ret.SecretDecryptable = resp

	return &ret, nil
}

type Core_SecretKey_AddArgs struct {
	Req *Core_SecretKey_AddRequestResolver
}

func (r *MutationResolver) Core_SecretKey_Add(ctx context.Context, args Core_SecretKey_AddArgs) (*Core_SecretKeyResolver, error) {
	var ret Core_SecretKeyResolver
	var err error

	llog.Info(ctx, "Resolving query for Add...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretKey().Add(ctx, args.Req.AddRequest)
	if err != nil {
		return nil, err
	}

	ret.SecretKey = resp

	return &ret, nil
}

type User_User_AddArgs struct {
	Req *User_User_AddRequestResolver
}

func (r *MutationResolver) User_User_Add(ctx context.Context, args User_User_AddArgs) (*User_UserResolver, error) {
	var ret User_UserResolver
	var err error

	llog.Info(ctx, "Resolving query for Add...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.User().User().Add(ctx, args.Req.AddRequest)
	if err != nil {
		return nil, err
	}

	ret.User = resp

	return &ret, nil
}

type Core_Task_AddArgs struct {
	Req *Core_Task_AddRequestResolver
}

func (r *MutationResolver) Core_Task_Add(ctx context.Context, args Core_Task_AddArgs) (*Core_TaskResolver, error) {
	var ret Core_TaskResolver
	var err error

	llog.Info(ctx, "Resolving query for Add...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().Task().Add(ctx, args.Req.AddRequest)
	if err != nil {
		return nil, err
	}

	ret.Task = resp

	return &ret, nil
}

type Core_SecretDecryptable_HookCreatePreArgs struct {
	Req *Core_SecretDecryptableResolver
}

func (r *MutationResolver) Core_SecretDecryptable_HookCreatePre(ctx context.Context, args Core_SecretDecryptable_HookCreatePreArgs) (*Core_SecretDecryptableResolver, error) {
	var ret Core_SecretDecryptableResolver
	var err error

	llog.Info(ctx, "Resolving query for HookCreatePre...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretDecryptable().HookCreatePre(ctx, args.Req.SecretDecryptable)
	if err != nil {
		return nil, err
	}

	ret.SecretDecryptable = resp

	return &ret, nil
}

type Core_TaskRun_HookCreatePreArgs struct {
	Req *Core_TaskRunResolver
}

func (r *MutationResolver) Core_TaskRun_HookCreatePre(ctx context.Context, args Core_TaskRun_HookCreatePreArgs) (*Core_TaskRunResolver, error) {
	var ret Core_TaskRunResolver
	var err error

	llog.Info(ctx, "Resolving query for HookCreatePre...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().TaskRun().HookCreatePre(ctx, args.Req.TaskRun)
	if err != nil {
		return nil, err
	}

	ret.TaskRun = resp

	return &ret, nil
}

type Core_File_HookCreatePreArgs struct {
	Req *Core_FileResolver
}

func (r *MutationResolver) Core_File_HookCreatePre(ctx context.Context, args Core_File_HookCreatePreArgs) (*Core_FileResolver, error) {
	var ret Core_FileResolver
	var err error

	llog.Info(ctx, "Resolving query for HookCreatePre...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().File().HookCreatePre(ctx, args.Req.File)
	if err != nil {
		return nil, err
	}

	ret.File = resp

	return &ret, nil
}

type Core_SecretKey_HookCreatePreArgs struct {
	Req *Core_SecretKeyResolver
}

func (r *MutationResolver) Core_SecretKey_HookCreatePre(ctx context.Context, args Core_SecretKey_HookCreatePreArgs) (*Core_SecretKeyResolver, error) {
	var ret Core_SecretKeyResolver
	var err error

	llog.Info(ctx, "Resolving query for HookCreatePre...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretKey().HookCreatePre(ctx, args.Req.SecretKey)
	if err != nil {
		return nil, err
	}

	ret.SecretKey = resp

	return &ret, nil
}

type Core_Task_HookSavePreArgs struct {
	Req *Core_TaskResolver
}

func (r *MutationResolver) Core_Task_HookSavePre(ctx context.Context, args Core_Task_HookSavePreArgs) (*Core_TaskResolver, error) {
	var ret Core_TaskResolver
	var err error

	llog.Info(ctx, "Resolving query for HookSavePre...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().Task().HookSavePre(ctx, args.Req.Task)
	if err != nil {
		return nil, err
	}

	ret.Task = resp

	return &ret, nil
}

type Auth_RegisterUserArgs struct {
	Req *Auth_RegisterUserRequestResolver
}

func (r *MutationResolver) Auth_RegisterUser(ctx context.Context, args Auth_RegisterUserArgs) (*Auth_AuthenticateResponseResolver, error) {
	var ret Auth_AuthenticateResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for RegisterUser...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Auth().RegisterUser(ctx, args.Req.RegisterUserRequest)
	if err != nil {
		return nil, err
	}

	ret.AuthenticateResponse = resp

	return &ret, nil
}

type Core_SecretDecryptabeAddArgs struct {
	Req *Core_SecretDecryptableAddRequestResolver
}

func (r *MutationResolver) Core_SecretDecryptabeAdd(ctx context.Context, args Core_SecretDecryptabeAddArgs) (*Core_SecretDecryptableResolver, error) {
	var ret Core_SecretDecryptableResolver
	var err error

	llog.Info(ctx, "Resolving query for SecretDecryptabeAdd...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretDecryptabeAdd(ctx, args.Req.SecretDecryptableAddRequest)
	if err != nil {
		return nil, err
	}

	ret.SecretDecryptable = resp

	return &ret, nil
}

type Core_Task_UpdateArgs struct {
	Req *Core_Task_UpdateRequestResolver
}

func (r *MutationResolver) Core_Task_Update(ctx context.Context, args Core_Task_UpdateArgs) (*Core_Task_UpdateResponseResolver, error) {
	var ret Core_Task_UpdateResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for Update...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().Task().Update(ctx, args.Req.UpdateRequest)
	if err != nil {
		return nil, err
	}

	ret.UpdateResponse = resp

	return &ret, nil
}

type Core_File_UpdateArgs struct {
	Req *Core_File_UpdateRequestResolver
}

func (r *MutationResolver) Core_File_Update(ctx context.Context, args Core_File_UpdateArgs) (*Core_File_UpdateResponseResolver, error) {
	var ret Core_File_UpdateResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for Update...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().File().Update(ctx, args.Req.UpdateRequest)
	if err != nil {
		return nil, err
	}

	ret.UpdateResponse = resp

	return &ret, nil
}

type Core_JobApplicant_UpdateArgs struct {
	Req *Core_JobApplicant_UpdateRequestResolver
}

func (r *MutationResolver) Core_JobApplicant_Update(ctx context.Context, args Core_JobApplicant_UpdateArgs) (*Core_JobApplicant_UpdateResponseResolver, error) {
	var ret Core_JobApplicant_UpdateResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for Update...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().JobApplicant().Update(ctx, args.Req.UpdateRequest)
	if err != nil {
		return nil, err
	}

	ret.UpdateResponse = resp

	return &ret, nil
}

type Core_SecretKey_UpdateArgs struct {
	Req *Core_SecretKey_UpdateRequestResolver
}

func (r *MutationResolver) Core_SecretKey_Update(ctx context.Context, args Core_SecretKey_UpdateArgs) (*Core_SecretKey_UpdateResponseResolver, error) {
	var ret Core_SecretKey_UpdateResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for Update...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretKey().Update(ctx, args.Req.UpdateRequest)
	if err != nil {
		return nil, err
	}

	ret.UpdateResponse = resp

	return &ret, nil
}

type User_Team_UpdateArgs struct {
	Req *User_Team_UpdateRequestResolver
}

func (r *MutationResolver) User_Team_Update(ctx context.Context, args User_Team_UpdateArgs) (*User_Team_UpdateResponseResolver, error) {
	var ret User_Team_UpdateResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for Update...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.User().Team().Update(ctx, args.Req.UpdateRequest)
	if err != nil {
		return nil, err
	}

	ret.UpdateResponse = resp

	return &ret, nil
}

type Core_TaskRun_UpdateArgs struct {
	Req *Core_TaskRun_UpdateRequestResolver
}

func (r *MutationResolver) Core_TaskRun_Update(ctx context.Context, args Core_TaskRun_UpdateArgs) (*Core_TaskRun_UpdateResponseResolver, error) {
	var ret Core_TaskRun_UpdateResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for Update...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().TaskRun().Update(ctx, args.Req.UpdateRequest)
	if err != nil {
		return nil, err
	}

	ret.UpdateResponse = resp

	return &ret, nil
}

type User_User_UpdateArgs struct {
	Req *User_User_UpdateRequestResolver
}

func (r *MutationResolver) User_User_Update(ctx context.Context, args User_User_UpdateArgs) (*User_User_UpdateResponseResolver, error) {
	var ret User_User_UpdateResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for Update...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.User().User().Update(ctx, args.Req.UpdateRequest)
	if err != nil {
		return nil, err
	}

	ret.UpdateResponse = resp

	return &ret, nil
}

type Core_SecretDecryptable_UpdateArgs struct {
	Req *Core_SecretDecryptable_UpdateRequestResolver
}

func (r *MutationResolver) Core_SecretDecryptable_Update(ctx context.Context, args Core_SecretDecryptable_UpdateArgs) (*Core_SecretDecryptable_UpdateResponseResolver, error) {
	var ret Core_SecretDecryptable_UpdateResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for Update...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Core().SecretDecryptable().Update(ctx, args.Req.UpdateRequest)
	if err != nil {
		return nil, err
	}

	ret.UpdateResponse = resp

	return &ret, nil
}

type Auth_Secret_UpdateArgs struct {
	Req *Auth_Secret_UpdateRequestResolver
}

func (r *MutationResolver) Auth_Secret_Update(ctx context.Context, args Auth_Secret_UpdateArgs) (*Auth_Secret_UpdateResponseResolver, error) {
	var ret Auth_Secret_UpdateResponseResolver
	var err error

	llog.Info(ctx, "Resolving query for Update...")

	// Get the method's server for this service
	c := r.getClientFn(ctx)

	resp, err := c.Auth().Secret().Update(ctx, args.Req.UpdateRequest)
	if err != nil {
		return nil, err
	}

	ret.UpdateResponse = resp

	return &ret, nil
}
