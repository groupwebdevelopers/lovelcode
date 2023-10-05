import React, { useContext } from 'react'
import { NavLink } from 'react-router-dom'
import { Link } from 'react-router-dom'
import MenuContext from '../../context/MenuContext'

export default function Menu() {
    const menuContext = useContext(MenuContext)
    const closeMenu = () => {
        menuContext.menuDisplayHandler(false)
    }
    return (
        <div className={`flex flex-col justify-between fixed w-64 sm:w-72 top-0 bottom-0 bg-white shadow-normal z-50 p-6 transition-all ${menuContext.menuDisplay ? 'left-0': '-left-64 sm:-left-72'}`}>
            <div>
                <div className='flex items-center justify-between mb-6'>
                    <span className='flex text-2xl text-main-dark-text-web cursor-pointer' onClick={closeMenu}>
                        <i className="bi bi-x"></i>
                    </span>
                    <Link to='/' className='text-main-blue-web font-Ray-ExtraBold text-xl'>
                        LovelCode
                    </Link>
                </div>
                <ul className='flex flex-col gap-y-4 text-base text-main-dark-text-web font-Ray-Bold'>
                    <li className='flex justify-end'>
                        <NavLink to='/' className={({ isActive }) => isActive ? "bg-main-blue-web text-white w-full text-left py-1 px-2 rounded-md" : ''}>
                            صفحه اصلی
                        </NavLink>
                    </li>
                    <li className='flex justify-end hover:text-main-blue-web transition-colors'>
                        <NavLink to='/a' className={({ isActive }) => isActive ? "bg-main-blue-web text-white" : ''}>
                            نمونه کارها
                        </NavLink>
                    </li>
                    <li className='flex justify-end hover:text-main-blue-web transition-colors'>
                        <NavLink to='/s' className={({ isActive }) => isActive ? "bg-main-blue-web text-white" : ''}>
                            تعرفه طراحی سایت
                        </NavLink>
                    </li>
                    <li className='flex justify-end hover:text-main-blue-web transition-colors'>
                        <NavLink to='/d' className={({ isActive }) => isActive ? "bg-main-blue-web text-white" : ''}>
                            وبلاگ
                        </NavLink>
                    </li>
                    <li className='flex justify-end hover:text-main-blue-web transition-colors'>
                        <NavLink to='/f' className={({ isActive }) => isActive ? "bg-main-blue-web text-white" : ''}>
                            درباره ما
                        </NavLink>
                    </li>
                    <li className='flex justify-end hover:text-main-blue-web transition-colors'>
                        <NavLink to='/g' className={({ isActive }) => isActive ? "bg-main-blue-web text-white" : ''}>
                            تماس با ما
                        </NavLink>
                    </li>
                </ul>
            </div>
            <Link to='/' className="flex items-center justify-center text-white rounded-md text-sm bg-main-blue-web h-10 w-full">
                ورود & ثبت نام
            </Link>
        </div>
    )
}
