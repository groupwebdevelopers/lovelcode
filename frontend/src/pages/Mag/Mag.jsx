import React, { useEffect, useState } from "react";
import NavBar from "../../components/NavBar/NavBar";
import AccordionComp from "./Accordion";

export default function Mag() {
  return (
    <div>
      <NavBar />
      <section className="container ">
        <div className="w-full grid grid-cols-12 gap-5">
          <div className="col-span-12 md:col-span-6 h-[12.5rem] lg:h-[15rem] rounded-3xl relative">
            <img
              className="w-full h-full object-cover rounded-3xl"
              src="/images/profile/benefits-of-deep-breathing-585x390.png"
              alt="img"
            />
            <div className="w-full h-full absolute top-0 rounded-3xl opacity-80 bg-gradient-to-t from-black"></div>
            <div className="w-full h-full absolute top-0 p-5 flex flex-col justify-between">
              <div className="w-44 p-2 bg-white  rounded-xl flex gap-2">
                <img
                  className="w-4"
                  src="../../../public/images/mainweb/Icons/graph 1.png"
                  alt=""
                />
                <span className="font-Ray-Bold text-xs text-main-dark-text-web">
                  دسته بندی : مقالات طراحی سایت
                </span>
              </div>
              <div className="w-44 rounded-xl flex flex-col gap-2 text-slate-50 text-xs font-Ray-Bold">
                <h4 className="text-lg">طراحی سایت واقعیت مجازی</h4>
                <div className="w-full flex gap-2">
                  <img
                    className="w-3 h-4"
                    src="../../../public/images/mainweb/Icons/User 2343.png"
                    alt=""
                  />
                  <span>علیرضا رحمانی</span>
                  <img
                    className="w-4 h-4 text-slate-50"
                    src="../../../public/images/mainweb/Icons/Group 2226.png"
                    alt=""
                  />
                  <span>1402/07/07</span>
                </div>
              </div>
            </div>
          </div>
          {/* 2 */}
          <div className="col-span-12 md:col-span-6 h-[12.5rem] lg:h-[15rem] rounded-3xl relative">
            <img
              className="w-full h-full object-cover rounded-3xl"
              src="/images/profile/benefits-of-deep-breathing-585x390.png"
              alt="img"
            />
            <div className="w-full h-full absolute top-0 rounded-3xl opacity-80 bg-gradient-to-t from-black"></div>
            <div className="w-full h-full absolute top-0 p-5 flex flex-col justify-between">
              <div className="w-44 p-2 bg-white  rounded-xl flex gap-2">
                <img
                  className="w-4"
                  src="../../../public/images/mainweb/Icons/graph 1.png"
                  alt=""
                />
                <span className="font-Ray-Bold text-xs text-main-dark-text-web">
                  دسته بندی : مقالات طراحی سایت
                </span>
              </div>
              <div className="w-44 rounded-xl flex flex-col gap-2 text-slate-50 text-xs font-Ray-Bold">
                <h4 className="text-lg">طراحی سایت واقعیت مجازی</h4>
                <div className="w-full flex gap-2">
                  <img
                    className="w-3 h-4"
                    src="../../../public/images/mainweb/Icons/User 2343.png"
                    alt=""
                  />
                  <span>علیرضا رحمانی</span>
                  <img
                    className="w-4 h-4 text-slate-50"
                    src="../../../public/images/mainweb/Icons/Group 2226.png"
                    alt=""
                  />
                  <span>1402/07/07</span>
                </div>
              </div>
            </div>
          </div>
          {/* 3 */}
          <div className="col-span-12 md:col-span-4 h-[12.5rem] lg:h-[15rem] rounded-3xl relative">
            <img
              className="w-full h-full object-cover rounded-3xl"
              src="/images/profile/benefits-of-deep-breathing-585x390.png"
              alt="img"
            />
            <div className="w-full h-full absolute top-0 rounded-3xl opacity-80 bg-gradient-to-t from-black"></div>
            <div className="w-full h-full absolute top-0 flex flex-col justify-between p-5">
              <div className="w-44 p-2 bg-white rounded-xl flex gap-2">
                <img
                  className="w-4"
                  src="../../../public/images/mainweb/Icons/graph 1.png"
                  alt=""
                />
                <span className="font-Ray-Bold text-xs text-main-dark-text-web">
                  دسته بندی : مقالات طراحی سایت
                </span>
              </div>
              <div className="w-44 rounded-xl flex flex-col gap-2 text-slate-50 text-xs font-Ray-Bold">
                <h4 className="text-lg">طراحی سایت واقعیت مجازی</h4>
                <div className="w-full flex gap-2">
                  <img
                    className="w-3 h-4"
                    src="../../../public/images/mainweb/Icons/User 2343.png"
                    alt=""
                  />
                  <span>علیرضا رحمانی</span>
                  <img
                    className="w-4 h-4 text-slate-50"
                    src="../../../public/images/mainweb/Icons/Group 2226.png"
                    alt=""
                  />
                  <span>1402/07/07</span>
                </div>
              </div>
            </div>
          </div>
          {/* 4 */}
          <div className="col-span-12 md:col-span-4 h-[12.5rem] lg:h-[15rem] rounded-3xl relative">
            <img
              className="w-full h-full object-cover rounded-3xl"
              src="/images/profile/benefits-of-deep-breathing-585x390.png"
              alt="img"
            />
            <div className="w-full h-full absolute top-0 rounded-3xl opacity-80 bg-gradient-to-t from-black"></div>
            <div className="w-full h-full absolute top-0 flex flex-col justify-between p-5">
              <div className="w-44 p-2 bg-white rounded-xl flex gap-2">
                <img
                  className="w-4"
                  src="../../../public/images/mainweb/Icons/graph 1.png"
                  alt=""
                />
                <span className="font-Ray-Bold text-xs text-main-dark-text-web">
                  دسته بندی : مقالات طراحی سایت
                </span>
              </div>
              <div className="w-44 rounded-xl flex flex-col gap-2 text-slate-50 text-xs font-Ray-Bold">
                <h4 className="text-lg">طراحی سایت واقعیت مجازی</h4>
                <div className="w-full flex gap-2">
                  <img
                    className="w-3 h-4"
                    src="../../../public/images/mainweb/Icons/User 2343.png"
                    alt=""
                  />
                  <span>علیرضا رحمانی</span>
                  <img
                    className="w-4 h-4 text-slate-50"
                    src="../../../public/images/mainweb/Icons/Group 2226.png"
                    alt=""
                  />
                  <span>1402/07/07</span>
                </div>
              </div>
            </div>
          </div>
          {/* 5 */}
          <div className="col-span-12 md:col-span-4 h-[12.5rem] lg:h-[15rem] rounded-3xl relative">
            <img
              className="w-full h-full object-cover rounded-3xl"
              src="/images/profile/benefits-of-deep-breathing-585x390.png"
              alt="img"
            />
            <div className="w-full h-full absolute top-0 rounded-3xl opacity-80 bg-gradient-to-t from-black"></div>
            <div className="w-full h-full absolute top-0 flex flex-col justify-between p-5">
              <div className="w-44 p-2 bg-white rounded-xl flex gap-2">
                <img
                  className="w-4"
                  src="../../../public/images/mainweb/Icons/graph 1.png"
                  alt=""
                />
                <span className="font-Ray-Bold text-xs text-main-dark-text-web">
                  دسته بندی : مقالات طراحی سایت
                </span>
              </div>
              <div className="w-44 rounded-xl flex flex-col gap-2 text-slate-50 text-xs font-Ray-Bold">
                <h4 className="text-lg">طراحی سایت واقعیت مجازی</h4>
                <div className="w-full flex gap-2">
                  <img
                    className="w-3 h-4"
                    src="../../../public/images/mainweb/Icons/User 2343.png"
                    alt=""
                  />
                  <span>علیرضا رحمانی</span>
                  <img
                    className="w-4 h-4 text-slate-50"
                    src="../../../public/images/mainweb/Icons/Group 2226.png"
                    alt=""
                  />
                  <span>1402/07/07</span>
                </div>
              </div>
            </div>
          </div>
        </div>
        {/* ------------------------------------- */}
        <div className="w-full grid grid-cols-12 gap-5 mt-12">
          <AccordionComp />
          <div className="col-span-12 md:col-span-8 lg:col-span-9 grid grid-cols-12 gap-5">
            {createEmptyArray(15).map((i) => {
              return (
                <div className="col-span-12  lg:col-span-6 xl:col-span-4 h-[23rem] bg-white rounded-3xl cursor-pointer">
                  <div className="w-full h-[14rem] rounded-t-xl">
                    <img
                      className="w-full h-full rounded-t-xl object-cover"
                      src="/images/profile/file 8.png"
                      alt="img"
                    />
                  </div>
                  <div className="w-full h-[9rem] flex flex-col justify-between px-3 py-5 text-main-dark-text-web">
                    <div className="w-full flex flex-col gap-2">
                      <h1 className="text-base font-Ray-Bold">
                        طراحی سایت و اپلیکیشن با فیگما
                      </h1>
                      <p className="text-xs font-Ray-Bold">
                        فیگما ابزاری برای طراحی رابط کاربری است که میتوان در آن
                        سایت و یا اپلیکیشن به صورت گرافیکی ایجاد کرد.
                      </p>
                    </div>
                    <div className="w-full flex gap-2 items-center text-xs font-Ray-Bold">
                      <img
                        className="w-3 h-4"
                        src="../../../public/images/mainweb/Icons/User 2343.png"
                        alt="user"
                      />
                      <span>علیرضا رحمانی</span>
                      <img
                        className="w-4 h-4 mr-4"
                        src="../../../public/images/mainweb/Icons/Group 2226.png"
                        alt="svg"
                      />
                      <span>1402/07/07</span>
                    </div>
                  </div>
                </div>
              );
            })}
          </div>
        </div>
      </section>
    </div>
  );
}

const createEmptyArray = (count) => {
  return new Array(count).fill("");
};
