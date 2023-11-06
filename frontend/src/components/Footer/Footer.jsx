import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";
export default function Footer() {
  const [data, setData] = useState([]);
  useEffect(() => {
    fetch("https://thlearn.iran.liara.run/api/v1/mainpage/footer")
      .then((res) => res.json())
      .then((data) => setData(data.data));
  }, []);
  const getOrder = (order) => {
    return data.filter((item) => item.order === order);
  };
  const splitTextArr = (text) => {
    return text.split("\\n");
  };
  return (
    <>
      <div className="bg-third-gray-text-web/20 relative">
        <a
          href="#top"
          className="hidden absolute -top-3 right-1/2 bg-white w-9 h-9 sm:flex justify-center items-center rounded-full hover:text-white hover:bg-main-blue-web duration-300"
        >
          <i className="bi bi-chevron-up"></i>
        </a>
        <div className="bg-gradient-to-r from-main-blue-web to-main-violet-web pt-12 pb-5 text-white md:rounded-t-[50px] rounded-t-[10px]">
          <div className="container leading-8">
            <div className="grid lg:grid-cols-4 md:grid-cols-2 gap-y-16 grid-cols-1 md: ">
              <div className="flex justify-center">
                <div className="1 flex flex-col  gap-6">
                  <Link to="/" className="flex items-center gap-x-2.5">
                    <div className="flex items-center justify-center shadow-normal lg:h-[55px] lg:w-[55px] h-12 w-12 bg-white rounded-xl lg:rounded-2xl">
                      <img
                        src="./images/mainweb/3D/Sec1/path28.svg"
                        className="h-7 lg:h-[31px]"
                        alt=""
                      />
                    </div>
                    <div className="flex flex-col">
                      <div className="text-lg font-Ray-Black ">LovelCode</div>
                      <div className="text-sm font-Ray-ExtraBold">لاول کد</div>
                    </div>
                  </Link>
                  <div className="font-Ray-Bold text-sm max-w-[249px] lg:max-w-[200px]">
                    <p>{data.length && getOrder(1)[0].body}</p>
                  </div>
                </div>
              </div>
              <div className="flex justify-center">
                <div className="2 flex flex-col gap-6">
                  <h2 className="font-Ray-ExtraBold text-[22px]">
                    {data.length && getOrder(2)[0].section}
                  </h2>
                  <ul className="list-disc lg:mr-4 font-Ray-Bold flex flex-col gap-2">
                    {data.length &&
                      splitTextArr(getOrder(2)[0].body).map((item) => (
                        <a href="#">
                          <li>{item}</li>
                        </a>
                      ))}
                  </ul>
                </div>
              </div>
              <div className="flex justify-center">
                <div className="3 flex flex-col gap-6">
                  <h2 className=" font-Ray-ExtraBold text-[22px]">
                    {data.length && getOrder(3)[0].section}
                  </h2>
                  <ul className="list-disc lg:mr-4 font-Ray-Bold flex flex-col gap-2">
                    {data.length &&
                      splitTextArr(getOrder(2)[0].body).map((item) => (
                        <a href="#">
                          <li>{item}</li>
                        </a>
                      ))}
                  </ul>
                </div>
              </div>
              <div className="flex justify-center">
                <div className="4 flex flex-col gap-6 ">
                  <h2 className=" font-Ray-ExtraBold text-[22px]">
                    {data.length && getOrder(4)[0].section}
                  </h2>
                  <ul className="lg:mr-4 font-Ray-Bold">
                    <li className="flex gap-2 items-center">
                      <i className="bi bi-geo-alt"></i>
                      <p className="font-ANJOMANFANUM-MEDIUM">
                        {data.length && splitTextArr(getOrder(4)[0].body)[0]}
                      </p>
                    </li>
                    <li className="flex gap-2 items-center">
                      <i className="bi bi-telephone"></i>
                      <p className="font-ANJOMANFANUM-MEDIUM">
                        {data.length && splitTextArr(getOrder(4)[0].body)[1]}
                      </p>
                    </li>
                    <li className="flex gap-2 items-center">
                      <i className="bi bi-envelope"></i>
                      <p>
                        {data.length && splitTextArr(getOrder(4)[0].body)[2]}
                      </p>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
            <div className="down flex flex-col-reverse gap-4 sm:flex-row  justify-between items-center border-t mt-7 py-5">
              <div className="r">
                <p className="font-Ray-Bold text-sm">
                  {data.length && splitTextArr(getOrder(5)[0].body)[0]}
                </p>
              </div>
              <div className="l flex items-center gap-1">
                <p className="font-Ray-Bold text-sm">
                  {data.length && splitTextArr(getOrder(5)[0].body)[1]}
                </p>
                <a href={data.length && splitTextArr(getOrder(5)[0].body)[2]}>
                  <img
                    className="w-6"
                    src="./images/mainweb/3D/Sec5/Instagram1.png"
                    alt=""
                  />
                </a>
                <a href={data.length && splitTextArr(getOrder(5)[0].body)[3]}>
                  <img
                    className="w-6"
                    src="/images/mainweb/3D/Sec5/telegram.png"
                    alt=""
                  />
                </a>
              </div>
            </div>
          </div>
        </div>
      </div>
    </>
  );
}
