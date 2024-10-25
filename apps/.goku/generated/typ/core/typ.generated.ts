/*eslint-disable */

import { EntityMinimal } from '@ongoku/app-lib/src/common/Entity'
import { convertTypeToWithMeta } from '@ongoku/app-lib/src/common/types'
import * as field from '@ongoku/app-lib/src/common/Field'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

import * as app_typ from '@sampleapp/goku-generated/typ/typ.generated'
import * as svccore_entfile_typ from '@sampleapp/goku-generated/typ/core/file/typ.generated'
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
export type EntityName = 'File' | 'Task' | 'TaskRun' | never

export type EntityNameCondition = GenericCondition<EntityName>

// MethodName: <description of the enum>
export type MethodName = 'FileUpload' | never

export type MethodNameCondition = GenericCondition<MethodName>

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

/**
 * Standard Request / Response Types
 */
