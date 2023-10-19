import React, { useEffect, useState } from "react";

export default function Register() {
  // inputs values
  const [email, setEmail] = useState("");
  const [phone, setPhone] = useState("");
  const [password, setPassword] = useState("");

  // events handler
  const emailChangHandler = (e) => {
    setEmail(e.target.value);
  };
  const phoneChangHandler = (e) => {
    setPhone(e.target.value);
  };
  const passwordChangHandler = (e) => {
    setPassword(e.target.value);
  };

  // submit Handler

  const submitHandler = () => {
    const newUser = {
      email,
      password,
      username: 'mohamadreza',
      name: 'mohammadreza',
      family: 'mohammadreza',
    };
    fetch("https://thlearn.iran.liara.run/api/v1/signup", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(newUser),
    })
      .then((response) => response.json())
      .then((data) => {
        console.log("پاسخ دریافت شد:", data);
      })
      .catch((error) => {
        console.error("خطا:", error);
      });
  };

  // use efect
  return (
    <div className="min-h-screen w-full flex items-center justify-center sm:my-12">
      <div className="flex flex-col w-full sm:w-96 bg-white sm:rounded-3xl p-10">
        <div className="w-full flex flex-col items-center gap-4 mb-6 sm:mb-12">
          <img src="/images/mainweb/3D/Sec1/path28.svg" alt="" />
          <p className="text-xl font-Ray-ExtraBold text-main-blue-web">
            ثبت نام در لاول کد
          </p>
        </div>
        <div className="flex flex-col gap-y-5 mb-6">
          <div>
            <input
              onChange={emailChangHandler}
              value={email}
              type="text"
              className="bg-gray-normal rounded-xl w-full h-10 sm:h-12 outline-none font-Ray-Medium text-sm px-5"
              placeholder="نام کاربری و یا ایمیل"
            />
          </div>
          <div>
            <input
              onChange={phoneChangHandler}
              value={phone}
              type="text"
              className="bg-gray-normal rounded-xl w-full h-10 sm:h-12 outline-none font-Ray-Medium text-sm px-5"
              placeholder="شماره تلفن"
            />
          </div>
          <div className="relative">
            <input
              onChange={passwordChangHandler}
              value={password}
              type="text"
              className="bg-gray-normal rounded-xl w-full h-10 sm:h-12 outline-none font-Ray-Medium text-sm px-5"
              placeholder="رمز عبور"
            />
            <i class="absolute top-0 bottom-0 left-4 flex items-center opacity-60 cursor-pointer bi bi-eye"></i>
          </div>
        </div>
        <div className="flex gap-2 text-main-gray-text-web mb-6">
          <input type="checkbox" name="" id="" />
          <span className="text-xs">
            ثبت نام در سایت به منزله پذیرش قوانین و مقررات است.
          </span>
        </div>
        <div
          onClick={submitHandler}
          className="bg-main-blue-web h-10 sm:h-12 rounded-xl flex justify-center items-center text-white text-xs sm:text-sm gap-2 mb-7"
        >
          ارسال کد به شماره
          <i class="text-lg bi bi-arrow-left-short"></i>
        </div>
        <span className="text-sm font-Ray-Bold text-main-dark-text-web text-center mb-5">
          از قبل وارد شده اید؟
          <a href="" className="text-main-blue-web">
            ورود به حساب کاربری
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
