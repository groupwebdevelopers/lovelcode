import React from 'react'

export default function Header() {
    return (
        <div className='container flex items-center justify-between pt-[30px]'>
            <div className="flex items-center gap-x-2.5">
                <div className="flex items-center justify-center shadow-normal h-[55px] w-[55px] bg-white rounded-2xl">
                    <img src="./images/mainweb/3D/Sec1/path28.svg" className='h-[31px]' alt="" />
                </div>
                <div className="flex flex-col">
                    <h3 className='text-lg font-Ray-ExtraBold text-main-blue-web'>
                        LovelCode
                    </h3>
                    <span className='text-sm font-Ray-Black text-main-gray-text-web'>
                        لاول کد
                    </span>
                </div>
            </div>
            <ul className="flex text-lg text-main-gray-text-web font-Ray-Bold gap-x-[34px]">
                <li>
                    صفحه اصلی
                </li>
                <li>
                    نمونه کارها
                </li>
                <li>
                    تعرفه طراحی سایت
                </li>
                <li>
                    وبلاگ
                </li>
                <li>
                    درباره ما
                </li>
                <li>
                    تماس با ما
                </li>
            </ul>
            <div className="flex items-center gap-x-8">
                <a className='text-main-blue-web'>
                    ورود
                </a>
                <a className='w-24 h-10 flex justify-center items-center rounded-xl bg-main-blue-web text-white'>
                    ثبت نام
                </a>
            </div>
        </div>
    )
}
