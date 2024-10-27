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
export type MethodName = 'Add' | 'Update' | 'Get' | 'List' | 'QueryByText' | 'HookCreatePre' | 'ActionDecrypt' | never

export type MethodNameCondition = GenericCondition<MethodName>

// SecretDecryptableAction: <description of the enum>
export type SecretDecryptableAction = 'Decrypt' | never

export type SecretDecryptableActionCondition = GenericCondition<SecretDecryptableAction>

// SecretDecryptableField: <description of the enum>
export type SecretDecryptableField = 'ID' | 'RawValue' | 'EncryptedValue' | 'Salt' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | 'SecretKeyID' | never

export type SecretDecryptableFieldCondition = GenericCondition<SecretDecryptableField>

/**
 * Local Types, Type Info Props, Type Infos
 */
// No Local Types

/**
 * Entity Type, Entity Info Props
 */
// SecretDecryptable: <description of the type>
export type SecretDecryptable = {
    id: scalars.ID
    rawValue: string | null
    encryptedValue: string
    salt: string
    createdAt: Date
    updatedAt: Date
    deletedAt: Date | null
    secretKeyID: scalars.ID
}

/**
 * Filter Types
 */
// SecretDecryptableFilter: <description of the type>
export type SecretDecryptableFilter = {
    id: IDCondition | null
    rawValue: StringCondition | null
    encryptedValue: StringCondition | null
    salt: StringCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    secretKeyID: IDCondition | null
    and: SecretDecryptableFilter[] | null
    or: SecretDecryptableFilter[] | null
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
    object: SecretDecryptable
}
