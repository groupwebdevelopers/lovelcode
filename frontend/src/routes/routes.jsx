import Home from "../pages/Home/Home"
import Portfolio from "../pages/Portfolio/Portfolio"
import Login from "../pages/Login/Login"

const routes = [
    { path: '/', element: <Home /> },
    { path: '/portfolio', element: <Portfolio /> },
    { path: '/login', element: <Login /> }
]
export default routes