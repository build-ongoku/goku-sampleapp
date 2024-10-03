/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BooleanCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const entityInfoReq: EntityInfoReq = {
    namespace: {
        service: 'Core',
        entity: 'JobApplicant',
    },
    associations: [],
    actions: [],
}

const typeReqs: TypeReq[] = [
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: ['JobApplicant'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('id'),
                isRepeated: false,
                isOptional: false,
                isMetaField: true,
                excludeFromForm: true,
                dtype: {
                    name: new Name('ID'),
                    kind: fieldkind.IDKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('name'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('String'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('userID'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_User'),
                    kind: fieldkind.ForeignEntityKind,
                    namespace: {
                        service: 'Core',
                        entity: 'User',
                    },
                },
            },
            {
                name: new Name('resumeID'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('File'),
                    kind: fieldkind.FileKind,
                    namespace: {
                        service: 'Core',
                        entity: 'File',
                    },
                },
            },
            {
                name: new Name('createdAt'),
                isRepeated: false,
                isOptional: false,
                isMetaField: true,
                excludeFromForm: true,
                dtype: {
                    name: new Name('Timestamp'),
                    kind: fieldkind.TimestampKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('updatedAt'),
                isRepeated: false,
                isOptional: false,
                isMetaField: true,
                excludeFromForm: true,
                dtype: {
                    name: new Name('Timestamp'),
                    kind: fieldkind.TimestampKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('deletedAt'),
                isRepeated: false,
                isOptional: true,
                isMetaField: true,
                excludeFromForm: true,
                dtype: {
                    name: new Name('Timestamp'),
                    kind: fieldkind.TimestampKind,
                    namespace: undefined,
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: ['JobApplicantFieldCondition'],
            enum: undefined,
            method: undefined,
        },
        fields: [],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: ['JobApplicantFilter'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('id'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('IDCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('name'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('StringCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('userID'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('IDCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('resumeID'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('IDCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('createdAt'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('TimestampCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('updatedAt'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('TimestampCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('deletedAt'),
                isRepeated: false,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('TimestampCondition'),
                    kind: fieldkind.ConditionKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('and'),
                isRepeated: true,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_JobApplicant_JobApplicantFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'JobApplicant',
                        types: ['JobApplicantFilter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('or'),
                isRepeated: true,
                isOptional: true,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_JobApplicant_JobApplicantFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'JobApplicant',
                        types: ['JobApplicantFilter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: ['StandardEntityRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('id'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_JobApplicant'),
                    kind: fieldkind.ForeignEntityKind,
                    namespace: {
                        service: 'Core',
                        entity: 'JobApplicant',
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: ['StandardEntityResponse'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('object'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_JobApplicant'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'JobApplicant',
                        types: ['JobApplicant'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: ['AddRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('object'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_JobApplicant'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'JobApplicant',
                        types: ['JobApplicant'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: ['UpdateRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('object'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_JobApplicant'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'JobApplicant',
                        types: ['JobApplicant'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('fields'),
                isRepeated: true,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_JobApplicant_JobApplicantField'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'Core',
                        entity: 'JobApplicant',
                        types: undefined,
                        enum: 'JobApplicantField',
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('excludeFields'),
                isRepeated: true,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_JobApplicant_JobApplicantField'),
                    kind: fieldkind.EnumKind,
                    namespace: {
                        service: 'Core',
                        entity: 'JobApplicant',
                        types: undefined,
                        enum: 'JobApplicantField',
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: ['UpdateResponse'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('object'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_JobApplicant'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'JobApplicant',
                        types: ['JobApplicant'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: ['GetRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('id'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('ID'),
                    kind: fieldkind.IDKind,
                    namespace: undefined,
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: ['ListRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('filter'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_JobApplicant_JobApplicantFilter'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'JobApplicant',
                        types: ['JobApplicantFilter'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: ['ListResponse'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('items'),
                isRepeated: true,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Core_JobApplicant'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: 'Core',
                        entity: 'JobApplicant',
                        types: ['JobApplicant'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('count'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Number'),
                    kind: fieldkind.NumberKind,
                    namespace: undefined,
                },
            },
        ],
    },
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: ['QueryByTextRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('queryText'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('String'),
                    kind: fieldkind.StringKind,
                    namespace: undefined,
                },
            },
        ],
    },
]

const enumReqs: EnumReq[] = [
    {
        namespace: {
            service: 'Core',
            entity: 'JobApplicant',
            types: undefined,
            enum: 'JobApplicantField',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'ID',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 2,
                value: 'Name',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 3,
                value: 'UserID',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 4,
                value: 'ResumeID',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 5,
                value: 'CreatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 6,
                value: 'UpdatedAt',
                description: undefined,
                displayValue: undefined,
            },
            {
                id: 7,
                value: 'DeletedAt',
                description: undefined,
                displayValue: undefined,
            },
        ],
    },
]
