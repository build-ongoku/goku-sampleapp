import { EnterScreen } from '@ongoku/app-lib/src/components/admin/mantine/EnterScreen'

const Home = () => {
    return <EnterScreen appName="hosting" authenticatedUserRedirectPath="/home" unauthenticatedUserRedirectPath="/login" />
}

export default Home
