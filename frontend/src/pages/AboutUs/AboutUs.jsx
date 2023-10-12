import React from "react";
import NavBar from "../../components/NavBar/NavBar";
import Footer from "../../components/Footer/Footer";
import { Link } from "react-router-dom";

const teamMembers = [
  {
    name: "محمد رضا گودرزی",
    experience: "5 سال سابقه کار",
    imgSrc: "/images/profile/photo_2023-09-28_23-12-16.jpg",
  },
  {
    name: "مهدی دلاور",
    experience: "5 سال سابقه کار",
    imgSrc: "/images/profile/IMG_20200820_103224-gradient-2020_11_06_13_04_28.jpg.jpg",
  },
  {
    name: "علیرضا رضایی",
    experience: "5 سال سابقه کار",
    imgSrc: "/images/profile/photo_2023-09-28_23-12-50.jpg",
  },
  {
    name: "عزیزالله پاینده",
    experience: "5 سال سابقه کار",
    imgSrc: "/images/profile/photo_2023-09-30_14-45-03.jpg",
  },
  {
    name: "علی ندرخانی",
    experience: "5 سال سابقه کار",
    imgSrc: "/images/profile/Rectangle 4635.png",
  },
  {
    name: "امیر حسین طباطبایی",
    experience: "5 سال سابقه کار",
    imgSrc: "/images/profile/maAvatar.jfif",
  },
  {
    name: "محمد امین یعقوبی",
    experience: "5 سال سابقه کار",
    imgSrc: "/images/profile/Ellipse 197.png",
  },
  {
    name: "علیرضا رحمانی",
    experience: "5 سال سابقه کار",
    imgSrc: "/images/profile/Rectangle 4636.png",
  },
];

