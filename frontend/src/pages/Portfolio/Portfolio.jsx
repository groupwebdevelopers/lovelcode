import React from 'react'
import NavBar from '../../components/NavBar/NavBar'
import Footer from '../../components/Footer/Footer'
export default function Portfolio() {
    return (
        <div className='relative'>
            <NavBar />
            <div className='all container text-main-dark-text-web'>
                <div className='p'>
                    <h2 className='font-Ray-ExtraBold text-3xl'>نمونه کار طراحی سایت</h2>
                    <p className='font-Ray-Medium text-lg mt-2'>نمونه کار طراحی سایت های فروشگاهی، خبری و طراحی سایت آموزشی و ... را که در این صفحه مشاهده می کنید همه پروژه های اجرا شده کسب و کارهای آنلاینی هست که درحال کار هستند. همه این نمونه کارهای طراحی سایت با سورسی بهینه و طبق استاندارهای گوگل طراحی شده است. این ادعا به پشتوانه طراحی بی نظیر و پیاده سازی قوی و باکیفیت سایت های بسیاری برای برندها و استارتاپ های معتبر ایران می باشد.</p>
                    <p className='font-Ray-Medium text-lg mt-2'> برای آشنایی شما با برندهایی که تاکنون با آنها همکاری داشته ایم و خدماتی که برای هریک انجام داده ایم، اطلاعات آن دسته از مشتریان مان که اجازه بیان نام پروژه هایشان را به ما داده اند و خدماتی که به آنها عرضه کرده ایم را، مختصرا عنوان نموده ایم تا بیشتر با بیلاودمارکتینگ و خدمات آن آشنا شوید.</p>
                </div>


                <div className='select hidden font-Ray-Bold lg:flex my-12 bg-white w-fit  rounded-xl items-center text-xs lg:text-base'>
                    <span className='bg-main-blue-web px-[22px] py-[14px] rounded-xl text-white'>همه</span>
                    <span className='px-[22px] py-[14px] rounded-xl'>سایت فروشگاهی</span>
                    <span className='px-[22px] py-[14px] rounded-xl'>سایت آموزشی</span>
                    <span className='px-[22px] py-[14px] rounded-xl'>سایت شرکتی</span>
                    <span className='px-[22px] py-[14px] rounded-xl'>سایت شخصی</span>
                    <span className='px-[22px] py-[14px] rounded-xl'>دیجیتال مارکتینگ</span>
                    <span className='px-[22px] py-[14px] rounded-xl'>برندینگ</span>
                    <span className='px-[22px] py-[14px] rounded-xl'>سایت خبری</span>
                </div>

                <div className='lg:hidden'>
                    <select name="" id="" className='w-36 h-14 p-2 rounded-lg mt-14 outline-none'>
                        {
                            Array.from({length:8},(_,i)=>i+1).map((item)=>(
                                <option key={item.length} value="همه">همه</option>
                            ))
                        }
                    </select>
                </div>
                

                <div className='grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 mb-[100px] gap-[20px] mt-12'>
                    {
                        Array.from({length:24} , (_,i)=>i+1).map(()=>(
                            <div className='flex justify-center'>
                                <div className='w-[270px] h-[270px] rounded-[25px] bg-white'></div>
                            </div>
                        ))
                    }
                </div>


            </div>
            <Footer></Footer>
        </div>
    )
}
