package app_generic

import (
	typeslib "github.com/teejays/gokutil/types"

	svcauth_entsecret_typ "sampleapp/backend/.goku/generated/service/auth/entity/secret/typ"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
	svccore_entjobapplicant_typ "sampleapp/backend/.goku/generated/service/core/entity/job_applicant/typ"
	svcuser_entteam_typ "sampleapp/backend/.goku/generated/service/user/entity/team/typ"
	svcuser_entuser_typ "sampleapp/backend/.goku/generated/service/user/entity/user/typ"
	app_typ "sampleapp/backend/.goku/generated/typ"
)

type Entity interface {
	svccore_entfile_typ.File | svccore_entjobapplicant_typ.JobApplicant | svcauth_entsecret_typ.Secret | svcuser_entteam_typ.Team | svcuser_entuser_typ.User
}

type BasicType interface {
	app_typ.Address | app_typ.Contact | app_typ.PersonName | app_typ.PhoneNumber
}

type BasicTypeWithMeta interface {
	app_typ.AddressWithMeta | app_typ.ContactWithMeta | app_typ.PersonNameWithMeta | app_typ.PhoneNumberWithMeta
}

type FilterType interface {
	app_typ.AddressFilter | app_typ.ContactFilter | svccore_entfile_typ.FileFilter | svccore_entjobapplicant_typ.JobApplicantFilter | app_typ.PersonNameFilter | app_typ.PhoneNumberFilter | svcauth_entsecret_typ.SecretFilter | svcuser_entteam_typ.TeamFilter | svcuser_entuser_typ.UserFilter
}

type FieldEnum interface {
	app_typ.AddressField | app_typ.ContactField | svccore_entfile_typ.FileField | svccore_entjobapplicant_typ.JobApplicantField | app_typ.PersonNameField | app_typ.PhoneNumberField | svcauth_entsecret_typ.SecretField | svcuser_entteam_typ.TeamField | svcuser_entuser_typ.UserField
	typeslib.Field
}
