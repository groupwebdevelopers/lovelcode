import React from "react";
import { Link } from "react-router-dom";

export default function Portfolio({ img, title, desc }) {
  return (
    <>
      <div className="flex flex-col items-start bg-white  rounded-[30px] overflow-hidden ">
        <div className="p-4">
          <img src={img} alt="" />
        </div>
        <div className="mt-4 px-4">
          <h3 className="text-[18px] font-Ray-ExtraBold text-main-dark-text-web">
            {title}
          </h3>
          <p className="text-main-gray-text-web font-Ray-Bold text-[14px]">
            {desc}
          </p>
        </div>
        <div className="w-full h-full flex text-main-gray-text-web font-Ray-Bold text-[14px] border-t mt-4 hover:bg-main-blue-web hover:text-white duration-300 py-4"> 
            <Link className="w-full flex justify-center items-center gap-2" href="#">
            <i class="bi bi-search"></i>
            <p>مشاهده سایت  {title}</p>
            </Link>
        </div>
      </div>
    </>
  );
}
