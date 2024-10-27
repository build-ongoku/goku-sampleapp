/*eslint-disable */

import { EntityMinimal } from '@ongoku/app-lib/src/common/Entity'
import { convertTypeToWithMeta } from '@ongoku/app-lib/src/common/types'
import * as field from '@ongoku/app-lib/src/common/Field'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

import * as app_typ from '@sampleapp/goku-generated/typ/typ.generated'
import * as svccore_entfile_typ from '@sampleapp/goku-generated/typ/core/file/typ.generated'
import * as svccore_entjobapplicant_typ from '@sampleapp/goku-generated/typ/core/job_applicant/typ.generated'
import * as svccore_entsecretdecryptable_typ from '@sampleapp/goku-generated/typ/core/secret_decryptable/typ.generated'
import * as svccore_entsecretkey_typ from '@sampleapp/goku-generated/typ/core/secret_key/typ.generated'
import * as svccore_enttask_typ from '@sampleapp/goku-generated/typ/core/task/typ.generated'
import * as svccore_enttaskrun_typ from '@sampleapp/goku-generated/typ/core/task_run/typ.generated'

/*eslint-disable */

/**
 * Enums
 */
// CronExpressionField: <description of the enum>
export type CronExpressionField = 'ParentID' | 'ID' | 'Second' | 'Minute' | 'Hour' | 'DayOfMonth' | 'Month' | 'DayOfWeek' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | never

export type CronExpressionFieldCondition = GenericCondition<CronExpressionField>

// EntityName: <description of the enum>
export type EntityName = 'JobApplicant' | 'File' | 'SecretKey' | 'SecretDecryptable' | 'Task' | 'TaskRun' | never

export type EntityNameCondition = GenericCondition<EntityName>

// MethodName: <description of the enum>
export type MethodName = 'FileUpload' | 'SecretDecryptabeAdd' | never

export type MethodNameCondition = GenericCondition<MethodName>

// SecretDecryptableAddRequestField: <description of the enum>
export type SecretDecryptableAddRequestField = 'ParentID' | 'ID' | 'Value' | 'SecretKeyID' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | never

export type SecretDecryptableAddRequestFieldCondition = GenericCondition<SecretDecryptableAddRequestField>

/**
 * Local Types, Type Info Props, Type Infos
 */

// CronExpression: <description of the type>
export type CronExpression = {
    second: string
    minute: string
    hour: string
    dayOfMonth: string
    month: string
    dayOfWeek: string
}
export type CronExpressionWithMeta = CronExpression & field.MetaFields

// CronExpressionInput: <description of the type>
export type CronExpressionInput = {
    second: string
    minute: string
    hour: string
    dayOfMonth: string
    month: string
    dayOfWeek: string
}
export type CronExpressionInputWithMeta = CronExpressionInput & field.MetaFields

// SecretDecryptableAddRequest: <description of the type>
export type SecretDecryptableAddRequest = {
    value: string
    secretKeyID: scalars.ID | null
}
export type SecretDecryptableAddRequestWithMeta = SecretDecryptableAddRequest & field.MetaFields

// SecretDecryptableAddRequestInput: <description of the type>
export type SecretDecryptableAddRequestInput = {
    value: string
    secretKeyID: scalars.ID | null
}
export type SecretDecryptableAddRequestInputWithMeta = SecretDecryptableAddRequestInput & field.MetaFields

/**
 * Entity Type, Entity Info Props
 */
// No Entities

/**
 * Filter Types
 */
// CronExpressionFilter: <description of the type>
export type CronExpressionFilter = {
    parentID: IDCondition | null
    id: IDCondition | null
    second: StringCondition | null
    minute: StringCondition | null
    hour: StringCondition | null
    dayOfMonth: StringCondition | null
    month: StringCondition | null
    dayOfWeek: StringCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: CronExpressionFilter[] | null
    or: CronExpressionFilter[] | null
}
// SecretDecryptableAddRequestFilter: <description of the type>
export type SecretDecryptableAddRequestFilter = {
    parentID: IDCondition | null
    id: IDCondition | null
    value: StringCondition | null
    secretKeyID: IDCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: SecretDecryptableAddRequestFilter[] | null
    or: SecretDecryptableAddRequestFilter[] | null
}

/**
 * Standard Request / Response Types
 */
