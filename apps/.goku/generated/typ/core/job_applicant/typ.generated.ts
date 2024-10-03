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
// JobApplicantField: <description of the enum>
export type JobApplicantField = 'ID' | 'Name' | 'UserID' | 'ResumeID' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | never

export type JobApplicantFieldCondition = GenericCondition<JobApplicantField>

/**
 * Local Types, Type Info Props, Type Infos
 */
// No Local Types

/**
 * Entity Type, Entity Info Props
 */
// JobApplicant: <description of the type>
export type JobApplicant = {
    id: scalars.ID
    name: string
    userID: scalars.ID
    resumeID: scalars.ID
    createdAt: Date
    updatedAt: Date
    deletedAt: Date | null
}

/**
 * Filter Types
 */
// JobApplicantFilter: <description of the type>
export type JobApplicantFilter = {
    id: IDCondition | null
    name: StringCondition | null
    userID: IDCondition | null
    resumeID: IDCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: JobApplicantFilter[] | null
    or: JobApplicantFilter[] | null
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
    object: JobApplicant
}
