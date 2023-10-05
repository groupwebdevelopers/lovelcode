import React from "react";
import { Link } from "react-router-dom";
export default function Footer() {
  return (
    <>
      <div className="bg-third-gray-text-web">
        <div className="bg-gradient-to-r from-main-blue-web to-main-violet-web pt-12 pb-5 text-white rounded-t-[50px]">
          <div className="container leading-8">
            <div className="grid lg:grid-cols-4 md:grid-cols-2 gap-y-16 grid-cols-1 md: ">
              <div className="flex justify-center">
                <div className="1 flex flex-col  gap-6">
                  <div className="flex items-center gap-x-2.5">
                    <Link
                      to="/"
                      className="flex items-center justify-center shadow-normal lg:h-[55px] lg:w-[55px] h-12 w-12 bg-white rounded-xl lg:rounded-2xl"
                    >
                      <img
                        src="./images/mainweb/3D/Sec1/path28.svg"
                        className="h-7 lg:h-[31px]"
                        alt=""
                      />
                    </Link>
                    <div className="flex flex-col">
                      <Link to="/" className="text-lg font-Ray-Black ">
                        LovelCode
                      </Link>
                      <Link to="/" className="text-sm font-Ray-ExtraBold">
                        لاول کد
                      </Link>
                    </div>
                  </div>
                  <div className="font-Ray-Bold text-sm max-w-[249px] lg:max-w-[200px]">
                    <p>
                      ما یک آژانس دیجیتال مارکتینگ تمام خدمت هستیم که کلیه
                      خدماتی که کسب و کارشما، برای رشد در فضای دیجیتال به آنها
                      نیاز دارد را برنامه ریزی و اجرا میکنیم.
                    </p>
                  </div>
                </div>
              </div>
              <div className="flex justify-center">
                <div className="2 flex flex-col gap-6">
                  <h2 className="font-Ray-ExtraBold text-[22px]">
                    خدمات اصلی ما
                  </h2>
                  <ul className="list-disc lg:mr-4 font-Ray-Bold">
                    <li>طراحی سایت</li>
                    <li>سئو و بهینه سازی سایت</li>
                    <li>پشتیبانی و نگهداری سایت</li>
                    <li>طراحی Ui/Ux</li>
                  </ul>
                </div>
              </div>
              <div className="flex justify-center">
                <div className="3 flex flex-col gap-6">
                  <h2 className=" font-Ray-ExtraBold text-[22px]">
                    خدمات اصلی ما
                  </h2>
                  <ul className="list-disc lg:mr-4 font-Ray-Bold">
                    <li>طراحی سایت</li>
                    <li>سئو و بهینه سازی سایت</li>
                    <li>پشتیبانی و نگهداری سایت</li>
                    <li>طراحی Ui/Ux</li>
                  </ul>
                </div>
              </div>
              <div className="flex justify-center">
                <div className="4 flex flex-col gap-6 ">
                  <h2 className=" font-Ray-ExtraBold text-[22px]">
                    راه های تماس
                  </h2>
                  <ul className="lg:mr-4 font-Ray-Bold">
                    <li className="flex gap-2 items-center">
                      <i className="bi bi-geo-alt"></i>
                      <p>
                        تهران ، شهرک غرب , خیابان غرب , طبقه <span>3</span>،
                        واحد <span>303</span>
                      </p>
                    </li>
                    <li className="flex gap-2 items-center">
                      <i className="bi bi-telephone"></i>
                      <p>09392848554 - 09392848554</p>
                    </li>
                    <li className="flex gap-2 items-center">
                      <i className="bi bi-envelope"></i>
                      <p>LovelCode@gmail.com</p>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
            <div className="down flex flex-col-reverse gap-4 sm:flex-row  justify-between items-center border-t mt-7 py-5">
              <div className="r">
                <p className="font-Ray-Bold text-sm">
                  تمام حقوق مادی و معنوی متعلق به لاول کد می‌باشد.
                </p>
              </div>
              <div className="l flex items-center gap-3">
                <p className="font-Ray-Bold text-sm">همراه ما باشید</p>
                <img
                  className="w-10"
                  src="./images/mainweb/3D/footer/insta48.svg"
                  alt=""
                />
                <img
                  className="w-10"
                  src="./images/mainweb/3D/footer/telegram.svg"
                  alt=""
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
