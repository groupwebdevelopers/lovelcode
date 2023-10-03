import { useState } from 'react'
import routes from './routes/routes'
import { useRoutes } from 'react-router-dom'
import MenuContext from './context/MenuContext'

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
            {router}
        </MenuContext.Provider>
    )
}

export default App
