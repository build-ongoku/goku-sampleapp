/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BooleanCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const serviceReq: ServiceReq = {
    namespace: {
        service: 'Auth',
        entity: undefined,
    },
}

const typeReqs: TypeReq[] = [
    {
        namespace: {
            service: 'Auth',
            entity: undefined,
            types: ['AuthenticateResponse'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('token'),
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
    {
        namespace: {
            service: 'Auth',
            entity: undefined,
            types: ['AuthenticateTokenRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('token'),
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
    {
        namespace: {
            service: 'Auth',
            entity: undefined,
            types: ['LoginRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('email'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Email'),
                    kind: fieldkind.EmailKind,
                    namespace: undefined,
                },
            },
            {
                name: new Name('password'),
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
    {
        namespace: {
            service: 'Auth',
            entity: undefined,
            types: ['RegisterUserRequest'],
            enum: undefined,
            method: undefined,
        },
        fields: [
            {
                name: new Name('email'),
                isRepeated: false,
                isOptional: false,
                isMetaField: false,
                excludeFromForm: undefined,
                dtype: {
                    name: new Name('Email'),
                    kind: fieldkind.EmailKind,
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
                    name: new Name('PersonName'),
                    kind: fieldkind.NestedKind,
                    namespace: {
                        service: undefined,
                        entity: undefined,
                        types: ['PersonName'],
                        enum: undefined,
                        method: undefined,
                    },
                },
            },
            {
                name: new Name('password'),
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
                name: new Name('teamID'),
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
]

const enumReqs: EnumReq[] = []
