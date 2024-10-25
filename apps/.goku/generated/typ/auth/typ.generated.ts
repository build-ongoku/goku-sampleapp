/*eslint-disable */

import { EntityMinimal } from '@ongoku/app-lib/src/common/Entity'
import { convertTypeToWithMeta } from '@ongoku/app-lib/src/common/types'
import * as field from '@ongoku/app-lib/src/common/Field'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

import * as app_typ from '@sampleapp/goku-generated/typ/typ.generated'
import * as svcauth_entsecret_typ from '@sampleapp/goku-generated/typ/auth/secret/typ.generated'

/*eslint-disable */

/**
 * Enums
 */
// EntityName: <description of the enum>
export type EntityName = 'Secret' | never

export type EntityNameCondition = GenericCondition<EntityName>

// MethodName: <description of the enum>
export type MethodName = 'LoginUser' | 'RegisterUser' | 'AuthenticateToken' | never

export type MethodNameCondition = GenericCondition<MethodName>

/**
 * Local Types, Type Info Props, Type Infos
 */
// No Local Types

/**
 * Entity Type, Entity Info Props
 */
// No Entities

/**
 * Filter Types
 */
// No Filter Types

/**
 * Standard Request / Response Types
 */
// AuthenticateResponse: <description of the type>
export type AuthenticateResponse = {
    token: string
}
// AuthenticateTokenRequest: <description of the type>
export type AuthenticateTokenRequest = {
    token: string
}
// LoginRequest: <description of the type>
export type LoginRequest = {
    email: scalars.Email
    password: string
}
// RegisterUserRequest: <description of the type>
export type RegisterUserRequest = {
    email: scalars.Email
    name: app_typ.PersonNameInput
    password: string
    teamID: scalars.ID
}
