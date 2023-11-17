import React from "react";

export default function Benefits({ icon, bg, title, desc }) {
  return (
    <>
      <div className="relative bg-white p-4 rounded-3xl text-main-dark-text-web flex flex-col">
        <div className="top ">
            <div className={`icon w-8 h-8 ${bg} rounded-lg flex justify-center items-center text-white`}>
                <i className={`bi ${icon}`}></i>
            </div>
        </div>
        <div className="down flex flex-col gap-3 mt-3">
            <h3 className="font-Ray-ExtraBold text-sm md:text-xl text-main-dark-text-web">{title}</h3>
            <p className="font-Ray-Bold text-xs md:text-sm">{desc}</p>
        </div>
        {/* <div className="absolute bg-[#f5f8fa] w-7 h-7 top-2 -left-1 rounded-md"></div> */}
      </div>
    </>
  );
}
