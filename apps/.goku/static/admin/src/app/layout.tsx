import '@ongoku/app-lib/src/components/admin/mantine/globals.css'
import { LayoutRoot } from '@ongoku/app-lib/src/components/admin/mantine/LayoutRoot'

export const metadata = {
    title: 'Admin App',
    description: 'Generated by Goku',
}

export default function RootLayout({ children }: { children: React.ReactNode }) {
    return <LayoutRoot>{children}</LayoutRoot>
}
