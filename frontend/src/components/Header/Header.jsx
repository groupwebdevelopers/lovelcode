import React from 'react'
import NavBar from '../NavBar/NavBar'

export default function Header() {
    return (
        <>
            <NavBar />
            <div className="container flex flex-col items-center mb-5 sm:mb-10 md:mb-20">
                <h4 className='text-xl sm:text-3xl lg:text-[40px] font-Ray-ExtraBold text-main-dark-text-web mb-2.5 text-center'>
                    آژانــس حرفــه ای دیــــجیتـــــال مــــــارکـــتـیـنـگ
                    <span className='text-main-blue-web inline-block'>
                        لاول کد
                    </span>
                </h4>
                <div className="flex items-center gap-x-10 lg:gap-x-12 mb-2 sm:mb-8">
                    <img src="/images/mainweb/3D/Sec1/2.png" className='hidden sm:block h-2 lg:h-3' alt="" />
                    <span className='text-main-green-web text-base sm:text-xl lg:text-2xl font-Ray-Bold text-center'>
                        در کنار راه اندازی کسب و کار مجازی شما هستیم :)
                    </span>
                    <img src="/images/mainweb/3D/Sec1/2.png" className='hidden sm:block h-2 lg:h-3' alt="" />
                </div>
                <div className="mb-6 sm:mb-8">
                    <img src="/images/mainweb/3D/Sec1/1.png" className='max-h-[320px] sm:max-h-[400px] lg:max-h-[554px]' alt="" />
                </div>
            </div>
        </>
    )
}
