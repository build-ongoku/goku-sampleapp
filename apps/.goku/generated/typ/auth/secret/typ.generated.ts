/*eslint-disable */

import { EntityMinimal } from '@ongoku/app-lib/src/common/Entity'
import { convertTypeToWithMeta } from '@ongoku/app-lib/src/common/types'
import * as field from '@ongoku/app-lib/src/common/Field'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

import * as app_typ from '@sampleapp/goku-generated/typ/typ.generated'
import * as svcauth_typ from '@sampleapp/goku-generated/typ/auth/typ.generated'

/*eslint-disable */

/**
 * Enums
 */
// MethodName: <description of the enum>
export type MethodName = 'Add' | 'Update' | 'Get' | 'List' | 'QueryByText' | never

export type MethodNameCondition = GenericCondition<MethodName>

// SecretField: <description of the enum>
export type SecretField = 'ID' | 'UserID' | 'Type' | 'Secret' | 'ExpiresAt' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | never

export type SecretFieldCondition = GenericCondition<SecretField>

// Type: <description of the enum>
export type Type = 'Password' | 'GithubToken' | never

export type TypeCondition = GenericCondition<Type>

/**
 * Local Types, Type Info Props, Type Infos
 */
// No Local Types

/**
 * Entity Type, Entity Info Props
 */
// Secret: <description of the type>
export type Secret = {
    id: scalars.ID
    userID: scalars.ID
    type: Type
    secret: string
    expiresAt: Date | null
    createdAt: Date
    updatedAt: Date
    deletedAt: Date | null
}

/**
 * Filter Types
 */
// SecretFilter: <description of the type>
export type SecretFilter = {
    id: IDCondition | null
    userID: IDCondition | null
    type: TypeCondition | null
    secret: StringCondition | null
    expiresAt: TimestampCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: SecretFilter[] | null
    or: SecretFilter[] | null
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
    object: Secret
}
