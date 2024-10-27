/*eslint-disable */

import { EntityMinimal } from '@ongoku/app-lib/src/common/Entity'
import { convertTypeToWithMeta } from '@ongoku/app-lib/src/common/types'
import * as field from '@ongoku/app-lib/src/common/Field'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

import * as svcauth_typ from '@sampleapp/goku-generated/typ/auth/typ.generated'
import * as svccore_typ from '@sampleapp/goku-generated/typ/core/typ.generated'
import * as svcuser_typ from '@sampleapp/goku-generated/typ/user/typ.generated'

/*eslint-disable */

/**
 * Enums
 */
// AddressField: <description of the enum>
export type AddressField = 'ParentID' | 'ID' | 'Line1' | 'Line2' | 'City' | 'State' | 'PostalCode' | 'Country' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | never

export type AddressFieldCondition = GenericCondition<AddressField>

// ContactField: <description of the enum>
export type ContactField =
    | 'ParentID'
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
    | 'Address'
    | 'Address_ParentID'
    | 'Address_ID'
    | 'Address_Line1'
    | 'Address_Line2'
    | 'Address_City'
    | 'Address_State'
    | 'Address_PostalCode'
    | 'Address_Country'
    | 'Address_CreatedAt'
    | 'Address_UpdatedAt'
    | 'Address_DeletedAt'
    | 'CreatedAt'
    | 'UpdatedAt'
    | 'DeletedAt'
    | never

export type ContactFieldCondition = GenericCondition<ContactField>

// Country: <description of the enum>
export type Country = 'USA' | 'Canada' | 'Mexico' | never

export type CountryCondition = GenericCondition<Country>

// EntityName: <description of the enum>
export type EntityName = 'Auth_Secret' | 'User_User' | 'User_Team' | 'Core_JobApplicant' | 'Core_File' | 'Core_SecretKey' | 'Core_SecretDecryptable' | 'Core_Task' | 'Core_TaskRun' | never

export type EntityNameCondition = GenericCondition<EntityName>

// Gender: <description of the enum>
export type Gender = 'Male' | 'Female' | 'Other' | never

export type GenderCondition = GenericCondition<Gender>

// MethodName: <description of the enum>
export type MethodName =
    | 'Auth_LoginUser'
    | 'Auth_RegisterUser'
    | 'Auth_AuthenticateToken'
    | 'Core_FileUpload'
    | 'Core_SecretDecryptabeAdd'
    | 'Auth_Secret_Add'
    | 'Auth_Secret_Update'
    | 'Auth_Secret_Get'
    | 'Auth_Secret_List'
    | 'Auth_Secret_QueryByText'
    | 'User_User_Add'
    | 'User_User_Update'
    | 'User_User_Get'
    | 'User_User_List'
    | 'User_User_QueryByText'
    | 'User_Team_Add'
    | 'User_Team_Update'
    | 'User_Team_Get'
    | 'User_Team_List'
    | 'User_Team_QueryByText'
    | 'Core_JobApplicant_Add'
    | 'Core_JobApplicant_Update'
    | 'Core_JobApplicant_Get'
    | 'Core_JobApplicant_List'
    | 'Core_JobApplicant_QueryByText'
    | 'Core_File_Add'
    | 'Core_File_Update'
    | 'Core_File_Get'
    | 'Core_File_List'
    | 'Core_File_QueryByText'
    | 'Core_SecretKey_Add'
    | 'Core_SecretKey_Update'
    | 'Core_SecretKey_Get'
    | 'Core_SecretKey_List'
    | 'Core_SecretKey_QueryByText'
    | 'Core_SecretDecryptable_Add'
    | 'Core_SecretDecryptable_Update'
    | 'Core_SecretDecryptable_Get'
    | 'Core_SecretDecryptable_List'
    | 'Core_SecretDecryptable_QueryByText'
    | 'Core_Task_Add'
    | 'Core_Task_Update'
    | 'Core_Task_Get'
    | 'Core_Task_List'
    | 'Core_Task_QueryByText'
    | 'Core_TaskRun_Add'
    | 'Core_TaskRun_Update'
    | 'Core_TaskRun_Get'
    | 'Core_TaskRun_List'
    | 'Core_TaskRun_QueryByText'
    | 'User_User_HookInit'
    | 'Core_File_HookCreatePre'
    | 'Core_SecretKey_HookInit'
    | 'Core_SecretKey_HookCreatePre'
    | 'Core_SecretDecryptable_HookCreatePre'
    | 'Core_Task_HookSavePre'
    | 'Core_TaskRun_HookCreatePre'
    | 'Core_SecretDecryptable_ActionDecrypt'
    | 'Core_Task_ActionRun'
    | 'Core_TaskRun_ActionRun'
    | never

