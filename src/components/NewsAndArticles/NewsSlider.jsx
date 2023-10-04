import React, { useRef, useState } from "react";
import { Swiper, SwiperSlide } from "swiper/react";
import "swiper/css";
import "swiper/css/pagination";
import "swiper/css/navigation";
import { Pagination } from "swiper/modules";
import { Autoplay, Navigation } from "swiper/modules";

const data = [
  {
    img: "images/mainweb/3D/Sec6/image 4.png",
    title: "نصب جاوا اسکریپ در اندروید با 4 روش جدید 2023",
    id: 1,
    para: "چطوری باید جاوا اسکریپت رو بر روی اندروید نصب کنیم ؟ همونطور که می‌دونید در دنیای برنامه نویسی همانند دنیای واقعی، با زبان‌های بسیاری روبه‌رو هستیم. یکی از محبوب‌ترین زبان های زبان جاوا اس...",
  },
  {
    img: "images/mainweb/3D/Sec6/image 5.png",
    title: "بک اند چیست؟",
    id: 2,
    para: "بک اند (Back End) چیست؟ چه اتفاقاتی اون پشت میفته؟ بک اند چیست؟ اگه به حوزه برنامه نویسی وب علاقه داشته باشید، حتما کلمه های فرانت اند و بک اند…...",
  },
  {
    img: "images/mainweb/3D/Sec6/image 6.png",
    title: "پایتون چیست ؟ چرا پایتون شایسته ترین زبان برنامه نویسی است؟",
    id: 3,
    para: "چطوری باید جاوا اسکریپت رو بر روی اندروید نصب کنیم ؟ همونطور که می‌دونید در دنیای برنامه نویسی همانند دنیای واقعی، با زبان‌های بسیاری روبه‌رو هستیم. یکی از محبوب‌ترین زبان های زبان جاوا اس...",
  },
  {
    img: "images/mainweb/3D/Sec6/image 4.png",
    title: "نصب جاوا اسکریپ در اندروید با 4 روش جدید 2023",
    id: 4,
    para: "چطوری باید جاوا اسکریپت رو بر روی اندروید نصب کنیم ؟ همونطور که می‌دونید در دنیای برنامه نویسی همانند دنیای واقعی، با زبان‌های بسیاری روبه‌رو هستیم. یکی از محبوب‌ترین زبان های زبان جاوا اس...",
  },
  {
    img: "images/mainweb/3D/Sec6/image 5.png",
    title: "بک اند چیست؟",
    id: 5,
    para: "بک اند (Back End) چیست؟ چه اتفاقاتی اون پشت میفته؟ بک اند چیست؟ اگه به حوزه برنامه نویسی وب علاقه داشته باشید، حتما کلمه های فرانت اند و بک اند…...",
  },
  {
    img: "images/mainweb/3D/Sec6/image 6.png",
    title: "پایتون چیست ؟ چرا پایتون شایسته ترین زبان برنامه نویسی است؟",
    id: 6,
    para: "چطوری باید جاوا اسکریپت رو بر روی اندروید نصب کنیم ؟ همونطور که می‌دونید در دنیای برنامه نویسی همانند دنیای واقعی، با زبان‌های بسیاری روبه‌رو هستیم. یکی از محبوب‌ترین زبان های زبان جاوا اس...",
  },
];
const NewsSlider = () => {
  const swiperRef = useRef(null);
  console.log(swiperRef.current);
  return (
    <div className="w-full auto relative lg:mt-10 mt-4">
      <button
        onClick={() => swiperRef.current.slidePrev()}
        className="prevEl absolute w-12 h-40 lg:w-16 lg:h-16 top-28 -right-16 lg:-right-24 z-50  bg-white lg:rounded-full rounded-l-full  text-violet-600 opacity-90 "
      >
        <i className="bi bi-chevron-compact-right text-2xl"></i>
      </button>
      <button
        onClick={() => swiperRef.current.slideNext()}
        className="nextEl absolute w-12 h-40 lg:w-16 lg:h-16 top-28 -left-16 lg:-left-24 z-50  bg-white lg:rounded-full rounded-r-full  text-violet-600 opacity-90  "
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
          640: {
            slidesPerView: 1,
            spaceBetween: 20,
          },
          768: {
            slidesPerView: 2,
            spaceBetween: 30,
          },
          1024: {
            slidesPerView: 3,
            spaceBetween: 30,
          },
          1280: {
            slidesPerView: 3,
            spaceBetween: 30,
          },
        }}
        loop={true}
        autoplay={{
          delay: 2500,
          disableOnInteraction: true,
        }}
        modules={[Autoplay, Navigation]}
        className="w-full h-full"
      >
        {data.map((item) => {
          return (
            <SwiperSlide className="w-full h-auto  " key={item.id}>
              <div className="w-auto h-full bg-white rounded-3xl flex flex-col items-center ">
                <div className="w-auto h-full flex flex-col items-center pt-4 px-4">
                  <img
                    src={item.img}
                    alt={item.title}
                    className="object-cover "
                  />
                  <div className="w-full flex flex-col items-start mt-4 text-second-gray-text-web">
                    <h2 className="font-Ray-ExtraBold text-main-dark-text-web">
                      {item.title}
                    </h2>
                    <p className="font-Ray-Medium mt-2 h-20">{item.para}</p>
                  </div>
                  <div className="w-full h-auto flex justify-start gap-8 mt-8">
                    <div className="flex gap-2 items-center">
                      <img
                        src="images/mainweb/Icons/Group 2226.png"
                        alt="png"
                        className="w-4 h-4"
                      />
                      <span className="font-Ray-Light">علیرضا رحمانی</span>
                    </div>
                    <div className="flex gap-2 items-center">
                      <img
                        src="images/mainweb/Icons/Group 2226.png"
                        alt="png"
                        className="w-4 h-4"
                      />
                      <span className="font-Ray-Light">1402/07/07</span>
                    </div>
                  </div>
                </div>
                <button className="w-full h-14 flex justify-center items-center gap-2 mt-8 border-t border-solid text-second-gray-text-web hover:bg-main-blue-web hover:text-main-bg-web transition-all duration-300 rounded-b-2xl">
                  <img
                    src="../../../public/images/mainweb/Icons/Group 2191.png"
                    alt=""
                    className="w-4"
                  />
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
