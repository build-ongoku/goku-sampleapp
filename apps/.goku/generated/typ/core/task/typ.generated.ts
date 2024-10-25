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
export type MethodName = 'Add' | 'Update' | 'Get' | 'List' | 'QueryByText' | 'HookSavePre' | 'ActionRun' | never

export type MethodNameCondition = GenericCondition<MethodName>

// TaskAction: <description of the enum>
export type TaskAction = 'Run' | never

export type TaskActionCondition = GenericCondition<TaskAction>

// TaskField: <description of the enum>
export type TaskField = 'ID' | 'Name' | 'Description' | 'Method' | 'MethodRequestData' | 'Enabled' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | never

export type TaskFieldCondition = GenericCondition<TaskField>

/**
 * Local Types, Type Info Props, Type Infos
 */
// No Local Types

/**
 * Entity Type, Entity Info Props
 */
// Task: <description of the type>
export type Task = {
    id: scalars.ID
    name: string
    description: string
    method: app_typ.MethodName
    methodRequestData: scalars.GenericData
    enabled: boolean
    createdAt: Date
    updatedAt: Date
    deletedAt: Date | null
}

/**
 * Filter Types
 */
// TaskFilter: <description of the type>
export type TaskFilter = {
    id: IDCondition | null
    name: StringCondition | null
    description: StringCondition | null
    method: app_typ.MethodNameCondition | null
    methodRequestData: GenericDataCondition | null
    enabled: BoolCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: TaskFilter[] | null
    or: TaskFilter[] | null
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
    object: Task
}
