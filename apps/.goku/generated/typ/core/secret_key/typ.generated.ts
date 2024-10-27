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
// Format: <description of the enum>
export type Format = 'PEM' | 'OpenSSH' | never

export type FormatCondition = GenericCondition<Format>

// MethodName: <description of the enum>
export type MethodName = 'Add' | 'Update' | 'Get' | 'List' | 'QueryByText' | 'HookInit' | 'HookCreatePre' | never

export type MethodNameCondition = GenericCondition<MethodName>

// SecretKeyField: <description of the enum>
export type SecretKeyField = 'ID' | 'Type' | 'PublicKey' | 'PrivateKeyFormat' | 'PublicKeyFormat' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | never

export type SecretKeyFieldCondition = GenericCondition<SecretKeyField>

// Type: <description of the enum>
export type Type = 'Ed25519' | 'RSA' | never

export type TypeCondition = GenericCondition<Type>

/**
 * Local Types, Type Info Props, Type Infos
 */
// No Local Types

/**
 * Entity Type, Entity Info Props
 */
// SecretKey: <description of the type>
export type SecretKey = {
    id: scalars.ID
    type: Type
    publicKey: string
    privateKeyFormat: Format
    publicKeyFormat: Format
    createdAt: Date
    updatedAt: Date
    deletedAt: Date | null
}

/**
 * Filter Types
 */
// SecretKeyFilter: <description of the type>
export type SecretKeyFilter = {
    id: IDCondition | null
    type: TypeCondition | null
    publicKey: StringCondition | null
    privateKeyFormat: FormatCondition | null
    publicKeyFormat: FormatCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: SecretKeyFilter[] | null
    or: SecretKeyFilter[] | null
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
    object: SecretKey
}
