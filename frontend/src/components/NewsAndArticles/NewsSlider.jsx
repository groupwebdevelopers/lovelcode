import React, { useEffect, useRef, useState } from "react";
import { Swiper, SwiperSlide } from "swiper/react";
import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/navigation";
import { Pagination } from "swiper/modules";
import { Autoplay, Navigation } from "swiper/modules";

const NewsSlider = () => {
  const [articles, setArticles] = useState([]);
  useEffect(() => {
    fetch("https://thlearn.iran.liara.run/api/v1/blog/get-all?page=1")
      .then((res) => {
        return res.json();
      })
      .then((res) => {
        setArticles(res.data);
        console.log(res.data)
      });
  }, []);
  const swiperRef = useRef(null);
  return (
    <div className="container lg:mt-10 mt-4 lg:px-24 2xl:relative">
      <button
        onClick={() => swiperRef.current.slidePrev()}
        className="prevEl absolute w-12 h-40 lg:w-12 lg:h-12 top-60 lg:top-72 2xl:top-52 right-0 lg:right-4 2xl:right-2  z-50  bg-white lg:rounded-2xl rounded-l-full  text-violet-600 opacity-80 "
      >
        <i className="bi bi-chevron-compact-right text-2xl"></i>
      </button>
      <button
        onClick={() => swiperRef.current.slideNext()}
        className="nextEl absolute w-12 h-40 lg:w-12 lg:h-12 top-60 lg:top-72 left-0 lg:left-4 2xl:top-52 2xl:left-2  z-50  bg-white lg:rounded-2xl rounded-r-full  text-violet-600 opacity-80  "
      >
        <i className=" bi bi-chevron-compact-left text-2xl"></i>
      </button>
      <Swiper
        ref={swiperRef}
        navigation={{
          prevEl: ".prevEl",
          nextEl: ".nextEl",
        }}
        breakpoints={{
          430: {
            slidesPerView: 1.5,
            spaceBetween: 20,
          },
          640: {
            slidesPerView: 2,
            spaceBetween: 30,
          },
          1024: {
            slidesPerView: 2.5,
            spaceBetween: 30,
          },
          1280: {
            slidesPerView: 3,
            spaceBetween: 30,
          },
          1760: {
            slidesPerView: 4,
            spaceBetween: 30,
          },
        }}
        loop={true}
        autoplay={{
          delay: 2500,
          disableOnInteraction: true,
        }}
        modules={[Autoplay, Navigation]}
      >
        {articles.map((item) => {
          return (
            <SwiperSlide className="w-full h-auto" key={item.title}>
              <div className="w-auto h-full bg-white rounded-3xl flex flex-col items-center shadow-md shadow-slate-600">
                <div className="w-full h-full flex flex-col items-center pt-4 px-4">
                  <img
                    src={item.imagePath}
                    alt={item.title}
                    className="object-cover xl:w-full xl:h-64 xl:rounded-3xl"
                  />
                  <div className="w-full flex flex-col items-start mt-4 text-second-gray-text-web">
                    <h2 className="font-Ray-ExtraBold text-main-dark-text-web h-12 md:h-8">
                      {item.title}
                    </h2>
                    <p className="font-Ray-Medium mt-2 h-32 md:h-28 text-sm md:text-base">
                      {}
                    </p>
                  </div>
                  <div className="w-full h-auto flex justify-start gap-8 ">
                    <div className="flex gap-2 items-center">
                      <img
                        src="images/mainweb/Icons/Group 2226.png"
                        alt="png"
                        className="w-4 h-4"
                      />
                      <span className="font-Ray-Light text-sm md:text-base">
                        علیرضا رحمانی
                      </span>
                    </div>
                    <div className="flex gap-2 items-center">
                      <img
                        src="images/mainweb/Icons/Group 2226.png"
                        alt="png"
                        className="w-4 h-4"
                      />
                      <span className="font-Ray-Light text-sm md:text-base">
                        1402/07/07
                      </span>
                    </div>
                  </div>
                </div>
                <button className="w-full h-14 flex justify-center items-center gap-2 mt-2 md:mt-4 border-t border-solid text-second-gray-text-web hover:bg-main-blue-web hover:text-main-bg-web transition-all duration-300 rounded-b-2xl">
                  <i className="bi bi-search"></i>
                  <span className="font-Ray-Medium ">مطالعه مقاله</span>
                </button>
              </div>
            </SwiperSlide>
          );
        })}
      </Swiper>
    </div>
  );
};

export default NewsSlider;
