import React, { useState } from "react";

const data = [
  {
    name: "محمدرضا گودرزی",
    img: "/images/profile/photo_2023-09-28_23-12-16.jpg",
    expert: "front",
  },
  {
    name: "مهدی دلاور",
    img: "/images/profile/IMG_20200820_103224-gradient-2020_11_06_13_04_28.jpg.jpg",
    expert: "front",
  },
  {
    name: "علیرضا رضایی",
    img: "/images/profile/photo_2023-09-28_23-12-50.jpg",
    expert: "front",
  },
  {
    name: " عزیزالله پاینده ",
    img: "/images/profile/photo_2023-09-30_14-45-03.jpg",
    expert: "front",
  },
  {
    name: "محمدرضا گودرزی",
    img: "/images/profile/photo_2023-09-28_23-12-16.jpg",
    expert: "back",
  },
  {
    name: "محمدرضا گودرزی",
    img: "/images/profile/photo_2023-09-28_23-12-16.jpg",
    expert: "back",
  },
  {
    name: "محمدرضا گودرزی",
    img: "/images/profile/photo_2023-09-28_23-12-16.jpg",
    expert: "designer",
  },
];

export default function Introduction() {
  const [active, setActive] = useState("front");
  const filteredItems = data.filter((item) => item.expert === active);
  return (
    <div className="w-full py-36  lg:px-36 md:py-48 lg:py-40 xl:py-48 container">
      <div className="w-full h-auto flex font-Ray-ExtraBold text-3xl justify-center gap-1">
        <span className="text-main-dark-text-web">معرفی اعضای تیم </span>
        <span className="text-main-blue-web">لاول کد </span>
      </div>
      <div className="w-full flex justify-center">
        <span className="text-main-green-web mt-2">کنارتون خوشحالیم</span>
      </div>
      <div className=" w-full h-auto lg:mt-16 mt-12">
        <div className="w-full  flex lg:flex-row lg:justify-between flex-col items-center">
          <div className="md:w-[30rem]  h-12 text-main-dark-text-web cursor-pointer flex justify-between items-center  bg-white rounded-xl">
            <span
              onClick={() => setActive("front")}
              className={`h-full px-4 font-Ray-Light lg:text-base text-[10px] md:text-sm  rounded-xl flex items-center transition-all duration-300 ${
                active === "front"
                  ? "bg-blue-500 text-white"
                  : "bg-white text-main-dark-text-web"
              }`}
            >
              برنامه نویسیان فرانت اند
            </span>
            <span
              onClick={() => setActive("back")}
              className={`h-full px-4 font-Ray-Light lg:text-base text-[10px] md:text-sm  rounded-xl flex items-center transition-all duration-300 ${
                active === "back"
                  ? "bg-blue-500 text-white"
                  : "bg-white text-main-dark-text-web"
              }`}
            >
              برنامه نویسان بک اند
            </span>
            <span
              onClick={() => setActive("designer")}
              className={`h-full px-4 font-Ray-Light lg:text-base text-[10px] md:text-sm  rounded-xl flex items-center transition-all duration-300 ${
                active === "designer"
                  ? "bg-blue-500 text-white"
                  : "bg-white text-main-dark-text-web"
              }`}
            >
              طراحان Ui/Ux
            </span>
          </div>
          <div className="hidden lg:inline">
            <button className="w-44 h-12 bg-white cursor-pointer flex font-Ray-Light text-xs lg:text-base justify-center items-center gap-2 rounded-xl hover:bg-main-blue-web hover:text-white transition-all duration-300">
              <span>مشاهده همه اعضای تیم</span>
              <i className="bi bi-arrow-left-short text-lg mt-2"></i>
            </button>
          </div>
        </div>
        <div className=" w-full px-5 md:px-0 h-auto grid grid-cols-12 mt-14 gap-y-10 lg:gap-12 pb-10 justify-center lg:flex lg:justify-center">
          {filteredItems.map((item) => {
            return (
              <div className=" flex flex-col items-center col-span-6 lg:col-span-3">
                <div className="rounded-full lg:w-48 lg:h-48 w-32 h-32 bg-red-400">
                  <img
                    className="w-full h-full rounded-full object-cover"
                    src={item.img}
                    alt="profile"
                  />
                </div>
                <span className="font-Ray-ExtraBold text-main-dark-text-web mt-3">
                  {item.name}
                </span>
                <span className="font-Ray-Light text-main-gray-text-web">
                  5 سال سابقه کار
                </span>
                <div className="flex">
                  <img
                    className="w-6 h-6 "
                    src="/images/mainweb/3D/Sec5/Instagram1.png"
                    alt="instagram"
                  />
                  <img
                    className="w-6 h-6 "
                    src="/images/mainweb/3D/Sec5/telegram.png"
                    alt="telegram"
                  />
                </div>
              </div>
            );
          })}
        </div>
        <div className="lg:hidden flex w-full justify-center">
            <button className="w-44 h-12 bg-white cursor-pointer flex font-Ray-Light text-xs lg:text-base justify-center items-center gap-2 rounded-xl hover:bg-main-blue-web hover:text-white transition-all duration-300">
              <span>مشاهده همه اعضای تیم</span>
              <i className="bi bi-arrow-left-short text-lg mt-2"></i>
            </button>
          </div>
      </div>
    </div>
  );
}
