package app_generic

import (
	typeslib "github.com/teejays/gokutil/types"

	svcauth_entsecret_typ "sampleapp/backend/.goku/generated/service/auth/entity/secret/typ"
	svccore_entfile_typ "sampleapp/backend/.goku/generated/service/core/entity/file/typ"
	svccore_enttask_typ "sampleapp/backend/.goku/generated/service/core/entity/task/typ"
	svccore_enttaskrun_typ "sampleapp/backend/.goku/generated/service/core/entity/task_run/typ"
	svccore_typ "sampleapp/backend/.goku/generated/service/core/typ"
	svcuser_entteam_typ "sampleapp/backend/.goku/generated/service/user/entity/team/typ"
	svcuser_entuser_typ "sampleapp/backend/.goku/generated/service/user/entity/user/typ"
	app_typ "sampleapp/backend/.goku/generated/typ"
)

type Entity interface {
	svccore_entfile_typ.File | svcauth_entsecret_typ.Secret | svccore_enttask_typ.Task | svccore_enttaskrun_typ.TaskRun | svcuser_entteam_typ.Team | svcuser_entuser_typ.User
}

type BasicType interface {
	app_typ.Address | app_typ.Contact | svccore_typ.CronExpression | app_typ.PersonName | app_typ.PhoneNumber
}

type BasicTypeWithMeta interface {
	app_typ.AddressWithMeta | app_typ.ContactWithMeta | svccore_typ.CronExpressionWithMeta | app_typ.PersonNameWithMeta | app_typ.PhoneNumberWithMeta
}

type FilterType interface {
	app_typ.AddressFilter | app_typ.ContactFilter | svccore_typ.CronExpressionFilter | svccore_entfile_typ.FileFilter | app_typ.PersonNameFilter | app_typ.PhoneNumberFilter | svcauth_entsecret_typ.SecretFilter | svccore_enttask_typ.TaskFilter | svccore_enttaskrun_typ.TaskRunFilter | svcuser_entteam_typ.TeamFilter | svcuser_entuser_typ.UserFilter
}

type FieldEnum interface {
	app_typ.AddressField | app_typ.ContactField | svccore_typ.CronExpressionField | svccore_entfile_typ.FileField | app_typ.PersonNameField | app_typ.PhoneNumberField | svcauth_entsecret_typ.SecretField | svccore_enttask_typ.TaskField | svccore_enttaskrun_typ.TaskRunField | svcuser_entteam_typ.TeamField | svcuser_entuser_typ.UserField
	typeslib.Field
}