export type MethodNameCondition = GenericCondition<MethodName>

// PersonNameField: <description of the enum>
export type PersonNameField = 'ParentID' | 'ID' | 'FirstName' | 'MiddleInitial' | 'LastName' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | never

export type PersonNameFieldCondition = GenericCondition<PersonNameField>

// PhoneNumberField: <description of the enum>
export type PhoneNumberField = 'ParentID' | 'ID' | 'CountryCode' | 'Number' | 'Extension' | 'CreatedAt' | 'UpdatedAt' | 'DeletedAt' | never

export type PhoneNumberFieldCondition = GenericCondition<PhoneNumberField>

// ServiceName: <description of the enum>
export type ServiceName = 'Auth' | 'User' | 'Core' | never

export type ServiceNameCondition = GenericCondition<ServiceName>

/**
 * Local Types, Type Info Props, Type Infos
 */

// Address: <description of the type>
export type Address = {
    line1: string
    line2: string | null
    city: string
    state: string
    postalCode: string
    country: Country
}
export type AddressWithMeta = Address & field.MetaFields

// AddressInput: <description of the type>
export type AddressInput = {
    line1: string
    line2: string | null
    city: string
    state: string
    postalCode: string
    country: Country
}
export type AddressInputWithMeta = AddressInput & field.MetaFields

// Contact: <description of the type>
export type Contact = {
    name: PersonName
    email: scalars.Email
    address: Address
}
export type ContactWithMeta = Contact & field.MetaFields

// ContactInput: <description of the type>
export type ContactInput = {
    name: PersonNameInput
    email: scalars.Email
    address: AddressInput
}
export type ContactInputWithMeta = ContactInput & field.MetaFields

// PersonName: <description of the type>
export type PersonName = {
    firstName: string
    middleInitial: string | null
    lastName: string
}
export type PersonNameWithMeta = PersonName & field.MetaFields

// PersonNameInput: <description of the type>
export type PersonNameInput = {
    firstName: string
    middleInitial: string | null
    lastName: string
}
export type PersonNameInputWithMeta = PersonNameInput & field.MetaFields

// PhoneNumber: <description of the type>
export type PhoneNumber = {
    countryCode: number
    number: string
    extension: string | null
}
export type PhoneNumberWithMeta = PhoneNumber & field.MetaFields

// PhoneNumberInput: <description of the type>
export type PhoneNumberInput = {
    countryCode: number
    number: string
    extension: string | null
}
export type PhoneNumberInputWithMeta = PhoneNumberInput & field.MetaFields

/**
 * Entity Type, Entity Info Props
 */
// No Entities

/**
 * Filter Types
 */
// AddressFilter: <description of the type>
export type AddressFilter = {
    parentID: IDCondition | null
    id: IDCondition | null
    line1: StringCondition | null
    line2: StringCondition | null
    city: StringCondition | null
    state: StringCondition | null
    postalCode: StringCondition | null
    country: CountryCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: AddressFilter[] | null
    or: AddressFilter[] | null
}
// ContactFilter: <description of the type>
export type ContactFilter = {
    parentID: IDCondition | null
    id: IDCondition | null
    name: PersonNameFilter | null
    email: EmailCondition | null
    address: AddressFilter | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: ContactFilter[] | null
    or: ContactFilter[] | null
}
// PersonNameFilter: <description of the type>
export type PersonNameFilter = {
    parentID: IDCondition | null
    id: IDCondition | null
    firstName: StringCondition | null
    middleInitial: StringCondition | null
    lastName: StringCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: PersonNameFilter[] | null
    or: PersonNameFilter[] | null
}
// PhoneNumberFilter: <description of the type>
export type PhoneNumberFilter = {
    parentID: IDCondition | null
    id: IDCondition | null
    countryCode: NumberCondition | null
    number: StringCondition | null
    extension: StringCondition | null
    createdAt: TimestampCondition | null
    updatedAt: TimestampCondition | null
    deletedAt: TimestampCondition | null
    and: PhoneNumberFilter[] | null
    or: PhoneNumberFilter[] | null
}

/**
 * Standard Request / Response Types
 */
