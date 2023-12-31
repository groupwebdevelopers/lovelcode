import React, { useRef } from "react";
import Portfolio from "../Portfolio/Portfolio";
import { Link } from "react-router-dom";

// Import Swiper React components
import { Swiper, SwiperSlide } from "swiper/react";

// Import Swiper styles
import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/navigation";
// import required modules
import { Autoplay, Pagination } from "swiper/modules";
import { Navigation } from "swiper/modules";

export default function Portfolios() {
  const swiperRef = useRef(null);
  return (
    <div className="overflow-x-hidden mt-[115px]">
      <div className="">
        <div className="col container grid grid-cols-1  xl:grid-cols-12 bg-gradient-to-br from-[#1875FF] to-[#8B18FF] xl:pr-16 py-20 rounded-[55px]">
          <div className="r text-white xl:col-span-4  xl:w-[350px]">
            <div className="text-center xl:text-right">
              <h2 className="xl:font-Ray-ExtraBold xl:text-3xl font-Ray-ExtraBold text-2xl ">
                نمونه کار های تیم لاول کد
              </h2>
              <p className="xl:text-lg xl:font-bold mt-3 text-sm">
                در بازاری که همه حرف میزنند، ترجیح میدهیم چیزی نگوییم، چرا که در
                انتها ما با کارهایمان قضاوت میشویم نه حرف هایمان
              </p>
            </div>
            <div className="flex flex-col md:items-start items-center">
              <div>
                <Link to={'./portfolio'} className="hover:bg-main-blue-web duration-300 hover:text-white flex gap-1 items-center mt-14 bg-white text-main-dark-text-web justify-center xl:w-32 xl:rounded-xl px-2 py-1 rounded-md xl:h-12 xl:font-Ray-Bold xl:text-sm">
                  <p>مشارهده همه</p>
                  <i className="bi bi-arrow-left"></i>
                </Link>
              </div>
              <div className="flex justify-center mr-6 mt-10 gap-2 ">
                <button
                  onClick={() => swiperRef.current.slideNext()}
                  className="bi bi-chevron-right bg-white text-main-dark-text-web w-9 h-9 rounded-[10px] nextEl hover:bg-main-blue-web hover:text-white duration-300"
                ></button>
                <button
                  onClick={() => swiperRef.current.slidePrev()}
                  className="bi bi-chevron-left bg-white text-main-dark-text-web w-9 h-9 rounded-[10px] prevEl hover:bg-main-blue-web hover:text-white duration-300"
                ></button>
              </div>
            </div>
          </div>
          <div className="l xl:col-span-8  px-10 md:px-0">
            <Swiper
              ref={swiperRef}
              // navigation={true}
              navigation={{
                prevEl: ".prevEl",
                nextEl: ".nextEl",
              }}
              autoplay={{
                delay: 2500,
                disableOnInteraction: true,
              }}
              modules={[Navigation, Autoplay]}
              className="mySwiper"
              breakpoints={{
                640: {
                  slidesPerView: 1,
                  spaceBetween: 20,
                },
                768: {
                  slidesPerView: 2,
                  spaceBetween: 30,
                },
                1024: {
                  slidesPerView: 2,
                  spaceBetween: 30,
                },
                1280: {
                  slidesPerView: 2.5,
                  spaceBetween: 30,
                },
              }}
            >
              <SwiperSlide>
                <Portfolio
                  img={"./images/mainweb/3D/Sec3/1.png"}
                  title={"شرکت آسازیست"}
                  desc={
                    "شرکت آسازیست اردم، ارائه دهنده راهکار های ارگانیک برای دفع آفات است. تیم ما افتخار همکاری با این مجموعه بین المللی را داشته است."
                  }
                ></Portfolio>
              </SwiperSlide>
              <SwiperSlide>
                <Portfolio
                  img={"./images/mainweb/3D/Sec3/2.png"}
                  title={"شرکت آسازیست"}
                  desc={
                    "شرکت آسازیست اردم، ارائه دهنده راهکار های ارگانیک برای دفع آفات است. تیم ما افتخار همکاری با این مجموعه بین المللی را داشته است."
                  }
                ></Portfolio>
              </SwiperSlide>
              <SwiperSlide>
                <Portfolio
                  img={"./images/mainweb/3D/Sec3/1.png"}
                  title={"شرکت آسازیست"}
                  desc={
                    "شرکت آسازیست اردم، ارائه دهنده راهکار های ارگانیک برای دفع آفات است. تیم ما افتخار همکاری با این مجموعه بین المللی را داشته است."
                  }
                ></Portfolio>
              </SwiperSlide>
              <SwiperSlide>
                <Portfolio
                  img={"./images/mainweb/3D/Sec3/2.png"}
                  title={"شرکت آسازیست"}
                  desc={
                    "شرکت آسازیست اردم، ارائه دهنده راهکار های ارگانیک برای دفع آفات است. تیم ما افتخار همکاری با این مجموعه بین المللی را داشته است."
                  }
                ></Portfolio>
              </SwiperSlide>
              
            </Swiper>
          </div>
        </div>
        {/* <Portfolio img={'./images/mainweb/3D/Sec3/2.png'} title={'شرکت آسازیست'} desc={'شرکت آسازیست اردم، ارائه دهنده راهکار های ارگانیک برای دفع آفات است. تیم ما افتخار همکاری با این مجموعه بین المللی را داشته است.'}></Portfolio> */}
      </div>
    </div>
  );
}
