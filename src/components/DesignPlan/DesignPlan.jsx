import React from "react";
import { Link } from "react-router-dom";
export default function DesignPlan() {
  return (
    <>
      <div className="flex justify-center">
        <div className="bg-white rounded-[30px] p-4 flex flex-col items-center w-[270px]">
          <div className="relative bg-main-violet-web w-[239px] h-48 rounded-[20px] overflow-hidden">
            <img
              className="absolute top-7 left-1/2 -translate-x-1/2"
              src="./images/mainweb/3D/Sec2/5.png"
              alt=""
            />
            <img
              className="absolute w-[187px] top-5 left-1/2 -translate-x-1/2"
              src="./images/mainweb/3D/Sec2/4.png"
              alt=""
            />
          </div>
          <div className="main text-[18px] flex flex-col items-center mt-4">
            <h3 className="font-Ray-Black text-main-dark-text-web">
              طرحی ui/ux از
            </h3>
            <h3 className="font-ANJOMANFANUM-ULTRABOLD text-main-red-web">
              5,000,000 تومان
            </h3>
          </div>
          <div className="ul text-main-dark-text-web mt-6">
            <div className="flex items-center gap-1">
              <i className="bi bi-check-lg text-green-600 text-xl"></i>
              <p className="text-[14px] font-Ray-ExtraBold">
                طراحی UI/UX بدون کد نویسی بک اند
              </p>
            </div>
            <div className="flex items-center gap-1">
              <i className="bi bi-check-lg text-green-600 text-xl"></i>
              <p className="text-[14px] font-Ray-ExtraBold">
                پیاده سازی مطابق اصول Core Wen Vitals
              </p>
            </div>
            <div className="flex items-center gap-1">
              <i className="bi bi-check-lg text-green-600 text-xl"></i>
              <p className="text-[14px] font-Ray-ExtraBold">
                طراحی مجزای نسخه موبایل و دکستاپ
              </p>
            </div>
          </div>
          <Link className="flex items-center justify-center gap-2 bg-main-blue-web text-white px-14 py-[14px] rounded-[12px] mt-[182px]" to={'#'}>
            <button className="text-[14px] font-Ray-Bold">ثبت درخاست</button>
            <i class="bi bi-arrow-left"></i>
          </Link>
        </div>
      </div>
    </>
  );
}
