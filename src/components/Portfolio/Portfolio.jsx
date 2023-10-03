import React from "react";
import { Link } from "react-router-dom";

export default function Portfolio({ img, title, desc }) {
  return (
    <>
      <div className="flex flex-col items-start bg-white  xl:rounded-[30px] rounded-lg overflow-hidden text-center md:text-right mt-5">
        <div className="lx:p-4 p-1">
          <img src={img} alt="" />
        </div>
        <div className="xl:mt-4 xl:px-4 px-1">
          <h3 className="xl:text-[18px] text-lg font-Ray-Bold xl:font-Ray-ExtraBold text-main-dark-text-web">
            {title}
          </h3>
          <p className="text-main-gray-text-web text-sm font-Ray-Bold xl:font-Ray-Bold xl:text-[14px]">
            {desc}
          </p>
        </div>
        <div className="w-full h-full flex text-main-gray-text-web xl:font-Ray-Bold xl:text-[14px] border-t mt-4 hover:bg-main-blue-web hover:text-white duration-300 text-xs text-center px-2 py-1 font-Ray-Bold xl:py-4"> 
            <Link className="w-full flex justify-center items-center xl:gap-2 gap-1" href="#">
            <i class="bi bi-search"></i>
            <p>مشاهده سایت  {title}</p>
            </Link>
        </div>
      </div>
    </>
  );
}
