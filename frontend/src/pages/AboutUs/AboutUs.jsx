import React from "react";
import NavBar from "../../components/NavBar/NavBar";
import Footer from "../../components/Footer/Footer";
import { Link } from "react-router-dom";

export default function AboutUs() {
  return (
    <>
      <div>
        <NavBar />
        <div className="all about-us container">
          <div className="breadcrumb ">
            <div className="w-full flex gap-1 items-center text-main-dark-text-web font-Ray-Bold text-xs md:text-sm">
              <i class="bi bi-geo-alt text-xs"></i>
              <Link to={"/"}>
                <span className="cursor-pointer">خانه </span>
              </Link>
              <i class="bi bi-chevron-left text-xs"></i>
              <span className="cursor-pointer">درباره ما</span>
            </div>
          </div>

          <div className="glimps  ">
            <div className="bg-white grid xl:grid-cols-2 rounded-3xl mb-[100px] mt-[50px]">
              <div className="paragraph m-[50px] ">
                <h1 className="font-Ray-ExtraBold text-[35px]">
                  {" "}
                  <span className="text-main-blue-web">لاول کد</span> در یک نگاه
                </h1>
                <p className=" font-Ray-Bold md:text-[18px] xl:text-[19px] 2xl:text-[20px] 2xl:leading-[47px] mt-[15px]">
                  <span className="text-main-blue-web">لاول کد</span> در سال
                  1402 ساخته شد!
                  <br />
                  طی فعالیت در حوزه­ بازاریابی و برندینگ با افراد و کسب و کارهای
                  کوچک و متوسط زیادی برخورد داشتیم که به دریافت مشاوره و راهکار
                  جهت توسعه و بهبود بیزینس خود نیاز داشتند، پاسخگویی به این نیاز
                  در حوزه ی بازاریابی و برندینگ منجر به شکل گیری{" "}
                  <span className="text-main-blue-web">لاول کد</span> شد. <br /> با در
                  نظر گرفتن دغدغه ­های معمول در کسب و کارهای ایران، تصمیم گرفتیم
                  که دانش، تخصص و تجربه خودمون رو توی بازاریابی، بازاریابی
                  دیجیتال، برندینگ و … را با برندهایی که به دنبال ارتقا جایگاه
                  خودشون هستند، کسب و کارهایی که به فکر بهبود و توسعه هستند و
                  افرادی که به دنبال دانش به روز در این حوزه هستند، به اشتراک
                  بگذاریم و مخاطبارو توی زمینه اجرا، مشاوره و آموزش بازاریابی و
                  برندینگ همراهی کنیم. به گونه ای که انجام اینکار هم برای خودمان
                  خوشحال کننده و رضایت بخش باشه و هم بوتنیم مشتریارو خوشحال و
                  راضی کنیم و همچنین نتیجه کار به گونه ای باشه که مصرف کننده
                  نهایی مشتری رضایت کامل رو داشته باشه. <br /> این اتفاقات باعث شد که
                  ما نام{" "}
                  <span className="text-main-blue-web">
                    لاول کد (LovelCode)
                  </span>{" "}
                  یعنی دوست داشتنی را انتخاب کنیم. بازاریابی دوست داشتنی ، در
                  واقع به دنبال ارائه خدماتیه که از نقطه نظر علم بازاریابی حرفه
                  ای و تاثیرگذار باشه و از طریق توسعه و پیشرفت کسب و کار منجر به
                  خشنودی مشتریان ختم به خیر بشه :)
                </p>
              </div>

              <div className="imgg   m-[50px] hidden xl:block">
                <img src="images/mainweb/3D/Sec7/two.png" alt="" className="w-[600px] 2xl:w-[700px]"/>
              </div>
            </div>
          </div>
        </div>
        <Footer />
      </div>
    </>
  );
}
