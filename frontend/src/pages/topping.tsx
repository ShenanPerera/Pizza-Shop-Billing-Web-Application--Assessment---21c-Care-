import Toppings from "../components/toppings";

const ToppingPage = () => {
  return (
    <div>
      <div className="px-10 pt-10 ">
        <h1 className="font-sans font-bold text-2xl leading-7">
          pizza
        </h1>
      </div>
      <br />
      <div>
        <Toppings />
      </div>
    </div>
  );
};

export default ToppingPage;