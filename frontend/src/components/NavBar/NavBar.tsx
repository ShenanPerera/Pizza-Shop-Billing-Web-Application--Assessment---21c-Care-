// import React from 'react';
// import pizza from '../../assets/pizza.svg';
// import topping from '../../assets/cheese.svg';
// import beverage from '../../assets/soft-drink.svg';
// import bill from '../../assets/bill.svg';
// // import pizzaShopLogo from '../../../public/pizza-shop.svg';
// import '../../App.css';
// import { Link } from 'react-router-dom';

// // Define the props interface
// interface SidenavProps {
//   activeItem: string;
//   onItemClick: (itemId: string) => void;
// }

// const Sidenav: React.FC<SidenavProps> = ({ activeItem, onItemClick }) => {
//   const navItems = [
//     { id: 'pizza', label: 'Pizza', icon: pizza , link: '/pizza' },
//     { id: 'toppings', label: 'Toppings', icon: topping , link: '/toppings' },
//     { id: 'beverages', label: 'Beverages', icon: beverage , link: '/beverages' },
//     { id: 'bill', label: 'Bill', icon: bill , link: '/bill' }
//   ];

//   return (
//     <div className='sidenav'>
//       <div>
//         <div>
//           {navItems.map((item) => (
//             <button
//               key={item.id}
//               onClick={() => onItemClick(item.id)}
//               className={`w-full flex items-center space-x-3 px-4 py-3 rounded-lg text-left bg-red-600 ${
//                 activeItem === item.id
//                   ? 'bg-red-600 text-white'
//                   : 'hover:bg-gray-700 text-gray-300'
//               }`}
//             >
//               <div className='w-6 h-6 sm:w-8 sm:h-8 flex-shrink-0'>
//                 <img 
//                   src={item.icon} 
//                   alt={item.label}
//                   className='object-contain' 
//                   style={{ width: '10%', height: '10%'}}
//                 />
//               <span className='text-sm sm:text-base md:text-lg'>{item.label}</span>
//             </div>
//             </button>
//           ))}
//         </div>
//       </div>
//     </div>
//   );
// };

// export default Sidenav;

import React from 'react';
import pizza from '../../assets/pizza.svg';
import topping from '../../assets/cheese.svg';
import beverage from '../../assets/soft-drink.svg';
import bill from '../../assets/bill.svg';
// import pizzaShopLogo from '../../../public/pizza-shop.svg';
import '../../App.css';
import { Link } from 'react-router-dom';

// Define the props interface
interface SidenavProps {
  activeItem: string;
  onItemClick: (itemId: string) => void;
}

const Sidenav: React.FC<SidenavProps> = ({ activeItem, onItemClick }) => {
  const navItems = [
    { id: 'pizza', label: 'Pizza', icon: pizza, link: '/pizza' },
    { id: 'toppings', label: 'Toppings', icon: topping, link: '/toppings' },
    { id: 'beverages', label: 'Beverages', icon: beverage, link: '/beverages' },
    { id: 'billing', label: 'Billing', icon: bill, link: '/billing' }
  ];

  return (
    <div className='sidenav'>
      <div>
        <div>
          {navItems.map((item) => (
            <Link
              key={item.id}
              to={item.link}
              onClick={() => onItemClick(item.id)}
              className={`w-full flex items-center space-x-3 px-4 py-3 rounded-lg text-left bg-red-600 no-underline ${
                activeItem === item.id
                  ? 'bg-red-600 text-white'
                  : 'hover:bg-gray-700 text-gray-300'
              }`}
              style={{ textDecoration: 'none' }}
            >
              <div className='w-6 h-6 sm:w-8 sm:h-8 flex-shrink-0'>
                <img
                  src={item.icon}
                  alt={item.label}
                  className='object-contain'
                  style={{ width: '10%', height: '10%'}}
                />
              <span className='text-sm sm:text-base md:text-lg'>{item.label}</span>
            </div>
            </Link>
          ))}
        </div>
      </div>
    </div>
  );
};

export default Sidenav;