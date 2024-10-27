package app_generic

import (
	typeslib "github.com/teejays/gokutil/types"

	svcauth_entsecret_typ "sampleapp/backend/.goku/generated/service/auth/entity/secret/typ"
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

type Entity interface {
	svccore_entfile_typ.File | svccore_entjobapplicant_typ.JobApplicant | svcauth_entsecret_typ.Secret | svccore_entsecretdecryptable_typ.SecretDecryptable | svccore_entsecretkey_typ.SecretKey | svccore_enttask_typ.Task | svccore_enttaskrun_typ.TaskRun | svcuser_entteam_typ.Team | svcuser_entuser_typ.User
}

type BasicType interface {
	app_typ.Address | app_typ.Contact | svccore_typ.CronExpression | app_typ.PersonName | app_typ.PhoneNumber | svccore_typ.SecretDecryptableAddRequest
}

type BasicTypeWithMeta interface {
	app_typ.AddressWithMeta | app_typ.ContactWithMeta | svccore_typ.CronExpressionWithMeta | app_typ.PersonNameWithMeta | app_typ.PhoneNumberWithMeta | svccore_typ.SecretDecryptableAddRequestWithMeta
}

type FilterType interface {
	app_typ.AddressFilter | app_typ.ContactFilter | svccore_typ.CronExpressionFilter | svccore_entfile_typ.FileFilter | svccore_entjobapplicant_typ.JobApplicantFilter | app_typ.PersonNameFilter | app_typ.PhoneNumberFilter | svccore_typ.SecretDecryptableAddRequestFilter | svccore_entsecretdecryptable_typ.SecretDecryptableFilter | svcauth_entsecret_typ.SecretFilter | svccore_entsecretkey_typ.SecretKeyFilter | svccore_enttask_typ.TaskFilter | svccore_enttaskrun_typ.TaskRunFilter | svcuser_entteam_typ.TeamFilter | svcuser_entuser_typ.UserFilter
}

type FieldEnum interface {
	app_typ.AddressField | app_typ.ContactField | svccore_typ.CronExpressionField | svccore_entfile_typ.FileField | svccore_entjobapplicant_typ.JobApplicantField | app_typ.PersonNameField | app_typ.PhoneNumberField | svcauth_entsecret_typ.SecretField | svccore_entsecretdecryptable_typ.SecretDecryptableField | svccore_typ.SecretDecryptableAddRequestField | svccore_entsecretkey_typ.SecretKeyField | svccore_enttask_typ.TaskField | svccore_enttaskrun_typ.TaskRunField | svcuser_entteam_typ.TeamField | svcuser_entuser_typ.UserField
	typeslib.Field
}
