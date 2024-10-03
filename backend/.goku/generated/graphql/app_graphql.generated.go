package app_graphql

import (
	"context"

	"github.com/teejays/gokutil/log"

	app_client "sampleapp/backend/.goku/generated/client"
	svcauth_entsecret_typ "sampleapp/backend/.goku/generated/service/auth/entity/secret/typ"
	svcauth_typ "sampleapp/backend/.goku/generated/service/auth/typ"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
	svccore_entjobapplicant_typ "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/typ"
	svcuser_entteam_typ "sampleapp/backend/.goku/generated/service/user/entity/team/typ"
	svcuser_entuser_typ "sampleapp/backend/.goku/generated/service/user/entity/user/typ"
	app_typ "sampleapp/backend/.goku/generated/typ"
)

var llog = log.GetLogger().WithHeading("Graphql Resolver")

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

type GenderConditionResolver struct {
	app_typ.GenderCondition
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

type Auth_LoginRequestResolver struct {
	svcauth_typ.LoginRequest
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

type User_Team_StandardEntityRequestResolver struct {
	svcuser_entteam_typ.StandardEntityRequest
}

type User_User_StandardEntityRequestResolver struct {
	svcuser_entuser_typ.StandardEntityRequest
}

type Core_File_StandardEntityRequestResolver struct {
	svccore_entfile_typ.StandardEntityRequest
}

type Core_JobApplicant_StandardEntityRequestResolver struct {
	svccore_entjobapplicant_typ.StandardEntityRequest
}

type Auth_Secret_StandardEntityRequestResolver struct {
	svcauth_entsecret_typ.StandardEntityRequest
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

type Core_File_StandardEntityResponseResolver struct {
	svccore_entfile_typ.StandardEntityResponse
}

func (r *Core_File_StandardEntityResponseResolver) Object() *Core_FileResolver {
	return &Core_FileResolver{File: r.StandardEntityResponse.Object}
}

type Core_JobApplicant_StandardEntityResponseResolver struct {
	svccore_entjobapplicant_typ.StandardEntityResponse
}

func (r *Core_JobApplicant_StandardEntityResponseResolver) Object() *Core_JobApplicantResolver {
	return &Core_JobApplicantResolver{JobApplicant: r.StandardEntityResponse.Object}
}

type Auth_Secret_StandardEntityResponseResolver struct {
	svcauth_entsecret_typ.StandardEntityResponse
}

func (r *Auth_Secret_StandardEntityResponseResolver) Object() *Auth_SecretResolver {
	return &Auth_SecretResolver{Secret: r.StandardEntityResponse.Object}
}

type Core_File_StorageClientConditionResolver struct {
	svccore_entfile_typ.StorageClientCondition
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

type Auth_Secret_TypeConditionResolver struct {
	svcauth_entsecret_typ.TypeCondition
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

type Core_File_AddRequestResolver struct {
	svccore_entfile_typ.AddRequest
}

func (r *Core_File_AddRequestResolver) Object() *Core_File_FileInputResolver {
	return &Core_File_FileInputResolver{FileInput: r.AddRequest.Object}
}

type Core_File_UpdateRequestResolver struct {
	svccore_entfile_typ.UpdateRequest
}

func (r *Core_File_UpdateRequestResolver) Object() *Core_FileResolver {
	return &Core_FileResolver{File: r.UpdateRequest.Object}
}

type Core_File_UpdateResponseResolver struct {
	svccore_entfile_typ.UpdateResponse
}

func (r *Core_File_UpdateResponseResolver) Object() *Core_FileResolver {
	return &Core_FileResolver{File: r.UpdateResponse.Object}
}

type Core_File_GetRequestResolver struct {
	svccore_entfile_typ.GetRequest
}

type Core_File_ListRequestResolver struct {
	svccore_entfile_typ.ListRequest
}

func (r *Core_File_ListRequestResolver) Filter() *Core_File_FileFilterResolver {
	return &Core_File_FileFilterResolver{FileFilter: r.ListRequest.Filter}
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

type Core_File_QueryByTextRequestResolver struct {
	svccore_entfile_typ.QueryByTextRequest
}

type Core_JobApplicant_AddRequestResolver struct {
	svccore_entjobapplicant_typ.AddRequest
}

func (r *Core_JobApplicant_AddRequestResolver) Object() *Core_JobApplicant_JobApplicantInputResolver {
	return &Core_JobApplicant_JobApplicantInputResolver{JobApplicantInput: r.AddRequest.Object}
}

type Core_JobApplicant_UpdateRequestResolver struct {
	svccore_entjobapplicant_typ.UpdateRequest
}

func (r *Core_JobApplicant_UpdateRequestResolver) Object() *Core_JobApplicantResolver {
	return &Core_JobApplicantResolver{JobApplicant: r.UpdateRequest.Object}
}

type Core_JobApplicant_UpdateResponseResolver struct {
	svccore_entjobapplicant_typ.UpdateResponse
}

func (r *Core_JobApplicant_UpdateResponseResolver) Object() *Core_JobApplicantResolver {
	return &Core_JobApplicantResolver{JobApplicant: r.UpdateResponse.Object}
}

type Core_JobApplicant_GetRequestResolver struct {
	svccore_entjobapplicant_typ.GetRequest
}

type Core_JobApplicant_ListRequestResolver struct {
	svccore_entjobapplicant_typ.ListRequest
}

func (r *Core_JobApplicant_ListRequestResolver) Filter() *Core_JobApplicant_JobApplicantFilterResolver {
	return &Core_JobApplicant_JobApplicantFilterResolver{JobApplicantFilter: r.ListRequest.Filter}
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

type Core_JobApplicant_QueryByTextRequestResolver struct {
	svccore_entjobapplicant_typ.QueryByTextRequest
}

type Auth_Secret_AddRequestResolver struct {
	svcauth_entsecret_typ.AddRequest
}

func (r *Auth_Secret_AddRequestResolver) Object() *Auth_Secret_SecretInputResolver {
	return &Auth_Secret_SecretInputResolver{SecretInput: r.AddRequest.Object}
}

type Auth_Secret_UpdateRequestResolver struct {
	svcauth_entsecret_typ.UpdateRequest
}

func (r *Auth_Secret_UpdateRequestResolver) Object() *Auth_SecretResolver {
	return &Auth_SecretResolver{Secret: r.UpdateRequest.Object}
}

type Auth_Secret_UpdateResponseResolver struct {
	svcauth_entsecret_typ.UpdateResponse
}

func (r *Auth_Secret_UpdateResponseResolver) Object() *Auth_SecretResolver {
	return &Auth_SecretResolver{Secret: r.UpdateResponse.Object}
}

type Auth_Secret_GetRequestResolver struct {
	svcauth_entsecret_typ.GetRequest
}

type Auth_Secret_ListRequestResolver struct {
	svcauth_entsecret_typ.ListRequest
}

func (r *Auth_Secret_ListRequestResolver) Filter() *Auth_Secret_SecretFilterResolver {
	return &Auth_Secret_SecretFilterResolver{SecretFilter: r.ListRequest.Filter}
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

type Auth_Secret_QueryByTextRequestResolver struct {
	svcauth_entsecret_typ.QueryByTextRequest
}

type User_Team_AddRequestResolver struct {
	svcuser_entteam_typ.AddRequest
}

func (r *User_Team_AddRequestResolver) Object() *User_Team_TeamInputResolver {
	return &User_Team_TeamInputResolver{TeamInput: r.AddRequest.Object}
}

type User_Team_UpdateRequestResolver struct {
	svcuser_entteam_typ.UpdateRequest
}

func (r *User_Team_UpdateRequestResolver) Object() *User_TeamResolver {
	return &User_TeamResolver{Team: r.UpdateRequest.Object}
}

type User_Team_UpdateResponseResolver struct {
	svcuser_entteam_typ.UpdateResponse
}

func (r *User_Team_UpdateResponseResolver) Object() *User_TeamResolver {
	return &User_TeamResolver{Team: r.UpdateResponse.Object}
}

type User_Team_GetRequestResolver struct {
	svcuser_entteam_typ.GetRequest
}

type User_Team_ListRequestResolver struct {
	svcuser_entteam_typ.ListRequest
}

func (r *User_Team_ListRequestResolver) Filter() *User_Team_TeamFilterResolver {
	return &User_Team_TeamFilterResolver{TeamFilter: r.ListRequest.Filter}
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

type User_Team_QueryByTextRequestResolver struct {
	svcuser_entteam_typ.QueryByTextRequest
}

type User_User_AddRequestResolver struct {
	svcuser_entuser_typ.AddRequest
}

func (r *User_User_AddRequestResolver) Object() *User_User_UserInputResolver {
	return &User_User_UserInputResolver{UserInput: r.AddRequest.Object}
}

type User_User_UpdateRequestResolver struct {
	svcuser_entuser_typ.UpdateRequest
}

func (r *User_User_UpdateRequestResolver) Object() *User_UserResolver {
	return &User_UserResolver{User: r.UpdateRequest.Object}
}

type User_User_UpdateResponseResolver struct {
	svcuser_entuser_typ.UpdateResponse
}

func (r *User_User_UpdateResponseResolver) Object() *User_UserResolver {
	return &User_UserResolver{User: r.UpdateResponse.Object}
}

type User_User_GetRequestResolver struct {
	svcuser_entuser_typ.GetRequest
}

type User_User_ListRequestResolver struct {
	svcuser_entuser_typ.ListRequest
}

func (r *User_User_ListRequestResolver) Filter() *User_User_UserFilterResolver {
	return &User_User_UserFilterResolver{UserFilter: r.ListRequest.Filter}
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

type User_User_QueryByTextRequestResolver struct {
	svcuser_entuser_typ.QueryByTextRequest
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

type Auth_LoginRequestArgs struct {
	Req *Auth_LoginRequestResolver
}

func (r *QueryResolver) Auth_LoginUser(ctx context.Context, args Auth_LoginRequestArgs) (*Auth_AuthenticateResponseResolver, error) {
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

type Auth_AuthenticateTokenRequestArgs struct {
	Req *Auth_AuthenticateTokenRequestResolver
}

func (r *QueryResolver) Auth_AuthenticateToken(ctx context.Context, args Auth_AuthenticateTokenRequestArgs) (*Auth_AuthenticateResponseResolver, error) {
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

type Core_File_GetRequestArgs struct {
	Req *Core_File_GetRequestResolver
}

func (r *QueryResolver) Core_File_Get(ctx context.Context, args Core_File_GetRequestArgs) (*Core_FileResolver, error) {
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

type Core_File_ListRequestArgs struct {
	Req *Core_File_ListRequestResolver
}

func (r *QueryResolver) Core_File_List(ctx context.Context, args Core_File_ListRequestArgs) (*Core_File_ListResponseResolver, error) {
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

type Core_File_QueryByTextRequestArgs struct {
	Req *Core_File_QueryByTextRequestResolver
}

func (r *QueryResolver) Core_File_QueryByText(ctx context.Context, args Core_File_QueryByTextRequestArgs) (*Core_File_ListResponseResolver, error) {
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

type Core_JobApplicant_GetRequestArgs struct {
	Req *Core_JobApplicant_GetRequestResolver
}

func (r *QueryResolver) Core_JobApplicant_Get(ctx context.Context, args Core_JobApplicant_GetRequestArgs) (*Core_JobApplicantResolver, error) {
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

type Core_JobApplicant_ListRequestArgs struct {
	Req *Core_JobApplicant_ListRequestResolver
}

func (r *QueryResolver) Core_JobApplicant_List(ctx context.Context, args Core_JobApplicant_ListRequestArgs) (*Core_JobApplicant_ListResponseResolver, error) {
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

type Core_JobApplicant_QueryByTextRequestArgs struct {
	Req *Core_JobApplicant_QueryByTextRequestResolver
}

func (r *QueryResolver) Core_JobApplicant_QueryByText(ctx context.Context, args Core_JobApplicant_QueryByTextRequestArgs) (*Core_JobApplicant_ListResponseResolver, error) {
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

type Auth_Secret_GetRequestArgs struct {
	Req *Auth_Secret_GetRequestResolver
}

func (r *QueryResolver) Auth_Secret_Get(ctx context.Context, args Auth_Secret_GetRequestArgs) (*Auth_SecretResolver, error) {
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

type Auth_Secret_ListRequestArgs struct {
	Req *Auth_Secret_ListRequestResolver
}

func (r *QueryResolver) Auth_Secret_List(ctx context.Context, args Auth_Secret_ListRequestArgs) (*Auth_Secret_ListResponseResolver, error) {
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

type Auth_Secret_QueryByTextRequestArgs struct {
	Req *Auth_Secret_QueryByTextRequestResolver
}

func (r *QueryResolver) Auth_Secret_QueryByText(ctx context.Context, args Auth_Secret_QueryByTextRequestArgs) (*Auth_Secret_ListResponseResolver, error) {
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

type User_Team_GetRequestArgs struct {
	Req *User_Team_GetRequestResolver
}

func (r *QueryResolver) User_Team_Get(ctx context.Context, args User_Team_GetRequestArgs) (*User_TeamResolver, error) {
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

type User_Team_ListRequestArgs struct {
	Req *User_Team_ListRequestResolver
}

func (r *QueryResolver) User_Team_List(ctx context.Context, args User_Team_ListRequestArgs) (*User_Team_ListResponseResolver, error) {
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

type User_Team_QueryByTextRequestArgs struct {
	Req *User_Team_QueryByTextRequestResolver
}

func (r *QueryResolver) User_Team_QueryByText(ctx context.Context, args User_Team_QueryByTextRequestArgs) (*User_Team_ListResponseResolver, error) {
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

type User_User_GetRequestArgs struct {
	Req *User_User_GetRequestResolver
}

func (r *QueryResolver) User_User_Get(ctx context.Context, args User_User_GetRequestArgs) (*User_UserResolver, error) {
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

type User_User_ListRequestArgs struct {
	Req *User_User_ListRequestResolver
}

func (r *QueryResolver) User_User_List(ctx context.Context, args User_User_ListRequestArgs) (*User_User_ListResponseResolver, error) {
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

type User_User_QueryByTextRequestArgs struct {
	Req *User_User_QueryByTextRequestResolver
}

func (r *QueryResolver) User_User_QueryByText(ctx context.Context, args User_User_QueryByTextRequestArgs) (*User_User_ListResponseResolver, error) {
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

type MutationResolver struct {
	getClientFn app_client.ClientGetterFn
}

type Auth_RegisterUserRequestArgs struct {
	Req *Auth_RegisterUserRequestResolver
}

func (r *MutationResolver) Auth_RegisterUser(ctx context.Context, args Auth_RegisterUserRequestArgs) (*Auth_AuthenticateResponseResolver, error) {
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

type Core_File_AddRequestArgs struct {
	Req *Core_File_AddRequestResolver
}

func (r *MutationResolver) Core_File_Add(ctx context.Context, args Core_File_AddRequestArgs) (*Core_FileResolver, error) {
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

type Core_File_UpdateRequestArgs struct {
	Req *Core_File_UpdateRequestResolver
}

func (r *MutationResolver) Core_File_Update(ctx context.Context, args Core_File_UpdateRequestArgs) (*Core_File_UpdateResponseResolver, error) {
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

type Core_JobApplicant_AddRequestArgs struct {
	Req *Core_JobApplicant_AddRequestResolver
}

func (r *MutationResolver) Core_JobApplicant_Add(ctx context.Context, args Core_JobApplicant_AddRequestArgs) (*Core_JobApplicantResolver, error) {
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

type Core_JobApplicant_UpdateRequestArgs struct {
	Req *Core_JobApplicant_UpdateRequestResolver
}

func (r *MutationResolver) Core_JobApplicant_Update(ctx context.Context, args Core_JobApplicant_UpdateRequestArgs) (*Core_JobApplicant_UpdateResponseResolver, error) {
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

type Auth_Secret_AddRequestArgs struct {
	Req *Auth_Secret_AddRequestResolver
}

func (r *MutationResolver) Auth_Secret_Add(ctx context.Context, args Auth_Secret_AddRequestArgs) (*Auth_SecretResolver, error) {
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

type Auth_Secret_UpdateRequestArgs struct {
	Req *Auth_Secret_UpdateRequestResolver
}

func (r *MutationResolver) Auth_Secret_Update(ctx context.Context, args Auth_Secret_UpdateRequestArgs) (*Auth_Secret_UpdateResponseResolver, error) {
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

type User_Team_AddRequestArgs struct {
	Req *User_Team_AddRequestResolver
}

func (r *MutationResolver) User_Team_Add(ctx context.Context, args User_Team_AddRequestArgs) (*User_TeamResolver, error) {
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

type User_Team_UpdateRequestArgs struct {
	Req *User_Team_UpdateRequestResolver
}

func (r *MutationResolver) User_Team_Update(ctx context.Context, args User_Team_UpdateRequestArgs) (*User_Team_UpdateResponseResolver, error) {
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

type User_User_AddRequestArgs struct {
	Req *User_User_AddRequestResolver
}

func (r *MutationResolver) User_User_Add(ctx context.Context, args User_User_AddRequestArgs) (*User_UserResolver, error) {
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

type User_User_UpdateRequestArgs struct {
	Req *User_User_UpdateRequestResolver
}

func (r *MutationResolver) User_User_Update(ctx context.Context, args User_User_UpdateRequestArgs) (*User_User_UpdateResponseResolver, error) {
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
