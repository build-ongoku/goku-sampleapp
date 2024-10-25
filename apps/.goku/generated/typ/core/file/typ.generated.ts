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
// FileField: <description of the enum>
export type FileField = 'ID' | 'FileName' | 'StorageClient' | 'StorageClientIdentifier' | 'SizeBytes' | 'MimeType' | 'FileHash' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | never

export type FileFieldCondition = GenericCondition<FileField>

// MethodName: <description of the enum>
export type MethodName = 'Add' | 'Update' | 'Get' | 'List' | 'QueryByText' | 'HookCreatePre' | never

export type MethodNameCondition = GenericCondition<MethodName>

// StorageClient: <description of the enum>
export type StorageClient = 'LocalFile' | 'Database' | 'CloudAmazonS3' | never

export type StorageClientCondition = GenericCondition<StorageClient>

/**
 * Local Types, Type Info Props, Type Infos
 */
// No Local Types

/**
 * Entity Type, Entity Info Props
 */
// File: <description of the type>
export type File = {
    id: scalars.ID
    fileName: string
    storageClient: StorageClient
    storageClientIdentifier: string
    sizeBytes: number
    mimeType: string
    fileHash: string
    createdAt: Date
    updatedAt: Date
    deletedAt: Date | null
}

/**
 * Filter Types
 */
// FileFilter: <description of the type>
export type FileFilter = {
    id: IDCondition | null
    fileName: StringCondition | null
    storageClient: StorageClientCondition | null
    storageClientIdentifier: StringCondition | null
    sizeBytes: NumberCondition | null
    mimeType: StringCondition | null
    fileHash: StringCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: FileFilter[] | null
    or: FileFilter[] | null
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
    object: File
}