const AboutUs = () => {
  return (
    <div>
      <NavBar />
      <div className="all about-us container">
        <div className="breadcrumb xl:-mt-[55px]">
          <div className="w-full flex gap-1 items-center text-main-dark-text-web font-Ray-Bold text-xs md:text-sm">
            <i className="bi bi-geo-alt text-xs"></i>
            <Link to={"/"}>
              <span className="cursor-pointer">خانه</span>
            </Link>
            <i className="bi bi-chevron-left text-xs"></i>
            <span className="cursor-pointer">درباره ما</span>
          </div>
        </div>

        <div className="glimps">
          <div className="bg-white grid xl:grid-cols-2 rounded-3xl mb-[70px] mt-[50px]">
            <div className="paragraph m-[50px]">
              <h1 className="font-Ray-ExtraBold text-[35px]">
                <span className="text-main-blue-web">لاول کد</span> در یک نگاه
              </h1>
              <p className=" font-Ray-Bold md:text-[18px] xl:text-[19px] 2xl:text-[20px] 2xl:leading-[47px] xl:leading-[30px] mt-[15px]">
                  <span className="text-main-blue-web">لاول کد</span> در سال
                  1402 ساخته شد!
                  <br />
                  طی فعالیت در حوزه­ بازاریابی و برندینگ با افراد و کسب و کارهای
                  کوچک و متوسط زیادی برخورد داشتیم که به دریافت مشاوره و راهکار
                  جهت توسعه و بهبود بیزینس خود نیاز داشتند، پاسخگویی به این نیاز
                  در حوزه ی بازاریابی و برندینگ منجر به شکل گیری{" "}
                  <span className="text-main-blue-web">لاول کد</span> شد. <br />{" "}
                  با در نظر گرفتن دغدغه ­های معمول در کسب و کارهای ایران، تصمیم
                  گرفتیم که دانش، تخصص و تجربه خودمون رو توی بازاریابی،
                  بازاریابی دیجیتال، برندینگ و … را با برندهایی که به دنبال
                  ارتقا جایگاه خودشون هستند، کسب و کارهایی که به فکر بهبود و
                  توسعه هستند و افرادی که به دنبال دانش به روز در این حوزه
                  هستند، به اشتراک بگذاریم و مخاطبارو توی زمینه اجرا، مشاوره و
                  آموزش بازاریابی و برندینگ همراهی کنیم. به گونه ای که انجام
                  اینکار هم برای خودمان خوشحال کننده و رضایت بخش باشه و هم
                  بوتنیم مشتریارو خوشحال و راضی کنیم و همچنین نتیجه کار به گونه
                  ای باشه که مصرف کننده نهایی مشتری رضایت کامل رو داشته باشه.{" "}
                  <br /> این اتفاقات باعث شد که ما نام{" "}
                  <span className="text-main-blue-web">
                    لاول کد (LovelCode)
                  </span>{" "}
                  یعنی دوست داشتنی را انتخاب کنیم. بازاریابی دوست داشتنی ، در
                  واقع به دنبال ارائه خدماتیه که از نقطه نظر علم بازاریابی حرفه
                  ای و تاثیرگذار باشه و از طریق توسعه و پیشرفت کسب و کار منجر به
                  خشنودی مشتریان ختم به خیر بشه :)
                </p>
            </div>

            <div className="imgg m-[50px] hidden xl:block">
              <img
                src="images/mainweb/3D/Sec7/two.png"
                alt=""
                className="w-[600px] 2xl:w-[700px]"
              />
            </div>
          </div>
        </div>

        <div className="thinking xl:flex md:flex md:flex-col md:items-center md:text-center xl:text-right xl:flex-row xl:items-center gap-12 mb-[100px]">
          <div className="img">
            <img
              src="images/mainweb/3D/Sec7/three.png"
              alt=""
              className="md:max-w-[400px]"
            />
          </div>

          <div className="paragraph2 flex flex-col gap-2">
            <span className="text-main-green-web md:text-[16px] font-Ray-Bold mt-[20px] md:mt-[0px] text-[16px]">
              قصه اینجوری شد که
            </span>
            <h1 className="font-Ray-ExtraBold md:text-[30px] text-[27px] ">
              فکر کردیم میتونیم دنیا رو عوض کنیم...
            </h1>
            <p className="font-Ray-Bold text-[18px]">
              آره خنده داره...ولی از چندتا جوون معتاد چه انتظاری دارید :) ما
              اعتیاد داشتیم به موفقیت، رویاهامون شدن انگیزه حرکت و راهنمایی
              داریم به اسم خلاقیت، از کمال طلبی انرژی گرفتیم و تجربه ها موفق
              قبلی مسیرمون رو روشن کرد. پس هم مسیر شدیم و دنبال همسفر میگردیم،
              مشتریانی همراه تا قله و همفکر، فکری از جنس پیشرفت
            </p>
          </div>
        </div>

        <div className="developers">
          <div className="head flex flex-col items-center">
            <h1 className="font-Ray-ExtraBold md:text-[30px] text-[27px] ">
              ما اینجا کنار هم کارمون رو زندگی میکنیم
            </h1>
            <p className="font-Ray-Bold md:text-[18px] text-main-green-web text-[16px] mt-[10px]">
              با خانوادمون آشنا بشید
            </p>
          </div>

          <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 xl:grid-cols-4 mb-[100px] gap-y-[60px] mt-12">
            {teamMembers.map((member, index) => (
              <div key={index} className="flex justify-center">
                <div className=" flex flex-col items-center col-span-6 lg:col-span-3">
                  <div className="rounded-full lg:w-48 lg:h-48 w-32 h-32 mt-[20px]">
                    <img
                      className="w-full h-full rounded-full object-cover"
                      src={member.imgSrc}
                      alt={member.name}
                    />
                  </div>
                  <span className="font-Ray-ExtraBold text-main-dark-text-web mt-3">
                    {member.name}
                  </span>
 

                </div>
              </div>
            ))}
          </div>
        </div>
      </div>
      <Footer />
    </div>
  );
};

export default AboutUs;
