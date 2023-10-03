import React from 'react'
import { Link } from 'react-router-dom'
import { NavLink } from 'react-router-dom'

export default function Header() {
    return (
        <div className='container flex items-center justify-between pt-[30px]'>
            <div className="flex items-center gap-x-2.5">
                <Link to='/' className="flex items-center justify-center shadow-normal lg:h-[55px] lg:w-[55px] h-12 w-12 bg-white rounded-xl lg:rounded-2xl">
                    <img src="./images/mainweb/3D/Sec1/path28.svg" className='h-7 lg:h-[31px]' alt="" />
                </Link>
                <div className="flex flex-col">
                    <Link to='/' className='text-base lg:text-lg/5 font-Ray-ExtraBold text-main-blue-web'>
                        LovelCode
                    </Link>
                    <Link to='/' className='text-xs lg:text-sm/4 font-Ray-Black text-main-gray-text-web'>
                        لاول کد
                    </Link>
                </div>
            </div>
            <ul className="hidden md:flex text-xs lg:text-lg text-main-gray-text-web font-Ray-Bold gap-x-6 lg:gap-x-[34px]">
                <li className='hover:text-main-blue-web transition-colors'>
                    <NavLink to='/' className={({ isActive }) => isActive ? "text-main-blue-web relative before:content-[''] before:absolute before:top-0 before:right-0 before:h-0.5 before:w-0.5 before:bg-main-blue-web" : ''}>
                        صفحه اصلی
                    </NavLink>
                </li>
                <li className='hover:text-main-blue-web transition-colors'>
                    <NavLink to='/portfolio' className={({ isActive }) => isActive ? "text-main-blue-web relative before:content-[''] before:absolute before:top-0 before:right-0 before:h-0.5 before:w-0.5 before:bg-main-blue-web" : ''}>
                        نمونه کارها
                    </NavLink>
                </li>
                <li className='hover:text-main-blue-web transition-colors'>
                    <NavLink to='/s' className={({ isActive }) => isActive ? "text-main-blue-web relative before:content-[''] before:absolute before:top-0 before:right-0 before:h-0.5 before:w-0.5 before:bg-main-blue-web" : ''}>
                        تعرفه طراحی سایت
                    </NavLink>
                </li>
                <li className='hover:text-main-blue-web transition-colors'>
                    <NavLink to='/d' className={({ isActive }) => isActive ? "text-main-blue-web relative before:content-[''] before:absolute before:top-0 before:right-0 before:h-0.5 before:w-0.5 before:bg-main-blue-web" : ''}>
                        وبلاگ
                    </NavLink>
                </li>
                <li className='hover:text-main-blue-web transition-colors'>
                    <NavLink to='/f' className={({ isActive }) => isActive ? "text-main-blue-web relative before:content-[''] before:absolute before:top-0 before:right-0 before:h-0.5 before:w-0.5 before:bg-main-blue-web" : ''}>
                        درباره ما
                    </NavLink>
                </li>
                <li className='hover:text-main-blue-web transition-colors'>
                    <NavLink to='/g' className={({ isActive }) => isActive ? "text-main-blue-web relative before:content-[''] before:absolute before:top-0 before:right-0 before:h-0.5 before:w-0.5 before:bg-main-blue-web" : ''}>
                        تماس با ما
                    </NavLink>
                </li>
            </ul>
            <div className="hidden md:flex items-center gap-x-2 lg:gap-x-6">
                <Link to='/' className='text-main-blue-web px-4 py-2 text-sm lg:text-base'>
                    ورود
                </Link>
                <Link to='/' className='w-20 h-10 lg:w-24 lg:h-12 flex justify-center items-center lg:rounded-xl rounded-lg text-sm lg:text-base bg-main-blue-web text-white'>
                    ثبت نام
                </Link>
            </div>
            <div className='block md:hidden text-xl cursor-pointer'>
                <i class="bi bi-list"></i>
            </div>
        </div>
    )
}
