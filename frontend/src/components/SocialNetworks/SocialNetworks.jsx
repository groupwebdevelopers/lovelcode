import React from "react";

const socialNetworksData = [
  {
    title: "مارا در اینستاگرام دنبال کنید.",
    buttonText: "دنبال کنید",
    imageSrc: "./images/mainweb/3D/Sec5/Instagram1.png",
    backgroundColor: "main-red-web",
  },
  {
    title: "ارسال پیام در تلگرام به لاول کد.",
    buttonText: "ارسال پیام",
    imageSrc: "./images/mainweb/3D/Sec5/telegram.png",
    backgroundColor: "main-blue-web",
  },
  {
    title: "ارسال پیام در واتساپ به لاول کد.",
    buttonText: "ارسال پیام",
    imageSrc: "./images/mainweb/3D/Sec5/Whatsapp4.png",
    backgroundColor: "main-green-web",
  },
];

export default function SocialNetworks() {
  return (
    <div className="container">
      <div className="all mt-10 xl:flex xl:justify-between xl:flex-row flex flex-col items-center lg:mr-[0] xl:mr-[0]">
        {socialNetworksData.map((network, id) => (
          <div
            key={id}
            className={`box bg-${network.backgroundColor} max-w-[385px] rounded-3xl h-[120px] flex mt-[50px]`}
          >
            <div className="pb w-[40%] m-[20px]">
              <h1 className="text-white font-Ray-Bold">{network.title}</h1>
              <button className="text-black bg-white w-[100px] h-[35px] rounded-xl mt-[5px] text-sm font-Ray-Bold">
                {network.buttonText}
              </button>
            </div>
            <div className="img flex w-[60%]">
              <img
                src={network.imageSrc}
                alt=""
                className="w-[150px] h-[150px] mr-[35px] xl:mr-[20px] -mt-[40px]"
              />
              <img src={network.imageSrc} alt="" className="w-[50px] h-[50px] mt-[40px]" />
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
