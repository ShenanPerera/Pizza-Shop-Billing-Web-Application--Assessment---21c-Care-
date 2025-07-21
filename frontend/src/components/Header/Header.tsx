// import { useNavigate } from 'react-router-dom';
import pizzashoplogo from '../../assets/pizza-shop.svg';

const HeaderComponent = () => {
  return (
    <nav className="flex items-center justify-between w-full px-6 py-4 bg-black">
      {/* Left side - Logo and Company Name */}
      <div className="flex items-center space-x-3">
        <img 
          src={pizzashoplogo}
          alt="Pizza Shop Logo"
          style={{ width: '50px', height: '50px' }}
        />
        <h1 className="text-red-500 text-xl font-bold">
          Pizzazz
        </h1>
      </div>
      
      {/* Right side - Button */}
      <button className="bg-red-500 hover:bg-red-600 text-white px-6 py-2 rounded-lg font-medium transition-colors duration-200">
        Order Now
      </button>
    </nav>
  );
};

export default HeaderComponent;