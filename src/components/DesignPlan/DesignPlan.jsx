import React from "react";
import { Link } from "react-router-dom";
export default function DesignPlan({color,vector,title,price,list}) {
  return (
    <>
      <div className="flex justify-center">
        <div className="bg-white rounded-[30px] p-4 flex flex-col items-center justify-between w-[270px] h-[664px]">
          <div>
          <div className={`relative ${color}  w-[239px] h-48 rounded-[20px] overflow-hidden`}>
            <img
              className="absolute top-7 left-1/2 -translate-x-1/2"
              src="./images/mainweb/3D/Sec2/5.png"
              alt=""
            />
            <img
              className="absolute w-[187px] top-5 left-1/2 -translate-x-1/2"
              src={vector}
              alt=""
            />
          </div>
          <div className="main text-[18px] flex flex-col items-center mt-4">
            <h3 className="font-Ray-Black text-main-dark-text-web">
              {title}
            </h3>
            <h3 className="font-ANJOMANFANUM-ULTRABOLD text-main-red-web">
              {price}
            </h3>
          </div>
          <div className="ul flex flex-col items-start justify-center text-main-dark-text-web mt-6">
            {
              list.map((li)=>(
            <div key={li.id} className="flex items-center gap-1">
              <i className={`bi ${li.checked} ${li.checked === 'bi-x' ? 'text-main-red-web text-2xl' : 'text-main-green-web'} text-xl`}></i>
              <p className="text-[14px] font-Ray-ExtraBold">
                {li.title}
              </p>
            </div>
              ))
            }
            
          </div>
          </div>
          <Link className="flex items-center justify-center gap-2 bg-main-blue-web text-white px-14 py-[14px] rounded-[12px] " to={'#'}>
            <button className="text-[14px] font-Ray-Bold">ثبت درخاست</button>
            <i class="bi bi-arrow-left"></i>
          </Link>
        </div>
      </div>
    </>
  );
}
