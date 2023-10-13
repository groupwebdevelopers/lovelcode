import React from "react";
import NavBar from "../../components/NavBar/NavBar";
import DesignPlan from "../../components/DesignPlan/DesignPlan";
import {
  Table,
  TableHeader,
  TableColumn,
  TableBody,
  TableRow,
  TableCell,
} from "@nextui-org/react";
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
          <div className="flex justify-center gap-[8px] font-ANJOMANFANUM-SEMIBOLD text-2xl mt-2">
            <h3 className="">09112121212</h3>
            <i class="bi bi-telephone"></i>
          </div>
          <h4 className="font-Ray-Bold text-xl mt-6">
            یا فرم دریافت تماس را پر کن:
          </h4>
          <a className="flex justify-center" href="#">
            <button className="font-Ray-Bold text-sm bg-main-blue-web text-main-light-web px-8 py-3 rounded-xl mt-[8px] flex items-center justify-center gap-2">
              <p>دریافت تماس و مشاوره</p>
              <i class="bi bi-arrow-left"></i>
            </button>
          </a>
        </div>
        <div className="mt-32 mb-36">
          <table className="w-full">
            <thead className="">
              <tr className="grid grid-cols-12 items-center">
                <div className="col-span-2">
                  <th className="font-Ray-ExtraBlack  text-[10px] sm:text-base lg:text-lg">
                    امکانات / پلن‌ها
                  </th>
                </div>
                <div className="col-span-10 flex justify-evenly gap-1 lg:gap-5">
                  <th className="font-Ray-ExtraBold text-[8px] bg-main-green-web text-white rounded-xl px-2 py-1 lg:px-11 lg:py-3 sm:text-base">
                    پلن طراحی پایه
                  </th>
                  <th className="font-Ray-ExtraBold text-[8px] bg-main-red-web text-white rounded-xl px-2 py-1 lg:px-11 lg:py-3 sm:text-base">
                    پلن طراحی پایه
                  </th>
                  <th className="font-Ray-ExtraBold text-[8px] bg-main-yellow-web text-white rounded-xl px-2 py-1 lg:px-11 lg:py-3 sm:text-base">
                    پلن طراحی پایه
                  </th>
                  <th className="font-Ray-ExtraBold text-[8px] bg-main-violet-web text-white rounded-xl px-2 py-1 lg:px-11 lg:py-3 sm:text-base">
                    پلن طراحی پایه
                  </th>
                </div>
              </tr>
            </thead>
          </table>
        </div>
      </div>
    </div>
  );
}
