import React from 'react'

export default function Login() {
    return (
        <div className='h-screen w-full flex items-center justify-center'>
            <div className="flex flex-col w-96 bg-white rounded-3xl p-10">
                <div className='w-full flex flex-col items-center gap-4 mb-12'>
                    <img src="/images/mainweb/3D/Sec1/path28.svg" alt="" />
                    <p className='text-xl font-Ray-ExtraBold text-main-blue-web'>ورود به لاول کد</p>
                </div>
                <div className='flex flex-col gap-y-5'>
                    <div>
                        <input type="text" className='bg-gray-normal rounded-xl w-full h-12 outline-none font-Ray-Medium text-sm px-5' placeholder='نشانی ایمیل' />
                    </div>
                    <div>
                        <input type="text" className='bg-gray-normal rounded-xl w-full h-12 outline-none font-Ray-Medium text-sm px-5' placeholder='نشانی ایمیل' />
                    </div>
                </div>
            </div>
        </div>
    )
}
