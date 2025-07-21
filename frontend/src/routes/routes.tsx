import { createBrowserRouter } from 'react-router-dom';
// import Layout from './components/Layout';
import PizzaPage from '../pages/pizzaPage';
import ToppingPage from '../pages/topping';
import BeveragePage from '../pages/beveragePage';
import BillingPage from '../pages/billingPage';
// import Bill from './pages/Bill';

export const router = createBrowserRouter([
  {
    path: '/',
    // element: <Layout />,
    children: [
      {
        path: 'pizza',
        element: <PizzaPage />,
      },
      {
        path: 'toppings',
        element: <ToppingPage />,
      },
      {
        path: 'beverages',
        element: <BeveragePage />,
      },
      {
        path: 'billing',
        element: <BillingPage />,
      },
    ],
  },
]);

// If you want to use React Router, also create a Layout component:
// Layout.jsx
/*
import React from 'react';
import { Outlet, useLocation } from 'react-router-dom';
import Sidenav from './Sidenav';

const Layout = () => {
  const location = useLocation();
  const currentPage = location.pathname === '/' ? 'pizza' : location.pathname.slice(1);

  return (
    <div className="min-h-screen bg-gray-100">
      <Sidenav activeItem={currentPage} />
      <div className="ml-64">
        <Outlet />
      </div>
    </div>
  );
};

export default Layout;
*/