import React, { useRef } from "react";
import NewsSlider from "./NewsSlider";

export default function NewsAndArticles() {
  return (
    <div className="w-full h-[39rem] relative  bg-gradient-to-l from-[#004EC2] to-[#8B18FF] flex flex-col items-center mt-20 py-2">
      <div className="container w-full flex justify-between px-6  text-main-light-web">
        <div className="flex flex-col h-20 md:h-32 md:w-[85%] md:items-center  justify-center">
          <span className="md:mr-20 font-Ray-ExtraBold text-xl md:text-2xl lg:text-[30px]">آخرین اخبار و مقالات</span>
          <span className="md:mr-20 hidden md:flex">
            در بلاگ لاوکد کد شما به آرشیو اخبار در مورد دنیای سایت و تکنولوژی
            دسترسی خواهید داشت.
          </span>
        </div>
        <div className="flex h-20 items-center">
          <button className="flex justify-center items-center w-24 h-10 lg:w-28 lg:h-12 bg-white rounded-xl font-Ray-Light text-sm text-main-dark-text-web lg:mt-6">
            <span>رفتن به بلاگ</span>
            <i className="bi bi-arrow-left-short text-lg mt-2"></i>
          </button>
        </div>
      </div>
      <NewsSlider />
    </div>
  );
}

{
  /* <div className="container w-full h-40 flex py-4 md:justify-center justify-start relative">
        <div className="flex flex-col text-main-light-web md:items-center gap-2 items-start ">
          <h3 className="font-Ray-ExtraBold text-xl md:text-2xl lg:text-[30px]  lg:mt-10 mt-4">
            آخرین اخبار و مقالات
          </h3>
          <p className="font-Ray-Medium text-xs md:text-lg lg:text-xl hidden md:flex">
            در بلاگ لاوکد کد شما به آرشیو اخبار در مورد دنیای سایت و تکنولوژی
            دسترسی خواهید داشت.
          </p>
        </div>
        <button className="flex absolute top-[4rem] -left-5 lg:top-16 lg:left-10 items-center justify-center gap-2 w-24 h-10 lg:w-28 lg:h-12 bg-white rounded-xl font-Ray-Light text-sm text-main-dark-text-web">
          <span>رفتن به بلاگ</span>
          <i className="bi bi-arrow-left-short text-lg mt-2"></i>
        </button>
      </div>*/
}
