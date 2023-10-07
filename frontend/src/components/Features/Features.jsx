import React from "react";

function FeatureCard({ imageSrc, title }) {
  return (
    <div className="bg-white md:w-[280px] md:h-[260px] w-[190px] h-[200px] text-center rounded-2xl">
      <img
        src={imageSrc}
        alt=""
        className="md:w-[200px] w-[140px] md:mr-10 mr-[25px]"
      />
      <h1 className="font-Ray-ExtraBold md:text-xl text-lg">{title}</h1>
    </div>
  );
}

export default function Features() {
  return (
    <div className="mt-10">
      <div className="header">
        <div className="flex justify-between">
          <img
            src="./images/mainweb/3D/Galaxy 1.png"
            alt=""
            className="w-[40px]"
          />
          <div className="flex flex-col justify-center text-center">
            <h1 className="text-2xl font-Ray-Black">
              ویژگی های شرکت <span className="text-main-blue-web">لاول کد</span>
            </h1>
            <span className="text-main-green-web text-sm">مختص شماست</span>
          </div>
          <img
            src="./images/mainweb/3D/Galaxy 2.png"
            alt=""
            className="w-[50px]"
          />
        </div>
      </div>
      <div className="allMain text-center">
        <div className="main xl:grid md:grid-cols-12 mt-10 container ">
          <div className="right md:col-span-6 md:col-start-1 md:col-end-7 md:mr-[70px] lg:mr-[200px] xl:mr-[53px] 2xl:mr-[96px]">
            <div className="box md:h-[540px] max-w-[580px] h-[500px] bg-[#FFEBD3] rounded-3xl flex justify-center flex-col items-center">
              <div className="img max-w-[300px] h-[250px]">
                <img src="./images/mainweb/3D/Sec4/5.png" alt="" />
              </div>
              <div className="box2 bg-white w-full md:w-[450px] md:h-[230px] h-[200px] rounded-3xl flex flex-col items-center gap-6 mt-[20px]">
                <div className="mt-4">
                  <h1 className="font-Ray-Black">
                    درصد رضایت و آمار سایت{" "}
                    <span className="text-main-blue-web">لاول کد </span> از زمان
                    ایجاد آن
                  </h1>
                </div>
                <div className="font-Ray-Bold">
                  <span className="font-ANJOMANFANUM-MEDIUM text-sm">
                    97.6 درصد
                  </span>
                  <span className="mt-2 md:mt-0  md:mr-[120px] mr-[50px]">نرخ رضایت مشتریان</span>
                  <hr className="w-[100%]" />
                </div>
                <div className="font-Ray-Bold">
                  <span className="font-ANJOMANFANUM-MEDIUM text-sm">
                    3000+
                  </span>
                  <span className="mt-2 md:mt-0 md:mr-[145px] mr-[80px]">مشتری از سراسر ایران</span>
                  <hr className="w-[100%]" />
                </div>
                <div className="font-Ray-Bold">
                  <span className="font-ANJOMANFANUM-MEDIUM text-sm">
                    50000+
                  </span>
                  <span className="mt-2 md:mt-0 md:mr-[110px] mr-[50px]">کاربران فعال محصولات ما</span>
                </div>
              </div>
            </div>
          </div>
          <div className="left md:col-start-7 md:col-end-13 mt-[50px] md:mt-[50px] xl:mt-[0] ">
            <div className="flex flex-wrap gap-5 lg:mr-[30px] justify-center items-center text-center ">
              <FeatureCard
                imageSrc="./images/mainweb/3D/Sec4/1.png"
                title="بهره مندی از متخصصین"
              />
              <FeatureCard
                imageSrc="./images/mainweb/3D/Sec4/2.png"
                title="امنیت بالا"
              />
              <FeatureCard
                imageSrc="./images/mainweb/3D/Sec4/3.png"
                title="پشتیبانی سریع"
              />
              <FeatureCard
                imageSrc="./images/mainweb/3D/Sec4/4.png"
                title="متعدد به زمان"
              />
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
