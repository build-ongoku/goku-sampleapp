/*eslint-disable */

import { EntityMinimal } from '@ongoku/app-lib/src/common/Entity'
import { convertTypeToWithMeta } from '@ongoku/app-lib/src/common/types'
import * as field from '@ongoku/app-lib/src/common/Field'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

import * as app_typ from '@sampleapp/goku-generated/typ/typ.generated'
import * as svcuser_entteam_typ from '@sampleapp/goku-generated/typ/user/team/typ.generated'
import * as svcuser_entuser_typ from '@sampleapp/goku-generated/typ/user/user/typ.generated'

/*eslint-disable */

/**
 * Enums
 */
// EntityName: <description of the enum>
export type EntityName = 'User' | 'Team' | never

export type EntityNameCondition = GenericCondition<EntityName>

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
