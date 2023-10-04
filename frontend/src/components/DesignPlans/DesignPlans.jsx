import React from "react";
import DesignPlan from "../DesignPlan/DesignPlan";

export default function DesignPlans() {
  const boxses = [
    {
      id:1,
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
      id:2,
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
      id:3,
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
      id:4,
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
    <div className="container grid grid-cols-1 md:grid-cols-2 gap-5 lg:grid-cols-4 xl:grid-cols-4">
      {boxses.map((item)=>(
        <div key={item.id}>
          <DesignPlan color={item.color} vector={item.vector} title={item.title} price={item.price} list={item.list}></DesignPlan>
        </div>
      ))}
    </div>
  );
}
