import React from "react";
import NavBar from "../../components/NavBar/NavBar";
import SocialNetworks from "../../components/SocialNetworks/SocialNetworks";
import Footer from "../../components/Footer/Footer";
export default function ContactUs() {
  return (
    <div>
      <NavBar />
      <div className="container w-full flex flex-col items-center">
        <h1 className="text-main-dark-text-web font-Ray-ExtraBold text-xl md:text-2xl lg:text-3xl">
          با پر کردن فرم زیر با ما در تماس باش
        </h1>
        <p className="text-main-green-web font-rey-bold lg:text-base md:text-sm text-xs mt-2">
          نظر و پیشنهاداتت رو حتما باهامون در میون بزار :)
        </p>
        <form className="grid grid-cols-12 justify-center gap-4 w-auto px-6 md:px-0 md:w-[40rem] mt-12">
          <input
            type="text"
            className="col-span-12 p-4 rounded-xl outline-none font-Ray-Bold text-sm"
            placeholder="نام نام خانوادگی"
          />
          <input
            type="text"
            className="col-span-12 md:col-span-6 p-4 rounded-xl outline-none font-Ray-Bold text-sm"
            placeholder="ایمیل"
          />
          <input
            type="text"
            className="col-span-12 md:col-span-6 p-4 rounded-xl outline-none font-Ray-Bold text-sm"
            placeholder="شماره تماس"
          />
          <input
            type="text"
            className="col-span-12 p-4 rounded-xl outline-none font-Ray-Bold text-sm"
            placeholder="موضوع پیام"
          />
          <textarea
            className="col-span-12 p-4 h-52 rounded-xl outline-none font-Ray-Bold text-sm"
            placeholder="متن پیام"
          ></textarea>
          <button className="w-32 h-12 rounded-xl bg-main-blue-web text-main-light-web mt-4 font-Ray-Bold text-sm">
            <span>ثبت درخواست</span>
            <i className="bi bi-arrow-left-short text-lg mt-2"></i>
          </button>
        </form>
      </div>
      <div className="container p-2 px-4  mt-14">
        <ul className="w-full flex justify-between md:px-[5rem] lg:px-[8rem] px-3 py-10 rounded-3xl bg-white flex-grow flex-wrap gap-y-4 gap-x-4">
          <li className=" flex flex-col items-start">
            <div className="font-Ray-Bold text-main-blue-web text-lg flex gap-2">
              <i class="bi bi-stopwatch"></i>
              <span>ساعت کاری</span>
            </div>
            <span className="font-Ray-Bold text-sm mt-1">
              ساعات کاری از ۱۲ ظهر الی ۲۴ شب
            </span>
          </li>
          <li className=" flex flex-col items-start">
            <div className="font-Ray-Bold text-main-blue-web text-lg flex gap-2">
              <i class="bi bi-envelope"></i>
              <span>ایمیل</span>
            </div>
            <span className="font-Ray-Bold text-sm mt-1">
              info@LovelCode.gmail.com
            </span>
          </li>
          <li className=" flex flex-col items-start">
            <div className="font-Ray-Bold text-main-blue-web text-lg flex gap-2">
              <i class="bi bi-telephone-forward"></i>
              <span>شماره پشتیبان واتساپ</span>
            </div>
            <span className="font-Ray-Bold text-sm mt-1">091212121212</span>
          </li>
          <li className=" flex flex-col items-start">
            <div className="font-Ray-Bold text-main-blue-web text-lg flex gap-2">
              <i class="bi bi-geo-alt"></i>
              <span>آدرس شرکت</span>
            </div>
            <span className="font-Ray-Bold text-sm mt-1">-------</span>
          </li>
        </ul>
      </div>
      <SocialNetworks />
      <Footer />
    </div>
  );
}
