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
    <div className="overflow-x-hidden">
      <div className="container">
        <div className="col grid grid-cols-12 bg-gradient-to-br from-[#1875FF] to-[#8B18FF] px-16 py-20 rounded-[55px]">
          <div className="r text-white col-span-4  w-[263px]">
            <div>
              <h2 className="font-Ray-ExtraBold text-3xl">
                نمونه کار های تیم لاول کد
              </h2>
              <p className="text-lg font-bold mt-3">
                در بازاری که همه حرف میزنند، ترجیح میدهیم چیزی نگوییم، چرا که در
                انتها ما با کارهایمان قضاوت میشویم نه حرف هایمان
              </p>
            </div>
            <div className="flex flex-col items-start">
              <div>
                <Link className="flex gap-1 items-center mt-14 bg-white text-main-dark-text-web justify-center w-32 rounded-xl h-12 font-Ray-Bold text-sm">
                  <p>مشارهده همه</p>
                  <i className="bi bi-arrow-left"></i>
                </Link>
              </div>
              <div className="flex justify-center mr-6 mt-6 gap-2">
                <button
                  onClick={() => swiperRef.current.slideNext()}
                  className="bi bi-chevron-right bg-white text-main-dark-text-web w-9 h-9 rounded-[10px] nextEl"
                ></button>
                <button
                  onClick={() => swiperRef.current.slidePrev()}
                  className="bi bi-chevron-left bg-white text-main-dark-text-web w-9 h-9 rounded-[10px] prevEl"
                ></button>
              </div>
            </div>
          </div>
          <div className="l col-span-8 -translate-x-60">
            <Swiper
              ref={swiperRef}
              // navigation={true}
              loop={true}
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
                  slidesPerView: 2,
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
                1200: {
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
