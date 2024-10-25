/*eslint-disable */

import { EntityMinimal } from '@ongoku/app-lib/src/common/Entity'
import { convertTypeToWithMeta } from '@ongoku/app-lib/src/common/types'
import * as field from '@ongoku/app-lib/src/common/Field'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

import * as app_typ from '@sampleapp/goku-generated/typ/typ.generated'
import * as svcuser_typ from '@sampleapp/goku-generated/typ/user/typ.generated'

/*eslint-disable */

/**
 * Enums
 */
// MethodName: <description of the enum>
export type MethodName = 'Add' | 'Update' | 'Get' | 'List' | 'QueryByText' | never

export type MethodNameCondition = GenericCondition<MethodName>

// TeamField: <description of the enum>
export type TeamField = 'ID' | 'Name' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | never

export type TeamFieldCondition = GenericCondition<TeamField>

/**
 * Local Types, Type Info Props, Type Infos
 */
// No Local Types

/**
 * Entity Type, Entity Info Props
 */
// Team: <description of the type>
export type Team = {
    id: scalars.ID
    name: string
    createdAt: Date
    updatedAt: Date
    deletedAt: Date | null
}

/**
 * Filter Types
 */
// TeamFilter: <description of the type>
export type TeamFilter = {
    id: IDCondition | null
    name: StringCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: TeamFilter[] | null
    or: TeamFilter[] | null
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
    object: Team
}
