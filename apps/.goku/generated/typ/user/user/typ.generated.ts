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
// UserField: <description of the enum>
export type UserField =
    | 'ID'
    | 'Name'
    | 'Name_ParentID'
    | 'Name_ID'
    | 'Name_FirstName'
    | 'Name_MiddleInitial'
    | 'Name_LastName'
    | 'Name_CreatedAt'
    | 'Name_UpdatedAt'
    | 'Name_DeletedAt'
    | 'Email'
    | 'Addresses'
    | 'Addresses_ParentID'
    | 'Addresses_ID'
    | 'Addresses_Line1'
    | 'Addresses_Line2'
    | 'Addresses_City'
    | 'Addresses_State'
    | 'Addresses_PostalCode'
    | 'Addresses_Country'
    | 'Addresses_CreatedAt'
    | 'Addresses_UpdatedAt'
    | 'Addresses_DeletedAt'
    | 'TeamID'
    | 'PastTeamIDs'
    | 'CreatedAt'
    | 'UpdatedAt'
    | 'DeletedAt'
    | never

export type UserFieldCondition = GenericCondition<UserField>

/**
 * Local Types, Type Info Props, Type Infos
 */
// No Local Types

/**
 * Entity Type, Entity Info Props
 */
// User: <description of the type>
export type User = {
    id: scalars.ID
    name: app_typ.PersonNameWithMeta
    email: scalars.Email
    addresses: app_typ.AddressWithMeta[] | null
    teamID: scalars.ID
    pastTeamIDs: scalars.ID[] | null
    createdAt: Date
    updatedAt: Date
    deletedAt: Date | null
}

/**
 * Filter Types
 */
// UserFilter: <description of the type>
export type UserFilter = {
    id: IDCondition | null
    name: app_typ.PersonNameFilter | null
    email: EmailCondition | null
    havingAddresses: app_typ.AddressFilter | null
    teamID: IDCondition | null
    havingPastTeamIDs: IDCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: UserFilter[] | null
    or: UserFilter[] | null
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
    object: User
}
