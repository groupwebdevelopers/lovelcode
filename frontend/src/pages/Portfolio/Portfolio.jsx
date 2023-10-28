import React, { useEffect, useState } from "react";
import NavBar from "../../components/NavBar/NavBar";
import Footer from "../../components/Footer/Footer";
export default function Portfolio() {
  const [Portfolios, setPortfolios] = useState([]);
  const [selector, setSelector] = useState("all");
  const [selectPortfolio, SetSelectPortfolio] = useState([]);
  const menuSelect = [
    { id: 1, name: "همه", type: "all" },
    { id: 2, name: "سایت فروشگاهی", type: "shop" },
    { id: 3, name: "سایت آموزشی", type: "edu" },
    { id: 4, name: "سایت شرکتی", type: "company" },
    { id: 5, name: "سایت شخصی", type: "Personal" },
    { id: 6, name: "دیجیتال مارکتینگ", type: "Digital" },
    { id: 7, name: "برندینگ", type: "Branding" },
    { id: 8, name: "سایت خبری", type: "news" },
    { id: 9, name: "ui/ux", type: "ui" },
  ];
  useEffect(() => {
    fetch("https://thlearn.iran.liara.run/api/v1/work-sample/get-featured")
      .then((res) => {
        return res.json();
      })
      .then((res) => {
        setPortfolios(res.data);
      });
  }, []);
  useEffect(() => {
    SetSelectPortfolio(
      [...Portfolios].filter((item) => item.type === selector)
    );
    console.log(Portfolios);
    console.log(selectPortfolio);
  }, [selector]);

  return (
    <div className="relative">
      <NavBar />
      <div className="all container text-main-dark-text-web">
        <div className="p">
          <h2 className="font-Ray-ExtraBold text-3xl">نمونه کار طراحی سایت</h2>
          <p className="font-Ray-Medium text-lg">
            نمونه کار طراحی سایت های فروشگاهی، خبری و طراحی سایت آموزشی و ... را
            که در توی این صفحه مشاهده می کنی همه پروژه های اجرا شده کسب و کارهای
            آنلاینی اند که درحال کار هستند. همه این نمونه کارهای طراحی سایت با
            سورسی بهینه و طبق استاندارهای گوگل طراحی شده. این ادعا به پشتوانه
            طراحی بی نظیر و پیاده سازی قوی و باکیفیت سایت های بسیاری برای برندها
            و استارتاپ های معتبر ایران انجام شده.
          </p>
          <p className="font-Ray-Medium text-lg">
            برای آشنایی با برندهایی که تاحالا با اونها همکاری داشته ایم و خدماتی
            که برای هریک انجام دادیم، اطلاعات اون دسته از مشتریامون که اجازه
            بیان اسم پروژه هاشون را به ما دادن و خدماتی که به اونها عرضه کردیم
            رو مختصرا اینجا معرفی کردیم هم برای آشنایی شما و هم قرار گرفتن توی
            لیست افتخارات تیم{" "}
            <span className="text-main-blue-web">لاول کد</span>
          </p>
        </div>
        <div className="flex flex-col items-center lg:items-start">
          <div className="select hidden font-Ray-Bold lg:flex my-12 bg-white w-fit  rounded-xl items-center text-xs lg:text-base">
            {menuSelect.map((item) => (
              <div
                key={item.id}
                onClick={() => setSelector(item.type)}
                className={`${
                  selector === item.type
                    ? "bg-main-blue-web text-white"
                    : "bg-white text-main-dark-text-web"
                } cursor-pointer px-[22px] py-[14px] rounded-xl `}
              >
                {item.name}
              </div>
            ))}
          </div>

          <div className="lg:hidden">
            <select
              className=" p-2 rounded-lg mt-14 outline-none ring"
              onChange={(e) => setSelector(e.target.value)}
            >
              {menuSelect.map((item) => (
                <option key={item.id} value={item.type}>{item.name}</option>
              ))}
            </select>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 mb-[100px] gap-[20px] mt-12">
            {selector === "all"
              ? Portfolios.map((item) => (
                  <a
                    href="#"
                    key={item.title}
                    className="w-[270px] h-[270px] rounded-[25px] bg-white"
                  >
                    <img src={item.imagePath} />
                  </a>
                ))
              : selectPortfolio.map((item) => (
                  <a
                    href="#"
                    key={item.title}
                    className="w-[270px] h-[270px] rounded-[25px] bg-white"
                  >
                    <img src={item.imagePath} />
                  </a>
                ))}
          </div>
        </div>
      </div>
      <Footer></Footer>
    </div>
  );
}
