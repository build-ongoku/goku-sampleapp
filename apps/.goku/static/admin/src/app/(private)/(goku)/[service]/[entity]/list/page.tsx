'use client'

import { AppInfoContext } from '@ongoku/app-lib/src/common/AppContext'
import { useContext } from 'react'
import { EntityMinimal } from '@ongoku/app-lib/src/common/Entity'
import { EntityListTable } from '@ongoku/app-lib/src/components/mantine/EntityList'

interface Props {
    params: {
        service: string
        entity: string
    }
}

const Page = <E extends EntityMinimal = any>(props: Props) => {
    const { service: serviceName, entity: entityName } = props.params
    const { appInfo } = useContext(AppInfoContext)
    if (!appInfo) {
        throw new Error('AppInfo not loaded')
    }

    const serviceInfo = appInfo.getServiceInfo(serviceName)
    const entityInfo = serviceInfo.getEntityInfo<E>(entityName)

    return (
        <div>
            <EntityListTable entityInfo={entityInfo} />
        </div>
    )
}

export default Page
