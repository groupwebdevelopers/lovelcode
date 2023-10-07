import React, { useRef } from "react";
import NewsSlider from "./NewsSlider";

export default function NewsAndArticles() {
  return (
    <div className=" w-auto h-[38rem] md:h-[43rem] bg-gradient-to-l from-[#004EC2] to-[#8B18FF] flex flex-col items-center mt-20 mb-40 lg:px-36 px-16">
      <div className="container w-full h-40 flex py-4 justify-center relative">
        <div className="flex flex-col text-main-light-web items-center gap-2 ">
          <h3 className="font-Ray-ExtraBold text-xl md:text-2xl lg:text-[30px]  lg:mt-10 mt-4">
            آخرین اخبار و مقالات
          </h3>
          <p className="font-Ray-Medium text-xs md:text-lg lg:text-xl">
            در بلاگ لاوکد کد شما به آرشیو اخبار در مورد دنیای سایت و تکنولوژی
            دسترسی خواهید داشت.
          </p>
        </div>
        <button className="flex absolute top-[5.5rem] -left-5 lg:top-16 lg:left-10 items-center justify-center gap-2 w-24 h-10 lg:w-28 lg:h-12 bg-white rounded-xl font-Ray-Light text-sm text-main-dark-text-web">
          <span>رفتن به بلاگ</span>
          <i className="bi bi-arrow-left-short text-lg mt-2"></i>
        </button>
      </div>
      <NewsSlider />
    </div>
  );
}
