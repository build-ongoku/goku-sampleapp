/*eslint-disable */

import * as fieldkind from '@ongoku/app-lib/src/common/fieldkind'
import * as scalars from '@ongoku/app-lib/src/common/scalars'
import { App, AppReq, ServiceReq, EntityInfoReq, TypeInfoReq, EnumReq, MethodReq } from '@ongoku/app-lib/src/common/app_v3'
import { Name, ITypeAppNamespace, IEnumAppNamespace, ITypeServiceNamespace, IEnumServiceNamespace, ITypeEntityNamespace, IEnumEntityNamespace } from '@ongoku/app-lib/src/common/namespacev2'
import { GenericCondition, StringCondition, NumberCondition, FloatCondition, BoolCondition, DateCondition, TimestampCondition, IDCondition, EmailCondition } from '@ongoku/app-lib/src/common/Filter'

const serviceReq: ServiceReq = {
    namespace: {
        service: 'user',
        entity: undefined,
    },
}

const typeReqs: TypeReq[] = [
    {
        namespace: {
            service: 'user',
            entity: undefined,
            types: ['entity_name_condition'],
            enum: undefined,
            method: undefined,
        },
        fields: [],
        getEmptyObjectFunc: (appInfo: App) => {
            // All nested types should be initialized with their empty instances

            // Return the empty object
            return {}
        },
    },
]

const enumReqs: EnumReq[] = [
    {
        namespace: {
            service: 'user',
            entity: undefined,
            types: undefined,
            enum: 'entity_name',
            method: undefined,
        },
        values: [
            {
                id: 1,
                value: 'User',
                description: undefined,
                displayValue: 'User',
            },
            {
                id: 2,
                value: 'Team',
                description: undefined,
                displayValue: 'Team',
            },
        ],
    },
]
