import React from "react";

function FeatureCard({ imageSrc, title }) {
  return (
    <div className="bg-white 2xl:bg-[#d6d6d68e]  md:w-[315px] xl:w-[300px] 2xl:w-[360px] 2xl:h-[310px] md:h-[270px] w-[190px] h-[200px] text-center rounded-2xl ">
      <img
        src={imageSrc}
        alt=""
        className="md:w-[200px] w-[140px] md:mr-[55px] xl:mr-[40px] mr-[25px] 2xl:mr-[75px] 2xl:mt-[20px]"
      />
      <h1 className="font-Ray-ExtraBold md:text-xl text-lg mt-[15px] ">{title}</h1>
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
        <div className="main xl:grid md:grid-cols-12 mt-7 container ">
          <div className="right md:col-span-6 md:col-start-1 md:col-end-7 md:mr-[42px] lg:mr-[185px] xl:mr-[0px] 2xl:mr-[px]">
            <div className="box md:h-[560px] 2xl:h-[645px] max-w-[735px] md:w-[655px] xl:w-[600px] 2xl:w-[720px] h-[500px] bg-[#FFEBD3] rounded-3xl flex justify-center flex-col items-center">
              <div className="img  w-[250px]  md:w-[350px] lg:w-[400px]  h-[250px] 2xl:w-[420px]  2xl:-mt-[40px]">
                <img src="./images/mainweb/3D/Sec4/5.png" alt="" />
              </div>
              <div className="box2 bg-white w-full md:w-[480px]  md:h-[230px] 2xl:h-[260px] 2xl:w-[550px] h-[200px] rounded-3xl flex flex-col items-center gap-6 2xl:gap-8 2xl:mt-[85px] mt-[50px]">
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
                  <span className="mt-2 md:mt-0  md:mr-[120px] mr-[100px] 2xl:mr-[220px]">نرخ رضایت مشتریان</span>
                  <hr className="w-[300px] 2xl:w-[400px]" />
                </div>
                <div className="font-Ray-Bold">
                  <span className="font-ANJOMANFANUM-MEDIUM text-sm">
                    3000+
                  </span>
                  <span className="mt-2 md:mt-0 md:mr-[145px] mr-[125px] 2xl:mr-[240px]">مشتری از سراسر ایران</span>
                  <hr className="w-[300px] 2xl:w-[400px]" />
                </div>
                <div className="font-Ray-Bold">
                  <span className="font-ANJOMANFANUM-MEDIUM text-sm">
                    50000+
                  </span>
                  <span className="mt-2 md:mt-0 md:mr-[110px] mr-[95px] 2xl:mr-[205px]">کاربران فعال محصولات ما</span>
                </div>
              </div>
            </div>
          </div>
          <div className="left md:col-start-7 md:col-end-13 mt-[50px] md:mt-[50px] xl:mt-[0] ">
            <div className="flex flex-wrap gap-6 2xl:gap-7 lg:mr-[30px] xl:mr-[0px] justify-center items-center text-center ">
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
