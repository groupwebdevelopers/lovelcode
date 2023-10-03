import React from "react";

export default function SocialNetworks() {
  return (
    <>
      <div>
        <div className="container">
          <div className="all mt-10 xl:flex xl:justify-between  md:mr-[200px] lg:mr-[0] xl:mr-[0] ">
            <div className="box bg-[#f61231] w-[400px] xl:w-[365px] rounded-3xl h-[120px] flex  mt-[50px]">
              <div className="pb w-[40%] m-[20px]  ">
                <h1 className="text-white font-Ray-Bold">مارا در اینستاگرام دنبال کنید.</h1>
                <button className="text-black bg-white w-[100px] h-[35px] rounded-xl mt-[5px] text-sm font-Ray-Bold">
                  دنبال کنید
                </button>
              </div>
              <div className="img flex w-[60%] ">
                <img
                  src="./images/mainweb/3D/Sec5/Instagram1.png"
                  alt=""
                  className="w-[150px] h-[150px] mr-[35px] xl:mr-[20px] -mt-[40px]"
                />
                <img
                  src="./images/mainweb/3D/Sec5/Instagram1.png"
                  alt=""
                  className="w-[50px] h-[50px] mt-[40px]"
                />
              </div>
            </div>

            <div className="box bg-main-blue-web w-[400px] xl:w-[365px] rounded-3xl h-[120px] flex  mt-[50px]">
              <div className="pb w-[40%] m-[20px]  ">
                <h1 className="text-white font-Ray-Bold">ارسال پیام در تلگرام به لاول کد.</h1>
                <button className="text-black bg-white w-[100px] h-[35px] rounded-xl mt-[5px] text-sm font-Ray-Bold">
                  ارسال پیام
                </button>
              </div>
              <div className="img flex w-[60%] ">
                <img
                  src="./images/mainweb/3D/Sec5/telegram.png"
                  alt=""
                  className="w-[150px] h-[150px] mr-[35px] xl:mr-[20px] -mt-[40px]"
                />
                <img
                  src="./images/mainweb/3D/Sec5/telegram.png"
                  alt=""
                  className="w-[50px] h-[50px] mt-[40px]"
                />
              </div>
            </div>

            <div className="box bg-[#45b945] w-[400px] xl:w-[365px] rounded-3xl h-[120px] flex  mt-[50px]">
              <div className="pb w-[40%] m-[20px]  ">
                <h1 className="text-white font-Ray-Bold">ارسال پیام در واتساپ به لاول کد.</h1>
                <button className="text-black bg-white w-[100px] h-[35px] rounded-xl mt-[5px] text-sm font-Ray-Bold">
                ارسال پیام
                </button>
              </div>
              <div className="img flex w-[60%] ">
                <img
                  src="./images/mainweb/3D/Sec5/Whatsapp4.png"
                  alt=""
                  className="w-[150px] h-[150px] mr-[35px] xl:mr-[20px] -mt-[40px]"
                />
                <img
                  src="./images/mainweb/3D/Sec5/Whatsapp4.png"
                  alt=""
                  className="w-[50px] h-[50px] mt-[40px]"
                />
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
