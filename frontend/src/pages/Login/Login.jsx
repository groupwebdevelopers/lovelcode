import React, { useEffect, useState } from "react";

export default function Login() {
  const [email, setEmail] = useState("");
  const emailChangHandler = (e) => {
    setEmail(e.target.value);
  };
  const [password, setPassword] = useState("");
  const passwordChangHandler = (e) => {
    setPassword(e.target.value);
  };
  const submitHandler = () => {
    const newUser = {
      email ,
      password 
    };
    fetch("http://127.0.0.1:3000/api/v1/signin", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newUser),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("پاسخ دریافت شد:", data);
        // در اینجا می‌توانید پاسخ را پردازش کنید.
      })
      .catch((error) => {
        console.error("خطا:", error);
        // در اینجا می‌توانید با خطاها برخورد کنید و پیام‌های خطا را نمایش دهید.
      });
  };
  return (
    <div className="min-h-screen w-full flex items-center justify-center sm:my-12">
      <div className="flex flex-col w-full sm:w-96 bg-white sm:rounded-3xl p-10">
        <div className="w-full flex flex-col items-center gap-4 mb-6 sm:mb-12">
          <img src="/images/mainweb/3D/Sec1/path28.svg" alt="" />
          <p className="text-xl font-Ray-ExtraBold text-main-blue-web">
            ورود به لاول کد
          </p>
        </div>
        <div className="flex flex-col gap-y-5 mb-6">
          <div>
            <input
              value={email}
              onChange={emailChangHandler}
              type="text"
              className="bg-gray-normal rounded-xl w-full h-10 sm:h-12 outline-none font-Ray-Medium text-sm px-5"
              placeholder="نشانی ایمیل"
            />
          </div>
          <div className="relative">
            <input
              value={password}
              onChange={passwordChangHandler}
              type="text"
              className="bg-gray-normal rounded-xl w-full h-10 sm:h-12 outline-none font-Ray-Medium text-sm px-5"
              placeholder="رمز عبور"
            />
            <i class="absolute top-0 bottom-0 left-4 flex items-center opacity-60 cursor-pointer bi bi-eye"></i>
          </div>
        </div>
        <div className="flex gap-2 text-main-gray-text-web mb-6">
          <input type="checkbox" name="" id="" />
          <span className="text-xs">مرا به خاطر بسپار</span>
        </div>
        <div onClick={submitHandler} className="bg-main-blue-web h-10 sm:h-12 rounded-xl flex justify-center items-center text-white text-xs sm:text-sm gap-2 mb-7">
          وارد شوید
          <i class="text-lg bi bi-arrow-left-short"></i>
        </div>
        <div className="grid grid-cols-2 gap-2 sm:gap-6 h-10 sm:h-12 mb-6 sm:mb-9">
          <div className="bg-gray-normal rounded-xl text-xs flex justify-center items-center text-main-gray-text-web font-Ray-Bold">
            ارسال کد یک بار مصرف
          </div>
          <div className="bg-gray-normal rounded-xl text-xs flex justify-center items-center text-main-gray-text-web font-Ray-Bold">
            فراموشی رمز عبور
          </div>
        </div>
        <span className="text-sm font-Ray-Bold text-main-dark-text-web text-center mb-5">
          هنوز حساب کاربری نساخته اید؟{" "}
          <a href="" className="text-main-blue-web">
            ایجاد حساب کاربری
          </a>
        </span>
        <div className="flex items-center justify-between mb-5">
          <div className="h-px w-full grow bg-main-dark-text-web"></div>
          <span className="shrink-0 text-sm text-main-gray-text-web px-2">
            ورود با
          </span>
          <div className="h-px w-full grow bg-main-dark-text-web"></div>
        </div>
        <div className="h-10 sm:h-12 bg-gray-normal rounded-xl flex justify-center items-center text-sm text-main-gray-text-web font-Ray-Bold gap-2">
          <i class="bi bi-google flex items-center justify-center"></i>
          ورود با گوگل
        </div>
      </div>
    </div>
  );
}
