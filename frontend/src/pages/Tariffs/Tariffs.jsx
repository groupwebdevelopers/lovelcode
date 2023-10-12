import React from "react";
import NavBar from "../../components/NavBar/NavBar";
import DesignPlan from "../../components/DesignPlan/DesignPlan";

export default function Tariffs() {
  const boxses = [
    {
      id: 1,
      color: "bg-main-violet-web",
      vector: "./images/mainweb/3D/Sec2/4.png",
      title: "طرحی ui/ux از",
      price: "5,000,000 تومان",
      list: [
        {
          id: 1,
          title: "طراحی UI/UX بدون کد نویسی بک اند",
          checked: "bi-check-lg",
        },
        {
          id: 2,
          title: "پیاده سازی مطابق اصول Core Wen Vitals",
          checked: "bi-check-lg",
        },
        {
          id: 3,
          title: "طراحی مجزای نسخه موبایل و دکستاپ",
          checked: "bi-check-lg",
        },
      ],
    },
    {
      id: 2,
      color: "bg-main-yellow-web",
      vector: "./images/mainweb/3D/Sec2/3.png",
      title: "طراحی سایت اختصتصی از ",
      price: "8,000,000 تومان",
      list: [
        {
          id: 1,
          title: "طراحی قالی اختصاصی",
          checked: "bi-check-lg",
        },
        {
          id: 2,
          title: "کد نویسی اختصاصی و ماژول های سفارش",
          checked: "bi-check-lg",
        },
        {
          id: 3,
          title: "سیستم مدیریت محتوا",
          checked: "bi-check-lg",
        },
        {
          id: 4,
          title: "امکان سفارشی سازی و ارتقا",
          checked: "bi-check-lg",
        },
        {
          id: 5,
          title: "مدیریت آسان محتوا بدون دانش فنی",
          checked: "bi-check-lg",
        },
        {
          id: 6,
          title: "مناسب برای تمامی مشاغل",
          checked: "bi-check-lg",
        },
        {
          id: 7,
          title: "پشتیبانی و هاست NVME رایگان",
          checked: "bi-check-lg",
        },
        {
          id: 8,
          title: "صفحات , محصولات و ایمیل نامحدود",
          checked: "bi-check-lg",
        },
      ],
    },
    {
      id: 3,
      color: "bg-main-red-web",
      vector: "./images/mainweb/3D/Sec2/2.png",
      title: "طراحی سایت آماده فروشگاهی از",
      price: "5,000,000 تومان",
      list: [
        {
          id: 1,
          title: "خرید آنلاین",
          checked: "bi-check-lg",
        },
        {
          id: 2,
          title: "کد نویسی اختصاصی",
          checked: "bi-check-lg",
        },
        {
          id: 3,
          title: "سیستم مدیریت محتوا",
          checked: "bi-check-lg",
        },
        {
          id: 4,
          title: "امکان سفارش سازی و ارتقا",
          checked: "bi-check-lg",
        },
        {
          id: 4,
          title: "مدیریت آسان محتوا بدون دانش فنی",
          checked: "bi-check-lg",
        },
        {
          id: 5,
          title: "مناسب برای تمامی مشاغل",
          checked: "bi-check-lg",
        },
        {
          id: 6,
          title: "پشتیبانی و هاست NVME رایگان",
          checked: "bi-check-lg",
        },
        {
          id: 7,
          title: "صفحات , محصولات و ایمیل نامحدود",
          checked: "bi-check-lg",
        },
      ],
    },
    {
      id: 4,
      color: "bg-main-green-web",
      vector: "./images/mainweb/3D/Sec2/1.png",
      title: "طراحی سایت آماذه پایه از",
      price: "1,200,000 تومان",
      list: [
        {
          id: 1,
          title: "خرید آنلاین",
          checked: "bi-x",
        },
        {
          id: 2,
          title: "کد نویسی اختصاصی",
          checked: "bi-check-lg",
        },
        {
          id: 3,
          title: "سیستم مدیریت محتوا",
          checked: "bi-check-lg",
        },
        {
          id: 4,
          title: "امکان سفارش سازی و ارتقا",
          checked: "bi-check-lg",
        },
        {
          id: 5,
          title: "مدیریت آسان محتوا بدون دانش فنی",
          checked: "bi-check-lg",
        },
        {
          id: 6,
          title: "مناسب برای تمامی مشاغل",
          checked: "bi-check-lg",
        },
        {
          id: 7,
          title: "پشتیبانی و هاست NVME رایگان",
          checked: "bi-check-lg",
        },
        {
          id: 8,
          title: "صفحات , محصولات و ایمیل نامحدود",
          checked: "bi-check-lg",
        },
      ],
    },
  ];
  return (
    <div>
      <NavBar />
      <div className="container bredcromp"></div>
      <div className="container">
        <div className="title text-center">
          <div className=" text-main-dark-text-web">
            <h1 className="font-Ray-ExtraBold text-[30px]">
              قیمت طراحی سایت و فروشگاه اینترنتی
            </h1>
          </div>
          <div className="mt-2">
            <p className="font-Ray-Bold text-xl text-main-green-web">
              تعرفه و امکانات بسته‌های مختلف پرتال را مقایسه کن و با توجه به
              نیازهای خودت یکی رو انتخاب کن.{" "}
            </p>
          </div>
        </div>
        <div className="select hidden font-Ray-Bold lg:flex my-12 bg-white w-fit  rounded-xl items-center text-xs lg:text-base">
          <a
            href="#"
            className="bg-main-blue-web px-[22px] py-[14px] rounded-xl text-white"
          >
            همه
          </a>
          <a href="#" className="px-[22px] py-[14px] rounded-xl">
            سایت فروشگاهی
          </a>
          <a href="#" className="px-[22px] py-[14px] rounded-xl">
            سایت آموزشی
          </a>
          <a href="#" className="px-[22px] py-[14px] rounded-xl">
            سایت شرکتی
          </a>
          <a href="#" className="px-[22px] py-[14px] rounded-xl">
            سایت شخصی
          </a>
          <a href="#" className="px-[22px] py-[14px] rounded-xl">
            دیجیتال مارکتینگ
          </a>
          <a href="#" className="px-[22px] py-[14px] rounded-xl">
            برندینگ
          </a>
          <a href="#" className="px-[22px] py-[14px] rounded-xl">
            سایت خبری
          </a>
          <a href="#" className="px-[22px] py-[14px] rounded-xl">
            ui/ux
          </a>
        </div>

        <div className="lg:hidden">
          <select
            name=""
            id=""
            className="w-36 h-14 p-2 rounded-lg mt-14 outline-none"
          >
            {Array.from({ length: 8 }, (_, i) => i + 1).map((item) => (
              <option key={item.length} value="همه">
                همه
              </option>
            ))}
          </select>
        </div>
        <div className="box flex flex-row-reverse flex-wrap justify-center gap-[22px] xl:justify-between mt-[50px] ">
          {boxses.map((item) => (
            <div key={item.id}>
              <DesignPlan
                color={item.color}
                vector={item.vector}
                title={item.title}
                price={item.price}
                list={item.list}
              ></DesignPlan>
            </div>
          ))}
        </div>
        <div className="hazine text-center mt-[100px] text-main-dark-text-web">
          <h2 className="text-4xl font-Ray-ExtraBold">
            هزینه طراحی سایت{" "}
            <span className="font-ANJOMANFANUM-MEDIUM">1402</span>
          </h2>
          <p className="font-Ray-Bold text-xl text-main-green-web mt-2">
            زمان کمی داریی؟ کارشناسان ما ایجان که بهت مشاوره بدن تا همین امروز
            پروژه طراحی سایت خودت رو شروع کنی.
          </p>
          <h3 className="font-Ray-ExtraBold text-2xl mt-6">
            الان با ما تماس بگیر :
          </h3>
          <div className="flex justify-center gap-[10px] font-ANJOMANFANUM-SEMIBOLD text-2xl mt-2">
            <h3 className="">09112121212</h3>
            <i class="bi bi-telephone"></i>
          </div>
          <h4 className="font-Ray-Bold text-xl mt-6">
            یا فرم دریافت تماس را پر کن:
          </h4>
          <a className="flex justify-center" href="#">
            <button className="font-Ray-Bold text-sm bg-main-blue-web text-main-light-web px-8 py-3 rounded-xl mt-[10px] flex items-center justify-center gap-2">
              <p>دریافت تماس و مشاوره</p>
              <i class="bi bi-arrow-left"></i>
            </button>
          </a>
        </div>
        <div className="tab mb-8 mt-24 container">
          <table className="w-full font-Ray-ExtraBold text-lg">
            <thead>
              <tr className="flex justify-between items-center bg-[#e4e4e4]/40 px-9 py-7 rounded-t-[30px]">
                <td className="w-56">امکانات / پلن‌ها</td>
                <div className="flex justify-between w-full">
                  <td className="px-11 py-3 text-white bg-main-green-web rounded-xl text-base">
                    پلن طراحی پایه
                  </td>
                  <td className="px-11 py-3 text-white bg-main-red-web rounded-xl text-base">
                    پلن طراحی پایه
                  </td>
                  <td className="px-11 py-3 text-white bg-main-yellow-web rounded-xl text-base">
                    پلن طراحی پایه
                  </td>
                  <td className="px-11 py-3 text-white bg-main-violet-web rounded-xl text-base">
                    پلن طراحی پایه
                  </td>
                </div>
              </tr>
            </thead>
            <tbody>
              <tr className="flex justify-between items-center bg-white px-9 py-7">
                <td className="w-56">طراحی تم اختصاصی برند</td>
                <div className="flex justify-between w-full">
                  <td className=" px-11 py-3 text-main-red-web text-3xl flex justify-center items-center">
                    <i class="bi bi-x"></i>
                  </td>
                  <td className=" px-11 py-3 text-main-red-web text-3xl flex justify-center items-center">
                    <i class="bi bi-x"></i>
                  </td>
                  <td className=" px-11 py-3 text-main-green-web text-3xl flex justify-center items-center">
                    <i class="bi bi-check-lg"></i>
                  </td>
                  <td className=" px-11 py-3 text-main-green-web text-3xl flex justify-center items-center">
                    <i class="bi bi-check-lg"></i>
                  </td>
                </div>
              </tr>
              <tr className="flex justify-between items-center bg-[#e4e4e4]/40 px-9 py-7">
                <td className="w-56">زمان تحویل (روز کاری)</td>
                <div className="flex justify-between w-full">
                  <td className=" px-11 py-3 font-Ray-ExtraBold flex justify-center items-center"><span className="font-ANJOMANFANUM-MEDIUM">5</span> روز</td>
                  <td className=" px-11 py-3 font-Ray-ExtraBold flex justify-center items-center"><span className="font-ANJOMANFANUM-MEDIUM">9</span> روز</td>
                  <td className=" px-11 py-3 font-Ray-ExtraBold flex justify-center items-center">
                  <span className="font-ANJOMANFANUM-MEDIUM">12</span> روز
                  </td>
                  <td className=" px-11 py-3 font-Ray-ExtraBold flex justify-center items-center">
                  <span className="font-ANJOMANFANUM-MEDIUM">17</span> روز
                  </td>
                </div>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
}
