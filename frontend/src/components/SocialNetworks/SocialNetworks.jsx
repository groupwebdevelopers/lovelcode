import React from "react";

const socialNetworksData = [
  {
    title: "مارا در اینستاگرام دنبال کنید.",
    buttonText: "دنبال کنید",
    imageSrc: "./images/mainweb/3D/Sec5/Instagram1.png",
    backgroundColor: "gradient-to-r from-[#E825D4] to-[#F87919]",
  },
  {
    title: "ارسال پیام در تلگرام به لاول کد.",
    buttonText: "ارسال پیام",
    imageSrc: "./images/mainweb/3D/Sec5/telegram.png",
    backgroundColor: "gradient-to-l from-[#00A0C5] to-[#00D2FC]",
  },
  {
    title: "ارسال پیام در واتساپ به لاول کد",
    buttonText: "ارسال پیام",
    imageSrc: "./images/mainweb/3D/Sec5/Whatsapp4.png",
    backgroundColor: "gradient-to-l from-[#4BB329] to-[#70E249]",
  },
];

export default function SocialNetworks() {
  return (
    <div className="container">
      <div className="all mt-10 xl:flex xl:justify-between xl:flex-row flex flex-col items-center ">
        {socialNetworksData.map((network, id) => (
          <div
            key={id}
            className={`box bg-${network.backgroundColor} w-[290px] md:w-[350px] rounded-3xl h-[120px] flex mt-[50px]`}
          >
            <div className="pb max-w-[40%] m-[20px]">
              <h1 className="text-white font-Ray-Bold">{network.title}</h1>
              <button className="text-black bg-white w-[100px] h-[35px] rounded-xl mt-[5px] text-sm font-Ray-Bold">
                {network.buttonText}
              </button>
            </div>
            <div className="img flex max-w-[60%]">
              <img
                src={network.imageSrc}
                alt=""
                className="w-[110px] md:w-[150px] md:h-[150px] h-[110px] mr-[0px] xl:mr-[0px] -mt-[30px] md:-mt-[50px]"
              />
              <img src={network.imageSrc} alt="" className="max-w-[50px] max-h-[50px] mt-[30px] -mr-[8px] md:-mr-[0]" />
            </div>
          </div>
        ))}
      </div>
    </div>
  );
}
