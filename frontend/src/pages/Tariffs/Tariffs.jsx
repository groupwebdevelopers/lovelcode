import React, { useEffect, useState } from "react";
import NavBar from "../../components/NavBar/NavBar";
import DesignPlan from "../../components/DesignPlan/DesignPlan";
import apiRequests from "../../Services/Axios/configs";
import Benefits from "../../components/Benefits/Benefits";
import FormTariffs from "../../components/FormTariffs/FormTariffs";
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
  const benefits = [
    {
      icon: "bi-star-fill",
      title: "ثبت رایگان پروژه",
      bg: "bg-main-blue-web",
      desc: "به رایگان و در چند کلیک پروژه خود را ثبت و از متخصصین مختلف پیشنهاد قیمت و زمان طراحی سایت مورد نظرتان را دریافت نمایید.",
    },
    {
      icon: "bi-headset",
      title: "پشتیبانی لحظه ای",
      bg: "bg-main-red-web",
      desc: "از طریق تلفن، ایمیل و گفتگوی آنلاین سوالات خود را مطرح و توسط پنل کاربری اختصاصی هر لحظه از وضعیت پروژه خود مطلع شوید.",
    },
    {
      icon: "bi-magic",
      title: "ظاهر چشم نواز و اختصاصی",
      bg: "bg-main-yellow-web",
      desc: "طراحی ظاهری سایت شما به صورت کاملا اختصاصی و جذاب و کاملا متناسب با حوزه کاری شما انتخاب و اجرا می شود.",
    },
    {
      icon: "bi-shield-check",
      title: "طراحی سایت استاندارد",
      bg: "bg-main-violet-web",
      desc: "تمام استانداردهای جهانی در طراحی سایت شما رعایت می شود تا وب سایت شما برای موتورهای جستجو و کاربران جذاب باشد.",
    },
  ];
  const [plans, setPlans] = useState([]);
  useEffect(() => {
    apiRequests
      .get("/plan/get-all-plan-types")
      .then((res) => setPlans(res.data.data));
  }, []);
  return (
    <div className="">
      <NavBar />
      <div className="container bredcromp"></div>
      <div className="container">
        <div className="flex flex-col-reverse md:flex-row">
          <div className="title text-main-dark-text-web">
            <div className=" text-main-dark-text-web">
              <h1 className="font-Ray-ExtraBold text-xl md:text-2xl lg:text-[30px]">
                قیمت طراحی سایت و فروشگاه اینترنتی
              </h1>
            </div>
            <div className="mt-2">
              <p className="font-Ray-Bold text-sm lg:text-xl text-main-green-web">
                طراحی سایت با قیمت مناسب روز و پکیج های گوناگو
              </p>
            </div>
            <div className="mt-2">
              <p className="font-Ray-Bold text-xs lg:text-lg ">
                در اینجا ما لیست قیمت طراحی سایت تمامی حوزه های طراحی سایت{" "}
                <span className="text-main-blue-web">لاول کد</span> ر را برای
                شما آماده کرده ایم. تنها کافیست که شما حوزه مورد نظر خودتان را
                انتخاب کنید و از هزینه طراحی سایت آن با خبر بشوید. در نظر داشته
                باشید که این قیمت ها، قیمت پایه هستند و برای طراحی سایت اختصاصی
                و وجود امکانات بیشتر می توانید با کارشناسان ما در ارتباط باشید.
              </p>
            </div>
          </div>
          <div>
            <img className="w-80" src="./images/mainweb/3D/Sec2/4.png" />
          </div>
        </div>
        <div className="flex items-center gap-5 mt-8">
          <h3 className="font-Ray-ExtraBold lg:text-2xl text-main-dark-text-web">
            مشاوره با کارشناس ما
          </h3>
          <a
            className="font-ANJOMANFANUM-SEMIBOLD text-white bg-main-blue-web px-3 py-1 lg:px-5 lg:py-3 rounded-lg lg:rounded-xl rounded-br-none flex items-center gap-2 text-sm lg:text-lg"
            href="tel:09305712252"
          >
            <span>09305712252</span>
            <i className="bi bi-telephone"></i>
          </a>
        </div>

        <div className="list-container flex flex-col items-center mt-24 text-center">
          <div className="title ">
            <h2 className="font-Ray-ExtraBold text-xl lg:text-3xl text-main-dark-text-web">
              لیست قیمت طراحی سایت در حوزه های مختلف
            </h2>
            <h4 className="text-main-green-web font-Ray-Bold text-sm lg:text-xl mt-3">
              برای مشاهده قیمت، حوزه طراحی سایت خود را انتخاب کنید !
            </h4>
          </div>
          <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-y-4 gap-x-16 py-9 px-16 bg-white rounded-[30px] mt-7 h-96 overflow-y-scroll md:overflow-auto md:h-auto">
            {plans.map((item) => (
              <a
                className="md:px-14 px-4 py-2 font-Ray-Bold text-xs md:text-sm lg:text-base xl:text-lg bg-main-blue-web rounded-xl text-white  "
                href=""
              >
                {item.translatedName}
              </a>
            ))}
          </div>
        </div>

        <div className="box flex flex-row-reverse flex-wrap justify-center gap-[22px] xl:justify-between mt-[56px] ">
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
        <div>
          <h2 className="font-Ray-ExtraBold text-xl lg:text-3xl text-center mt-24">مزیت ها و تفاوت های <span className="text-main-blue-web">لاول کد</span> با سایر شرکت ها</h2>
        </div>
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 mt-10 gap-5">
          {
            benefits.map((item , index)=>(
              <Benefits bg={item.bg} desc={item.desc} icon={item.icon} title={item.title} key={index}></Benefits>
            ))
          }
        </div>
        <div className="form mt-28">
          <div className="text-center">
            <h2 className="font-Ray-ExtraBold text-xl md:text-2xl lg:text-3xl text-main-dark-text-web  ">فرم سفارش طراحی سایت</h2>
            <h4 className="text-main-green-web text-sm md:text-base lg:text-xl font-Ray-Bold">اگه از لیست های بالا نتونسی پلن مورد نظرت رو پیدا کنی با پر کردن فرم پایین و درخواست شخصی سازی سایت خودت رو برامون توضیح بده :)</h4>
          </div>
          <FormTariffs></FormTariffs>
        </div>
      </div>
    </div>
  );
}
