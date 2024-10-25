'use client'

import { LayoutRootPrivateAppInfo } from '@ongoku/app-lib/src/components/admin/mantine/LayoutRootPrivateAppInfo'
import { appInfoReq } from '@sampleapp/goku-generated/types/types.generated'

const Layout = (props: { children: React.ReactNode }) => {
    return <LayoutRootPrivateAppInfo appInfoReq={appInfoReq}>{props.children}</LayoutRootPrivateAppInfo>
}

export default Layout
