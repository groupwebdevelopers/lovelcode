import axios from "axios";
import { Field, Form, Formik } from "formik";
import React from "react";
import apiRequests from "../../Services/Axios/configs";

export default function FormTariffs() {
  return (
    <>
      <div className="mt-12">
        <Formik
          initialValues={{
            name: "",
            phone: "",
            email: "",
            typeOfWebSite: "",
            storage: "",
            meet: "",
            desc: "",
          }}
          onSubmit={(value) => {
            const newRequest = {
              name: value.name,
              phone: +(value.phone),
              email: value.email,
              typeOfWebSite: value.typeOfWebSite,
              storage: value.storage,
              meet: value.meet,
              desc: value.desc,
            };
            console.log(newRequest);
            apiRequests.post('/order-plan/create' , newRequest)
            .then(res=>console.log(res))
            .catch(res=>console.log(res))

            value.name = ''
          }}
        >
          {({ values, handleSubmit, handleChange, errors, touched }) => (
            <Form>
              <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-y-9 gap-x-5 ">
                <div>
                  <Field
                    type="text"
                    name="name"
                    id="name"
                    className="py-3 px-5 w-full rounded-xl placeholder:font-Ray-Bold placeholder:text-main-gray-text-web text-sm md:text-base outline-main-blue-web"
                    placeholder="نام ونام خانوادگی"
                  ></Field>
                </div>
                <div>
                  <Field
                    type="text"
                    name="phone"
                    id="phone"
                    className="py-3 px-5 w-full rounded-xl placeholder:font-Ray-Bold placeholder:text-main-gray-text-web text-sm md:text-base outline-main-blue-web"
                    placeholder="شماره تماس"
                  ></Field>
                </div>
                <div>
                  <Field
                    type="text"
                    name="email"
                    id="email"
                    className="py-3 px-5 w-full rounded-xl placeholder:font-Ray-Bold placeholder:text-main-gray-text-web text-sm md:text-base outline-main-blue-web"
                    placeholder="ایمیل"
                  ></Field>
                </div>
                <div>
                  <Field
                    type="text"
                    name="typeOfWebSite"
                    id="typeOfWebSite"
                    className="py-3 px-5 w-full rounded-xl placeholder:font-Ray-Bold placeholder:text-main-gray-text-web text-sm md:text-base outline-main-blue-web"
                    placeholder="قصد راه اندازی چه نوع وبسایتی رو داری؟"
                  ></Field>
                </div>
                <div className="relative">
                  <Field
                    as="select"
                    name="storage"
                    id="storage"
                    className=" w-full appearance-none py-3 px-5 rounded-xl placeholder:font-Ray-Bold placeholder:text-main-gray-text-web text-sm md:text-base outline-main-blue-web text-main-gray-text-web font-Ray-Bold "
                  >
                    <option value="0">فضای میزبانی</option>
                    <option value="100">100</option>
                    <option value="200">200</option>
                    <option value="300">300</option>
                  </Field>
                  <div className="absolute top-1/2 -translate-y-1/2 left-5">
                    <i className="bi bi-chevron-down"></i>
                  </div>
                </div>
                <div className="relative">
                  <Field
                    as="select"
                    name="meet"
                    id="meet"
                    className="appearance-none  py-3 px-5 rounded-xl w-full placeholder:font-Ray-Bold placeholder:text-main-gray-text-web text-sm md:text-base outline-main-blue-web  text-main-gray-text-web font-Ray-Bold "
                  >
                    <option value="0">نحوه آشنایی با سایت لاول کد</option>
                    <option value="friends">دوست</option>
                    <option value="family">فامیل</option>
                    <option value="online">فضای مجازی</option>
                  </Field>
                  <div className="absolute top-1/2 -translate-y-1/2 left-5">
                    <i className="bi bi-chevron-down"></i>
                  </div>
                </div>
                <div className="col-span-1 md:col-span-2 lg:col-span-3 ">
                  <Field
                    as="textarea"
                    name="desc"
                    id="desc"
                    className=" py-3 px-5 w-full h-52 rounded-xl placeholder:font-Ray-Bold placeholder:text-main-gray-text-web text-sm md:text-base outline-main-blue-web"
                    placeholder="هر توضیحی که گمان می کنی می تونه به ما توی بررسی پروژه کمک کنه رو برامون بنویس."
                  ></Field>
                </div>
              </div>
              <button
                type="submit"
                className="mt-8 bg-main-blue-web text-white px-6 py-3 rounded-xl"
              >
                ثبت درخاست
              </button>
            </Form>
          )}
        </Formik>
      </div>
    </>
  );
}
