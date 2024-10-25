/*eslint-disable */

import { EntityMinimal } from '@ongoku/app-lib/src/common/Entity'
import { convertTypeToWithMeta } from '@ongoku/app-lib/src/common/types'
import * as field from '@ongoku/app-lib/src/common/Field'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

import * as app_typ from '@sampleapp/goku-generated/typ/typ.generated'
import * as svccore_typ from '@sampleapp/goku-generated/typ/core/typ.generated'

/*eslint-disable */

/**
 * Enums
 */
// MethodName: <description of the enum>
export type MethodName = 'Add' | 'Update' | 'Get' | 'List' | 'QueryByText' | 'HookCreatePre' | 'ActionRun' | never

export type MethodNameCondition = GenericCondition<MethodName>

// Status: <description of the enum>
export type Status = 'Pending' | 'Running' | 'Success' | 'Failed' | never

export type StatusCondition = GenericCondition<Status>

// TaskRunAction: <description of the enum>
export type TaskRunAction = 'Run' | never

export type TaskRunActionCondition = GenericCondition<TaskRunAction>

// TaskRunField: <description of the enum>
export type TaskRunField = 'ID' | 'StartedAt' | 'CompletedAt' | 'Status' | 'MethodRequestData' | 'MethodResponseData' | 'TriggeredBy' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | 'TaskID' | never

export type TaskRunFieldCondition = GenericCondition<TaskRunField>

// TriggeredBy: <description of the enum>
export type TriggeredBy = 'Cron' | 'Manual' | never

export type TriggeredByCondition = GenericCondition<TriggeredBy>

/**
 * Local Types, Type Info Props, Type Infos
 */
// No Local Types

/**
 * Entity Type, Entity Info Props
 */
// TaskRun: <description of the type>
export type TaskRun = {
    id: scalars.ID
    startedAt: Date
    completedAt: Date
    status: Status
    methodRequestData: scalars.GenericData
    methodResponseData: scalars.GenericData
    triggeredBy: TriggeredBy
    createdAt: Date
    updatedAt: Date
    deletedAt: Date | null
    taskID: scalars.ID
}

/**
 * Filter Types
 */
// TaskRunFilter: <description of the type>
export type TaskRunFilter = {
    id: IDCondition | null
    startedAt: TimestampCondition | null
    completedAt: TimestampCondition | null
    status: StatusCondition | null
    methodRequestData: GenericDataCondition | null
    methodResponseData: GenericDataCondition | null
    triggeredBy: TriggeredByCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    taskID: IDCondition | null
    and: TaskRunFilter[] | null
    or: TaskRunFilter[] | null
}

/**
 * Standard Request / Response Types
 */
// StandardEntityRequest: <description of the type>
export type StandardEntityRequest = {
    id: scalars.ID
}
// StandardEntityResponse: <description of the type>
export type StandardEntityResponse = {
    object: TaskRun
}
