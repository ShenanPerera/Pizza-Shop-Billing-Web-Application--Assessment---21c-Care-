import Beverages from "../components/beverages";

const BeveragePage = () => {
  return (
    <div>
      <div className="px-10 pt-10 ">
        <h1 className="font-sans font-bold text-2xl leading-7">
          pizza
        </h1>
      </div>
      <br />
      <div>
        <Beverages/>
      </div>
    </div>
  );
};

export default BeveragePage;