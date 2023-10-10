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
          <button className="flex justify-center items-center w-24 h-10 lg:w-28 lg:h-12 bg-white rounded-xl font-Ray-Light text-sm text-main-dark-text-web lg:mt-6 lg:ml-[4.6rem]">
            <span>رفتن به بلاگ</span>
            <i className="bi bi-arrow-left-short text-lg mt-2"></i>
          </button>
        </div>
      </div>
      <NewsSlider />
    </div>
  );
}