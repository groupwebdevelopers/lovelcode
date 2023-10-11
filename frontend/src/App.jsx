import { useState } from 'react'
import routes from './routes/routes'
import { useRoutes } from 'react-router-dom'
import MenuContext from './context/MenuContext'
import Menu from './components/Menu/Menu'
import Overlay from './components/Overlay/Overlay'

function App() {
    const router = useRoutes(routes)
    const [menuDisplay, setMenuDisplay] = useState(false)
    const menuDisplayHandler = (display) => {
        setMenuDisplay(display)
    }
    return (
        <MenuContext.Provider
            value={{ menuDisplay, menuDisplayHandler }}
        >
            <Menu />
            <Overlay />
            {router}
        </MenuContext.Provider>
    )
}

export default App
