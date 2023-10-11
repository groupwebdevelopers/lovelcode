import Home from "../pages/Home/Home"
import Portfolio from "../pages/Portfolio/Portfolio"
import Login from "../pages/Login/Login"
import Register from "../pages/Register/Register"
import AboutUs from '../pages/AboutUs/AboutUs'
import ContactUs from '../pages/ContactUs/ContactUs'

const routes = [
    { path: '/', element: <Home /> },
    { path: '/portfolio', element: <Portfolio /> },
    { path: '/login', element: <Login /> },
    { path: '/register', element: <Register /> },
    { path: '/about-us', element: <AboutUs /> },
    { path: '/contact-us', element: <ContactUs /> }
]
export default routes