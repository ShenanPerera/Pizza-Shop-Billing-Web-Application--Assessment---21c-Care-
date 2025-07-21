import React, { useState } from 'react';
// import PizzaPage from './pages/pizzaPage';
// import ToppingPage from './pages/topping';
// import BeveragePage from './pages/beveragePage';
import BillingPage from './pages/billingPage';
import Sidenav from './components/NavBar/NavBar';
import HeaderComponent from './components/Header/Header';
import Pizza from './components/pizza';
import ToppingsTable from './components/toppings';
import BeveragesTable from './components/beverages';

const App = () => {
  const [currentPage, setCurrentPage] = useState('pizza');

  const renderPage = () => {
    switch (currentPage) {
      case 'pizza':
        return <Pizza onNavigateToBilling={() => setCurrentPage('billing')} />;
      case 'toppings':
        return <ToppingsTable onNavigateToBilling={() => setCurrentPage('billing')} />;
      case 'beverages':
        return <BeveragesTable onNavigateToBilling={() => setCurrentPage('billing')} />;
      case 'billing':
        return <BillingPage />;
      default:
        return <Pizza onNavigateToBilling={() => setCurrentPage('billing')} />;
    }
  };

  return (
    <div>
      <HeaderComponent />
      <div className="flex flex-1">
        <div className="w-64 bg-black">
          <Sidenav activeItem={currentPage} onItemClick={setCurrentPage} />
        </div>
        <div className="flex-1 bg-gray-100 p-8 overflow-auto">{renderPage()}</div>
      </div>
    </div>
  );
};

export default App;
