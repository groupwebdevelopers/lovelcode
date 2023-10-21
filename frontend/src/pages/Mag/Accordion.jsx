import { Accordion, AccordionItem } from "@nextui-org/react";
const AccordionComp = () => {
  const defaultContent =
    "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.";
  const itemClasses = {
    base: "py-0 w-full font-Ray-Bold text-sm",
    title: "font-normal text-medium",
    trigger:
      "px-5 py-0  data-[hover=true]:bg-default-100 rounded-lg h-14 flex items-center justify-between gap-2 bg-[#f3f3f3]",
    indicator:
      "text-medium -rotate-90 data-[open=true]:rotate-90 transition-all duration-300 ",
    content: "text-small px-5 pb-6",
  };
  return (
    <div className="col-span-12 md:col-span-4 lg:col-span-3">
      <div className="w-full  bg-white px-3 py-5 rounded-xl h-[45rem] hidden md:block">
        <div className="px-2 mb-2">
          <button className="w-full bg-main-blue-web text-white py-4 font-Ray-Bold text-sm rounded-xl">
            <i class="bi bi-funnel ml-2"></i>
            <span>دسته بندی مقالات</span>
          </button>
        </div>
        <Accordion variant="splitted" itemClasses={itemClasses}>
          <AccordionItem
            key="1"
            aria-label="Accordion 1"
            title={
              <div className="flex gap-2">
                <img
                  className="w-[12px] h-[17px]"
                  src="/images/mainweb/Icons/note 2418.png"
                />
                <span>اخبار</span>
              </div>
            }
          >
            {defaultContent}
          </AccordionItem>
          <AccordionItem
            key="2"
            aria-label="Accordion 2"
            title={
              <div className="flex gap-2">
                <img
                  className="w-[12px] h-[17px]"
                  src="/images/mainweb/Icons/note 2418.png"
                />
                <span>آموزش</span>
              </div>
            }
          >
            {defaultContent}
          </AccordionItem>
          <AccordionItem
            key="3"
            aria-label="Accordion 3"
            title={
              <div className="flex gap-2">
                <img
                  className="w-[12px] h-[17px]"
                  src="/images/mainweb/Icons/note 2418.png"
                />
                <span>آموزش های طراحی سایت</span>
              </div>
            }
          >
            {defaultContent}
          </AccordionItem>
          <AccordionItem
            key="4"
            aria-label="Accordion 4"
            title={
              <div className="flex gap-2">
                <img
                  className="w-[12px] h-[17px]"
                  src="/images/mainweb/Icons/note 2418.png"
                />
                <span>مقالات دیجیتال مارکتینگ</span>
              </div>
            }
          >
            {defaultContent}
          </AccordionItem>
          <AccordionItem
            key="5"
            aria-label="Accordion 5"
            title="مقالات برنامه نویسی"
          >
            {defaultContent}
          </AccordionItem>
          <AccordionItem
            key="6"
            aria-label="Accordion 6"
            title={
              <div className="flex gap-2">
                <img
                  className="w-[12px] h-[17px]"
                  src="/images/mainweb/Icons/note 2418.png"
                />
                <span>مقالات هاستینگ و سرور</span>
              </div>
            }
          >
            {defaultContent}
          </AccordionItem>
          <AccordionItem
            key="7"
            aria-label="Accordion 7"
            title={
              <div className="flex gap-2">
                <img
                  className="w-[12px] h-[17px]"
                  src="/images/mainweb/Icons/note 2418.png"
                />
                <span>مقالات سیو</span>
              </div>
            }
          >
            {defaultContent}
          </AccordionItem>
          <AccordionItem
            key="8"
            aria-label="Accordion 8"
            title={
              <div className="flex gap-2">
                <img
                  className="w-[12px] h-[17px]"
                  src="/images/mainweb/Icons/note 2418.png"
                />
                <span>مقالات رابط کاربری</span>
              </div>
            }
          >
            {defaultContent}
          </AccordionItem>
          <AccordionItem
            key="9"
            aria-label="Accordion 9"
            title={
              <div className="flex gap-2">
                <img
                  className="w-[12px] h-[17px]"
                  src="/images/mainweb/Icons/note 2418.png"
                />
                <span>مقالات تجربه کاربری</span>
              </div>
            }
          >
            {defaultContent}
          </AccordionItem>
          <AccordionItem
            key="10"
            aria-label="Accordion 10"
            title={
              <div className="flex gap-2">
                <img
                  className="w-[12px] h-[17px]"
                  src="/images/mainweb/Icons/note 2418.png"
                />
                <span>مقالات گرافیکی</span>
              </div>
            }
          >
            {defaultContent}
          </AccordionItem>
        </Accordion>
      </div>
      <div className="col-span-12 px-2 mb-2 md:hidden">
        <button className="w-full bg-main-blue-web text-white py-4 font-Ray-Bold text-sm rounded-xl">
          <i class="bi bi-funnel ml-2"></i>
          <span>دسته بندی مقالات</span>
        </button>
      </div>
    </div>
  );
};

export default AccordionComp;
