import React from 'react'
import { Link } from 'react-router-dom';
import { Swiper, SwiperSlide } from 'swiper/react'
import { Autoplay, Pagination, Navigation } from 'swiper/modules';
import 'swiper/css'
import 'swiper/css/pagination';
import 'swiper/css/navigation';

export default function Customers() {
    return (
        <div className='container flex flex-col items-center'>
            <h4 className='text-lg sm:text-2xl lg:text-[30px] font-Ray-ExtraBold text-main-dark-text-web mb-2.5 text-center'>
                انتخاب درست این عزیزان بوده ایم
            </h4>
            <span className='text-main-green-web text-sm sm:text-base lg:text-lg font-Ray-Bold text-center mb-2'>
                برخی از مشتریان رضایتمند آژانس دیجیتال مارکتینگ
                <span className='text-main-blue-web inline-block'>
                    لاول کد
                </span>
            </span>
            <div className='flex items-center gap-x-16 xl:gap-x-24 w-full'>
                <button className='hidden sm:flex arrow-right h-7 w-10 lg:h-9 lg:w-12 rounded-lg border border-main-blue-web items-center justify-center text-main-blue-web text-lg hover:bg-main-blue-web hover:text-white transition-colors'>
                    <i class="bi bi-arrow-right-short"></i>
                </button>
                <Swiper
                    loop={true}
                    slidesPerView={3}
                    spaceBetween={30}
                    navigation={{ nextEl: ".arrow-left", prevEl: ".arrow-right" }}
                    autoplay={{
                        delay: 20000,
                        disableOnInteraction: false,
                    }}
                    pagination={{
                        clickable: true,
                        enabled: false
                    }}
                    breakpoints={{
                        640: {
                            slidesPerView: 4,
                            pagination:{
                                enabled: true
                            }
                        },
                        768: {
                            slidesPerView: 4,
                            pagination:{
                                enabled: true
                            }
                        },
                        1024: {
                            slidesPerView: 5,
                            pagination:{
                                enabled: true
                            }
                        },
                    }}
                    modules={[Pagination, Autoplay, Navigation]}
                    className="w-full relative customers h-28 sm:h-40 flex items-center"
                >
                    <SwiperSlide className='flex items-center'>
                        <Link>
                            <img src="/images/mainweb/3D/Secs5/unique.png" className='max-h-16 opacity-50 hover:opacity-100 transition-all' alt="" />
                        </Link>
                    </SwiperSlide>
                    <SwiperSlide className='flex items-center'>
                        <Link>
                            <img src="/images/mainweb/3D/Secs5/Vector.png" className='max-h-16 opacity-50 hover:opacity-100 transition-all' alt="" />
                        </Link>
                    </SwiperSlide>
                    <SwiperSlide className='flex items-center'>
                        <Link>
                            <img src="/images/mainweb/3D/Secs5/greenweb.png" className='max-h-16 opacity-50 hover:opacity-100 transition-all' alt="" />
                        </Link>
                    </SwiperSlide>
                    <SwiperSlide className='flex items-center'>
                        <Link>
                            <img src="/images/mainweb/3D/Secs5/Elabord.png" className='max-h-16 opacity-50 hover:opacity-100 transition-all' alt="" />
                        </Link>
                    </SwiperSlide>
                    <SwiperSlide className='flex items-center'>
                        <Link>
                            <img src="/images/mainweb/3D/Secs5/Elabord.png" className='max-h-16 opacity-50 hover:opacity-100 transition-all' alt="" />
                        </Link>
                    </SwiperSlide>
                    <SwiperSlide className='flex items-center'>
                        <Link>
                            <img src="/images/mainweb/3D/Secs5/Elabord.png" className='max-h-16 opacity-50 hover:opacity-100 transition-all' alt="" />
                        </Link>
                    </SwiperSlide>
                    <SwiperSlide className='flex items-center'>
                        <Link>
                            <img src="/images/mainweb/3D/Secs5/Elabord.png" className='max-h-16 opacity-50 hover:opacity-100 transition-all' alt="" />
                        </Link>
                    </SwiperSlide>
                    <SwiperSlide className='flex items-center'>
                        <Link>
                            <img src="/images/mainweb/3D/Secs5/Elabord.png" className='max-h-16 opacity-50 hover:opacity-100 transition-all' alt="" />
                        </Link>
                    </SwiperSlide>
                    <SwiperSlide className='flex items-center'>
                        <Link>
                            <img src="/images/mainweb/3D/Secs5/Elabord.png" className='max-h-16 opacity-50 hover:opacity-100 transition-all' alt="" />
                        </Link>
                    </SwiperSlide>
                    <SwiperSlide className='flex items-center'>
                        <Link>
                            <img src="/images/mainweb/3D/Secs5/Elabord.png" className='max-h-16 opacity-50 hover:opacity-100 transition-all' alt="" />
                        </Link>
                    </SwiperSlide>
                    <SwiperSlide className='flex items-center'>
                        <Link>
                            <img src="/images/mainweb/3D/Secs5/Elabord.png" className='max-h-16 opacity-50 hover:opacity-100 transition-all' alt="" />
                        </Link>
                    </SwiperSlide>
                    <SwiperSlide className='flex items-center'>
                        <Link>
                            <img src="/images/mainweb/3D/Secs5/Elabord.png" className='max-h-16 opacity-50 hover:opacity-100 transition-all' alt="" />
                        </Link>
                    </SwiperSlide>
                </Swiper>
                <button className='hidden sm:flex arrow-left h-7 w-10 lg:h-9 lg:w-12 rounded-lg border border-main-blue-web items-center justify-center text-main-blue-web text-lg hover:bg-main-blue-web hover:text-white transition-colors'>
                    <i class="bi bi-arrow-left-short"></i>
                </button>
            </div>
        </div>
    )
}
